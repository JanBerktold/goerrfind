package tests

type SomeStruct struct{}

func (s *SomeStruct) Pointer() {
}

func (s SomeStruct) NoPointer() {
}

func (s *SomeStruct) hiddenPointer() {
}

func (s SomeStruct) hiddenNoPointer() {
}

func SomeMethod() {
}

func hiddenMethod() {
}
