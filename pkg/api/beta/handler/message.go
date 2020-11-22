package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dwethmar/atami/pkg/api/middleware"
	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/message"
)

// MessageUser output
type MessageUser struct {
	UID      string `json:"uid"`
	Username string `json:"username"`
}

// Message output
type Message struct {
	UID       string      `json:"uid"`
	Text      string      `json:"text"`
	User      MessageUser `json:"user"`
	CreatedAt time.Time   `json:"created_at"`
}

// CreatMessageInput input
type CreatMessageInput struct {
	Text string `json:"text"`
}

// CreatMessageSuccess input
type CreatMessageSuccess struct {
	UID string `json:"uid"`
}

// ListMessages handler
func ListMessages(ms *message.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		usr, err := middleware.UserFromContext(r.Context())
		if err != nil || usr == nil {
			fmt.Print(err)
			response.SendServerError(w, r)
			return
		}

		if result, err := ms.Find(); err == nil {
			response.SendJSON(w, r, result, http.StatusOK)
		} else {
			response.SendBadRequestError(w, r, err)
		}

		// fmt.Fprintf(w, "Hi there %v, I love %s!", usr.Username, r.URL.Path[1:])
	})
}

// CreateMessages handler
func CreateMessages(ms *message.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		usr, err := middleware.UserFromContext(r.Context())
		if err != nil || usr == nil {
			fmt.Print(err)
			response.SendServerError(w, r)
			return
		}

		var input CreatMessageInput
		if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
			fmt.Printf("Error while decoding entry: %v", err)
			response.SendServerError(w, r)
			return
		}

		newMsg := message.CreateMessage{
			Text:            input.Text,
			CreatedByUserID: usr.ID,
		}

		if err := ms.ValidateCreateMessage(newMsg); err == nil {
			if msg, err := ms.Create(newMsg); err == nil {
				response.SendJSON(w, r, CreatMessageSuccess{
					UID: msg.UID,
				}, http.StatusOK)
			} else {
				fmt.Print(err)
				response.SendServerError(w, r)
			}
		} else {
			response.SendBadRequestError(w, r, err)
		}

		// fmt.Fprintf(w, "CREATE NEW MSG %v!", usr.Username)
	})
}
