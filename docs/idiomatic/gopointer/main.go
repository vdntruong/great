package main

import "fmt"

/**

 */

func main() {
	var v int
	v = 1
	fmt.Printf("variable v contains value = %v\n", v)
	fmt.Printf("variable v held by a memory has address = %v\n", &v)

	var p *int
	p = &v

	fmt.Printf("variable p contains value = %v\n", p)
	fmt.Printf("variable p held by a memory has address = %v (it's v's address)\n", &p)
	fmt.Printf("value of the pointer value that p helding is = %v (it's v's value)\n", *p)

	v = 2
	fmt.Printf("variable v contains value = %v\n", v)
	fmt.Printf("variable p contains value = %v\n", p)
	fmt.Printf("value of the pointer value that p helding is = %v\n", *p)

	*p = 3 // ~ we change v to 3
	fmt.Printf("variable v contains value = %v\n", v)
	fmt.Printf("variable p contains value = %v\n", p)
	fmt.Printf("value of the pointer value that p helding is = %v\n", *p)
}
