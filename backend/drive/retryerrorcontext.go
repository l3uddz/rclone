package drive

type ErrorWithRetryContext struct {
	ServiceAccountFile string
	err                error
}

func (e *ErrorWithRetryContext) Error() string {
	return e.err.Error()
}

func NewErrorWithRetryContext(err error, serviceAccountFile string) *ErrorWithRetryContext {
	return &ErrorWithRetryContext{
		ServiceAccountFile: serviceAccountFile,
		err:                err,
	}
}
