package main

import (
	"fmt"

	"github.com/KalleDK/go-money/money"
)

func main() {

	f := money.Formatter{
		Decimal:  ',',
		Thousand: '.',
	}

	n := int64(0)
	for i := int64(1); i < 19; i++ {
		n = n*10 + (i % 10)
		a := money.FromCenti(n)
		b := money.FromCenti(-n)
		fmt.Printf("%27s\n", f.Sprint(a))
		fmt.Printf("%27s\n", f.Sprint(b))
	}
}
