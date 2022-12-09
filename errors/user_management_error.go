package errors

import "fmt"

type UserManagementError struct {
	StatusCode int
	Err        error
}

func (r *UserManagementError) Error() string {
	return fmt.Sprintf("err %v", r.Err)
}

var (
	ErrorNotFound = fmt.Errorf("resource not found")
)
