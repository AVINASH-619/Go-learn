package main

import "fmt"

//Definig the structures
type Address struct {
	name string 
	street string
	city string
	state string
	Pincode int
}

func anonomousStruct(){
	Element := struct { 
        name      string 
        branch    string 
        language  string 
        Particles int
    }{ 
        name:      "Pikachu", 
        branch:    "ECE", 
        language:  "C++", 
        Particles: 498, 
    } 
  
    // Display the anonymous structure 
    fmt.Println(Element) 
}


func main(){
// Assign values to the obj of the structures
var obj = Address{name:"Akshay", street:"PremNagar", state:"Uttarakhand", Pincode:252636}
var ptr=&obj
fmt.Printf("%s\n",(*ptr).name)
fmt.Printf("%s\n",obj.name)
anonomousStruct()
}