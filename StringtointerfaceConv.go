package main

import (
    "fmt"
)

func Display(vals []interface{}) {
    for _, val := range vals {
        fmt.Println(val)
    }
}

func main() {
    names := []string{"Gopal", "Mohan", "Bikesh"}
 
//conversion of string to interface type 
  vals := make([]interface{}, len(names))
    for i, v := range names {
        vals[i] = v
    }
    Display(vals)
}

