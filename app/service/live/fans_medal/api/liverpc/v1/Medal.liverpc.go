// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/Medal.proto

package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "go-common/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports

// ===============
// Medal Interface
// ===============

type Medal interface {
	// * up主开启勋章
	//
	Create(context.Context, *MedalCreateReq) (*MedalCreateResp, error)

	// * 获取up主自己的勋章，更严格的类型
	//
	Query(context.Context, *MedalQueryReq) (*MedalQueryResp, error)

	// * 根据ID查询勋章详情，返回严格的类型
	//
	QueryId(context.Context, *MedalQueryIdReq) (*MedalQueryIdResp, error)
}

// =====================
// Medal Live Rpc Client
// =====================

type medalRpcClient struct {
	client *liverpc.Client
}

// NewMedalRpcClient creates a Rpc client that implements the Medal interface.
// It communicates using Rpc and can be configured with a custom HTTPClient.
func NewMedalRpcClient(client *liverpc.Client) Medal {
	return &medalRpcClient{
		client: client,
	}
}

func (c *medalRpcClient) Create(ctx context.Context, in *MedalCreateReq) (*MedalCreateResp, error) {
	out := new(MedalCreateResp)
	err := doRpcRequest(ctx, c.client, 1, "Medal.create", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medalRpcClient) Query(ctx context.Context, in *MedalQueryReq) (*MedalQueryResp, error) {
	out := new(MedalQueryResp)
	err := doRpcRequest(ctx, c.client, 1, "Medal.query", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medalRpcClient) QueryId(ctx context.Context, in *MedalQueryIdReq) (*MedalQueryIdResp, error) {
	out := new(MedalQueryIdResp)
	err := doRpcRequest(ctx, c.client, 1, "Medal.query_id", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
