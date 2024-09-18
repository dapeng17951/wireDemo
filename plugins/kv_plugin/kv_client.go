/*
 * @Description:
 * @Author: shangke
 * @Date: 2024-09-18 21:50:28
 * @LastEditTime: 2024-09-18 22:10:55
 * @LastEditors: shangke
 */
package kv_plugin

import (
	"fmt"
)

//具体插件实现

type Kv_client struct{}

func (Kv_client) Put(key string, value []byte) error {
	fmt.Println("client")
	return nil
}

func (Kv_client) Get(key string) ([]byte, error) {
	s := "kv_client"
	return []byte(s), nil
}
