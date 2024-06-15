package main

import "fmt"

func main() {

	var name1 = "top"
	var name2 string = "ratchav"
	name3 := "kmm"

	fmt.Printf("value of `name1` = [%s]\r\n", name1)
	fmt.Printf("value of `name2` = [%s]\r\n", name2)
	fmt.Printf("value of `name3` = [%s]\r\n\r\n", name3)

	bit := false
	fmt.Printf("value of `bit` = [%t]\r\n\r\n", bit)

	var number int = 5
	fmt.Printf("value of `number` = [%d]\r\n\r\n", number)

	var decimal float32 = 3.14

	fmt.Printf("value of `float32` = [%.3f]\r\n\r\n", decimal)

}
