package server

import (
	"net/http"
	"sort"
	"time"

	"github.com/lenvendo/ig-absolut-api/configs"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// NewServer инициализирует сервер.
func NewServer(ops ...Option) (svc *Server, err error) {
	svc = new(Server)

	for _, o := range ops {
		o(svc)
	}

	return svc, nil
}

func SetLogger(logger log.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

func SetConfig(cfg *configs.Config) Option {
	return func(s *Server) {
		s.cfg = cfg
	}
}

func SetHandler(handlers map[string]http.Handler) Option {

	sts := make([]string, 0)
	for s := range handlers {
		sts = append(sts, s)
	}
	sort.Strings(sts)

	return func(s *Server) {
		mux := mux.NewRouter().StrictSlash(false)

		for i := len(sts) - 1; i >= 0; i-- {
			name := sts[i]

			if handler, ok := handlers[name]; ok {
				mux.PathPrefix("/" + name).Handler(handler)
			}
		}

		s.handler = mux
	}
}

func SetGRPC(joins ...func(grpc *grpc.Server)) Option {
	return func(s *Server) {
		grpcServer := grpc.NewServer(
			grpc.UnaryInterceptor(grpctransport.Interceptor),
			grpc.ConnectionTimeout(time.Second*time.Duration(s.cfg.Server.GRPC.TimeoutSec)),
			grpc.KeepaliveParams(keepalive.ServerParameters{
				MaxConnectionIdle: 5 * time.Minute,
			}),
		)
		for _, j := range joins {
			j(grpcServer)
		}
		s.grpc = grpcServer
	}
}
