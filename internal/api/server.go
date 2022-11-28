package api

import (
	"google.golang.org/grpc"
	"net"
)

type Deps struct {
	FileHandler FileServiceServer
}

type Server struct {
	deps Deps
	srv  *grpc.Server
}

func NewServer(deps Deps) *Server {
	return &Server{
		deps: deps,
		srv:  grpc.NewServer(),
	}
}

func (s *Server) ListenAndServe(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	RegisterFileServiceServer(s.srv, s.deps.FileHandler)

	if err := s.srv.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.srv.GracefulStop()
}
