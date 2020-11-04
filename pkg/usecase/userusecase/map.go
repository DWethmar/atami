package userusecase

import (
	"github.com/dwethmar/atami/pkg/model"
	"github.com/dwethmar/atami/pkg/user"
)

func toUser(user *user.User) *model.User {
	return &model.User{
		ID:        int64(user.ID),
		UID:       user.UID.String(),
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
}

func toUsers(users []*user.User) []*model.User {
	result := make([]*model.User, len(users))
	for i, user := range users {
		result[i] = toUser(user)
	}
	return result
}
