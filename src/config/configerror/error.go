package configerror

type configError struct {
	msg     string
	isError bool
}

func (e *configError) Msg() string {
	if e.isError {
		return e.Error()
	}
	return e.Warning()
}

func (e *configError) Error() string {
	return e.msg
}

func (e *configError) Warning() string {
	return e.msg
}

func (e *configError) IsError() bool {
	return e.isError
}

func (e *configError) IsWarning() bool {
	return !e.isError
}
