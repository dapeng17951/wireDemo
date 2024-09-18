/*
 * @Description:
 * @Author: shangke
 * @Date: 2024-09-18 21:20:26
 * @LastEditTime: 2024-09-18 21:46:12
 * @LastEditors: shangke
 */
package kv

import (
	"context"
	"wireDemo/protos/kv"
)

// GRPCClient is an implementation of KV that talks over RPC.
//1.定义客户端
//引入 protos/kv
//实现interface中的接口定义
//内部实现interface接口定义至grpc server定义的转换

type GRPCClient struct{ client kv.KVClient }

func (m *GRPCClient) Put(key string, value []byte) error {
	_, err := m.client.Put(context.Background(), &kv.PutRequest{
		Key:   key,
		Value: value,
	})
	return err
}

func (m *GRPCClient) Get(key string) ([]byte, error) {
	resp, err := m.client.Get(context.Background(), &kv.GetRequest{
		Key: key,
	})
	if err != nil {
		return nil, err
	}

	return resp.Value, nil
}

// Here is the gRPC server that GRPCClient talks to.
// 2.定义服务端
// 引入 protos/kv
// 实现interface中的接口定义
// 内部实现interface接口定义至grpc server定义的转换
type GRPCServer struct {
	// This is the real implementation
	Impl                     KV
	kv.UnimplementedKVServer //重点：新版，必须加上，因为生成的code中多了该项
}

func (m *GRPCServer) Put(
	ctx context.Context,
	req *kv.PutRequest) (*kv.Empty, error) {
	return &kv.Empty{}, m.Impl.Put(req.Key, req.Value)
}

func (m *GRPCServer) Get(
	ctx context.Context,
	req *kv.GetRequest) (*kv.GetResponse, error) {
	v, err := m.Impl.Get(req.Key)
	return &kv.GetResponse{Value: v}, err
}
