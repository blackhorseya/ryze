package er

// Error declare custom error
type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"msg,omitempty"`
	Log     string `json:"-"`
}

func (e *Error) Error() string {
	return e.Message
}

// New a error
func New(code int, msg string, log string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
		Log:     log,
	}
}
