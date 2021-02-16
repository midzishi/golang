package main
import (
	"fmt"
	"math/rand"
	"strings"
)
// print println
// printf работает как formated
// const менять нельзя
// переменные объявляем var
// strings.Contains -- содержится ли слово в строке
func main() {
	fmt.Println("Hello")
	var SomeText = "Text"
	var num = rand.Intn(401000000-56000000) + 56000000
	fmt.Printf("%v -- смысл жизни, мира, вселенной и вообще всего\n", 42)
	fmt.Printf("%s -- some text\n", SomeText)
	fmt.Printf("%v is random number\n", num)
	var boolean = strings.Contains(SomeText, "ext")
	fmt.Println(boolean)
	var number = 15
	var predict = 0
	var n = 101
	var t = 0
	var count = 0
	for {
		fmt.Printf("Looking from %v to %v\n", t, n)
		predict = rand.Intn(n)+t
		if predict == number{
			fmt.Printf("%v == %v. Found for %v steps\n", predict, number, count)
			break
		}
		if predict > number {
			fmt.Printf("%v more than %v\n", predict, number)
			n--
		}
		if predict < number {
			fmt.Printf("%v less than %v\n", predict, number)
			t++
		}
		count ++
	}
}
