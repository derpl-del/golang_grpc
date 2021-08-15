package MovieController

import (
	"context"
	"fmt"
	"simple-api/apps/controller/LogsController"
	"simple-api/apps/helper/MovieHelper"
	"simple-api/apps/proto/MovieProto"
	"simple-api/config/env/MovieEnv"
	"time"
)

type MovieServer struct {
}

//Pipeline Get Movie
func (s *MovieServer) GetMovie(ctx context.Context, req *MovieProto.SeachRequest) (*MovieProto.SearchResponse, error) {
	t := time.Now()
	trxid := t.Format("20060102150405")
	LogsController.CreatePoint(trxid, "START", "GetMovie", req.String())
	res := MovieProto.SearchResponse{}
	res.Trxid = trxid
	if req.Pagination <= 0 {
		res.ErrorCode = "1"
		res.ErrorMessage = "pagination cannot lower than 0 or equal 0"
		LogsController.CreatePoint(trxid, "Finish", "GetMovie", res.String())
		return &res, nil
	} else if len(req.Searchword) == 0 {
		res.ErrorCode = "2"
		res.ErrorMessage = "searchword cannot be null"
		LogsController.CreatePoint(trxid, "Finish", "GetMovie", res.String())
		return &res, nil
	}

	var env MovieEnv.Parameter
	err := env.GetParameter()
	if err != nil {
		fmt.Println(err)
	}
	env.Trxid = trxid
	var rc = make(chan MovieHelper.Movies)

	go MovieHelper.SendMovieRequestAsync(env, int(req.Pagination), req.Searchword, rc)

	list := <-rc
	close(rc)
	if list.Response != "True" {
		res.ErrorCode = "3"
		res.ErrorMessage = list.Response
		LogsController.CreatePoint(trxid, "Error", "GetMovie", res.String())
		return &res, nil
	}
	var movies []*MovieProto.Movie
	for _, data := range list.Search {
		row := MovieProto.Movie{
			Title:  data.Title,
			Year:   data.Year,
			ImdbID: data.ImdbID,
			Type:   data.Type,
			Poster: data.Poster,
		}
		movies = append(movies, &row)
	}
	res.ErrorCode = "0"
	res.ErrorMessage = "success"
	res.List = movies
	LogsController.CreatePoint(env.Trxid, "Finish", "GetMovie", res.String())

	return &res, nil
}
