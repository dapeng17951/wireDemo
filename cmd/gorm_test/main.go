/*
 * @Description:
 * @Author: shangke
 * @Date: 2024/10/7
 * @LastEditTime: 2024/10/7
 * @LastEditors: shangke
 */
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//相对路径问题：针对工程目录下的文件
//gorm open 时，针对sqlite会自动创建数据库文件

func main() {
	constr := "db/s.db"
	_, err := gorm.Open(sqlite.Open(constr), &gorm.Config{})
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("db open ok")
	}
	for {
	}

}
