package middleware

import (
	"net/http"

	"git.juancwu.dev/juancwu/budgit/internal/ctxkeys"
	"git.juancwu.dev/juancwu/budgit/internal/service"
)

// TODO: implement clearing jwt token in auth service

// AuthMiddleware checks for JWT token and adds user + profile + subscription to context if valid
func AuthMiddleware(authService *service.AuthService, userService *service.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: get auth cookie and verify value
			// TODO: fetch user information from database if cookie value is valid
			// TODO: add user to context if valid
			next.ServeHTTP(w, r)
		})
	}
}

// RequireGuest ensures request is not authenticated
func RequireGuest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := ctxkeys.User(r.Context())
		if user != nil {
			if r.Header.Get("HX-Request") == "true" {
				w.Header().Set("HX-Redirect", "/app/dashboard")
				w.WriteHeader(http.StatusSeeOther)
				return
			}
			http.Redirect(w, r, "/app/dashboard", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	}
}

// RequireAuth ensures the user is authenticated and has completed onboarding
func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := ctxkeys.User(r.Context())
		if user == nil {
			// For HTMX requests, use HX-Redirect header to force full page redirect
			if r.Header.Get("HX-Request") == "true" {
				w.Header().Set("HX-Redirect", "/auth")
				w.WriteHeader(http.StatusSeeOther)
				return
			}
			// For regular requests, use standard redirect
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}

		// Check if user has completed onboarding
		// Uses profile.Name as indicator (empty = incomplete onboarding)
		profile := ctxkeys.Profile(r.Context())
		if profile.Name == "" && r.URL.Path != "/auth/onboarding" {
			// User hasn't completed onboarding, redirect to onboarding
			if r.Header.Get("HX-Request") == "true" {
				w.Header().Set("HX-Redirect", "/auth/onboarding")
				w.WriteHeader(http.StatusSeeOther)
				return
			}
			http.Redirect(w, r, "/auth/onboarding", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	}
}
