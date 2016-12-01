package uerrors

type uerror struct {
	s  string
	us string
}

func New(text, usertext string) *uerror {
	ue := &uerror{
		s:  text,
		us: usertext,
	}
	return ue
}

// FromErrors accepts err and uerr which satisfy the error interface and creates a new uerror
func FromErrors(err, uerr error) *uerror {
	ue := &uerror{
		s:  err.Error(),
		us: uerr.Error(),
	}
	return ue
}

func (ue *uerror) Error() string {
	return ue.s
}

func (ue *uerror) UserError() string {
	return ue.us
}
