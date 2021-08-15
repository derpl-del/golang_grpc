package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"simple-api/apps/models/LogModels"
	"simple-api/apps/proto/MovieProto"
	"simple-api/config/dbadapter"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var (
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:4040", "gRPC server endpoint")
)

func run() error {
	var Adapter = dbadapter.Adapter{}.New()
	Adapter.Table.AutoMigrate(&LogModels.Logs{})
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := MovieProto.RegisterGetMovieInfoHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	log.Println("proxy is running...")
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
