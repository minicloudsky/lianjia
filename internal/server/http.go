package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/http/pprof"
	v1 "github.com/minicloudsky/lianjia/api/lianjia/v1"
	"github.com/minicloudsky/lianjia/internal/conf"
	"github.com/minicloudsky/lianjia/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, lianjia *service.Service, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.Handle("/debug/pprof/", pprof.NewHandler())
	srv.Handle("/debug/pprof/cmdline", pprof.NewHandler())
	srv.Handle("/debug/pprof/profile", pprof.NewHandler())
	srv.Handle("/debug/pprof/symbol", pprof.NewHandler())
	srv.Handle("/debug/pprof/trace", pprof.NewHandler())
	v1.RegisterLianjiaHTTPServer(srv, lianjia)
	return srv
}
