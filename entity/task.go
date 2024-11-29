package entity

type Task struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DateTime    string `json:"dateTime"`
}

var TaskChan chan Task = make(chan Task, 1000)
