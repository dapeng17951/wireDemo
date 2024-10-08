/*
 * @Description:
 * @Author: shangke
 * @Date: 2024/10/8
 * @LastEditTime: 2024/10/8
 * @LastEditors: shangke
 */
package main

import (
	"fmt"
	duck "github.com/scottlepp/go-duck/duck"
)

func main() {
	tempopt := duck.Opts{Exe: "duckdb.exe"}
	db := duck.NewDuckDB("foo.db", tempopt)

	commands := []string{
		"CREATE TABLE t1 (i INTEGER, j INTEGER);",
		"INSERT INTO t1 VALUES (1, 5);",
	}
	_, err := db.RunCommands(commands)
	if err != nil {
		fmt.Println(err.Error())
	}
	res, err := db.Query("SELECT * from t1;")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf(res)
	}

	for {
	}

}
