package models

// IsNotFoundError returns whether an error represents a "not found" error.
func IsNotFoundError(err error) bool {
	switch err.(type) {
	case UserNotFoundError:
		return true

	}
	return false

}

// UserNotFoundError represents when a user is not found.
type UserNotFoundError struct{}

func (e UserNotFoundError) Error() string {
	return "User not found"
}
