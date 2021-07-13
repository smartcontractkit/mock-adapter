package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ExternalAdapter struct {
	LocalAddr        string
	InsideDockerAddr string
}

type ExternalAdapterResponse struct {
	JobRunId string              `json:"id"`
	Data     ExternalAdapterData `json:"data"`
	Error    error               `json:"error"`
}

type ExternalAdapterData struct {
	Result int `json:"result"`
}

const adapterPort = ":6060"

type OkResult struct{}

var variableData int

// starts an external adapter on specified port
func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/random", randomNumber)
	router.GET("/five", five)
	router.GET("/variable", variable)
	router.POST("/set_variable", setVariable)

	log.Info().Str("Port", adapterPort).Msg("Starting external adapter")
	log.Fatal().AnErr("Error", http.ListenAndServe(adapterPort, router)).Msg("Error occured while running external adapter")
}

// index allows a status check on the adapter
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Adapter listening!")
	result := &ExternalAdapterResponse{
		JobRunId: "",
		Data:     ExternalAdapterData{Result: 0},
		Error:    nil,
	}
	log.Info().Str("Endpoint", "/").Msg("Index")
	_ = json.NewEncoder(w).Encode(result)
}

// RandomNumber returns a random int from 0 to 100
func randomNumber(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	num := rand.Intn(100)
	result := &ExternalAdapterResponse{
		JobRunId: "",
		Data:     ExternalAdapterData{Result: num},
		Error:    nil,
	}
	log.Info().Str("Endpoint", "/random").Int("Result", num).Msg("Random Number")
	_ = json.NewEncoder(w).Encode(result)
}

// five returns five
func five(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result := &ExternalAdapterResponse{
		JobRunId: "",
		Data:     ExternalAdapterData{Result: 5},
		Error:    nil,
	}
	log.Info().Str("Endpoint", "/five").Int("Result", 5).Msg("Five")
	_ = json.NewEncoder(w).Encode(result)
}

func setVariable(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	q := r.URL.Query()
	log.Info().Interface("query", q).Msg("params received")
	v := q.Get("var")
	data, _ := strconv.Atoi(v)
	variableData = data
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result := &OkResult{}
	log.Info().Str("Endpoint", "/set_variable").Int("Variable", variableData).Msg("Set Variable")
	_ = json.NewEncoder(w).Encode(result)
}

func variable(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	log.Info().Int("data", variableData).Msg("variable response")
	result := &ExternalAdapterResponse{
		JobRunId: "",
		Data:     ExternalAdapterData{Result: variableData},
		Error:    nil,
	}
	log.Info().Str("Endpoint", "/variable").Int("Variable", variableData).Msg("Get Variable")
	_ = json.NewEncoder(w).Encode(result)
}
