package main

import "fmt"

func main() {
	var name, region string = "NOVERIKO", "Jakarta"
	var age, class int = 21, 7
	i := 21
	j := true
	var a int = 21
	var b int = 15
	var k float64 = 123.456
	var x int = 1022

	fmt.Println("This is a student profile page")
	fmt.Println("Name : ", name)
	fmt.Println("Age : ", age)
	fmt.Println("Region : ", region)
	fmt.Println("Golang Class : ", class, "\n")

	//variabel i
	fmt.Printf("%d \n", i)
	//tipe var i
	fmt.Printf("%T \n", i)
	//karakter %
	fmt.Println("%")
	//nilai bool j
	fmt.Printf("%t \n", j)

	//nilai unicode rus
	fmt.Printf("%c \n", b)
	//unknown output
	fmt.Println("10101")
	//unknown output
	fmt.Println("?")

	//nilai base 10
	fmt.Printf("%d \n", a)
	//nilai base 8
	fmt.Printf("%o \n", a)
	//nilai base 16 lower
	fmt.Printf("%x \n", b)
	//nilai base 16 upper
	fmt.Printf("%X \n", b)
	//nilai unicode kar U+042F
	fmt.Printf("%#U \n", x)

	//nilai float desimal full
	fmt.Printf("%f \n", k)
	//nilai float desimal scientific
	fmt.Printf("%e \n", k)

	fmt.Println("\u042F", "(ya)")

}
