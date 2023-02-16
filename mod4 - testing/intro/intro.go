package main

import (
	"fmt"
	"intro/addresses"
)

func main() {
	addressType := addresses.AddressType("Street banana")
	fmt.Println(addressType)

	addressType2 := addresses.AddressType("XXStreet banana")
	fmt.Println(addressType2)
}
