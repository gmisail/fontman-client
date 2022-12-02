package errors

type InstallationError struct {
	Message string
}

// TODO: replace with a standard error format
func (i InstallationError) Error() string {
	panic(i.Message)
}
