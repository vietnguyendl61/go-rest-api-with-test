package handlers

import (
	"net/http"
)

type JobHandler struct {
}

func NewJobHandler() JobHandler {
	return JobHandler{}
}

func (h JobHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h JobHandler) GetOne(w http.ResponseWriter, r *http.Request) {

}
