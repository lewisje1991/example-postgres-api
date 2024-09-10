package diary

type Response struct {
	ID    string `json:"id"`
	Date  string `json:"day"`
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
