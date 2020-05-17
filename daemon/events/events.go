package events // import "github.com/demonoid81/moby/daemon/events"

import (
	"sync"
	"time"

	eventtypes "github.com/demonoid81/moby/api/types/events"
	"github.com/demonoid81/moby/pkg/pubsub"
)

const (
	eventsLimit = 256
	bufferSize  = 1024
)

// Events is pubsub channel for events generated by the engine.
type Events struct {
	mu     sync.Mutex
	events []eventtypes.Message
	pub    *pubsub.Publisher
}

// New returns new *Events instance
func New() *Events {
	return &Events{
		events: make([]eventtypes.Message, 0, eventsLimit),
		pub:    pubsub.NewPublisher(100*time.Millisecond, bufferSize),
	}
}

// Subscribe adds new listener to events, returns slice of 256 stored
// last events, a channel in which you can expect new events (in form
// of interface{}, so you need type assertion), and a function to call
// to stop the stream of events.
func (e *Events) Subscribe() ([]eventtypes.Message, chan interface{}, func()) {
	eventSubscribers.Inc()
	e.mu.Lock()
	current := make([]eventtypes.Message, len(e.events))
	copy(current, e.events)
	l := e.pub.Subscribe()
	e.mu.Unlock()

	cancel := func() {
		e.Evict(l)
	}
	return current, l, cancel
}

// SubscribeTopic adds new listener to events, returns slice of 256 stored
// last events, a channel in which you can expect new events (in form
// of interface{}, so you need type assertion).
func (e *Events) SubscribeTopic(since, until time.Time, ef *Filter) ([]eventtypes.Message, chan interface{}) {
	eventSubscribers.Inc()
	e.mu.Lock()

	var topic func(m interface{}) bool
	if ef != nil && ef.filter.Len() > 0 {
		topic = func(m interface{}) bool { return ef.Include(m.(eventtypes.Message)) }
	}

	buffered := e.loadBufferedEvents(since, until, topic)

	var ch chan interface{}
	if topic != nil {
		ch = e.pub.SubscribeTopic(topic)
	} else {
		// Subscribe to all events if there are no filters
		ch = e.pub.Subscribe()
	}

	e.mu.Unlock()
	return buffered, ch
}

// Evict evicts listener from pubsub
func (e *Events) Evict(l chan interface{}) {
	eventSubscribers.Dec()
	e.pub.Evict(l)
}

// Log creates a local scope message and publishes it
func (e *Events) Log(action, eventType string, actor eventtypes.Actor) {
	now := time.Now().UTC()
	jm := eventtypes.Message{
		Action:   action,
		Type:     eventType,
		Actor:    actor,
		Scope:    "local",
		Time:     now.Unix(),
		TimeNano: now.UnixNano(),
	}

	// fill deprecated fields for container and images
	switch eventType {
	case eventtypes.ContainerEventType:
		jm.ID = actor.ID
		jm.Status = action
		jm.From = actor.Attributes["image"]
	case eventtypes.ImageEventType:
		jm.ID = actor.ID
		jm.Status = action
	}

	e.PublishMessage(jm)
}

// PublishMessage broadcasts event to listeners. Each listener has 100 milliseconds to
// receive the event or it will be skipped.
func (e *Events) PublishMessage(jm eventtypes.Message) {
	eventsCounter.Inc()

	e.mu.Lock()
	if len(e.events) == cap(e.events) {
		// discard oldest event
		copy(e.events, e.events[1:])
		e.events[len(e.events)-1] = jm
	} else {
		e.events = append(e.events, jm)
	}
	e.mu.Unlock()
	e.pub.Publish(jm)
}

// SubscribersCount returns number of event listeners
func (e *Events) SubscribersCount() int {
	return e.pub.Len()
}

// loadBufferedEvents iterates over the cached events in the buffer
// and returns those that were emitted between two specific dates.
// It uses `time.Unix(seconds, nanoseconds)` to generate valid dates with those arguments.
// It filters those buffered messages with a topic function if it's not nil, otherwise it adds all messages.
func (e *Events) loadBufferedEvents(since, until time.Time, topic func(interface{}) bool) []eventtypes.Message {
	var buffered []eventtypes.Message
	if since.IsZero() && until.IsZero() {
		return buffered
	}

	var sinceNanoUnix int64
	if !since.IsZero() {
		sinceNanoUnix = since.UnixNano()
	}

	var untilNanoUnix int64
	if !until.IsZero() {
		untilNanoUnix = until.UnixNano()
	}

	for i := len(e.events) - 1; i >= 0; i-- {
		ev := e.events[i]

		if ev.TimeNano < sinceNanoUnix {
			break
		}

		if untilNanoUnix > 0 && ev.TimeNano > untilNanoUnix {
			continue
		}

		if topic == nil || topic(ev) {
			buffered = append([]eventtypes.Message{ev}, buffered...)
		}
	}
	return buffered
}
