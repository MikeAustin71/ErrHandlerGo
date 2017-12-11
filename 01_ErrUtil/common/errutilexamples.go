package common

import (
	"errors"
	"fmt"
)

// TestErrUtilityExample001 - Test Example designed to print
// multiple parent data associated with a
// single error
func TestErrUtilityExample001() {
	bi := ErrBaseInfo{}

	f := bi.New("TestSourceFileName", "TestObject", "TestFuncName", 9000)
	g := bi.New("TestSrcFileName2", "TestObject2", "TestFuncName2", 14000)
	h := bi.New("TestSrcFileName3", "TestObject3", "TestFuncName3", 15000)

	ex1 := make([]ErrBaseInfo, 0, 10)
	ex1 = append(ex1, f, g, h)

	ex21 := "TestSrcFileName99"
	ex21ParentObj := "TestObject99"
	ex22 := "TestFuncName99"
	ex23 := int64(16000)
	ex2 := bi.New(ex21, ex21ParentObj, ex22, ex23)

	ex4 := "This is the Error Message"
	err := errors.New(ex4)
	ex6 := int64(22)

	x := SpecErr{}.Initialize(ex1, ex2, err, SpecErrTypeERROR, ex6)

	fmt.Println(x.Error())

}

// TestErrorUtilityExample_002 - SpeErr example
func TestErrorUtilityExample002() {
	// Set up Parent Info
	bi := ErrBaseInfo{}

	f := bi.New("TestSourceFileName", "TestObject", "TestFuncName", 9000)
	g := bi.New("TestSrcFileName2", "TestObject2", "TestFuncName2", 14000)
	h := bi.New("TestSrcFileName3", "TestObject3", "TestFuncName3", 15000)

	ex1 := make([]ErrBaseInfo, 0, 10)
	ex1 = append(ex1, f, g, h)

	// Set up BaseInfo
	ex21 := "TestSrcFileName99"
	ex21ParentObj := "TestObject99"
	ex22 := "TestFuncName99"
	ex23 := int64(16000)
	ex2 := bi.New(ex21, ex21ParentObj, ex22, ex23)

	ex4 := "This is the Error Msg"
	err := errors.New(ex4)
	ex6 := int64(22)

	x := SpecErr{}.Initialize(ex1, ex2, err, SpecErrTypeFATAL, ex6)

	panic(x)
}
