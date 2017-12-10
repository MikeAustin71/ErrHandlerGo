package main

import (
	"MikeAustin71/ErrHandlerGo/01_ErrUtil/common"
	"fmt"
)

/*

import (
	"MikeAustin71/ErrHandlerGo/01_ErrUtil/common"
	"errors"
	"fmt"
)

*/

func main() {

	common.TestErrUtilityExample001()


}

/*
func testOpsMsgDtoCreateContextInfoObj() common.OpsMsgContextInfo {
	ci := common.OpsMsgContextInfo{}
	return ci.New("TSource06", "PObj06", "Func006", 6000)
}

func testOpsMsgDtoCreateParentHistory() []common.OpsMsgContextInfo {
	ci := common.OpsMsgContextInfo{}

	x1 := ci.New("TSource01", "PObj01", "Func001", 1000)
	x2 := ci.New("TSource02", "PObj02", "Func002", 2000)
	x3 := ci.New("TSource03", "PObj03", "Func003", 3000)
	x4 := ci.New("TSource04", "PObj04", "Func004", 4000)
	x5 := ci.New("TSource05", "PObj05", "Func005", 5000)

	parent := make([]common.OpsMsgContextInfo,0,10)

	parent = append(parent, x1)
	parent = append(parent, x2)
	parent = append(parent, x3)
	parent = append(parent, x4)
	parent = append(parent, x5)

	return parent
}
*/