package grpcweb_reflection_v1alpha

import (
	"errors"

	"github.com/ilackarms/grpc-web-go-client/grpcweb"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	pb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

type serverReflectionClient struct {
	cc *grpcweb.Client
}

// NewServerReflectionClient instantiates a new server reflection client.
// most part of the implementation is same as the original grpc_reflection_v1alpha package's.
//
// the version (like v1alpha) is corrensponding to grpc_reflection_v1alpha package
func NewServerReflectionClient(cc *grpcweb.Client) pb.ServerReflectionClient {
	return &serverReflectionClient{cc}
}

func (c *serverReflectionClient) ServerReflectionInfo(ctx context.Context, opts ...grpc.CallOption) (pb.ServerReflection_ServerReflectionInfoClient, error) {
	if len(opts) != 0 {
		return nil, errors.New("currently, ilackarms/grpc-web-go-client does not support grpc.CallOption")
	}

	req := newRequest(nil)

	stream, err := c.cc.BidiStreaming(ctx, req)
	if err != nil {
		return nil, err
	}

	return &serverReflectionServerReflectionInfoClient{cc: stream}, nil
}

type serverReflectionServerReflectionInfoClient struct {
	cc grpcweb.BidiStreamClient

	// To satisfy pb.ServerReflection_ServerReflectionInfoClient
	grpc.ClientStream
}

func (x *serverReflectionServerReflectionInfoClient) Send(m *pb.ServerReflectionRequest) error {
	req := newRequest(m)
	return x.cc.Send(req)
}

func (x *serverReflectionServerReflectionInfoClient) Recv() (*pb.ServerReflectionResponse, error) {
	res, err := x.cc.Receive()
	if err != nil {
		return nil, err
	}
	return res.Content.(*pb.ServerReflectionResponse), nil
}

func (x *serverReflectionServerReflectionInfoClient) CloseSend() error {
	return x.cc.CloseSend()
}

func newRequest(in *pb.ServerReflectionRequest) *grpcweb.Request {
	out := &pb.ServerReflectionResponse{}
	return grpcweb.NewRequest("/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo", in, out)
}
