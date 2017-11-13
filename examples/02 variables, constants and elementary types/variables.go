package main

//syntax for variable
//var variableName type

var i int
var vname1, vname2, vname3 int //all three variables are of type int
var x = 5
var vnam1, vnam2, vnam3 int = 1, 2, 3

func main() {
	i = 2
	x = x + i

	/*
		Define three variables without type "type" and without keyword "var", and initialize their values.
		vname1 is v1，vname2 is v2，vname3 is v3
	*/
	vna1, vna2, vna3 := 1, 2, 3
	vna1 = vna1 + 1
	vna1 = vna2
	vna2 = vna3
	vna3 = vna1

}
