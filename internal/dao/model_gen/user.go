package model_gen

type User struct {
	ID           int64  `json:"id,omitempty"`
	JobID        int64  `json:"job_id"`
	DepartmentID int64  `json:"department_id"`
	Seniority    int    `json:"seniority"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Birthday     string `json:"birthday"`
	Phone        string `json:"phone"`
	UUID         string `json:"uuid"`
}

type NewUser struct {
	JobID        int64  `json:"job_id"`
	DepartmentID int64  `json:"department_id"`
	Seniority    int    `json:"seniority"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Birthday     string `json:"birthday"`
	Phone        string `json:"phone"`
}

type UserID struct {
	UserID int64 `json:"user_id"`
}
