package egrpc

import (
	"context"

	"github.com/alibaba/sentinel-golang/core/base"
	"google.golang.org/grpc"

	"github.com/1477921168/ego/core/elog"
)

// Option overrides a Container's default configuration.
type Option func(c *Container)

// WithServerOption inject server option to grpc server
// User should not inject interceptor option, which is recommend by WithStreamInterceptor
// and WithUnaryInterceptor
func WithServerOption(options ...grpc.ServerOption) Option {
	return func(c *Container) {
		if c.config.serverOptions == nil {
			c.config.serverOptions = make([]grpc.ServerOption, 0)
		}
		c.config.serverOptions = append(c.config.serverOptions, options...)
	}
}

// WithStreamInterceptor inject stream interceptors to server option
func WithStreamInterceptor(interceptors ...grpc.StreamServerInterceptor) Option {
	return func(c *Container) {
		if c.config.streamInterceptors == nil {
			c.config.streamInterceptors = make([]grpc.StreamServerInterceptor, 0)
		}

		c.config.streamInterceptors = append(c.config.streamInterceptors, interceptors...)
	}
}

// WithUnaryInterceptor inject unary interceptors to server option
func WithUnaryInterceptor(interceptors ...grpc.UnaryServerInterceptor) Option {
	return func(c *Container) {
		if c.config.unaryInterceptors == nil {
			c.config.unaryInterceptors = make([]grpc.UnaryServerInterceptor, 0)
		}
		c.config.unaryInterceptors = append(c.config.unaryInterceptors, interceptors...)
	}
}

// WithUnaryServerResourceExtractor sets the resource extractor of unary server request.
func WithUnaryServerResourceExtractor(fn func(context.Context, interface{}, *grpc.UnaryServerInfo) string) Option {
	return func(c *Container) {
		c.config.unaryServerResourceExtract = fn
	}
}

// WithUnaryServerBlockFallback sets the block fallback handler of unary server request.
func WithUnaryServerBlockFallback(fn func(context.Context, interface{}, *grpc.UnaryServerInfo, *base.BlockError) (interface{}, error)) Option {
	return func(c *Container) {
		c.config.unaryServerBlockFallback = fn
	}
}

// WithNetwork inject network
func WithNetwork(network string) Option {
	return func(c *Container) {
		c.config.Network = network
	}
}

// WithLogger inject logger
func WithLogger(logger *elog.Component) Option {
	return func(c *Container) {
		c.logger = logger
	}
}
