package models

type ToDo struct {
	ID        uint
	Task      string
	Completed bool
	UserID    uint
}

func (t *ToDo) Sanitize() map[string]interface{} {
	return map[string]interface{}{
		"id":        t.ID,
		"task":      t.Task,
		"completed": t.Completed,
	}
}
