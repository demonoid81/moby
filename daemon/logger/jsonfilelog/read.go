package jsonfilelog // import "github.com/demonoid81/moby/daemon/logger/jsonfilelog"

import (
	"context"
	"encoding/json"
	"io"

	"github.com/demonoid81/moby/api/types/backend"
	"github.com/demonoid81/moby/daemon/logger"
	"github.com/demonoid81/moby/daemon/logger/jsonfilelog/jsonlog"
	"github.com/demonoid81/moby/daemon/logger/loggerutils"
	"github.com/demonoid81/moby/pkg/tailfile"
	"github.com/sirupsen/logrus"
)

const maxJSONDecodeRetry = 20000

// ReadLogs implements the logger's LogReader interface for the logs
// created by this driver.
func (l *JSONFileLogger) ReadLogs(config logger.ReadConfig) *logger.LogWatcher {
	logWatcher := logger.NewLogWatcher()

	go l.readLogs(logWatcher, config)
	return logWatcher
}

func (l *JSONFileLogger) readLogs(watcher *logger.LogWatcher, config logger.ReadConfig) {
	defer close(watcher.Msg)

	l.mu.Lock()
	l.readers[watcher] = struct{}{}
	l.mu.Unlock()

	l.writer.ReadLogs(config, watcher)

	l.mu.Lock()
	delete(l.readers, watcher)
	l.mu.Unlock()
}

func decodeLogLine(dec *json.Decoder, l *jsonlog.JSONLog) (*logger.Message, error) {
	l.Reset()
	if err := dec.Decode(l); err != nil {
		return nil, err
	}

	var attrs []backend.LogAttr
	if len(l.Attrs) != 0 {
		attrs = make([]backend.LogAttr, 0, len(l.Attrs))
		for k, v := range l.Attrs {
			attrs = append(attrs, backend.LogAttr{Key: k, Value: v})
		}
	}
	msg := &logger.Message{
		Source:    l.Stream,
		Timestamp: l.Created,
		Line:      []byte(l.Log),
		Attrs:     attrs,
	}
	return msg, nil
}

type decoder struct {
	rdr io.Reader
	dec *json.Decoder
	jl  *jsonlog.JSONLog
}

func (d *decoder) Reset(rdr io.Reader) {
	d.rdr = rdr
	d.dec = nil
	if d.jl != nil {
		d.jl.Reset()
	}
}

func (d *decoder) Close() {
	d.dec = nil
	d.rdr = nil
	d.jl = nil
}

func (d *decoder) Decode() (msg *logger.Message, err error) {
	if d.dec == nil {
		d.dec = json.NewDecoder(d.rdr)
	}
	if d.jl == nil {
		d.jl = &jsonlog.JSONLog{}
	}
	for retries := 0; retries < maxJSONDecodeRetry; retries++ {
		msg, err = decodeLogLine(d.dec, d.jl)
		if err == nil || err == io.EOF {
			break
		}

		logrus.WithError(err).WithField("retries", retries).Warn("got error while decoding json")
		// try again, could be due to a an incomplete json object as we read
		if _, ok := err.(*json.SyntaxError); ok {
			d.dec = json.NewDecoder(d.rdr)
			continue
		}

		// io.ErrUnexpectedEOF is returned from json.Decoder when there is
		// remaining data in the parser's buffer while an io.EOF occurs.
		// If the json logger writes a partial json log entry to the disk
		// while at the same time the decoder tries to decode it, the race condition happens.
		if err == io.ErrUnexpectedEOF {
			d.rdr = io.MultiReader(d.dec.Buffered(), d.rdr)
			d.dec = json.NewDecoder(d.rdr)
			continue
		}
	}
	return msg, err
}

// decodeFunc is used to create a decoder for the log file reader
func decodeFunc(rdr io.Reader) loggerutils.Decoder {
	return &decoder{
		rdr: rdr,
		dec: nil,
		jl:  nil,
	}
}

func getTailReader(ctx context.Context, r loggerutils.SizeReaderAt, req int) (io.Reader, int, error) {
	return tailfile.NewTailReader(ctx, r, req)
}
