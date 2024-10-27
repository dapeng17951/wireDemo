/*
 * @Description:
 * @Author: shangke
 * @Date: 2024/10/27
 * @LastEditTime: 2024/10/27
 * @LastEditors: shangke
 */
package main

import "fmt"

func main() {
	user := UserServer{}
	user.Server.Name = "xx"
	user.Server.Age = 18
	fmt.Print(user)
	fmt.Println()
	user2 := UserServer2{}
	user2.Name = "xxx"
	user2.Age = 19
	fmt.Print(user2)

}
