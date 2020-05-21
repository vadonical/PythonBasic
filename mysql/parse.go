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
	// 常用的
	base := []int{5}
	if tools.ArrayContainsInteger(base, m.index) && m.code == 0 {
		sql := string(m.packet[m.index:])
		newSql := replace(sql)
		return m.construct(newSql)
	}

	return packet
}

func (m *mysqlParser) construct(sql string) []byte {
	m.sTwo = append(m.sTwo, m.packet[4])
	m.rSql = []byte(sql)
	m.lOne = tools.Number2Bytes(len(sql)+1, 4, true)

	return m.composite()
}

func (m *mysqlParser) composite() []byte {
	var packet []byte

	packet = append(packet, m.lOne...)
	packet = append(packet, m.sTwo...)
	packet = append(packet, m.rSql...)

	return packet
}
