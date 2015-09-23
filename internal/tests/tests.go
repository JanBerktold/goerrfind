package tests

func testSwitch(a int) error {
	e := errors.New("yagdg")
	switch a {
	case 1:
		return nil
	case 2:
		return testSwitch(1)
	case 3:
		return e
	}
	return err
}

func testIf() error {
	e := errors.New("yagdg")
	if 1 == 2 {
		return nil
	} else if 2 == 1 {
		return e
	} else {
		fmt.Println("adgadg")
		return nil
	}
	return errors.New("hey")
}
