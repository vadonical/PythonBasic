package mysql

import (
	"bytes"
	"gomask/tools"
)

type mysqlPacket struct {
	lOne []byte
	sTwo []byte
	rSql []byte
}

type mysqlParser struct {
	mysqlPacket
	packet []byte
	sql    string
	index  int
	code   int
}

var passKeys = []string{
	"INFORMATION_SCHEMA",
}

func (m *mysqlParser) exclude() bool {
	for _, v := range passKeys {
		if bytes.Contains(bytes.ToUpper(m.packet), bytes.ToUpper([]byte(v))) {
			return true
		}
	}
	return false
}

func (m *mysqlParser) determine() {
	m.index = tools.BytesIndexIgnoreCase(m.packet, []byte("select"))
}

func (m *mysqlParser) Dispatch(packet []byte) []byte {
	m.packet = packet
	m.determine()
	if m.exclude() {
		return packet
	}
	if m.index == 5 && m.code == 0 {

	}
	return packet
}
