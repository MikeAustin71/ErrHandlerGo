package common

import (
	"errors"
	"testing"
)

func TestErrorUtility(t *testing.T) {

	ex1 := "errprefix"
	ex2 := "errutil_test.go"
	ex3 := "TestErrorUtility"
	ex4 := int64(334)
	ex99 := "Test Error #1"

	err := errors.New(ex99)

	var se SpecErr

	x := se.New(ex1, err, true, ex2, ex3, ex4)

	if x.PrefixMsg != ex1 {
		t.Error("Expected 'errprefix' got", x.PrefixMsg)
	}

	if x.ErrMsg != ex99 {
		t.Error("Expected 'Test Error #1' got", x.ErrMsg)
	}

	if x.SrcFile != ex2 {
		t.Error("Expected 'errutil_test.go' got", x.SrcFile)
	}

	if x.FuncName != ex3 {
		t.Error("Expected 'TestErrorUtility' got", x.FuncName)
	}

	if x.ErrNo != ex4 {
		t.Error("Expected '334' got", x.ErrNo)
	}

}

func TestUninitializedBaseInfo(t *testing.T) {
	var se SpecErr

	if se.BaseInfo.SourceFileName != "" {
		t.Error("String SourceFileName was uninitialized. Was expecting empty string, got", se.BaseInfo.SourceFileName)
	}

	if se.BaseInfo.FuncName != "" {
		t.Error("String FuncName was uninitialized. Was expecting empty string, got", se.BaseInfo.FuncName)
	}

	if se.BaseInfo.BaseErrorID != 0 {
		t.Error("Int64 BaseErrorID was uninitialized. Was expecting value of zero, got", se.BaseInfo.BaseErrorID)
	}

}

func TestInitializeParentInfo(t *testing.T) {

	var bi ErrBaseInfo
	x := bi.New("TestSourceFileName", "TestFuncName", 9000)
	y := bi.New("TestSrcFileName2", "TestFuncName2", 14000)
	z := bi.New("TestSrcFileName3", "TestFuncName3", 15000)

	var se SpecErr

	se.ParentInfo = append(se.ParentInfo, x)
	se.ParentInfo = append(se.ParentInfo, y)
	se.ParentInfo = append(se.ParentInfo, z)

	l := len(se.ParentInfo)

	if l != 3 {
		t.Error("Expected ParentInfo Length of 3, got", l)
	}

	if se.ParentInfo[1].FuncName != "TestFuncName2" {
		t.Error("Expected 2nd Element 'TestFuncName2', got", se.ParentInfo[1].FuncName)
	}

}

func TestAddSlicesParentInfo(t *testing.T) {
	var bi ErrBaseInfo
	x := bi.New("TestSourceFileName", "TestFuncName", 9000)
	y := bi.New("TestSrcFileName2", "TestFuncName2", 14000)
	z := bi.New("TestSrcFileName3", "TestFuncName3", 15000)

	a := make([]ErrBaseInfo, 0, 30)

	a = append(a, x, y, z)

	var se SpecErr

	se.ParentInfo = a

	l := len(se.ParentInfo)

	if l != 3 {
		t.Error("Expected ParentInfo Length of 3, got", l)
	}

	if se.ParentInfo[1].FuncName != "TestFuncName2" {
		t.Error("Expected 2nd Element 'TestFuncName2', got", se.ParentInfo[1].FuncName)
	}

}

func TestSetParentInfo(t *testing.T) {
	var bi ErrBaseInfo
	x := bi.New("TestSourceFileName", "TestFuncName", 9000)
	y := bi.New("TestSrcFileName2", "TestFuncName2", 14000)
	z := bi.New("TestSrcFileName3", "TestFuncName3", 15000)

	a := make([]ErrBaseInfo, 0, 30)

	a = append(a, x, y, z)

	var se SpecErr

	se.ParentInfo = se.SetParentInfo(a)

	l := len(se.ParentInfo)

	if l != 3 {
		t.Error("Expected ParentInfo length of 3, go length of ", l)
	}

	if se.ParentInfo[1].FuncName != "TestFuncName2" {
		t.Error("Expected 2nd Element 'TestFuncName2', got", se.ParentInfo[1].FuncName)
	}
}

func TestSetErrDetail(t *testing.T) {
	var bi ErrBaseInfo
	x := bi.New("TestSrcFileName2", "TestFuncName2", 14000)

	ex1 := "errprefix"
	ex4 := int64(334)
	err := errors.New("Test Error #1")

	se := x.GetBaseSpecErr().NewDetailErr(ex1, err, false, ex4)

	if se.ErrNo != 14334 {
		t.Error("Expected Err No 14334, go", se.ErrNo)
	}

	if se.SrcFile != "TestSrcFileName2" {
		t.Error("Expected Source File: 'TestSrcFileName2',got", se.SrcFile)
	}

	if se.FuncName != "TestFuncName2" {
		t.Error("Expected FuncName: 'TestFuncName2', got", se.FuncName)
	}

}

func TestIsSpecErrNo(t *testing.T) {
	var se SpecErr
	s := se.SetNoError()

	isErr := CheckIsSpecErr(s)

	if isErr {
		t.Error("Expected CheckIsSpecErr() to return false, go", isErr)
	}

}

func TestIsSpecErrYes(t *testing.T) {
	ex1 := "errprefix"
	ex2 := "errutil_test.go"
	ex3 := "TestErrorUtility"
	ex4 := int64(334)
	ex99 := "Test Error #1"

	err := errors.New(ex99)

	var se SpecErr

	x := se.New(ex1, err, true, ex2, ex3, ex4)

	isErr := CheckIsSpecErr(x)

	if !isErr {
		t.Error("Expected CheckIsSpecErr() to return true, go", isErr)
	}

}

func TestSetNoErr(t *testing.T) {

	var se SpecErr

	x := se.SetNoError()

	if x.IsErr {
		t.Error("Expected IsErr = 'false', got", x.IsErr)
	}
}
