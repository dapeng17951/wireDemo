/*
 * @Description:
 * @Author: shangke
 * @Date: 2024-09-18 21:20:21
 * @LastEditTime: 2024-09-18 21:47:10
 * @LastEditors: shangke
 */
package kv

import (
	"context"

	"wireDemo/protos/kv"

	"github.com/hashicorp/go-plugin"

	"google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
// 6.协议版本信息
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "kv",
}

// PluginMap is the map of plugins we can dispense.
// 5. 映射
var PluginMap = map[string]plugin.Plugin{
	"kv_grpc": &KVGRPCPlugin{},
}

// 1.定义接口
type KV interface {
	Put(key string, value []byte) error
	Get(key string) ([]byte, error)
}

// 2.定义grpc-plugin
// 引入 go-plugin
type KVGRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl KV
}

// 3.定义服务
// 引入：kv_grpc.pb.go
// 引入：shared 下的 grpc.go
func (p *KVGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	kv.RegisterKVServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

// 4.定义服务
// 引入：kv_grpc.pb.go
// 引入：shared 下的 grpc.go
func (p *KVGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: kv.NewKVClient(c)}, nil
}
