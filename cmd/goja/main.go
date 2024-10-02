/*
 * @Description:
 * @Author: shangke
 * @Date: 2024-10-02 17:16:28
 * @LastEditTime: 2024-10-02 18:13:17
 * @LastEditors: shangke
 */
package main

import (
	"fmt"

	"github.com/dop251/goja"
)

func main() {

	//params := "var a=8; var b=10;"
	const SCRIPT = `	
		function sum(a, b) {
			return a + b;
		}
		function show(a){
			return a
		}
		function p(p)  {
			var temp=JSON.parse(p)
			return temp[1]
		}
		function pp(p)  {
			var temp=JSON.parse(p)
			return temp.p.name
		}
		`

	vm := goja.New()
	_, err := vm.RunString(SCRIPT)
	if err != nil {
		panic(err)
	}
	sum, ok := goja.AssertFunction(vm.Get("sum"))
	if !ok {
		panic("Not a function")
	}
	res, err := sum(goja.Undefined(), vm.ToValue(40), vm.ToValue(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	show, ok := goja.AssertFunction(vm.Get("show"))
	if !ok {
		panic("Not a function")
	}
	ret, err := show(goja.Undefined(), vm.ToValue("hello goja"))
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)

	//解析json测试
	p, ok := goja.AssertFunction(vm.Get("p"))
	if !ok {
		panic("Not a function")
	}
	ret, err = p(goja.Undefined(), vm.ToValue("[100,\"hello\",3.14]"))
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)

	fmt.Println(ret.ExportType())

	//json测试：对象字符串
	pp, ok := goja.AssertFunction(vm.Get("pp"))
	if !ok {
		panic("Not a function")
	}
	ret, err = pp(goja.Undefined(), vm.ToValue("{\"name\":\"sk\",\"age\":18,\"sm\":[10,\"sm\",3.14],\"p\":{\"name\":\"x\",\"age\":18}}"))
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)

	fmt.Println(ret.ExportType())
	for {
	}
}
