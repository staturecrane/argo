package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	argocd "github.com/argoproj/argo-cd"
	"github.com/argoproj/argo-cd/server/cluster"
	"github.com/argoproj/argo-cd/server/version"
	"github.com/cockroachdb/cockroach/pkg/util/protoutil"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

const (
	port = 8080
)

var (
	endpoint = fmt.Sprintf("localhost:%d", port)
)

// ArgoCDServer is the API server for ArgoCD
type ArgoCDServer struct {
}

// NewServer returns a new instance of the ArgoCD API server
func NewServer() *ArgoCDServer {
	return &ArgoCDServer{}
}

// Run runs the API Server
// We use k8s.io/code-generator/cmd/go-to-protobuf to generate the .proto files from the API types.
// k8s.io/ go-to-protobuf uses protoc-gen-gogo, which comes from gogo/protobuf (a fork of
// golang/protobuf).
func (a *ArgoCDServer) Run() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// Cmux is used to support servicing gRPC and HTTP1.1+JSON on the same port
	m := cmux.New(conn)
	grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	// gRPC Server
	grpcS := grpc.NewServer()
	version.RegisterVersionServiceServer(grpcS, &version.Server{})
	cluster.RegisterClusterServiceServer(grpcS, &cluster.Server{})

	// HTTP 1.1+JSON Server
	// grpc-ecosystem/grpc-gateway is used to proxy HTTP requests to the corresponding gRPC call
	mux := http.NewServeMux()
	// The following option is used with grpc-gateway to work around the following issue:
	// https://github.com/gogo/protobuf/issues/212#issuecomment-253709966
	gwMuxOpts := runtime.WithMarshalerOption(runtime.MIMEWildcard, new(protoutil.JSONPb))
	gwmux := runtime.NewServeMux(gwMuxOpts)
	mux.Handle("/", gwmux)
	dOpts := []grpc.DialOption{grpc.WithInsecure()}
	mustRegisterGWHandler(version.RegisterVersionServiceHandlerFromEndpoint, ctx, gwmux, endpoint, dOpts)
	mustRegisterGWHandler(cluster.RegisterClusterServiceHandlerFromEndpoint, ctx, gwmux, endpoint, dOpts)
	httpS := &http.Server{
		Addr:    endpoint,
		Handler: mux,
	}

	log.Infof("argocd %s serving on port %d", argocd.GetVersion(), port)

	// Use the muxed listeners for your servers.
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	err = m.Serve()
	if err != nil {
		panic(err)
	}
}

type registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

// mustRegisterGWHandler is a convenience function to register a gateway handler
func mustRegisterGWHandler(register registerFunc, ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	err := register(ctx, mux, endpoint, opts)
	if err != nil {
		panic(err)
	}
}