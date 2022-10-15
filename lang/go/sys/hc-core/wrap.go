package hc_core

import (
	"fmt"
	"net"
	"os"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative hc-core.proto

var SocketPath = "/run/lzc-sys/hc-core.sock"

func Serve(srv CoreSystemServer) error {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)

	RegisterCoreSystemServer(s, srv)

	err := os.RemoveAll(SocketPath)
	if err != nil {
		return err
	}
	lis, err := net.Listen("unix", SocketPath)
	if err != nil {
		return err
	}
	fmt.Printf("[GRPC-server] Listening and serving GRPC on %s://%s\n", "unix", lis.Addr().String())

	_ = os.Chmod(SocketPath, 0777)

	return s.Serve(lis)
}

type Client struct {
	CoreSystemClient
	conn *grpc.ClientConn
}

func (c *Client) Close() error { return c.conn.Close() }

func NewClient() (*Client, error) {
	conn, err := grpc.Dial("unix://"+SocketPath, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Client{CoreSystemClient: NewCoreSystemClient(conn), conn: conn}, nil
}
