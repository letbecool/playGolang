package main 
import(
"fmt"
"encoding/json"
"os"
)
type name struct{
Fname string 
Lname string  
Fabfood []string
}

func main(){


 
 collection := name{
 	Fname	: "Gopal",
 	Lname	: "Ojha",
 	Fabfood	: []string {"MOMO","Pizza","chicken roll"},

 } 
marshalleddata,err := json.Marshal(collection)
if err !=nil {
 fmt.Println("error:",err)
}
fmt.Println(collection)

fmt.Println(marshalleddata)
os.Stdout.Write(marshalleddata)


}