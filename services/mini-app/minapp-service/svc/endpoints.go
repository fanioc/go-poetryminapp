// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 574fb16d86
// Version Date: 2019年 04月 12日 星期五 00:42:59 UTC

package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats.

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	pb "github.com/fanioc/go-poetryminapp/services/mini-app"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	LoginEndpoint            endpoint.Endpoint
	CheckUserSessionEndpoint endpoint.Endpoint
	GetUserInfoEndpoint      endpoint.Endpoint
	UpdateUserInfoEndpoint   endpoint.Endpoint
	GetUserConfigEndpoint    endpoint.Endpoint
	SetUserConfigEndpoint    endpoint.Endpoint
}

// Endpoints

func (e Endpoints) Login(ctx context.Context, in *pb.LoginParams) (*pb.Session, error) {
	response, err := e.LoginEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Session), nil
}

func (e Endpoints) CheckUserSession(ctx context.Context, in *pb.CheckSessionParams) (*pb.CheckSession, error) {
	response, err := e.CheckUserSessionEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.CheckSession), nil
}

func (e Endpoints) GetUserInfo(ctx context.Context, in *pb.GetUserInfoParams) (*pb.UserInfo, error) {
	response, err := e.GetUserInfoEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.UserInfo), nil
}

func (e Endpoints) UpdateUserInfo(ctx context.Context, in *pb.UpdateUserInfoParams) (*pb.ErrCode, error) {
	response, err := e.UpdateUserInfoEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.ErrCode), nil
}

func (e Endpoints) GetUserConfig(ctx context.Context, in *pb.GetUserConfigParams) (*pb.UserCofing, error) {
	response, err := e.GetUserConfigEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.UserCofing), nil
}

func (e Endpoints) SetUserConfig(ctx context.Context, in *pb.SetUserConfigParams) (*pb.ErrCode, error) {
	response, err := e.SetUserConfigEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.ErrCode), nil
}

// Make Endpoints

func MakeLoginEndpoint(s pb.MinAppServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.LoginParams)
		v, err := s.Login(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeCheckUserSessionEndpoint(s pb.MinAppServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.CheckSessionParams)
		v, err := s.CheckUserSession(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeGetUserInfoEndpoint(s pb.MinAppServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.GetUserInfoParams)
		v, err := s.GetUserInfo(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeUpdateUserInfoEndpoint(s pb.MinAppServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.UpdateUserInfoParams)
		v, err := s.UpdateUserInfo(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeGetUserConfigEndpoint(s pb.MinAppServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.GetUserConfigParams)
		v, err := s.GetUserConfig(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeSetUserConfigEndpoint(s pb.MinAppServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.SetUserConfigParams)
		v, err := s.SetUserConfig(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

// WrapAllExcept wraps each Endpoint field of struct Endpoints with a
// go-kit/kit/endpoint.Middleware.
// Use this for applying a set of middlewares to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllExcept(middleware, "Status", "Ping")
func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"Login":            struct{}{},
		"CheckUserSession": struct{}{},
		"GetUserInfo":      struct{}{},
		"UpdateUserInfo":   struct{}{},
		"GetUserConfig":    struct{}{},
		"SetUserConfig":    struct{}{},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc, _ := range included {
		if inc == "Login" {
			e.LoginEndpoint = middleware(e.LoginEndpoint)
		}
		if inc == "CheckUserSession" {
			e.CheckUserSessionEndpoint = middleware(e.CheckUserSessionEndpoint)
		}
		if inc == "GetUserInfo" {
			e.GetUserInfoEndpoint = middleware(e.GetUserInfoEndpoint)
		}
		if inc == "UpdateUserInfo" {
			e.UpdateUserInfoEndpoint = middleware(e.UpdateUserInfoEndpoint)
		}
		if inc == "GetUserConfig" {
			e.GetUserConfigEndpoint = middleware(e.GetUserConfigEndpoint)
		}
		if inc == "SetUserConfig" {
			e.SetUserConfigEndpoint = middleware(e.SetUserConfigEndpoint)
		}
	}
}

// LabeledMiddleware will get passed the endpoint name when passed to
// WrapAllLabeledExcept, this can be used to write a generic metrics
// middleware which can send the endpoint name to the metrics collector.
type LabeledMiddleware func(string, endpoint.Endpoint) endpoint.Endpoint

// WrapAllLabeledExcept wraps each Endpoint field of struct Endpoints with a
// LabeledMiddleware, which will receive the name of the endpoint. See
// LabeldMiddleware. See method WrapAllExept for details on excluded
// functionality.
func (e *Endpoints) WrapAllLabeledExcept(middleware func(string, endpoint.Endpoint) endpoint.Endpoint, excluded ...string) {
	included := map[string]struct{}{
		"Login":            struct{}{},
		"CheckUserSession": struct{}{},
		"GetUserInfo":      struct{}{},
		"UpdateUserInfo":   struct{}{},
		"GetUserConfig":    struct{}{},
		"SetUserConfig":    struct{}{},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc, _ := range included {
		if inc == "Login" {
			e.LoginEndpoint = middleware("Login", e.LoginEndpoint)
		}
		if inc == "CheckUserSession" {
			e.CheckUserSessionEndpoint = middleware("CheckUserSession", e.CheckUserSessionEndpoint)
		}
		if inc == "GetUserInfo" {
			e.GetUserInfoEndpoint = middleware("GetUserInfo", e.GetUserInfoEndpoint)
		}
		if inc == "UpdateUserInfo" {
			e.UpdateUserInfoEndpoint = middleware("UpdateUserInfo", e.UpdateUserInfoEndpoint)
		}
		if inc == "GetUserConfig" {
			e.GetUserConfigEndpoint = middleware("GetUserConfig", e.GetUserConfigEndpoint)
		}
		if inc == "SetUserConfig" {
			e.SetUserConfigEndpoint = middleware("SetUserConfig", e.SetUserConfigEndpoint)
		}
	}
}