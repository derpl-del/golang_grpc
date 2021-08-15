package main

import (
	"context"
	"fmt"
	"simple-api/apps/controller/MovieController"
	"simple-api/apps/helper/MovieHelper"
	"simple-api/apps/proto/MovieProto"
	"simple-api/config/env/MovieEnv"
	"testing"
)

var (
	parameter = MovieEnv.Parameter{}
	rc        = make(chan MovieHelper.Movies)
	server    = MovieController.MovieServer{}
	ctx       = context.Background()
	List      = TestData{
		Request: []MovieProto.SeachRequest{
			{Pagination: 1, Searchword: "bat"},
			{Pagination: -1, Searchword: "bat"},
			{Pagination: 1, Searchword: ""},
		},
		Expectation: []string{"success", "pagination cannot lower than 0 or equal 0", "searchword cannot be null"},
	}
)

type TestData struct {
	Request     []MovieProto.SeachRequest
	Expectation []string
}

func TestGetEnvParameter(t *testing.T) {
	err := parameter.GetParameter()
	if err != nil {
		t.Fatal(err)
	}
	if len(parameter.URL) < 1 {
		t.Fatalf("Empty URL yml")
	}
	if len(parameter.Apikey) < 1 {
		t.Fatalf("Empty Apikey yml")
	}
}

func TestSendPostAsync(t *testing.T) {
	parameter.GetParameter()
	parameter.Trxid = "test"
	go MovieHelper.SendMovieRequestAsync(parameter, int(List.Request[0].Pagination), List.Request[0].Searchword, rc)

	list := <-rc
	close(rc)
	if list.Response != "True" {
		t.Fatalf(list.Response)
	}
}

func TestGetMovieController(t *testing.T) {
	for i := range List.Request {
		res, err := server.GetMovie(ctx, &List.Request[i])
		if err != nil {
			fmt.Println("1")
			t.Fatal(err)
		} else if res.ErrorCode != "0" && res.ErrorMessage != List.Expectation[i] {
			fmt.Println("2")
			t.Fatalf(res.ErrorMessage)
		}
	}
}
