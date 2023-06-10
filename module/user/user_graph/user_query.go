package user_graph

import (
	"context"
	"xyu-graphql/internal/dao/model"
)

type QueryResolver struct{}

func (r *QueryResolver) UserList(ctx context.Context) ([]*model.User, error) {
	bk := &model.User{
		ID:           0,
		JobID:        0,
		DepartmentID: 0,
		Seniority:    0,
		Name:         "XXX",
		Address:      "XXX",
		Birthday:     "XXX-XXX-XX XXX:XXX:XXX",
		Phone:        "12324354365",
		UUID:         "SVSV089V0AOIHBD0BS9BFD0B",
	}
	var bks []*model.User
	bks = append(bks, bk)
	return bks, nil
}

func (r *QueryResolver) GetUserByID(ctx context.Context, id model.UserID) ([]*model.User, error) {
	bk := &model.User{
		ID:           0,
		JobID:        0,
		DepartmentID: 0,
		Seniority:    0,
		Name:         "XXX",
		Address:      "XXX",
		Birthday:     "XXX-XXX-XX XXX:XXX:XXX",
		Phone:        "12324354365",
		UUID:         "SVSV089V0AOIHBD0BS9BFD0B",
	}
	var bks []*model.User
	bks = append(bks, bk)
	return bks, nil
}
