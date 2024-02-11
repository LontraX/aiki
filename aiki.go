package aiki

// Task represents a unit of work that needs to be processed.
type Task struct {
	ID      string // Unique identifier for the task.
	Payload []byte // Payload contains the data associated with the task.

}

// NewTask creates a new task with the given ID and payload.
func NewTask(id string, payload []byte) *Task {
	return &Task{
		ID:      id,
		Payload: payload,
	}
}
