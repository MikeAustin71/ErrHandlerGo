package common

import (
	"errors"
	"fmt"
)

// TestEx001 - Test Example designed to print
// multiple parent data associated with a
// single error
func TestEx001() {
	bi := ErrBaseInfo{}

	f := bi.New("TestSourceFileName", "TestFuncName", 9000)
	g := bi.New("TestSrcFileName2", "TestFuncName2", 14000)
	h := bi.New("TestSrcFileName3", "TestFuncName3", 15000)

	ex1 := make([]ErrBaseInfo, 0, 10)
	ex1 = append(ex1, f, g, h)

	ex21 := "TestSrcFileName99"
	ex22 := "TestFuncName99"
	ex23 := int64(16000)
	ex2 := bi.New(ex21, ex22, ex23)

	ex3 := "prefixString"
	ex4 := "Error Msg 99"
	err := errors.New(ex4)
	ex6 := int64(22)

	x := SpecErr{}.Initialize(ex1, ex2, ex3, err, false, ex6)

	fmt.Println(x.Error())

}

// TestEx002 -
func TestEx002() {
	// Set up Parent Info
	bi := ErrBaseInfo{}

	f := bi.New("TestSourceFileName", "TestFuncName", 9000)
	g := bi.New("TestSrcFileName2", "TestFuncName2", 14000)
	h := bi.New("TestSrcFileName3", "TestFuncName3", 15000)

	ex1 := make([]ErrBaseInfo, 0, 10)
	ex1 = append(ex1, f, g, h)

	// Set up BaseInfo
	ex21 := "TestSrcFileName99"
	ex22 := "TestFuncName99"
	ex23 := int64(16000)
	ex2 := bi.New(ex21, ex22, ex23)

	ex3 := "prefixString"
	ex4 := "This is the Error Msg"
	err := errors.New(ex4)
	ex6 := int64(22)

	x := SpecErr{}.Initialize(ex1, ex2, ex3, err, true, ex6)

	panic(x)
}
