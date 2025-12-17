package handler

import "net/http"

type dashboardHandler struct{}

func NewDashboardHandler() *dashboardHandler {
	return &dashboardHandler{}
}

func (h *dashboardHandler) DashboardPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Dashboard page"))
}
