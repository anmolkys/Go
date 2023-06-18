package main

import (
	"bufio"
	"fmt"
	"os"
)

const LoginToken string = "aghhgdfhfhf"

//lexer : removes ; , assigns types
//types
// func main() {
// 	var flo float64 = 255.47834278755789
// 	fmt.Print(flo)

// 	//aliases
// 	var st = "ABC"
// 	fmt.Printf(st)

// 	//no var , walrus op(:=) is used (cannot be used in global)

// 	num := 300000
// 	fmt.Println(num)

// 	//in public declaration the first letter is caps

//		fmt.Print(LoginToken)
//	}
func main() {
	welcome := "Welcome "
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating: ")

	//comma ok //err ok
	//works as try //works as catch

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your rating", input)
	fmt.Printf("Type of rating is %T \n", input)

}
