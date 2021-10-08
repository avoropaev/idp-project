package config

import "errors"

var (
	ErrEmptyS1Url = errors.New("s1 url is empty")
	ErrEmptyS2Url = errors.New("s2 url is empty")
)

type External struct {
	S1 string
	S2 string
}

func (e External) Validate() (err error) {
	switch {
	case e.S1 == "":
		return ErrEmptyS1Url
	case e.S2 == "":
		return ErrEmptyS2Url
	}

	return nil
}
