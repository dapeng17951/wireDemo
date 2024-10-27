/*
 * @Description:
 * @Author: shangke
 * @Date: 2024/10/27
 * @LastEditTime: 2024/10/27
 * @LastEditors: shangke
 */
package main

type Server struct {
	Name string
	Age  int
}

type UserServer struct {
	Server Server
}

type UserServer2 struct {
	Server
}
