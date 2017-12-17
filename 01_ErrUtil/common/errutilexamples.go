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

	x := SpecErr{}.Initialize(ex1, ex2, "", err, SpecErrTypeERROR, ex6)

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

	x := SpecErr{}.Initialize(ex1, ex2, "", err, SpecErrTypeFATAL, ex6)

	panic(x)
}

func TestSpecErrStandardError_001(){

	se := SpecErr{}.InitializeBaseInfo(testExampleSpecErrParentInfo(), testExampleSpecErrCreateErrBaseInfoObj())
	se.SetStdError("This is Standard Error Message.", 974 )
	fmt.Printf(se.String())
}

func TestSpecErrStandardError_002(){

	se := SpecErr{}.InitializeCurrentBaseInfo(testExampleSpecErrCreateErrBaseInfoObj())
	se.SetStdError("This is Standard Error Message.", 974 )
	fmt.Printf(se.String())
}

func TestSpecErrStandardError_003(){

	se := SpecErr{}
	se.SetStdError("This is Standard Error Message.", 974 )
	fmt.Printf(se.String())
}

func TestSpecErrStandardError_004(){

	se := SpecErr{}.InitializeBaseInfo(testExampleSpecErrParentInfo(), testExampleSpecErrCreateErrBaseInfoObj())
	se.SetStdError("", 974 )
	fmt.Printf(se.String())
}

func TestSpecErrFatalError_001(){

	se := SpecErr{}.InitializeBaseInfo(testExampleSpecErrParentInfo(), testExampleSpecErrCreateErrBaseInfoObj())
	se.SetFatalError("This is Fatal Error Message.", 224 )
	fmt.Printf(se.String())
}

func TestSpecErrFatalError_002(){

	se := SpecErr{}.InitializeBaseInfo(testExampleSpecErrParentInfo(), testExampleSpecErrCreateErrBaseInfoObj())
	se.SetFatalError("This is Fatal Error Message.", 224 )
	se.SetMessageOutputMode(false)
	fmt.Printf(se.String())
}



/*
=======================================================================================================
								Private Methods
=======================================================================================================
 */

func testExampleSpecErrCreateErrBaseInfoObj() ErrBaseInfo {
	ebi := ErrBaseInfo{}
	return ebi.New("TSource06", "PObj06", "Func006", 6000)
}

func testExampleSpecErrParentInfo() []ErrBaseInfo {
	ebi := ErrBaseInfo{}

	x1 := ebi.New("TSource01", "PObj01", "Func001", 1000)
	x2 := ebi.New("TSource02", "PObj02", "Func002", 2000)
	x3 := ebi.New("TSource03", "PObj03", "Func003", 3000)
	x4 := ebi.New("TSource04", "PObj04", "Func004", 4000)
	x5 := ebi.New("TSource05", "PObj05", "Func005", 5000)

	parent := make([]ErrBaseInfo,0,10)

	parent = append(parent, x1)
	parent = append(parent, x2)
	parent = append(parent, x3)
	parent = append(parent, x4)
	parent = append(parent, x5)

	return parent
}


func testExampleSpecErrCreateStdErrorMsg() SpecErr {
	se := SpecErr{}.InitializeBaseInfo(testExampleSpecErrParentInfo(), testExampleSpecErrCreateErrBaseInfoObj())
	se.SetStdError("This is Standard Error Msg for test object", 429)
	return se
}

func testExampleSpecErrCreateFatalErrorMsg() SpecErr {
	se := SpecErr{}.InitializeBaseInfo(testExampleSpecErrParentInfo(), testExampleSpecErrCreateErrBaseInfoObj())
	se.SetFatalError("This is FATAL Error Msg for test object", 152)
	return se
}

func testExampleSpecErrCreateInfoMsg() SpecErr {
	se := SpecErr{}.InitializeBaseInfo(testExampleSpecErrParentInfo(), testExampleSpecErrCreateErrBaseInfoObj())
	se.SetInfoMessage("This is Information Message for test object", 19)
	return se
}

func testExampleSpecErrCreateWarningMsg() SpecErr {
	se := SpecErr{}.InitializeBaseInfo(testExampleSpecErrParentInfo(), testExampleSpecErrCreateErrBaseInfoObj())
	se.SetWarningMessage("This is Warning Message for test object.", 67)
	return se
}


func testExampleSpecErrCreateSuccessfulCompletionMsg() SpecErr {
	se := SpecErr{}.InitializeBaseInfo(testExampleSpecErrParentInfo(), testExampleSpecErrCreateErrBaseInfoObj())
	se.SetSuccessfulCompletion("", 64)
	return se
}

func testExampleSpeErrCreateNoErrorsNoMessagesMsg() SpecErr {
	se := SpecErr{}.InitializeBaseInfo(testExampleSpecErrParentInfo(), testExampleSpecErrCreateErrBaseInfoObj())
	se.SetNoErrorsNoMessages("",28)
	return se
}
