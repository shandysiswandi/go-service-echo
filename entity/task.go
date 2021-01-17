package entity

// Task is
type Task struct {
	UUID
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Timestamps
}

// TableName is
func (Task) TableName() string {
	return "tasks"
}
