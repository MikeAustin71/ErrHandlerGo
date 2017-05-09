package main

import (
	common "MikeAustin71/ErrHandlerGo/01_ErrUtil/common"
	"errors"
	"fmt"
)

func main() {

	bi := common.ErrBaseInfo{}

	f := bi.New("TestSourceFileName", "TestFuncName", 9000)
	g := bi.New("TestSrcFileName2", "TestFuncName2", 14000)
	h := bi.New("TestSrcFileName3", "TestFuncName3", 15000)

	ex1 := make([]common.ErrBaseInfo, 0, 10)
	ex1 = append(ex1, f, g, h)

	ex2_1 := "TestSrcFileName99"
	ex2_2 := "TestFuncName99"
	ex2_3 := int64(16000)
	ex2 := bi.New(ex2_1, ex2_2, ex2_3)

	ex3 := "prefixString"
	ex4 := "Error Msg 99"
	err := errors.New(ex4)
	ex5 := false
	ex6 := int64(22)

	x := common.SpecErr{}.Initialize(ex1, ex2, ex3, err, ex5, ex6)

	fmt.Println(x.Error())

}
