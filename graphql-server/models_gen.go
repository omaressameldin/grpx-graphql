// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql_server

type DeleteTodo struct {
	TodoID int `json:"todoId"`
}

type NewTodo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      int    `json:"userId"`
}

type ReadTodo struct {
	TodoID int `json:"todoId"`
}

type UpdateTodo struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	TodoID      int     `json:"todoId"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
