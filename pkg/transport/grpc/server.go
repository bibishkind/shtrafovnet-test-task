package grpc

import (
	"net"

	"github.com/bibishkind/shtrafovnet-test-task/pb"
	"github.com/bibishkind/shtrafovnet-test-task/pkg/service"
	"google.golang.org/grpc"
)

type Server struct {
	pb.SearchServer
	grpcServer *grpc.Server
	service    service.Service
}

func NewServer(srvc service.Service) *Server {
	return &Server{
		grpcServer: grpc.NewServer(),
		service:    srvc,
	}
}

func (srv *Server) ListenAndServe(port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	pb.RegisterSearchServer(srv.grpcServer, srv)

	return srv.grpcServer.Serve(lis)
}

func (srv *Server) Stop() {
	srv.grpcServer.GracefulStop()
}
