package wrapper

import (
	"encoding/json"
	"net/http"
	"time"
)

type Wrapper struct {
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	ServerTime int64       `json:"serverTime"`
}

func (w *Wrapper) New(code, msg string, data interface{}) *Wrapper {
	w.Code = code
	w.Message = msg
	w.Data = data
	w.ServerTime = time.Now().UTC().UnixNano() / 1000000

	return w
}

func (w *Wrapper) MarshalJson() ([]byte, error) {
	return json.Marshal(*w)
}

type HttpStatus int

const (
	StatusOK                  = HttpStatus(http.StatusOK)
	StatusCreated             = HttpStatus(http.StatusCreated)
	StatusAccepted            = HttpStatus(http.StatusAccepted)
	StatusNoContent           = HttpStatus(http.StatusNoContent)
	StatusBadRequest          = HttpStatus(http.StatusBadRequest)
	StatusUnauthorized        = HttpStatus(http.StatusUnauthorized)
	StatusForbidden           = HttpStatus(http.StatusForbidden)
	StatusNotFound            = HttpStatus(http.StatusNotFound)
	StatusConflict            = HttpStatus(http.StatusConflict)
	StatusInternalServerError = HttpStatus(http.StatusInternalServerError)
	StatusNotImplemented      = HttpStatus(http.StatusNotImplemented)
	StatusServiceUnavailable  = HttpStatus(http.StatusServiceUnavailable)
)

func (h HttpStatus) New(msg string, data interface{}) (int, *Wrapper) {
	w := &Wrapper{}
	return int(h), w.New(http.StatusText(int(h)), msg, data)
}
