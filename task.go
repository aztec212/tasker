package tasker

// Task interface
type Task interface {
	Execute() error
	HandleError(e error)
}
