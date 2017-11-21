/*

A field declared with a type but no explicit field name is an anonymous field, 
also called an embedded field or an embedding of the type in the struct. 
An embedded type must be specified as a type name T or as a pointer
 to a non-interface type name *T, and T itself may not be a pointer type. 
The unqualified type name acts as the field name.

*/
import "fmt"

type name struct {
	fname string
	lname string
}
type agent struct {
	name
	condition bool
}

func (n name) check() {
	fmt.Println("my name is ", n.fname, n.lname)
}
func (a agent) check() {
	fmt.Println(a.fname, a.lname, a.condition)
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
}
