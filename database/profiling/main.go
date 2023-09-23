package main

import "fmt"

const (
	GUEST uint = 1 << iota
	ADMIN
	PROFILE1
	PROFILE2
	PROFILE3
)

func main() {

	fmt.Printf("GUEST %d\nADMIN %d\nPROFILE1 %d\nPROFILE2 %d\nPROFILE3 %d\n", GUEST, ADMIN, PROFILE1, PROFILE2, PROFILE3)

	var profile uint = 0

	profile = PROFILE1 | PROFILE2
	fmt.Printf("profile %d\n", profile)

	test := ADMIN & profile
	fmt.Printf("profile %d\n", test)

	test = GUEST & profile
	fmt.Printf("profile %d\n", test)

	test = PROFILE1 & profile
	fmt.Printf("profile %d\n", test)

	test = PROFILE2 & profile
	fmt.Printf("profile %d\n", test)
}
