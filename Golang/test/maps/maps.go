package main

import "fmt"

var product = make(map[string]float32)

func main() {
	fmt.Println("product =", product)
	//aad
	product["Macbook"] = 40000
	product["Iphone"] = 30000
	product["Ipad"] = 20000
	fmt.Println("product=", product)
	//delete
	delete(product, "Ipad")
	fmt.Println("product=", product)
	//update
	product["Iphone"] = 20000
	fmt.Println("product=", product)

	value1 := product["Iphone"]
	fmt.Println("value1=", value1)

	courseName := map[string]string{"101": "Java", "102": "python"}
	fmt.Println("courseName =", courseName)
}
