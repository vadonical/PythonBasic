package analysis

import (
	"bytes"
	"gomask/tools"
	"log"
)

type pgSQLPacket struct {
	sOne []byte
	lTwo []byte
	sThr []byte
	rSql []byte
	sPos []byte
}

type PgSQLParser struct {
	pgSQLPacket
	packet []byte
	sql    string
	index  int
	code   int
}

var PassKeys = []string{
	"pg_",
	"INFORMATION_SCHEMA",
}

func (p PgSQLParser) exclude() bool {
	for _, v := range PassKeys {
		if bytes.Contains(bytes.ToUpper(p.packet), bytes.ToUpper([]byte(v))) {
			return true
		}
	}
	return false
}

func (p *PgSQLParser) determine() {
	key := []byte("select")
	p.index = tools.BytesIndexIgnoreCase(p.packet, key)
}

func (p *PgSQLParser) GetSql(packet []byte) string {
	// 为结构体赋值
	p.packet = packet
	// 确定种类
	p.determine()
	// 判断类型进行不同的处理
	if p.exclude() {
		return ""
	}
	if p.index == 6 && p.code == 0 {
		p.rSql = p.packet[6 : len(p.packet)-38]
		p.sql = string(p.rSql)
		return p.sql
	}
	log.Fatalln("new protocol not resolved in pgsql")
	return ""
}

func (p *PgSQLParser) GetPacket(statement string) []byte {
	p.sql = statement
	if statement[len(statement)-1] == 59 {
		p.sql = statement[:len(statement)-1]
	}
	if p.index == 6 && p.code == 0 {
		return p.constructOne()
	}
	return nil
}

// for dbv
func (p *PgSQLParser) constructOne() []byte {
	p.sOne = append(p.sOne, p.packet[0])
	p.sThr = append(p.sThr, p.packet[5])
	p.sPos = p.packet[len(p.packet)-38:]

	p.lTwo = tools.Number2Bytes(len(p.sql)+8, 4, false)
	p.rSql = []byte(p.sql)

	return p.construct()
}

func (p PgSQLParser) construct() []byte {
	var packet []byte

	packet = append(packet, p.sOne...)
	packet = append(packet, p.lTwo...)
	packet = append(packet, p.sThr...)
	packet = append(packet, p.rSql...)
	packet = append(packet, p.sPos...)

	return packet
}
