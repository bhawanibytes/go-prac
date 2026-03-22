package router

import (
	"go-prac/internal/handler"
	"net/http"
)

func UserRoutes() *http.ServeMux {
	userMux := http.NewServeMux()

	userMux.HandleFunc("/", handler.UserHome)
	userMux.HandleFunc("/profile", handler.UserProfile)
	userMux.HandleFunc("/profile/edit", handler.UserProfileEdit)
	userMux.HandleFunc("/profile/edit/password", handler.UserProfileEditPassword)
	userMux.HandleFunc("/profile/edit/password/confirm", handler.UserProfileEditPasswordConfirm)
	userMux.HandleFunc("/profile/edit/password/confirm/submit", handler.UserProfileEditPasswordConfirmSubmit)

	return userMux
}
