package models

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	State       string `json:"state"`
}

type GeneralResponse struct {
	Success      bool   `json:"success"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type GetTaskResponse struct {
	Task Task `json:"task"`
	GeneralResponse
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
