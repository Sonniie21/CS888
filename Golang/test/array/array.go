package main

import "fmt"

var productName [4]string
var price [4]float64

func main() {
	productName[0] = "Macbook"
	productName[1] = "Ipad"
	productName[2] = "Airpods"

	price := [3]float32{40000, 20000, 3000}
	fmt.Println(price)
}
