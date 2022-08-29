package rest

import (
	"context"
	"github.com/bibishkind/shtrafovnet-test-task/pb"
	"github.com/bibishkind/shtrafovnet-test-task/pkg/transport/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type Server struct {
	pb.SearchServer
	serveMux   *runtime.ServeMux
	httpServer *http.Server
	grpcServer *grpc.Server
}

func NewServer(grpcServer *grpc.Server) *Server {
	mux := runtime.NewServeMux()
	httpServer := &http.Server{Handler: mux}
	return &Server{
		serveMux:   mux,
		httpServer: httpServer,
		grpcServer: grpcServer,
	}
}

func (srv *Server) ListenAndServe(addr string) error {
	srv.httpServer.Addr = addr

	err := pb.RegisterSearchHandlerServer(context.Background(), srv.serveMux, srv.grpcServer)
	if err != nil {
		return err
	}

	return srv.httpServer.ListenAndServe()
}

func (srv *Server) Stop(ctx context.Context) error {
	return srv.httpServer.Shutdown(ctx)
}
