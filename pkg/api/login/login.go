package login

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/auth"
)

// NewUser struct definition
type Credentials struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login handles login requests
func Login(service auth.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// if err := r.ParseForm(); err != nil {
		// 	fmt.Fprintf(w, "ParseForm() err: %v", err)
		// 	return
		// }

		// fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)

		// username := r.FormValue("username")
		// password := r.FormValue("password")

	})
}
