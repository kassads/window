package main

import (
	//"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {

	MainWindow{
		Title:   "buh",
		MinSize: Size{600, 400},
		Layout:  VBox{},
	}.Run()
}