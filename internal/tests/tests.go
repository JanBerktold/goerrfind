package tests

import (
	"errors"
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
		return err
	} else {
		return nil
	}
	return errors.New("hey")
}
