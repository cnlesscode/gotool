package main

import (
	"fmt"

	"github.com/cnlesscode/gotool/paginator"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	pager := paginator.Run(98, 1, 100, 10)
	fmt.Printf("CurrentPage: %v\n", pager.CurrentPage)
	fmt.Printf("FirstPage: %v\n", pager.FirstPage)
	fmt.Printf("PrePage: %v\n", pager.PrePage)
	fmt.Printf("Pages: %v\n", pager.Pages)
	fmt.Printf("NextPage: %v\n", pager.NextPage)
	fmt.Printf("LastPage: %v\n", pager.LastPage)
	fmt.Printf("TotalPages: %v\n", pager.TotalPages)
}
