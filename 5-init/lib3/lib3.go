package lib3

import (
	"GolangStudy/5-init/lib4"
	"fmt"
)

func init() {
	fmt.Println("lib3.init ...")
}

func Lib3Test() {
	fmt.Println("lib3.Lib3Test ...")
	lib4.Lib4Test()
}
