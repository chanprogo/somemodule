package params

import "math/big"

var Ether = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
var Gwei = new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)
