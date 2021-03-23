package future

type SuccessFunc func(string)
type FailFunc func(error)
type ExecuteStringFunc func() (string, error)

type MaybeString struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

func (ms *MaybeString) Success(f SuccessFunc) *MaybeString {
	ms.successFunc = f
	return ms
}

func (ms *MaybeString) Fail(f FailFunc) *MaybeString {
	ms.failFunc = f
	return ms
}

func (ms *MaybeString) Execute(f ExecuteStringFunc) {
	go func(ms *MaybeString) {
		str, err := f()
		if err != nil {
			ms.failFunc(err)
		} else {
			ms.successFunc(str)
		}
	}(ms)
}
