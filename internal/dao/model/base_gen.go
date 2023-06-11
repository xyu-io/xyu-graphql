// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Author struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Phone   *string `json:"phone,omitempty"`
	Address string  `json:"address"`
}

type AuthorType struct {
	Name    string  `json:"name"`
	Phone   *string `json:"phone,omitempty"`
	Address string  `json:"address"`
}

type Book struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Author *Author `json:"author"`
	UUID   string  `json:"uuid"`
	Data   *string `json:"data,omitempty"`
}

type BookID struct {
	ID int64 `json:"id"`
}

type NewBook struct {
	Name   string      `json:"name"`
	Author *AuthorType `json:"author,omitempty"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewUser struct {
	Name         string `json:"name"`
	JobID        int64  `json:"job_id"`
	DepartmentID *int64 `json:"department_id,omitempty"`
	Seniority    *int64 `json:"seniority,omitempty"`
	Address      string `json:"address"`
	Birthday     string `json:"birthday"`
	Phone        string `json:"phone"`
}

type Test struct {
	Title string `json:"title"`
}

type Todo struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type TodoID struct {
	ID *int64 `json:"id,omitempty"`
}

type User struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	JobID        int64  `json:"job_id"`
	DepartmentID int64  `json:"department_id"`
	Seniority    int64  `json:"seniority"`
	Address      string `json:"address"`
	Birthday     string `json:"birthday"`
	Phone        string `json:"phone"`
	UUID         string `json:"uuid"`
}

type UserID struct {
	UserID int64 `json:"user_id"`
}