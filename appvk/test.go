pacdkage main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter a string")
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(line))
}
