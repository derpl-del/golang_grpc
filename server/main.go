package main

import (
	"net"
	"simple-api/apps/controller/MovieController"
	"simple-api/apps/models/LogModels"
	"simple-api/apps/proto/MovieProto"
	"simple-api/config/dbadapter"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var Adapter = dbadapter.Adapter{}.New()
	Adapter.Table.AutoMigrate(&LogModels.Logs{})
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	movie := MovieController.MovieServer{}
	srv := grpc.NewServer()
	MovieProto.RegisterGetMovieInfoServer(srv, &movie)
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
