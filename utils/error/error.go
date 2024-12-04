package error_utils

type CustomErr struct {
	HttpCode int
	Message  string
	Detail   string
	Data     interface{}
}

func (slf *CustomErr) Error() string {
	return slf.Message
}
