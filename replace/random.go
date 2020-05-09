package replace

import "strings"

var Statements []string = []string{
	"SELECT * FROM \"zx_schema\".\"zx_table\" WHERE \"id\" = 1",
	"select * from data",
	"select * from data1",
}

func test(statement string, mode int, length int, repl bool) {
	/*
	mode: 1 select statement
	mode: 2 create table with

	*/

	// extract statement
	if strings.HasPrefix(strings.ToUpper(statement), "SELECT") {
		// select statement
	} else {
		// create statement
	}

	// select

	// create table

	// create view
}
