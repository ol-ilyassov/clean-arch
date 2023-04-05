package surname

import (
	"github.com/pkg/errors"
)

var (
	MaxLength      = 100
	ErrWrongLength = errors.Errorf("surname must be less than or equal to %d characters", MaxLength)
)

type Surname string

func (s Surname) String() string {
	return string(s)
}

func New(name string) (*Surname, error) {
	if len([]rune(name)) > MaxLength {
		return nil, ErrWrongLength
	}
	s := Surname(name)
	return &s, nil
}
