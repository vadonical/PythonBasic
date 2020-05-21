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

// 某些查询类语句可能是系统语句，这个时候不能改变
func (o *oracleParser) exclude() bool {
	for _, v := range passKeys {
		if bytes.Contains(bytes.ToUpper(o.packet), bytes.ToUpper([]byte(v))) {
			return true
		}
	}
	return false
}

// 判断语句的位置，code字段再此主要用作区别
func (o *oracleParser) determine() {
	key := []byte("SELECT")
	o.index = tools.BytesIndexIgnoreCase(o.packet, key)
}

// 查询语句分发
func (o *oracleParser) Dispatch(packet []byte) []byte {
	o.packet = packet
	o.determine()

	if o.exclude() {
		return packet
	}

	return packet
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
