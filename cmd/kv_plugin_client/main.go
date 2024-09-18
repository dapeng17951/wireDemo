/*
 * @Description:
 * @Author: shangke
 * @Date: 2024-09-18 21:49:11
 * @LastEditTime: 2024-09-18 22:11:04
 * @LastEditors: shangke
 */
package main

import (
	"wireDemo/plugins/kv_plugin"
	"wireDemo/shared/kv"

	"github.com/hashicorp/go-plugin"
)

// 1.插件client
func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: kv.Handshake,
		Plugins: map[string]plugin.Plugin{
			"kv": &kv.KVGRPCPlugin{Impl: &kv_plugin.Kv_client{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
