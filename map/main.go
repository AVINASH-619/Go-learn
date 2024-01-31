// In the maps, a key must be unique and always in the type which is comparable using == operator or the type which support != operator. So, most of the built-in type can be used as a key like an int, float64, rune, string, comparable array and structure, pointer, etc. The data types like slice and noncomparable arrays and structs or the custom data types which are not comparable don’t use as a map key.
// In maps, the values are not unique like keys and can be of any type like int, float64, rune, string, pointer, reference type, map type, etc.
// The type of keys and type of values must be of the same type, different types of keys and values in the same maps are not allowed. But the type of key and the type values can differ.
// The map is also known as a hash map, hash table, unordered map, dictionary, or associative array.
// In maps, you can only add value when the map is initialized if you try to add value in the uninitialized map, then the compiler will throw an error.


package main

import "fmt"

func main() {
    //Itinitalize the map
	var My_map = make(map[int]string)   //It is used to create an empty intialized map
	// var mp1 map[string][string] It will create an unintialized map and we can't push the key value pair to map
	if My_map == nil {
        fmt.Println("True")
    } else {
     
        fmt.Println("False")
    }

	//Updating and adding values to map
    My_map[1]="Avinash"
	My_map[2]="Vivek"

	//To check weather key exist or not
	value,ok:=My_map[3]
	if ok{
		fmt.Printf("The value exists and it is %s",value)
	} else{
		fmt.Println("The value not exist")
	}

	//Iterating throught the map
	for key,value:=range My_map{
		fmt.Printf("Key: %d Value: %s\n",key,value)
    }
	
	//Deleting the keys  in map
	delete(My_map,1)
    for key,value:=range My_map{
		fmt.Printf("Key: %d Value: %s\n",key,value)
	}

	
}
