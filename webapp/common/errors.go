package common

type UserExistError struct {
	Code int
}
func (u UserExistError) Error() string {
	return "User Has Already Registed before!"
}


type ServerSideError struct {
	Code int
}
func (u ServerSideError) Error() string {
	return "Server Fail"
}


type AuthenticationError struct {
	Code int
}
func (u AuthenticationError) Error() string {
	return "User Authentication Error!"
}


type UserNotExistError struct {
	Code int
}
func (u UserNotExistError) Error() string {
	return "User Not Exist!"
}