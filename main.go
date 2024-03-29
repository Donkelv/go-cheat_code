package main

import (
	"errors"
	"fmt"
	"github.com/pluralsight/webservice/models"
)

const (
	first  = 1
	second = "second"
)

const (
	one = iota
	two = iota
)

func main() {

	userStruct := models.User{
		ID:        2,
		FirstName: "Kevin",
		LastName:  "Gonzalez"}

	fmt.Println(userStruct)

	port := 3000
	numberOfRetries := 3

	goCleanSheet()
	p, err := startwebserver(port, numberOfRetries)

	///if i need only one of the values returned from a function
	///I can use the underscore to ignore the other value
	///p, _ := startwebserver(port, numberOfRetries)

	fmt.Println(err)
	fmt.Println(p)

}

func startwebserver(port int, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")
	//do important things
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)

	return port, errors.New("something went wrong")
}

func goCleanSheet() {
	/// Declaring variables with Primitive Data types
	var i = 64
	fmt.Println(i)

	var f float32 = 3.24
	fmt.Println(f)

	firstName := "Kevin"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	fmt.Println("hey Gophers from a module")

	///Declaring pointers
	///Pointers hold the address in the memory for a particular var

	var pointer1 *string = new(string)
	*pointer1 = "Kevin"
	fmt.Println(pointer1)

	//Declare a pointers
	ptr := &firstName
	fmt.Println(ptr, *ptr)

	///Declaring constants
	///You can't assign a new value to a constant
	///So you need to initialize them at the time that you are initializing them
	const pi = 3.142
	fmt.Println(pi)

	///Print declared multiple constants
	fmt.Println(first, second)

	///Print iota declared constants
	fmt.Println(one, two)

	///Declaring Arrays
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr2 := [3]int{2, 4, 3}
	fmt.Println(arr2)

	///Slice pointing to the underlying array.. i think its a more dynamic way of working with an array
	slice := arr[:]
	arr[1] = 45
	slice[2] = 74

	fmt.Println(slice)

	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)

	//slice is not a fixed size of array

	slice2 = append(slice2, 67, 56, 123)
	fmt.Println(slice2)

	sc1 := slice2[1:]
	sc2 := slice2[1:3]
	sc3 := slice2[:2]
	fmt.Println(sc1, sc2, sc3)

	///Declring map data

	m := map[string]int{"Kevin": 100}
	fmt.Println(m)
	fmt.Println(m["Kevin"])

	//Reassign a value
	m["Kevin"] = 200
	fmt.Println(m["Kevin"])
	delete(m, "Kevin")
	fmt.Println(m["Kevin"])

	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	//initialize a struct
	var u user
	u.ID = 1
	u.FirstName = "Kevin"
	u.LastName = "Gonzalez"
	fmt.Println(u)

	u2 := user{
		ID:        2,
		FirstName: "Kevin",
		LastName:  "Gonzalez"}

	fmt.Println(u2)
}
