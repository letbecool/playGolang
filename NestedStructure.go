
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
