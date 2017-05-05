package main

import (
	common "MikeAustin71/ErrHandlerGo/01_ErrUtil/common"
	"fmt"
)

func main() {

	var bi common.ErrBaseInfo
	x := bi.New("TestSourceFileName", "TestFuncName", 9000)
	y := bi.New("TestSrcFileName2", "TestFuncName2", 14000)
	z := bi.New("TestSrcFileName3", "TestFuncName3", 15000)

	a := make([]common.ErrBaseInfo, 0, 30)

	a = append(a, x, y, z)

	var se common.SpecErr

	se.ParentInfo = se.SetParentInfo(a)

	l := len(se.ParentInfo)

	if l != 3 {
		fmt.Println("Length of se.ParentInfo is NOT 3")
		return
	}

	if se.ParentInfo[1].FuncName != "TestFuncName2" {
		fmt.Println("Expected 2nd Element 'TestFuncName2', got", se.ParentInfo[1].FuncName)
		return
	}

	fmt.Println("Successful Completion!")

}
