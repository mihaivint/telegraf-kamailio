package kamailio

import (
	binrpc "github.com/florentchauveau/go-kamailio-binrpc/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

var ProcessRecords = processRecords

func TestProcessRecords(t *testing.T) {
	fields := make(map[string]interface{})
	var testStructItems = make([]binrpc.StructItem, 3)
	testStructItems[0].Key = "max_connections"
	record, _ := binrpc.CreateRecord(2048)
	testStructItems[0].Value = *record
	testStructItems[1].Key = "opened_connections"
	record, _ = binrpc.CreateRecord(0)
	testStructItems[1].Value = *record
	testStructItems[2].Key = "clear_text_write_queued_bytes"
	record, _ = binrpc.CreateRecord(0)
	testStructItems[2].Value = *record
	var testRecords []binrpc.Record = make([]binrpc.Record, 1)
	testRecords[0].Type = uint8(0x3)
	testRecords[0].Value = testStructItems

	ProcessRecords(testRecords, "tls.info", fields)
	require.Equal(t, map[string]interface{}{"tls.info.clear_text_write_queued_bytes": 0, "tls.info.max_connections": 2048, "tls.info.opened_connections": 0}, fields)
}
