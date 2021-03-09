package drive

type errorWithRetryContext struct {
	ServiceAccountFile string
	err                error
}

func (e *errorWithRetryContext) Error() string {
	return e.err.Error()
}

func newErrorWithRetryContext(err error, serviceAccountFile string) *errorWithRetryContext {
	return &errorWithRetryContext{
		ServiceAccountFile: serviceAccountFile,
		err:                err,
	}
}
