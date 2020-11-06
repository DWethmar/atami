package registration

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/auth"
)

// NewUser struct definition
type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register handler handles the request to create new user
func Register(service auth.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// if err := r.ParseForm(); err != nil {
		// 	fmt.Fprintf(w, "ParseForm() err: %v", err)
		// 	response.SendServerError(w, r)
		// 	return
		// }

		var newUser = NewUser{}
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			fmt.Printf("Error while decoding entry: %v\n", err)
			response.SendBadRequestError(w, r, errors.New("Invalid input"))
			return
		}

		user, err := service.Register(auth.RegisterUser{
			Username:      newUser.Username,
			Email:         newUser.Email,
			PlainPassword: newUser.Password,
		})

		if err != nil || user == nil {
			fmt.Printf("Error while registering user: %v\n", err)
			response.SendBadRequestError(w, r, err)
			return
		}

		response.SendJSON(w, r, toRespond(user), http.StatusCreated)
	})
}
