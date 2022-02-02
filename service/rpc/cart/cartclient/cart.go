// Code generated by goctl. DO NOT EDIT!
// Source: cart.proto

package cartclient

import (
	"context"

	"github.com/skywalkerwei/pluton-faster/service/rpc/cart/cart"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddReq   = cart.AddReq
	AddResp  = cart.AddResp
	CartItem = cart.CartItem
	DelReq   = cart.DelReq
	DelResp  = cart.DelResp
	EditReq  = cart.EditReq
	EditResp = cart.EditResp
	ListReq  = cart.ListReq
	ListResp = cart.ListResp

	Cart interface {
		List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error)
		Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddResp, error)
		Edit(ctx context.Context, in *EditReq, opts ...grpc.CallOption) (*EditResp, error)
		Del(ctx context.Context, in *DelReq, opts ...grpc.CallOption) (*DelResp, error)
	}

	defaultCart struct {
		cli zrpc.Client
	}
)

func NewCart(cli zrpc.Client) Cart {
	return &defaultCart{
		cli: cli,
	}
}

func (m *defaultCart) List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.List(ctx, in, opts...)
}

func (m *defaultCart) Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddResp, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.Add(ctx, in, opts...)
}

func (m *defaultCart) Edit(ctx context.Context, in *EditReq, opts ...grpc.CallOption) (*EditResp, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.Edit(ctx, in, opts...)
}

func (m *defaultCart) Del(ctx context.Context, in *DelReq, opts ...grpc.CallOption) (*DelResp, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.Del(ctx, in, opts...)
}
