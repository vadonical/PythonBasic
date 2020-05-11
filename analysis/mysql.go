package analysis

import (
	"bytes"
	"gomask/tools"
	"log"
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

func (m mysqlParser) exclude() bool {
	for _, v := range passKeys {
		if bytes.Contains(bytes.ToUpper(m.packet), bytes.ToUpper([]byte(v))) {
			return true
		}
	}
	return false
}

func (m *mysqlParser) determine() {
	key := []byte("select")
	m.index = tools.BytesIndexIgnoreCase(m.packet, key)
}

func (m *mysqlParser) GetSql(packet []byte) string {
	m.packet = packet
	m.determine()
	if m.exclude() {
		return ""
	}
	if m.index == 5 && m.code == 0 {
		m.rSql = m.packet[5:]
		m.sql = string(m.rSql)
		return m.sql
	}
	log.Fatalln("new protocol not resolved in mysql")

	return ""
}

func (m *mysqlParser) GetPacket(sql string) []byte {
	m.sql = sql
	if sql[len(sql)-1] == 59 {
		m.sql = sql[:len(sql)-1]
	}
	if m.index == 5 && m.code == 0 {
		return m.constructOne()
	}
	return nil
}

func (m *mysqlParser) constructOne() []byte {
	m.lOne = tools.Number2Bytes(len(m.sql)+1, 4, true)
	m.sTwo = append(m.sTwo, m.packet[4])
	m.rSql = []byte(m.sql)

	return m.construct()

}

func (m mysqlParser) construct() []byte {
	var packet []byte

	packet = append(packet, m.lOne...)
	packet = append(packet, m.sTwo...)
	packet = append(packet, m.rSql...)

	return packet
}
