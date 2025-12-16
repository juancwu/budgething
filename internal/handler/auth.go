package handler

import (
	"net/http"

	"git.juancwu.dev/juancwu/budgit/internal/ui"
	"git.juancwu.dev/juancwu/budgit/internal/ui/pages"
)

type authHandler struct {
}

func NewAuthHandler() *authHandler {
	return &authHandler{}
}

func (h *authHandler) AuthPage(w http.ResponseWriter, r *http.Request) {
	ui.Render(w, r, pages.Auth(""))
}

func (h *authHandler) PasswordPage(w http.ResponseWriter, r *http.Request) {
	ui.Render(w, r, pages.AuthPassword(""))
}
