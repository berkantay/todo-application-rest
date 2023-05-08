package entity

type Todo struct {
	Id          int64  `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Deadline    string `json:"deadline,omitempty"`
	Priority    int64  `json:"priority,omitempty"` //0-3 scale
}

type CreateTodoRequest struct {
	Description string `json:"description,omitempty"`
	Deadline    string `json:"deadline,omitempty"`
	Priority    int64  `json:"priority,omitempty"` //0-3 scale
}
