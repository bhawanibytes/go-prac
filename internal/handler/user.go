package handler

import "net/http"

func UserHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User"))
}

func UserProfile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Profile"))
}

func UserProfileEdit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Profile Edit"))
}

func UserProfileEditPassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Profile Edit Password"))
}

func UserProfileEditPasswordConfirm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Profile Edit Password Confirm"))
}

func UserProfileEditPasswordConfirmSubmit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Profile Edit Password Confirm Submit"))
}
