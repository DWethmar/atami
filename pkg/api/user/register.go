package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/usecase/userusecase"
)

// NewUser struct definition
type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterUser handler
func RegisterUser(usecase *userusecase.Usecase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			response.SendServerError(w, r)
			return
		}

		var newUser = NewUser{}
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			fmt.Printf("Error while decoding entry: %v", err)
			response.SendBadRequestError(w, r, "Invalid input")
			return
		}

		user, err := usecase.Register(newUser.Username, newUser.Email, newUser.Password)
		if err != nil {
			response.SendBadRequestError(w, r, err)
			return
		}

		response.SendJSON(w, r, user, http.StatusCreated)
	})
}
