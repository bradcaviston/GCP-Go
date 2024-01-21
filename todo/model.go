package todo

type Todo struct {
	Id      string `json:"id"`
	Text    string `json:"text" validate:"required"`
	Checked bool   `json:"checked" validate:"required"`
}