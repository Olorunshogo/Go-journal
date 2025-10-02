
package main
import ("fmt")

// var (
// 	a int = 2
// 	b int = 4
// 	c, d string = "c is: Hello", "d is: World"
// )

func goBy() {
	// a = 1
	// fmt.Println(a,b)
	// fmt.Printf("a has value: %v and type: %T\n", a, a)
	// fmt.Printf("a has value: %#v and type: %T\n", a, a) // This should output the value using Go-syntax format.
	// fmt.Printf("b has value: %v and type: %T\n", b, b)

	// fmt.Println(c,"\n",d)
	// fmt.Printf("c has value: %v and type: %T\n", c, d)
	// fmt.Printf("d has value: %v and type: %T\n", d, d)

	// var i = 15

	// fmt.Printf("%b\n", i)
	// fmt.Printf("%d\n", i)
	// fmt.Printf("%+d\n", i)
	// fmt.Printf("%o\n", i)
	// // fmt.Printf("%0\n", i)
	// fmt.Printf("%x\n", i)
	// fmt.Printf("%X\n", i)
	// fmt.Printf("%#x\n", i)
	// fmt.Printf("%4d\n", i)
	// fmt.Printf("%-4d\n", i)
	// fmt.Printf("%04d\n", i)

	// // New Line
	// fmt.Printf("\n")

	// // Go Data Types
	// fmt.Printf("Go DATA TYPES \n \n")

	// var k bool = true
	// var l int = 5
	// var m float32 = 3.14
	// var n string = "Hi! I'm a string."

	// fmt.Println("Boolean: ", k)
	// fmt.Println("Integer: ", l)
	// fmt.Println("Float: ", m)
	// fmt.Println("String: ", n)

	// New Line
	// fmt.Printf("\n")

	// Go Strings
	// fmt.Printf("Go STRINGS \n \n")

	// var array1 = [4]int{0,1,2,3}
	// var array2 = [5]int{4,5,6,7,8}
	// array3 := [6]int{9, 10, 11, 12, 13, 14,}
	// var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	// var prices = [3]int{10,20,30}

	// fmt.Println("Array 1 is: ", array1)
	// fmt.Println("Array 2 is: ", array2)
	// fmt.Print("The array of cars is: ", cars, "\n")

	// fmt.Println(prices[0])
	// fmt.Println(prices[2])

	// prices[2] = 50
	// fmt.Println("List prices: ", prices)

	// fmt.Println("The length os cars is: ", len(cars))
	// fmt.Println("The length os prices is: ", len(prices))

	// New Line
	// fmt.Printf("\n")

	// // Go Slices
	// fmt.Printf("Go SLICES \n \n")

	// var mySlice1 = []int{0,1,2,3}
	// fmt.Println(len(mySlice1))
	// fmt.Println(cap(mySlice1))
	// fmt.Println(mySlice1)

	// var mySlice2 = []string{"Go", "Slices", "Are", "Powerful"}
	// fmt.Println(len(mySlice2))
	// fmt.Println(cap(mySlice2))
	// fmt.Println(mySlice2)

	// // // Creating a Slice from an Array
	// sliceArray1 := array1[0:3]
	// sliceArray2 := array2[1:4]
	// sliceArray3 := array3[0:6]

	// fmt.Println("Slice Array 1 is: ", sliceArray1)
	// fmt.Println("Slice Array 2 is: ", sliceArray2)

	// fmt.Printf("mySlice = %v\n", sliceArray3)
	// fmt.Printf("Length = %d\n", len(sliceArray3))
	// fmt.Printf("Capacity = %d\n", cap(sliceArray3))

	// array1[2] = 50
	// fmt.Println(array1) 
	// fmt.Println(array1[0])
	// fmt.Println(array1[2])

	// fmt.Println("sliceArray1 is: ", sliceArray1)
	// fmt.Println(len(sliceArray1))

	// sliceArray1 = append(sliceArray1, 30, 32)

	// println("sliceArray1 is now: ", sliceArray1)
	// fmt.Println(len(sliceArray1))

	// numbers := append(array1[:], array2[:]...)
	// numbers = append(numbers, array3[:]...)

	// fmt.Println("Combined numbers: ", numbers)

	// fmt.Printf("Numbers = %v\n", numbers)
	// fmt.Printf("Length = %d\n", len(numbers))
	// fmt.Printf("Capacity = %d\n", cap(numbers))

	// // // Create a copy with only needed numbers
	// neededNumbers := numbers[:len(numbers)-10]
	// numbersCopy := make([]int, len(neededNumbers))
	// copy(numbersCopy, neededNumbers)

	// fmt.Printf("numbersCopy = %v\n", numbersCopy)
	// fmt.Printf("Lenght = %d\n", len(numbersCopy))
	// fmt.Printf("Capacity = %d\n", cap(numbersCopy))

	// // New Line
	// fmt.Printf("\n")

	// // Go Operators
	// fmt.Printf("Go OPEARATORS \n \n")

	// p1 := 15 + 15
	// p2 := 15 - 15
	// p3 := 15 * 15
	// p4 := 15 / 15
	// p5 := 15 % 15

	// sumP := p1 + p2 + p3 + p4 + p5	
	// fmt.Println("The value of p1-p5 is: ", p1,p2,p3,p4,p5, " and its sum is: ", sumP)

	// // Go Conditions
	// fmt.Printf("Go CONDITIONS \n \n")


	// if p1 > p2 {
	// 	fmt.Println(p1, " is greater than ", p2)
	// }

	// if (sumP > 250) {
	// 	fmt.Println("The sum of all p is greater than ", 300)
	// } else {
	// 	fmt.Println("The sum of all p is less than 300.")
	// }

	// temperature := 28
	// if (temperature > 30) {
	// 	fmt.Println("It is very hot out there.")
	// 	if (temperature > 39) {
	// 		fmt.Println("Don't even try to go out there.")
	// 	}
	// } else if temperature < 16 {
	// 	fmt.Println("It is very cold out there.")
	// } else {
	// 	fmt.Println("It is warm out there.")
	// }


	// // New Line
	// fmt.Printf("\n")

	// // Go Switch Statement
	// fmt.Printf("Go SWITCH STATEMENT \n \n")

	// day := 5

	// switch day {
	// case 1:
	// 	fmt.Println("Sunday")
	// case 2:
	// 	fmt.Println("Monday")
	// case 3:
	// 	fmt.Println("Tuesday")
	// case 4:
	// 	fmt.Println("Wednesday")
	// case 5:
	// 	fmt.Println("Thursday")
	// case 6:
	// 	fmt.Println("Friday")
	// case 7:
	// 	fmt.Println("Saturday")
	// default:
	// 	fmt.Println("Not a weekday")
	// }

	// switch day {
	// 	case 1,3,5:
	// 		fmt.Println("Odd weekday")
	// 	case 2,4:
	// 		fmt.Println("Even weekday")
	// 	case 6,7:
	// 		fmt.Println("Weekend")
	// 	default:
	// 		fmt.Println("Invalid day of day number")
	// }

	// New Line
	fmt.Printf("\n")

	// Go Loops
	fmt.Printf("Go LOOPS \n \n")

	for i := 0; i <= 100; i+=5 {
		if (i == 55) {
			continue
		}

		if (i == 75) {
			break
		}
		fmt.Println(i)
	}

	adj := [2]string{"Big", "Tasty"}
	fruits := [3]string{"Apple", "Orange", "Banana"}
	for i:=0; i < len(adj); i++ {
		for j:=0; j < len(fruits); j++ {
			fmt.Println(adj[1],fruits[j])
		}
	}

	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}








}

func main() {
	goBy()
}