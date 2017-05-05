package common

import (
	"fmt"
	"testing"
)

func TestUninitializedErrBaseInfo(t *testing.T) {
	var se ErrBaseInfo

	if se.SourceFileName != "" {
		t.Error("String SourceFileName was uninitialized. Was expecting empty string, got", se.SourceFileName)
	}

	if se.FuncName != "" {
		t.Error("String FuncName was uninitialized. Was expecting empty string, got", se.FuncName)
	}

	if se.BaseErrorID != 0 {
		t.Error("Int64 BaseErrorID was uninitialized. Was expecting value of zero, got", se.BaseErrorID)
	}

}

// Creates new BaseInfo Structure
func TestNewBaseInfo(t *testing.T) {
	ex1 := "TestSourceFileName"
	ex2 := "TestFuncName"
	ex3 := int64(9000)

	var bi ErrBaseInfo
	x := bi.New(ex1, ex2, ex3)

	if x.SourceFileName != ex1 {
		t.Error(fmt.Sprintf("Expected %v got,", ex1), x.SourceFileName)
	}

	if x.FuncName != ex2 {
		t.Error(fmt.Sprintf("Expected %v got,", ex2), x.FuncName)
	}

	if x.BaseErrorID != ex3 {
		t.Error(fmt.Sprintf("Expected %v got,", ex3), x.BaseErrorID)
	}

}

// Changes Function Name Only
func TestNewFunc(t *testing.T) {
	ex1 := "TestSourceFileName"
	ex2 := "TestFuncName"
	ex3 := int64(9000)
	ex4 := "NewFuncName"

	var bi ErrBaseInfo
	x := bi.New(ex1, ex2, ex3)

	if x.FuncName != ex2 {
		t.Error(fmt.Sprintf("Expected %v got,", ex2), x.FuncName)
	}

	y := x.NewFunc(ex4)

	if y.SourceFileName != ex1 {
		t.Error(fmt.Sprintf("Expected %v got,", ex1), y.SourceFileName)
	}

	if y.FuncName != ex4 {
		t.Error(fmt.Sprintf("Expected %v got,", ex4), y.FuncName)
	}

	if y.BaseErrorID != ex3 {
		t.Error(fmt.Sprintf("Expected %v got,", ex3), y.BaseErrorID)
	}
}

func TestBaseInfoDeepCopy(t *testing.T) {
	ex1 := "TestSourceFileName"
	ex2 := "TestFuncName"
	ex3 := int64(9000)

	var bi ErrBaseInfo
	x := bi.New(ex1, ex2, ex3)

	y := x.NewBaseInfo()

	if y.SourceFileName != ex1 {
		t.Error(fmt.Sprintf("Expected %v got,", ex1), y.SourceFileName)
	}

	if y.FuncName != ex2 {
		t.Error(fmt.Sprintf("Expected %v got,", ex2), y.FuncName)
	}

	if y.BaseErrorID != ex3 {
		t.Error(fmt.Sprintf("Expected %v got,", ex3), y.BaseErrorID)
	}

}

func TestGetSpecErrFromBaseInfo(t *testing.T) {
	ex1 := "TestSourceFileName"
	ex2 := "TestFuncName"
	ex3 := int64(9000)
	var bi ErrBaseInfo

	x := bi.New(ex1, ex2, ex3)

	se := x.GetBaseSpecErr()

	if se.IsBaseInfoSet == false {
		t.Error("IsBaseInfoSet should be true. Instead, value was ", se.IsBaseInfoSet)
	}

	if se.BaseInfo.SourceFileName != "TestSourceFileName" {
		t.Error("Expected BaseInfo.SourceFileName 'TestSourceFileName', got ", se.BaseInfo.SourceFileName)
	}

	if se.BaseInfo.FuncName != "TestFuncName" {
		t.Error("Expected BaseInfo.FuncName 'TestFuncName', got ", se.BaseInfo.FuncName)
	}

	if se.BaseInfo.BaseErrorID != 9000 {
		t.Error("Expected BaseInfo.BaseErrorID '9000', got ", se.BaseInfo.BaseErrorID)
	}

}
