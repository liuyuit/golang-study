package main

import (
	//"GolangStudy/5-init/lib1" // "GolangStudy/5-init/lib3"
	_ "GolangStudy/5-init/lib1"
	mylib2 "GolangStudy/5-init/lib2"
	. "GolangStudy/5-init/lib4"
)

func main() {
	mylib2.Lib2Test()
	Lib4Test()
}
