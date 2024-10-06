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
	"os/exec"
)

func main() {
	command := exec.Command("go", "env")
	//err := command.Run()			//仅执行
	//if err != nil {
	//	fmt.Printf("cmd error")
	//}
	output, err := command.CombinedOutput() //执行并反馈结果
	if err != nil {
		fmt.Printf("cmd error")
	}
	fmt.Println(string(output))
	for {
	}
}
