package analysis

import (
	"bytes"
	"fmt"
	"gomask/replace"
	"net"
	"strings"
)

var LimitLength = 20

const (
	Oracle = iota
	Mysql
	Postgres
	SqlServer
	DB2
	DM
)

type Stream struct {
	conn      net.IPConn
	packet    []byte
	uppPacket string
}

// byte流转大写字符，主要用来判断数据类型
//func (s *Stream) upperPacket() {
//	s.uppPacket = strings.ToUpper(string(s.packet))
//}

// byte流转字符串
func (s *Stream) strPacket() string {
	return string(s.packet)
}

// 分发数据流，根据不同的数据种类来处理
func (s *Stream) Distribute(p []byte) []byte {
	s.uppPacket = string(bytes.ToUpper(p))

	if len(p) > LimitLength && s.isQuery() && s.isCheck() {
		fmt.Println("ININININIINNININININI")
		s.packet = p
		// 确定数据库种类
		kind := s.determineType()
		switch kind {
		case 1:
			fmt.Println("PgSQL-----------------------------------")
			var pg PgSQLParser
			sql := pg.GetSql(p)
			fmt.Println("SQL:", sql)
			if len(sql) > 0 {
				// 随机获取到了值
				statement := replace.Statements[0]
				fmt.Println("EXECUTE:", statement)
				s.packet = pg.GetPacket(statement)
			}
		case 2:
			fmt.Println("MySQL-----------------------------------")
			var my mysqlParser
			sql := my.GetSql(p)
			fmt.Println("MySQL SQL:", sql)
			if len(sql) > 0 {
				sql := replace.Statements[1]
				fmt.Println("EXECUTE:", sql)
				s.packet = my.GetPacket(sql)
			}

		case 3:
			fmt.Println("Oracle")
		case 4:
			fmt.Println("SQLite")
		default:
			fmt.Println("new sql not resolved")
		}

		return s.packet
	}
	return p
}

// 判断是否是插叙类语句
func (s *Stream) isQuery() bool {
	fmt.Println(1111111111)
	keys := []string{"SELECT", "FROM"}
	fmt.Println(s.uppPacket)
	fmt.Println(strings.Contains(s.uppPacket, "SELECT"))
	fmt.Println(strings.Contains(s.uppPacket, "FROM"))
	for _, v := range keys {
		if !strings.Contains(s.uppPacket, v) {
			fmt.Println(333332333)
			return false
		}
	}
	fmt.Println(2222222222)
	return true
}

// 校验校验IP是否正确且具有权限
func (s *Stream) isCheck() bool {
	return true
}

// 判断语句中数据库的种类
// 1 PgSQL
// 2 MySQL
// 3 Oracle
// 4 SQLite
func (s *Stream) determineType() int {
	return 2
}
