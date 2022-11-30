package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"go-study.com/m/core"
)

type HttpServer struct {
	logger     *core.Logger
	httpServer *http.ServeMux
}

type HttpHandler func(w http.ResponseWriter, r *http.Request)

func (s HttpServer) Init(port int) error {
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), s.httpServer)
	if err != nil {
		return err
	}
	return nil
}

func (s HttpServer) HandleGet(path string, callback HttpHandler) {
	s.httpServer.HandleFunc(path, callback)
}

func (s HttpServer) ParseProductFromReader(reader io.Reader) (*core.Product, error) {
	reqBody, err := ioutil.ReadAll(reader)
	if err != nil {
		s.logger.Error(fmt.Sprintf("server: could not read request body: %s\n", err))
	}
	product := &core.Product{}
	err = json.Unmarshal(reqBody, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func NewHttpServer(logger *core.Logger) *HttpServer {
	return &HttpServer{
		logger:     logger,
		httpServer: http.NewServeMux(),
	}
}
