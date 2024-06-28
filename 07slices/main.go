package main

import (
	"fmt"
	"sort"
)

func main() {	
	fmt.Println("welcome Slices")
	var fruitList=[]string{"Apple","Banana","Tomato","Grapes","Mango"}
	fmt.Printf("type of %T\n",fruitList)
	fruitList= append(fruitList,"Kiwi")
	fmt.Println(fruitList)
	fruitList= append(fruitList[1:])
	fmt.Println(fruitList)
	highScores:=make([]int,4)
	highScores[0]=234
	highScores[1]=945	
	highScores[2]=465
	highScores[3]=867
	highScores=append(highScores, 222,224,333)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
}