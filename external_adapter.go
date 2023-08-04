package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

const jsonHeader = "application/json; charset=UTF-8"

type OkResult struct{}

var variableData int
var jsonVariableData []byte

// starts an external adapter on specified port
func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	adapterAddress := ":6060"
	if len(os.Args) > 1 {
		adapterAddress = os.Args[1]
	}

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/random", randomNumber)
	router.GET("/five", five)
	router.GET("/variable", variable)
	router.GET("/json_variable", jsonVariable)
	router.POST("/", index)
	router.POST("/random", randomNumber)
	router.POST("/five", five)
	router.POST("/variable", variable)
	router.POST("/set_variable", setVariable)
	router.POST("/json_variable", jsonVariable)
	router.POST("/set_json_variable", setJsonVariable)

	log.Info().Str("Address", adapterAddress).Msg("Starting external adapter")
	log.Fatal().AnErr("Error", http.ListenAndServe(adapterAddress, router)).Msg("Error occured while running external adapter")
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
	w.Header().Set("Content-Type", jsonHeader)
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
	w.Header().Set("Content-Type", jsonHeader)
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
	w.Header().Set("Content-Type", jsonHeader)
	result := &OkResult{}
	log.Info().Str("Endpoint", "/set_variable").Int("Variable", variableData).Msg("Set Variable")
	_ = json.NewEncoder(w).Encode(result)
}

func variable(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", jsonHeader)
	log.Info().Int("data", variableData).Msg("variable response")
	result := &ExternalAdapterResponse{
		JobRunId: "",
		Data:     ExternalAdapterData{Result: variableData},
		Error:    nil,
	}
	log.Info().Str("Endpoint", "/variable").Int("Variable", variableData).Msg("Get Variable")
	_ = json.NewEncoder(w).Encode(result)
}

func setJsonVariable(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	header := r.Header.Get("Content-Type")
	if header != "application/json" {
		http.Error(w, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not read the body: %v", err), http.StatusBadRequest)
		return
	}
	var dataMap map[string]interface{}
	var dataArray []interface{}
	if err := json.Unmarshal(body, &dataMap); err != nil {
		// could not parse into map, maybe it is an array
		if err := json.Unmarshal(body, &dataArray); err != nil {
			http.Error(w, fmt.Sprintf("Could not parse the json into an array or map: %v", err), http.StatusBadRequest)
			return
		}
	}

	jsonVariableData = body

	log.Info().Msg("New json body received")
	w.Header().Set("Content-Type", jsonHeader)
	log.Info().Str("Endpoint", "/set_json_variable").Msg("Set Variable")
	fmt.Fprint(w, string(jsonVariableData))
}

func jsonVariable(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", jsonHeader)
	log.Info().Str("Endpoint", "/json_variable").Msg("Get Json Variable")
	fmt.Fprint(w, string(jsonVariableData))
}
