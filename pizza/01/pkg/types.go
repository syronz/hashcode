package pkg

import "math/big"

type Ingredient struct {
	Num     *big.Int
	Counter int
}

type Pizza struct {
	Bin    *big.Int //this is binary code
	IngQty int
	Index  int
}
