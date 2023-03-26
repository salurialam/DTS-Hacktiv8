package main

func main() {
	// var currentYear = 2023

	// if age := currentYear - 2000; age < 17 {
	// 	fmt.Println("Kamu belum boleh membuat KTP")
	// } else {
	// 	fmt.Println("Kamu sudah bisa membuat KTP")
	// }

	// var score = 8

	// switch score {
	// case 8:
	// 	fmt.Println("Perfect")
	// case 7:
	// 	fmt.Println("Awesome")
	// default:
	// 	fmt.Println("Not bad")
	// }

	// var score = 6

	// switch {
	// case score == 8:
	// 	fmt.Println("Perfect")
	// case (score < 8) && (score > 3):
	// 	fmt.Println("Not bad")
	// 	fallthrough
	// case score < 5:
	// 	fmt.Println("It's ok, but please study harder")
	// default:
	// 	{
	// 		fmt.Println("Study harder")
	// 		fmt.Println("U need to learn more")
	// 	}
	// }

	// var score = 6

	// if score > 7 {
	// 	switch score {
	// 	case 10:
	// 		fmt.Println("Perfect")
	// 	default:
	// 		fmt.Println("Nice")
	// 	}
	// } else {
	// 	if score >= 5 {
	// 		fmt.Println("Not bad")
	// 	} else if score >= 3 {
	// 		fmt.Println("Keep trying")
	// 	} else {
	// 		fmt.Println("you can do it")
	// 		if score == 0 {
	// 			fmt.Println("Try harder")
	// 		}
	// 	}
	// }

	// LOOPINGS

	// for i := 1; i <= 3; i++ {
	// 	fmt.Println("Angka ", i)
	// }

	// LOOPINGS (SECOND WAY OF LOOPING)

	// var i = 1

	// for i <= 3 {
	// 	fmt.Println("angka ", i)
	// 	i++
	// }

	// LOOPINGS (THIRD WAY OF LOOPING)

	// var i = 0

	// for {
	// 	fmt.Println("Angka ", i)
	// 	i++
	// 	if i == 3 {
	// 		break
	// 	}
	// }

	// LOOPINGS (BREAK AND CONTINUE KEYWORD)

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}
	// 	if i > 8 {
	// 		break
	// 	}
	// 	fmt.Println("Angka ", i)
	// }

	// LOOPINGS (NESTED LOOPING)

	// for i := 0; i < 5; i++ {
	// 	for j := i; j < 5; j++ {
	// 		fmt.Print(j, " ")
	// 	}

	// 	fmt.Println()
	// }

	// LOOPINGS (LABEL)

	// outerLoop:
	// 	for i := 0; i <= 3; i++ {
	// 		fmt.Println("Perulangan ke - ", i+1)
	// 		for j := 0; j <= 3; j++ {
	// 			if i == 3 {
	// 				break outerLoop
	// 			}
	// 			fmt.Print(j, " ")
	// 		}
	// 		fmt.Println()
	// 	}

	// ARRAY

	// var numbers [4]int
	// numbers = [4]int{1, 2, 3, 4}

	// var strings = [3]string{"Airell", "Nanda", "Mailo"}

	// fmt.Printf("%#v\n", numbers)
	// fmt.Printf("%#v\n", strings)

	// ARRAY (MODIFY ELEMENT THROUGH INDEX)

	// var fruits = [3]string{"apel", "pisang", "mangga"}
	// fruits[0] = "apple"
	// fruits[1] = "banana"
	// // fruits[2] = "mango"

	// fmt.Printf("%#v", fruits)

	// ARRAY (LOOP THROUGH ELEMENT)

	// var fruits = [3]string{"apple", "banana", "mango"}
	// //cara pertama
	// for i, v := range fruits {
	// 	fmt.Printf("index: %d, value: %s\n", i, v)
	// }

	// fmt.Println(strings.Repeat("#", 25))

	// //cara kedua
	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("index: %d, value: %s\n", i, fruits[i])
	// }

	// ARRAY (MULTIDIMENSIONAL ARRAY)

	// balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	// for _, arr := range balances {
	// 	for _, value := range arr {
	// 		fmt.Printf("%d ", value)
	// 	}
	// 	fmt.Println()
	// }

	// SLICE

	// var fruits = []string{"apple", "banana", "mango"}

	// _ = fruits

	// fmt.Printf("%#v", fruits)

	// SLICE (APPEND FUNCTION)

	// var fruits = make([]string, 3)
	// _ = fruits

	// fruits[0] = "apple"
	// fruits[1] = "banana"
	// fruits[2] = "mango"

	// fruits = append(fruits, "apple", "banana", "mango")

	// fmt.Printf("%#v", fruits)

	// SLICE (APPEND FUNCTION WITH ELLIPSIS)

	// var fruits1 = []string{"apple", "banana", "mango"}

	// var fruits2 = []string{"durian", "pineapple", "starfruit"}

	// fruits1 = append(fruits1, fruits2...)

	// fmt.Printf("%#v", fruits1)

	// SLICE (COPY FUNCTION)

	// Contoh 1

	// var fruits1 = []string{"apple", "banana", "mango"}
	// var fruits2 = []string{"durian", "pineapple", "starfruit"}

	// fruits1 = append(fruits1, fruits2...)
	// fmt.Printf("%#v", fruits1)

	// Contoh 2

	// var fruits1 = []string{"apple", "banana", "mango"}
	// var fruits2 = []string{"durian", "pineapple", "starfruit"}

	// nn := copy(fruits1, fruits2)

	// fmt.Println("Fruits1 =>", fruits1)
	// fmt.Println("Fruits2 =>", fruits2)
	// fmt.Println("Copied elements =>", nn)

	// SLICE (SLICING)

	// var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	// var fruits2 = fruits1[1:4]
	// fmt.Printf("%#v\n", fruits2)

	// var fruits3 = fruits1[0:]
	// fmt.Printf("%#v\n", fruits3)

	// var fruits4 = fruits1[:3]
	// fmt.Printf("%#v\n", fruits4)

	// var fruits5 = fruits1[:] //sama dengan fruits[:len(fruits)]
	// fmt.Printf("%#v\n", fruits5)

	// SLICING (COMBINING SLICING AND APPEND)

	// var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	// fruits1 = append(fruits1[:3], "rambutan")

	// fmt.Printf("%#v\n", fruits1)

	// SLICE (BACKING ARRAY)

	// var fruits1 = []string{"apple", "mango", "durian", "banana", "starfruit"}

	// var fruits2 = fruits1[2:4]

	// fruits2[0] = "rambutan"

	// fmt.Println("fruits1 => ", fruits1)
	// fmt.Println("fruits2 => ", fruits2)

	// SLICE (CAP FUNCTION)

	// var fruits1 = []string{"apple", "mango", "durian", "banana"}

	// fmt.Println("Fruits1 cap: ", cap(fruits1))
	// fmt.Println("Fruits1 len: ", len(fruits1))

	// fmt.Println(strings.Repeat("#", 20))

	// var fruits2 = fruits1[0:3]

	// fmt.Println("Fruits2 cap: ", cap(fruits2))
	// fmt.Println("Fruits2 len: ", len(fruits2))

	// fmt.Println(strings.Repeat("#", 20))

	// var fruits3 = fruits1[1:]

	// fmt.Println("Fruits cap: ", cap(fruits3))
	// fmt.Println("Fruits len: ", len(fruits3))

	// SLICE (CREATING NEW BACKING ARRAY)

	// cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	// newCars := []string{}

	// newCars = append(newCars, cars[0:2]...)

	// cars[0] = "Nissan"
	// fmt.Println("cars:", cars)
	// fmt.Println("newCars:", newCars)

	// MAP
	//
	//var person = map[string]string{
	//	"name":    "Airell",
	//	"age":     "23",
	//	"address": "Jakarta",
	//}
	//
	//fmt.Println("name:", person["name"])
	//fmt.Println("age:", person["age"])
	//fmt.Println("address:", person["address"])

	// MAP (LOOPING WITH MAP)

	//var person = map[string]string{
	//	"name":    "Airell",
	//	"age":     "23",
	//	"address": "Jakarta",
	//}
	//for key, value := range person {
	//	fmt.Println(key, ":", value)
	//}

	// MAP (DELETING VALUE)

	// var person = map[string]string{
	// 	"name":    "Airell",
	// 	"age":     "23",
	// 	"address": "Jakarta",
	// }

	// fmt.Println("Before deleting:", person)

	// delete(person, "age")
	// fmt.Println("After deleting:", person)

	// MAP (DETECTING VALUE)

	// var person = map[string]string{
	// 	"name":    "Airell",
	// 	"age":     "23",
	// 	"address": "Jakarta",
	// }

	// value, exist := person["age"]

	// if exist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Value doesn't exist")
	// }

	// delete(person, "age")

	// value, exist = person["age"]

	// if exist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Value has been deleted")
	// }

	// MAP (COMBINING SLICE WITH MAP)

	//var people = []map[string]string{
	//	{"name": "Airell", "age": "23"},
	//	{"name": "Nanda", "age": "23"},
	//	{"name": "Mailo", "age": "15"},
	//}
	//
	//for i, person := range people {
	//	fmt.Printf("index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	//}

}
