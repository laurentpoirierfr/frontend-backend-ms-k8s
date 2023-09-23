package main

import "fmt"

const (
	GUEST uint64 = 1 << iota
	ADMIN
	PROFILE1
	PROFILE2
	PROFILE3
)

func main() {

	fmt.Printf("GUEST %d\nADMIN %d\nPROFILE1 %d\nPROFILE2 %d\nPROFILE3 %d\n", GUEST, ADMIN, PROFILE1, PROFILE2, PROFILE3)

	var profiles uint64 = 0

	profiles = PROFILE1 | PROFILE2
	fmt.Printf("profile %d\n", profiles)

	test := ADMIN & profiles
	fmt.Printf("Test %d\n", test)

	test = GUEST & profiles
	fmt.Printf("Test %d\n", test)

	test = PROFILE1 & profiles
	fmt.Printf("Test %d\n", test)

	test = PROFILE2 & profiles
	fmt.Printf("Test %d\n", test)
}
