package model

type Author struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Book struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Author *Author `json:"author"`
	UUID   string  `json:"uuid"`
}

type BookID struct {
	ID int64 `json:"id"`
}

type NewBook struct {
	Name   string      `json:"name"`
	Author *AuthorType `json:"author"`
}

type AuthorType struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
