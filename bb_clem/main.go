package main

import (
	"fmt"
	"github.com/m0dd3r/h2go"
	"log"
)

var (
	user_secret    = "Qhj7YIkvdMrv4n7lKXlbNN5+eeHgJCxXzswgxhSmzeA="
	org_secret     = "6ecdd0f52c429b050300a9d0cfb5b424"
	element_secret = "0A0lN1hzsDFjemr6PfynDsgBogh7auT7Ttb3ZEB6Ysk="
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func main() {
	// Start Session
	s := h2go.NewSession(user_secret, org_secret, element_secret)

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
