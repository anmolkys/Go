package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
// func main() {
// 	welcome := "Welcome "
// 	fmt.Println(welcome)
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter rating: ")

// 	//comma ok //err ok
// 	//works as try //works as catch

// 	input, _ := reader.ReadString('\n')
// 	fmt.Println("Thanks for your rating", input)
// 	fmt.Printf("Type of rating is %T \n", input)

func main() {
	fmt.Println("Welcome to Pizza App")
	fmt.Println("Rate our Pizza")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1", numRating)
	}

}
