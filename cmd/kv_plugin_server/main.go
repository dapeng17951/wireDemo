/*
 * @Description:
 * @Author: shangke
 * @Date: 2024-09-18 21:48:52
 * @LastEditTime: 2024-09-22 17:21:23
 * @LastEditors: shangke
 */
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"wireDemo/shared/kv"

	"github.com/hashicorp/go-plugin"
)

func run(path string) error {
	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: kv.Handshake,
		Plugins:         kv.PluginMap,
		//Cmd:             exec.Command("sh", "-c", os.Getenv("KV_PLUGIN")), //调用 plugin 进程，这里使用环境变量执行的
		Cmd: exec.Command(path + "\\" + "client.exe"),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC},
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

	result, err := kv.Get("1")
	if err != nil {
		return err
	}

	fmt.Println(string(result))

	err2 := kv.Put("1", []byte("2"))
	if err2 != nil {
		return err2
	}

	return fmt.Errorf("Please only use 'get' or 'put', given: %q", "1")
}

// 服务端
func main() {
	// We don't want to see the plugin logs.
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	//https://learnku.com/go/wikis/23398

	if err := run(exPath); err != nil {
		fmt.Printf("error: %+v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
