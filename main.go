package main 

import "fmt"

type Car struct{
	name string
	model string 
	make string 
	price int 
	color string 
}


func main(){

	var mycar Car 

	mycar.name = "BMW"
	mycar.model = "M4"
	mycar.make = "M450d"
	mycar.price = 12000000
	mycar.color = "Blue"


	fmt.Println(mycar.name)
	fmt.Println(mycar.model)
	fmt.Println(mycar.make)
	fmt.Println(mycar.price)
	fmt.Println(mycar.color)


}