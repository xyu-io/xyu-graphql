package user_graph

import (
	"context"
	"xyu-graphql/internal/dao/model"
)

type MutationResolver struct{}

func (r *MutationResolver) RegisterUser(ctx context.Context, user model.NewUser) (*model.User, error) {
	bk := &model.User{
		ID:           110,
		JobID:        user.JobID,
		DepartmentID: *user.DepartmentID,
		Seniority:    *user.Seniority,
		Name:         user.Name,
		Address:      user.Address,
		Birthday:     user.Birthday,
		Phone:        user.Phone,
	}
	return bk, nil
}
