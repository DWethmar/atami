package beta

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dwethmar/atami/pkg/api/middleware"
	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
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

func mapMessage(msg *message.Message) *Message {
	var user MessageUser
	if msg.User != nil {
		user = MessageUser{
			UID:      msg.User.UID,
			Username: msg.User.Username,
		}
	}

	return &Message{
		UID:       msg.UID,
		Text:      msg.Text,
		CreatedAt: msg.CreatedAt,
		User:      user,
	}
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
		usr, err := middleware.GetUser(r.Context())
		if err != nil || usr == nil {
			fmt.Print(err)
			response.ServerError(w, r)
			return
		}

		messages := make([]*Message, 0)

		if result, err := ms.Find(0, 100); err == nil {
			for _, msg := range result {
				messages = append(messages, mapMessage(msg))
			}
		} else {
			response.BadRequestError(w, r, err)
			return
		}

		response.JSON(w, r, messages, http.StatusOK)
	})
}

// GetMessages handler
func GetMessages(ms *message.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usr, err := middleware.GetUser(r.Context())
		if err != nil || usr == nil {
			fmt.Print(err)
			response.ServerError(w, r)
			return
		}

		uid, err := middleware.UIDFromContext(r.Context())
		if err != nil {
			fmt.Print(err)
			response.ServerError(w, r)
			return
		}

		if msg, err := ms.FindByUID(uid); err == nil {
			if msg == nil {
				response.NotFoundError(w, r)
				return
			}
			response.JSON(w, r, mapMessage(msg), http.StatusOK)
		} else {
			fmt.Print(err)
			response.ServerError(w, r)
			return
		}
	})
}

// CreateMessage handler
func CreateMessage(ms *message.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usr, err := middleware.GetUser(r.Context())
		if err != nil || usr == nil {
			response.UnauthorizedError(w, r, errors.New("unauthorized"))
			return
		}

		var input CreatMessageInput
		if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
			fmt.Printf("Error while decoding entry: %v", err)
			response.ServerError(w, r)
			return
		}

		newMsg := message.CreateMessageRequest{
			Text:            input.Text,
			CreatedByUserID: usr.ID,
		}

		if err := ms.ValidateCreateMessage(newMsg); err == nil {
			if msg, err := ms.Create(newMsg); err == nil {
				response.JSON(w, r, CreatMessageSuccess{
					UID: msg.UID,
				}, http.StatusCreated)
				return
			}
			response.BadRequestError(w, r, err)
			return
		}

		fmt.Print(err)
		response.ServerError(w, r)
	})
}

// DeleteMessage handler
func DeleteMessage(ms *message.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usr, err := middleware.GetUser(r.Context())
		if err != nil || usr == nil {
			fmt.Print(err)
			response.ServerError(w, r)
			return
		}

		uid, err := middleware.UIDFromContext(r.Context())
		if err != nil {
			fmt.Print(err)
			response.ServerError(w, r)
			return
		}

		if msg, err := ms.FindByUID(uid); err == nil {
			if msg == nil {
				response.NotFoundError(w, r)
				return
			}

			if msg.CreatedByUserID == usr.ID {
				if err := ms.Delete(msg.ID); err == nil {
					response.JSON(w, r, mapMessage(msg), http.StatusOK)
				} else {
					response.ServerError(w, r)
				}
			} else {
				// Not authorized
				response.UnauthorizedError(w, r, errors.New("unauthorized"))
			}
		} else {
			fmt.Print(err)
			response.ServerError(w, r)
			return
		}
	})
}

// NewMessageRouter creates new message router
func NewMessageRouter(userService *user.Service, messageService *message.Service) http.Handler {
	r := chi.NewRouter()

	logger := httplog.NewLogger("message", httplog.Options{})
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Authenticated(userService))
	r.Use(cors.Handler(cors.Options{
		AllowOriginFunc: func(r *http.Request, origin string) bool {
			return true
		},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/", ListMessages(messageService))
	r.Get("/{uid}", middleware.RequireUID(GetMessages(messageService)))
	r.Post("/", CreateMessage(messageService))
	r.Delete("/{uid}", middleware.RequireUID(DeleteMessage(messageService)))

	return r
}
