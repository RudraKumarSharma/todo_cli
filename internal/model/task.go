package model

type Task struct {
	ID int `json:"int"`
	Title string `json:"title"`
	Done bool `json:"done"`
}