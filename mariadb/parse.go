package mariadb

import (
	"bytes"
	"gomask/tools"
)

type mariadbPacket struct {
	lOne []byte
	sTwo []byte
	rSql []byte
}

type mariadbParser struct {
	mariadbPacket
	packet []byte
	sql    string
	index  int
	code   int
}

var passKeys = []string{
	"INFORMATION_SCHEMA",
}

func (m *mariadbParser) exclude() bool {
	for _, v := range passKeys {
		if bytes.Contains(bytes.ToUpper(m.packet), bytes.ToUpper([]byte(v))) {
			return true
		}
	}
	return false
}

func (m *mariadbParser) determine() {
	m.index = tools.BytesIndexIgnoreCase(m.packet, []byte("select"))
}

func (m *mariadbParser) Dispatch(packet []byte) []byte {
	m.packet = packet
	m.determine()
	if m.exclude() {
		return packet
	}
	base := []int{5}
	if tools.ArrayContainsInteger(base, m.index) && m.code == 0 {
		//sql := string(m.packet[m.index:])
		newSql := "select * from data"
		return m.construct(newSql)

	}
	return packet
}

func (m *mariadbParser) construct(sql string) []byte {
	m.sTwo = append(m.sTwo, m.packet[4])
	m.rSql = []byte(sql)
	m.lOne = tools.Number2Bytes(len(sql)+1, 4, true)
	return m.composite()
}

func (m *mariadbParser) composite() []byte {
	var packet []byte

	packet = append(packet, m.lOne...)
	packet = append(packet, m.sTwo...)
	packet = append(packet, m.rSql...)

	return packet
}
