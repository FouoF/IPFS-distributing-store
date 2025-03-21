// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.3
// - protoc             v4.25.1
// source: admin-service/v1/operation.proto

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

const OperationOperationAddEndpoint = "/adminservice.v1.Operation/AddEndpoint"
const OperationOperationAddNode = "/adminservice.v1.Operation/AddNode"
const OperationOperationCreateIndex = "/adminservice.v1.Operation/CreateIndex"
const OperationOperationGetEndpoint = "/adminservice.v1.Operation/GetEndpoint"
const OperationOperationListEndpoint = "/adminservice.v1.Operation/ListEndpoint"
const OperationOperationListIndex = "/adminservice.v1.Operation/ListIndex"
const OperationOperationListLeaf = "/adminservice.v1.Operation/ListLeaf"
const OperationOperationListNode = "/adminservice.v1.Operation/ListNode"
const OperationOperationRemoveEndpoint = "/adminservice.v1.Operation/RemoveEndpoint"
const OperationOperationRemoveNode = "/adminservice.v1.Operation/RemoveNode"

type OperationHTTPServer interface {
	AddEndpoint(context.Context, *AddEndpointRequest) (*AddEndpointReply, error)
	AddNode(context.Context, *AddNodeRequest) (*AddNodeReply, error)
	CreateIndex(context.Context, *CreateIndexRequest) (*CreateIndexReply, error)
	GetEndpoint(context.Context, *GetEndpointRequest) (*GetEndpointReply, error)
	ListEndpoint(context.Context, *ListEndpointRequest) (*ListEndpointReply, error)
	ListIndex(context.Context, *ListIndexRequest) (*ListIndexReply, error)
	ListLeaf(context.Context, *ListLeafRequest) (*ListLeafReply, error)
	ListNode(context.Context, *ListNodeRequest) (*ListNodeReply, error)
	RemoveEndpoint(context.Context, *RemoveEndpointRequest) (*RemoveEndpointReply, error)
	RemoveNode(context.Context, *RemoveNodeRequest) (*RemoveNodeReply, error)
}

func RegisterOperationHTTPServer(s *http.Server, srv OperationHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/node/list", _Operation_ListNode0_HTTP_Handler(srv))
	r.POST("/admin/node/add", _Operation_AddNode0_HTTP_Handler(srv))
	r.POST("/admin/node/remove", _Operation_RemoveNode0_HTTP_Handler(srv))
	r.GET("/admin/endpoint/list", _Operation_ListEndpoint0_HTTP_Handler(srv))
	r.GET("/admin/endpoint/get", _Operation_GetEndpoint0_HTTP_Handler(srv))
	r.POST("/admin/endpoint/add", _Operation_AddEndpoint0_HTTP_Handler(srv))
	r.POST("/admin/endpoint/remove", _Operation_RemoveEndpoint0_HTTP_Handler(srv))
	r.POST("/admin/index/create", _Operation_CreateIndex0_HTTP_Handler(srv))
	r.GET("/admin/index/list", _Operation_ListIndex0_HTTP_Handler(srv))
	r.GET("/admin/index/listleaf", _Operation_ListLeaf0_HTTP_Handler(srv))
}

func _Operation_ListNode0_HTTP_Handler(srv OperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListNodeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOperationListNode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListNode(ctx, req.(*ListNodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListNodeReply)
		return ctx.Result(200, reply)
	}
}

func _Operation_AddNode0_HTTP_Handler(srv OperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddNodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOperationAddNode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddNode(ctx, req.(*AddNodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddNodeReply)
		return ctx.Result(200, reply)
	}
}

func _Operation_RemoveNode0_HTTP_Handler(srv OperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RemoveNodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOperationRemoveNode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RemoveNode(ctx, req.(*RemoveNodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RemoveNodeReply)
		return ctx.Result(200, reply)
	}
}

func _Operation_ListEndpoint0_HTTP_Handler(srv OperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListEndpointRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOperationListEndpoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListEndpoint(ctx, req.(*ListEndpointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListEndpointReply)
		return ctx.Result(200, reply)
	}
}

func _Operation_GetEndpoint0_HTTP_Handler(srv OperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetEndpointRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOperationGetEndpoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetEndpoint(ctx, req.(*GetEndpointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetEndpointReply)
		return ctx.Result(200, reply)
	}
}

func _Operation_AddEndpoint0_HTTP_Handler(srv OperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddEndpointRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOperationAddEndpoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddEndpoint(ctx, req.(*AddEndpointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddEndpointReply)
		return ctx.Result(200, reply)
	}
}

func _Operation_RemoveEndpoint0_HTTP_Handler(srv OperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RemoveEndpointRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOperationRemoveEndpoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RemoveEndpoint(ctx, req.(*RemoveEndpointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RemoveEndpointReply)
		return ctx.Result(200, reply)
	}
}

