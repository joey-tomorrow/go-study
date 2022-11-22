package main_test

import (
	"fmt"
	"testing"
)

func Test_Slice(t *testing.T) {
	//address := make([]int, 10)
	var address []int
	address = append(address, 1)
	address = append(address, 2)
	//address = append(address, 1)

	fmt.Println(address[2:])

	var StreetComplement = "123456789012345678901"
	fmt.Println(StreetComplement[:20])
	fmt.Println(fmt.Sprintf("%-20v", StreetComplement))

}
