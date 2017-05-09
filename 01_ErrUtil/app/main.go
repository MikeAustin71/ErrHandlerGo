package main

import (
	common "MikeAustin71/ErrHandlerGo/01_ErrUtil/common"
	"fmt"
)

func main() {

	parent := common.ErrBaseInfo{}.GetNewParentInfo("main.go", "main", 1000)

	fmt.Println("Length of 'parent' is:", len(parent))

	fmt.Println("SourceFile Name should be 'main.go'. Actual=: ", parent[0].SourceFileName)

}

func testErr(parent []ErrBaseInfo){

}
