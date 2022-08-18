//go:generate ../../../tools/readme_config_includer/generator
package kamailio

import (
	binrpc "github.com/florentchauveau/go-kamailio-binrpc/v3"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"net"
	"strconv"
	"strings"
	_ "embed"
)

type Kamailio struct {
	Socket  string `toml:"socket"`
	Modules string `toml:"modules"`
	Nodeid  string `toml:"nodeid"`
}

//go:embed sample.conf
var sampleConfig string

func (k *Kamailio) Description() string {
	return "Kamailio telegraf input plugin"
}

func (k *Kamailio) SampleConfig() string {
	return sampleConfig
}
func recordGet(conn net.Conn, method string) ([]binrpc.Record, error) {
	// WritePacket returns the cookie generated
	cookie, err := binrpc.WritePacket(conn, method)
	if err != nil {
		panic(err)
	}

	records, err := binrpc.ReadPacket(conn, cookie)
	if err != nil {
		panic(err)
	}
	return records, err
}

func processRecords(records []binrpc.Record, module string, fields map[string]interface{}) {
	suffix := ""
	nrRecords := len(records)
	for i, record := range records {
		items, _ := record.StructItems()
		if nrRecords > 1 {
			suffix = "." + strconv.Itoa(i)
		}
		for _, item := range items {
			value, err := item.Value.Int()
			if err == nil {
				fields[module+"."+item.Key+suffix] = value
				continue
			}
			valued, err := item.Value.Double()
			if err == nil {
				fields[module+"."+item.Key+suffix] = valued
				continue
			}
			values, err := item.Value.String()
			if err == nil {
				fields[module+"."+item.Key+suffix] = values
			}

		}
	}
}

func (k *Kamailio) Gather(acc telegraf.Accumulator) error {
	var conn net.Conn
	var err error
	var records []binrpc.Record

	tags := make(map[string]string)
	fields := make(map[string]interface{})
	//establish connection to Kamailio server
	if len(k.Socket) > 0 {
		sock := strings.Split(k.Socket, ":")
		conn, err = net.Dial(sock[0], sock[1])
	} else {
		//fall back to default
		conn, err = net.Dial("unix", "/tmp/kamailio_ctl")
	}

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for _, module := range strings.Fields(k.Modules) {
		records, _ = recordGet(conn, module)
		processRecords(records, module, fields)
	}

	if k.Nodeid != "" {
		tags["nodeid"] = k.Nodeid
	}

	acc.AddFields("kamailio", fields, tags)
	return nil
}

func init() {
	inputs.Add("kamailio", func() telegraf.Input {
		return &Kamailio{Socket: "unix:/tmp/kamailio_ctl", Modules: "tm.stats core.shmmem core.tcp_info sl.stats tls.info"}
	})
}