func _Operation_CreateIndex0_HTTP_Handler(srv OperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateIndexRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOperationCreateIndex)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateIndex(ctx, req.(*CreateIndexRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateIndexReply)
		return ctx.Result(200, reply)
	}
}

func _Operation_ListIndex0_HTTP_Handler(srv OperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListIndexRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOperationListIndex)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListIndex(ctx, req.(*ListIndexRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListIndexReply)
		return ctx.Result(200, reply)
	}
}

func _Operation_ListLeaf0_HTTP_Handler(srv OperationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListLeafRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOperationListLeaf)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListLeaf(ctx, req.(*ListLeafRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListLeafReply)
		return ctx.Result(200, reply)
	}
}

type OperationHTTPClient interface {
	AddEndpoint(ctx context.Context, req *AddEndpointRequest, opts ...http.CallOption) (rsp *AddEndpointReply, err error)
	AddNode(ctx context.Context, req *AddNodeRequest, opts ...http.CallOption) (rsp *AddNodeReply, err error)
	CreateIndex(ctx context.Context, req *CreateIndexRequest, opts ...http.CallOption) (rsp *CreateIndexReply, err error)
	GetEndpoint(ctx context.Context, req *GetEndpointRequest, opts ...http.CallOption) (rsp *GetEndpointReply, err error)
	ListEndpoint(ctx context.Context, req *ListEndpointRequest, opts ...http.CallOption) (rsp *ListEndpointReply, err error)
	ListIndex(ctx context.Context, req *ListIndexRequest, opts ...http.CallOption) (rsp *ListIndexReply, err error)
	ListLeaf(ctx context.Context, req *ListLeafRequest, opts ...http.CallOption) (rsp *ListLeafReply, err error)
	ListNode(ctx context.Context, req *ListNodeRequest, opts ...http.CallOption) (rsp *ListNodeReply, err error)
	RemoveEndpoint(ctx context.Context, req *RemoveEndpointRequest, opts ...http.CallOption) (rsp *RemoveEndpointReply, err error)
	RemoveNode(ctx context.Context, req *RemoveNodeRequest, opts ...http.CallOption) (rsp *RemoveNodeReply, err error)
}

type OperationHTTPClientImpl struct {
	cc *http.Client
}

func NewOperationHTTPClient(client *http.Client) OperationHTTPClient {
	return &OperationHTTPClientImpl{client}
}

func (c *OperationHTTPClientImpl) AddEndpoint(ctx context.Context, in *AddEndpointRequest, opts ...http.CallOption) (*AddEndpointReply, error) {
	var out AddEndpointReply
	pattern := "/admin/endpoint/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOperationAddEndpoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OperationHTTPClientImpl) AddNode(ctx context.Context, in *AddNodeRequest, opts ...http.CallOption) (*AddNodeReply, error) {
	var out AddNodeReply
	pattern := "/admin/node/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOperationAddNode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OperationHTTPClientImpl) CreateIndex(ctx context.Context, in *CreateIndexRequest, opts ...http.CallOption) (*CreateIndexReply, error) {
	var out CreateIndexReply
	pattern := "/admin/index/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOperationCreateIndex))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OperationHTTPClientImpl) GetEndpoint(ctx context.Context, in *GetEndpointRequest, opts ...http.CallOption) (*GetEndpointReply, error) {
	var out GetEndpointReply
	pattern := "/admin/endpoint/get"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOperationGetEndpoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OperationHTTPClientImpl) ListEndpoint(ctx context.Context, in *ListEndpointRequest, opts ...http.CallOption) (*ListEndpointReply, error) {
	var out ListEndpointReply
	pattern := "/admin/endpoint/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOperationListEndpoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OperationHTTPClientImpl) ListIndex(ctx context.Context, in *ListIndexRequest, opts ...http.CallOption) (*ListIndexReply, error) {
	var out ListIndexReply
	pattern := "/admin/index/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOperationListIndex))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OperationHTTPClientImpl) ListLeaf(ctx context.Context, in *ListLeafRequest, opts ...http.CallOption) (*ListLeafReply, error) {
	var out ListLeafReply
	pattern := "/admin/index/listleaf"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOperationListLeaf))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OperationHTTPClientImpl) ListNode(ctx context.Context, in *ListNodeRequest, opts ...http.CallOption) (*ListNodeReply, error) {
	var out ListNodeReply
	pattern := "/admin/node/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOperationListNode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OperationHTTPClientImpl) RemoveEndpoint(ctx context.Context, in *RemoveEndpointRequest, opts ...http.CallOption) (*RemoveEndpointReply, error) {
	var out RemoveEndpointReply
	pattern := "/admin/endpoint/remove"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOperationRemoveEndpoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OperationHTTPClientImpl) RemoveNode(ctx context.Context, in *RemoveNodeRequest, opts ...http.CallOption) (*RemoveNodeReply, error) {
	var out RemoveNodeReply
	pattern := "/admin/node/remove"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOperationRemoveNode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
