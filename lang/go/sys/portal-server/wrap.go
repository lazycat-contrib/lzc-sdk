package portal_server

import (
	"net"
	"os"
	"path"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	grpc "google.golang.org/grpc"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative portal-server.proto

var (
	SocketPath = "/run/lzc-sys/portal-server.socket"
	LzcAppSocketPath = "/lzcapp/run/sys/portal-server.socket"
)

func Serve(srv HPortalSysServer) error {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)

	RegisterHPortalSysServer(s, srv)

	err := os.MkdirAll(path.Dir(SocketPath), 0755)
	if err != nil {
		return err
	}
	os.Remove(SocketPath)

	lis, err := net.Listen("unix", SocketPath)
	if err != nil {
		return err
	}
	os.Chmod(SocketPath, 0777)

	return s.Serve(lis)
}

type Client struct {
	HPortalSysClient
	conn *grpc.ClientConn
}

func (c Client) Close() error { return c.conn.Close() }

func NewClient() (*Client, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	}
	conn, err := grpc.Dial("unix://"+LzcAppSocketPath, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{HPortalSysClient: NewHPortalSysClient(conn), conn: conn}, nil
}
