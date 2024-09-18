/*
 * @Description:
 * @Author: shangke
 * @Date: 2024-09-18 21:48:52
 * @LastEditTime: 2024-09-18 22:17:29
 * @LastEditors: shangke
 */
package main

import (
	"fmt"
	"os"
	"os/exec"
	"wireDemo/shared/kv"

	"github.com/hashicorp/go-plugin"
)

func run() error {
	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: kv.Handshake,
		Plugins:         kv.PluginMap,
		//Cmd:             exec.Command("sh", "-c", os.Getenv("KV_PLUGIN")), //调用 plugin 进程，这里使用环境变量执行的
		Cmd: exec.Command("sh", "-c", "client.exe"),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		return err
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("kv_grpc") //interface.go 中注册的名
	if err != nil {
		return err
	}

	// We should have a KV store now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	kv := raw.(kv.KV)
	os.Args = os.Args[1:]
	switch os.Args[0] {
	case "get":
		result, err := kv.Get(os.Args[1])
		if err != nil {
			return err
		}

		fmt.Println(string(result))

	case "put":
		err := kv.Put(os.Args[1], []byte(os.Args[2]))
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("Please only use 'get' or 'put', given: %q", os.Args[0])
	}

	return nil
}

// 服务端
func main() {
	// We don't want to see the plugin logs.

	if err := run(); err != nil {
		fmt.Printf("error: %+v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
