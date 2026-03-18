package routers

import "net/http"

func UserRoutes() *http.ServeMux {
	UserRouter := http.NewServeMux()

	UserRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User"))
	})
	UserRouter.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile"))
	})
	UserRouter.HandleFunc("/profile/edit", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile Edit"))
	})
	UserRouter.HandleFunc("/profile/edit/password", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile Edit Password"))
	})
	UserRouter.HandleFunc("/profile/edit/password/confirm", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile Edit Password Confirm"))
	})
	UserRouter.HandleFunc("/profile/edit/password/confirm/submit", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile Edit Password Confirm Submit"))
	})

	return UserRouter
}
