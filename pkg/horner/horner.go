package horner

import "math/big"

// PolyEval returns the evaluation of the polynomial given its coefficients, x value and a prime p.
// coefficients[0] is the coefficient of x^0, coefficients[1] is the coefficient of x^1, and so on.
// Example:
//
//	fmt.Println(PolyEval(
//		[]*big.Int{big.NewInt(12), big.NewInt(11), big.NewInt(2)},
//		big.NewInt(2),
//		big.NewInt(79),
//	)) // 42
func PolyEval(coefficients []*big.Int, x *big.Int, p *big.Int) *big.Int {
	poly := big.NewInt(coefficients[len(coefficients)-1].Int64())
	for i := 1; i < len(coefficients); i++ {
		poly.Mul(poly, x).
			Add(poly, coefficients[(len(coefficients)-1)-i]).
			Mod(poly, p)
	}
	return poly
}
