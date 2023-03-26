package main

import "fmt"

func main() {

	var i = 21
	var k float64 = 123.456

	fmt.Println("Menampilkan nilai i : ")
	fmt.Println(i)
	fmt.Println()

	fmt.Println("Menampilkan tipe data dari variable i")
	fmt.Printf("Tipe data variable i adalah %T \n", i)
	fmt.Println()

	fmt.Println("Menampilkan tanda %")
	fmt.Printf("%% \n")
	fmt.Println()

	fmt.Println("Menampilkan nilai boolean j = ")
	fmt.Printf("Nilai boolean j = %t \n", true)
	fmt.Println()

	fmt.Println("Menampilkan nilai boolean j = ")
	fmt.Printf("Nilai boolean j = %t \n", false)
	fmt.Println()

	fmt.Println("menampilkan unicode russia : ")
	fmt.Printf("%c \n", '\u042F')
	fmt.Println()

	fmt.Println("Menampilkan nilai base 10 :")
	fmt.Printf("%d \n", i)
	fmt.Println()

	fmt.Println("Menampilkan nilai base 8 :")
	fmt.Printf("%o \n", i)
	fmt.Println()

	fmt.Println("menampilkan nilai base 16 : ")
	fmt.Printf("%x \n", 15)
	fmt.Println()

	fmt.Println("menampilkan nilai base 16 : ")
	fmt.Printf("%X \n", 15)
	fmt.Println()

	fmt.Println("menampilkan unicode  karakter Я : ")
	fmt.Printf("%U \n", '\u042F')
	fmt.Println()

	fmt.Println("Menampilkan float : ")
	fmt.Printf("%.6f \n", k)
	fmt.Println()

	fmt.Println("Menampilkan float : ")
	fmt.Printf("%e \n", k)
}
