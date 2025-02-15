package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"java,phyton"}
	fmt.Println(courseName)
	courseName = append(courseName, "C", "C#", "HTML", "CSS", "JavaScript")
	fmt.Println(courseName)

	courseweb := courseName[:4]
	fmt.Println(courseweb)
}
