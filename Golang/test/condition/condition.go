package main

import (
	"fmt"
)

var myScore int

func main() {
	//myMoney := 100
	//if myMoney >= 100 {
	//fmt.Println("คุณสามารถซื้อกระเป๋าใบนี้ได้")
	//} else {
	//fmt.Println("ขออภัยคุณไม่สามารถซื้อกระเป๋าใบนี้ได้")

	//workshop
	fmt.Println("grade calculator")
	fmt.Scanf("%d", &myScore)
	fmt.Println("myScore=", myScore)
	if myScore >= 80 {
		fmt.Println("you got A grade")
	} else if myScore >= 70 {
		fmt.Println("you got B grade")
	} else if myScore >= 60 {
		fmt.Println("you got C grade")
	} else if myScore >= 50 {
		fmt.Println("you got D grade")
	} else if myScore <= 50 {
		fmt.Println("you got F grade T_T")
	}

}
