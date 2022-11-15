package errors

type PermissionError struct {
	Message string
}

// TODO: replace with a standard error format
func (i PermissionError) Error() string {
	panic(i.Message)
}
