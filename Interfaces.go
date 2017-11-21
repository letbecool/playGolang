package main

import "fmt"

type name struct {
	fname string
	lname string
}
type agent struct {
	name
	condition bool
}
//interface
type anyone interface{
check()
}

func (n name) check() {
	fmt.Println("my name is ", n.fname, n.lname)
}
func (a agent) check() {
	fmt.Println(a.fname, a.lname, a.condition)
}
func letstry(any anyone){
any.check()
}

func main() {

	p1 := name{
		"ram",
		"shyam",
	}
	p2 := agent{
		name{"Bikesh",
			"shrestha"},
		true,
	}
	p1.check()
	p2.check()
	//in talking minimum , if we want to call only the name which belongs to agent 
	p2.name.check()
	// anything inside the struct name and agent will be the value of interface anyone
	letstry(p1)
	letstry(p2)
	letstry(p2.name)
}

/*
output
my name is  ram shyam
Bikesh shrestha true
my name is  Bikesh shrestha
my name is  ram shyam
Bikesh shrestha true
my name is  Bikesh shrestha

*/
