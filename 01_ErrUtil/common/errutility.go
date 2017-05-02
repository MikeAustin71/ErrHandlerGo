package common

// SpecError - A data structure used
// to hold custom error information
type SpecError struct {
	PrefixMsg string
	SuffixMsg string
	ErrMsg    string
}

func (s *SpecError) NewError(prefix string, suffix string, err error) {

	s.PrefixMsg = prefix
	s.SuffixMsg = suffix
	if err != nil {
		s.ErrMsg = err.Error()
	} else {
		s.ErrMsg = ""
	}

}

func (s *SpecError) Error() string {
	return s.PrefixMsg + s.ErrMsg + s.SuffixMsg
}

func CheckErrPanic(e error) {
	if e != nil {
		panic(e)
	}
}

func SpecCheckErrPanic(prefix string, suffix string, err error) {
	if err == nil {
		return
	}

	e := SpecError{PrefixMsg: prefix, SuffixMsg: suffix, ErrMsg: err.Error()}

	panic(e)
}
