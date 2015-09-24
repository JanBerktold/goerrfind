package tests

import (
	"errors"
)

const (
	Err1 = errors.New("error 1")
)

func testSwitch(a int) error {
	err := errors.New("yagdg")
	switch a {
	case 1:
		return nil
	case 2:
		return testSwitch(1)
	case 3:
		return err
	}
	return err
}

func testIf() error {
	err := errors.New("yagdg")
	if 1 == 2 {
		return nil
	} else if 2 == 1 {
		err = errors.New("another")
		return err
	} else {
		return Err1
	}
	return errors.New("hey")
}
