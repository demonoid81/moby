package daemon // import "github.com/demonoid81/moby/daemon"

import (
	// Importing packages here only to make sure their init gets called and
	// therefore they register themselves to the logdriver factory.
	_ "github.com/demonoid81/moby/daemon/logger/awslogs"
	_ "github.com/demonoid81/moby/daemon/logger/etwlogs"
	_ "github.com/demonoid81/moby/daemon/logger/fluentd"
	_ "github.com/demonoid81/moby/daemon/logger/gcplogs"
	_ "github.com/demonoid81/moby/daemon/logger/gelf"
	_ "github.com/demonoid81/moby/daemon/logger/jsonfilelog"
	_ "github.com/demonoid81/moby/daemon/logger/logentries"
	_ "github.com/demonoid81/moby/daemon/logger/loggerutils/cache"
	_ "github.com/demonoid81/moby/daemon/logger/splunk"
	_ "github.com/demonoid81/moby/daemon/logger/syslog"
)
