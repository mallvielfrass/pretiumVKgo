package main
import (
"fmt"
)
func main() {
 
ids:="410747880"
fields:="bdate"
res:=send(ids,fields) 
fmt.Println(res)
}
