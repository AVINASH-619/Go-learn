package main

import (
	"fmt"
)

func main(){

	//Declartion of array=> Var array_name[length]Type
	//In an array, if an array does not initialize explicitly, then the default value of this array is 0.
	//We can't change the defalut value in go 
	var arr [10]int
	fmt.Printf("%T\n",arr)
	//Intialize the array=>array_name:= [length]Type{item1, item2, item3,...itemN}
   // brr:=[5]int{1}

   //Iterating through array using len
    fmt.Println(len(crr))
	for i:=0;i<len(crr);i++{
		//fmt.Println(crr[i])
	}
	//Case 2
	for _,val:=range arr{
		val++
		//fmt.Printf("%d %d\n",idx,value)
	}

    
}