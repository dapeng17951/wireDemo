/*
 * @Description:
 * @Author: shangke
 * @Date: 2024/10/6
 * @LastEditTime: 2024/10/6
 * @LastEditors: shangke
 */
package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"os"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	for {

	}
}
