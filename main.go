package main

import "fmt"

func main() {
	b := NewBox(&Location{42.714620, -73.311469}, &Location{42.048728, -71.070818})
	fmt.Println(b.VirtualCenter)
}
