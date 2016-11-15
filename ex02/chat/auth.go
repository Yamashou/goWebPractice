package main

import (
	"net/http"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		//未認証
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		panic(err.Error())
	} else {
		//成功。ラップされた、ハンドラを呼び出す
		h.next.ServeHTTP(w, r)
	}
}

//MustAuth ...
func MustAuth(hander http.Handler) http.Handler {
	return &authHandler{next: hander}
}
