// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.21.12
// source: go-vroom/v1/solve.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	request "go-vroom/api/go-vroom/v1/request"
	response "go-vroom/api/go-vroom/v1/response"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSolveSolve = "/go_vroom.v1.Solve/Solve"

type SolveHTTPServer interface {
	Solve(context.Context, *request.SolveRequest) (*response.SolveResponse, error)
}

func RegisterSolveHTTPServer(s *http.Server, srv SolveHTTPServer) {
	r := s.Route("/")
	r.POST("/solve", _Solve_Solve0_HTTP_Handler(srv))
}

func _Solve_Solve0_HTTP_Handler(srv SolveHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in request.SolveRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSolveSolve)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Solve(ctx, req.(*request.SolveRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*response.SolveResponse)
		return ctx.Result(200, reply)
	}
}

type SolveHTTPClient interface {
	Solve(ctx context.Context, req *request.SolveRequest, opts ...http.CallOption) (rsp *response.SolveResponse, err error)
}

type SolveHTTPClientImpl struct {
	cc *http.Client
}

func NewSolveHTTPClient(client *http.Client) SolveHTTPClient {
	return &SolveHTTPClientImpl{client}
}

func (c *SolveHTTPClientImpl) Solve(ctx context.Context, in *request.SolveRequest, opts ...http.CallOption) (*response.SolveResponse, error) {
	var out response.SolveResponse
	pattern := "/solve"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSolveSolve))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
