package oracle

import (
	"bytes"
	"gomask/tools"
)

type oraclePacket struct {
	lOne []byte
	sTwo []byte
	sPos []byte
}

type oracleParser struct {
	oraclePacket
	packet []byte
	sql    string
	index  int
	code   int
}

var passKeys = []string{
	"DUAL",
}

func (o oracleParser) exclude() bool {
	for _, v := range passKeys {
		if bytes.Contains(bytes.ToUpper(o.packet), bytes.ToUpper([]byte(v))) {
			return true
		}
	}
	return false
}

func (o *oracleParser) judge() {
	key := []byte("SELECT")
	o.index = tools.BytesIndexIgnoreCase(o.packet, key)
}

func (o *oracleParser) GetSql(packet []byte) string {
	o.packet = packet
	o.judge()
	if o.exclude() {
		return ""
	}

	return ""
}

func (o *oracleParser) GetPacket(sql string) []byte {
	o.sql = sql
	if sql[len(sql)-1] == 59 {
		o.sql = sql[:len(sql)-1]
	}
	if tools.ArrayContainsInteger([]int{70, 71, 72}, o.index) {

	}
	return []byte("Hello")
}
