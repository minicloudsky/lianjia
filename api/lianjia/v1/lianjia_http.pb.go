// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.3.1
// - protoc             v4.25.1
// source: lianjia/v1/lianjia.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationLianjiaListErshoufang = "/lianjia.v1.Lianjia/ListErshoufang"

type LianjiaHTTPServer interface {
	ListErshoufang(context.Context, *ListErshoufangRequest) (*ListErshoufangReply, error)
}

func RegisterLianjiaHTTPServer(s *http.Server, srv LianjiaHTTPServer) {
	r := s.Route("/")
	r.GET("/lianjia/ershoufang/list", _Lianjia_ListErshoufang0_HTTP_Handler(srv))
}

func _Lianjia_ListErshoufang0_HTTP_Handler(srv LianjiaHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListErshoufangRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLianjiaListErshoufang)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListErshoufang(ctx, req.(*ListErshoufangRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListErshoufangReply)
		return ctx.Result(200, reply)
	}
}

type LianjiaHTTPClient interface {
	ListErshoufang(ctx context.Context, req *ListErshoufangRequest, opts ...http.CallOption) (rsp *ListErshoufangReply, err error)
}

type LianjiaHTTPClientImpl struct {
	cc *http.Client
}

func NewLianjiaHTTPClient(client *http.Client) LianjiaHTTPClient {
	return &LianjiaHTTPClientImpl{client}
}

func (c *LianjiaHTTPClientImpl) ListErshoufang(ctx context.Context, in *ListErshoufangRequest, opts ...http.CallOption) (*ListErshoufangReply, error) {
	var out ListErshoufangReply
	pattern := "/lianjia/ershoufang/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLianjiaListErshoufang))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}