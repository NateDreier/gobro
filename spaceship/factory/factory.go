package main

import (
	"code/gobro/spaceship/factory/appliances"
	"fmt"
)

func main() {
	fmt.Println("enter 0 or 1")

	var myType int
	fmt.Scan(&myType)

	myAppliance, err := appliances.CreateAppliance(myType)
	if err == nil {
		myAppliance.Start()
		fmt.Println(myAppliance.GetPurpose())
	} else {
		fmt.Println(err)
	}
}
