package main

import "fmt"

// FUNCTION

// func main() {
// 	greet("Airell", 23)
// }

// func greet(name string, age int) {
// 	fmt.Printf("Hello there! My name is %s and I'm %d years old", name, age)
// }

// FUNCTION (PARAMETER)

// func main() {
// 	greet("Airell", "Jalan Ronggo Waluyo")
// }

// func greet(name, address string) {
// 	fmt.Printf("Hi my name is %s\n", name)
// 	fmt.Printf("I live in %s\n", address)
// }

// // FUNCTION (RETURN)

// func main() {
// 	var names = []string{"Airell", "Jordan"}
// 	var printMsg = greet("Heii", names)

// 	fmt.Println(printMsg)
// }

// func greet(msg string, names []string) string {
// 	var joinStr = strings.Join(names, " ")

// 	var result string = fmt.Sprintf("%s %s", msg, joinStr)

// 	return result
// }

// FUNCTION (RETURNING MULTIPLE VALUES)

// func main() {
// 	var diameter float64 = 15
// 	var area, circumference = calculate(diameter)

// 	fmt.Println("Area :", area)
// 	fmt.Println("Circumference :", circumference)
// }

// func calculate(d float64) (area float64, circumference float64) {
// 	// Menghitung luas
// 	area = math.Pi * math.Pow(d/2, 2)
// 	// Menghitung keliling
// 	circumference = math.Pi * d

// 	return area, circumference
// }

// FUNCTION (VARADIC FUNCTION)

// func main() {
// 	studentList := print("Naruto", "Luffy", "Zoro", "Kaneki", "Deku")
// 	fmt.Printf("%v", studentList)
// }

// func print(names ...string) []map[string]string {
// 	var result []map[string]string

// 	for i, v := range names {
// 		key := fmt.Sprintf("student%d", i+1)
// 		temp := map[string]string{
// 			key: v,
// 		}
// 		result = append(result, temp)
// 	}
// 	return result
// }

// FUNCTION (VARADIC FUNCTION #2)

// func main() {
// 	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	result := sum(numberLists...)

// 	fmt.Println("Result:", result)
// }

// func sum(numbers ...int) int {
// 	total := 0

// 	for _, v := range numbers {
// 		total += v
// 	}
// 	return total
// }

// FUNCTION (VARADIC FUNCTION #3)

// func main() {
// 	profile("Alam", "Mie ayam", "Dimsum", "Kentang")
// }

// func profile(name string, favFoods ...string) {
// 	mergeFavfoods := strings.Join(favFoods, ", ")

// 	fmt.Println("Halo, nama saya", name)
// 	fmt.Println("Makanan favorit saya", mergeFavfoods)
// }

// CLOSURE (DECLARE CLOSURE IN VARIABLE)

// func main() {
// 	var evenNumbers = func(numbers ...int) []int {
// 		var result []int
// 		for _, v := range numbers {
// 			fmt.Println(v)
// 			if v%2 == 0 {
// 				result = append(result, v)
// 			}
// 		}
// 		return result
// 	}
// 	var numbers = []int{4, 93, 77, 10, 52, 22, 34}

// 	fmt.Println(evenNumbers(numbers...))
// }

// CLOSURE (IIFE)
//
//func main() {
//
//	var isPalindrome = func(str string) bool {
//		var temp string = ""
//
//		for i := len(str) - 1; i >= 0; i-- {
//			temp += string(byte(str[i]))
//		}
//		return temp == str
//	}("katak")
//	fmt.Println(isPalindrome)
//}

// CLOSURE (CLOSURE AS A RETURN VALUE)
//
//func main() {
//	var studentLists = []string{"luffy", "Deku", "Boruto", "Killua", "Ace"}
//	var find = findStudent(studentLists)
//
//	fmt.Println(find("luffy"))
//}
//
//func findStudent(students []string) func(string) string {
//
//	return func(s string) string {
//		var student string
//		var position int
//
//		for i, v := range students {
//			if strings.ToLower(v) == strings.ToLower(s) {
//				student = v
//				position = i
//				break
//			}
//		}
//		if student == "" {
//			return fmt.Sprintf("%s doesn't exist!!!", s)
//		}
//		return fmt.Sprintf("We found %s position %d", s, position+1)
//	}
//}

//// CLOSURE (CALLBACK)
//type isOddNum func(int) bool
//
//func main() {
//	var numbers = []int{2, 5, 8, 10, 3, 99, 23}
//
//	var find = findOddNumbers(numbers, func(number int) bool {
//		return number%2 != 0
//	})
//	fmt.Println("Total odd numbers:", find)
//}
//
//func findOddNumbers(numbers []int, callback isOddNum) int {
//	var totalOddNumbers int
//	for _, v := range numbers {
//		if callback(v) {
//			totalOddNumbers++
//		}
//	}
//	return totalOddNumbers
//}

// POINTER

func main() {
	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("firstNumber (value) :", firstNumber)
	fmt.Println("firstNumber (memori address) :", &firstNumber)

	fmt.Println("secondNumber (value) :", *secondNumber)
	fmt.Println("secondNumber (memori address) :", secondNumber)
}
