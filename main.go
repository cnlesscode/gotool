package main

import (
	"fmt"

	"github.com/cnlesscode/gotool/datetime"
)

func main() {
	fmt.Printf("%v\n", datetime.TimeStampToDatatimeSlice(1658997290))
	// [2022 07 28 16 34 50]
	fmt.Printf("%v\n", datetime.TimeStampToDatatimeSlice(-1))
	// [2023 01 05 12 15 58]
}
