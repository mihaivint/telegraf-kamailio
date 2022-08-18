//go:build !custom || inputs || inputs.kamailio

package all

import _ "github.com/influxdata/telegraf/plugins/inputs/kamailio" // register plugin
