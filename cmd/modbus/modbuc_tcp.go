/*
 * @Description:
 * @Author: shangke
 * @Date: 2024/10/5
 * @LastEditTime: 2024/10/5
 * @LastEditors: shangke
 */
package main

import (
	"fmt"
	modbus "github.com/grid-x/modbus"
	"strconv"
	"time"
)

func main() {
	handler := modbus.NewTCPClientHandler("127.0.0.1:502")
	handler.SlaveID = 1
	handler.Timeout = 100 * time.Millisecond
	err := handler.Connect()
	defer handler.Close()
	if err != nil {
		fmt.Printf("连接错误")
	}
	client := modbus.NewClient(handler)

	//读线圈:批量
	bres, err := client.ReadCoils(0, 5)
	if err != nil {
		fmt.Printf("线圈批量读取错误", err.Error())
	}
	b := bres[0]
	bs := strconv.FormatInt(int64(b), 2)
	bvs := []bool{}
	for index, bv := range bs {
		fmt.Print(index)
		fmt.Print(bv)
		if bs[len(bs)-1-index] == '0' {
			bvs = append(bvs, false)
		}
		if bs[len(bs)-1-index] == '1' {
			bvs = append(bvs, true)
		}
	}
	fmt.Print(bvs)
	//写线圈
	//写线圈 off 0x0000 on 0xff00
	bw, err := client.WriteSingleCoil(0, 0xFF00)
	if err != nil {
		fmt.Printf("单写线圈错误", err.Error())
	}
	fmt.Print(bw)
	//写线圈，批量
	bws, err := client.WriteMultipleCoils(0, 5, []byte{0x01, 0x01, 0x01, 0x01, 0x01})
	if err != nil {
		fmt.Printf("多写线圈错误", err.Error())
	}
	fmt.Print(bws)
	for {

	}
}
