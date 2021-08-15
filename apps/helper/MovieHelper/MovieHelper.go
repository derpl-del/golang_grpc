package MovieHelper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"simple-api/apps/controller/LogsController"
	"simple-api/config/env/MovieEnv"
	"strconv"
	"time"
)

//Movie Models for Response OMDB Api
type Movies struct {
	Search []struct {
		Title  string `json:"Title"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
		Type   string `json:"Type"`
		Poster string `json:"Poster"`
	} `json:"Search"`
	TotalResults string `json:"totalResults"`
	Response     string `json:"Response"`
}

//Function to http get request
func SendMovieRequestAsync(env MovieEnv.Parameter, pagination int, searchword string, rc chan Movies) {
	LogsController.CreatePoint(env.Trxid, "START", "SendRequest", env.URL)
	var response Movies
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("GET", env.URL, nil)
	if err != nil {
		LogsController.CreatePoint(env.Trxid, "Error", "SendRequest", fmt.Sprintf("%v", err))
		response.Response = "false"
		rc <- response
	}

	q := req.URL.Query()
	q.Add("apikey", env.Apikey)
	q.Add("s", searchword)
	q.Add("page", strconv.Itoa(pagination))
	req.URL.RawQuery = q.Encode()
	LogsController.CreatePoint(env.Trxid, "Out", "SendRequest", req.URL.String())
	resp, err2 := client.Get(req.URL.String())
	if err2 != nil {
		LogsController.CreatePoint(env.Trxid, "Error", "SendRequest", fmt.Sprintf("%v", err2))
		response.Response = fmt.Sprintf("%v", err2)
		fmt.Println(response)
		rc <- response
		return
	} else if resp.StatusCode < 200 || resp.StatusCode >= 299 {
		errorlog := fmt.Sprintf("Got status_code : %v", resp.StatusCode)
		LogsController.CreatePoint(env.Trxid, "Error", "SendRequest", errorlog)
		response.Response = errorlog
		fmt.Println(response)
		rc <- response
		return
	}
	byteValue, _ := ioutil.ReadAll(resp.Body)
	LogsController.CreatePoint(env.Trxid, "In", "SendRequest", string(byteValue))
	json.Unmarshal(byteValue, &response)
	LogsController.CreatePoint(env.Trxid, "Finish", "SendRequest", fmt.Sprintf("%v", response))
	rc <- response

}
