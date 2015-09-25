package main

import (
	"fmt"
	"github.com/m0dd3r/cloudelements"
	"log"
)

var (
	user_secret    = ""
	org_secret     = ""
	element_secret = ""
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func main() {
	// Start Session
	s := cloudelements.NewSession(user_secret, org_secret, element_secret)

	customers, err := s.GetEcommerceCustomers("", 1, 25)

	println("")
	if err != nil {
		log.Fatal(err)
	} else {
		for _, c := range customers {
			fmt.Println(*c)
		}
	}
	//	else {
	//		fmt.Println("Bad response status from CloudElements server")
	//		fmt.Printf("\t Status:  %v\n", resp.Status())
	//		fmt.Printf("\t Message: %v\n", errors.Message)
	//	}
	//	fmt.Println("--------------------------------------------------------------------------------")
	println("")

}
