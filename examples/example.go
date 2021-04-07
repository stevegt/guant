package main

import (
	"fmt"

	// . "github.com/stevegt/goadapt"
	"github.com/stevegt/guant"
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	var currPrice float64 = 19.7
	var strike float64 = 22
	t := 13 / 365.0 // years
	fmt.Println("years to expiry:", t)

	// With provided contract price calculate Sigma
	var mark float64 = 0.675
	fmt.Println("mark:", mark)
	x := guant.Derivative{
		N: distuv.Normal{Mu: 0, Sigma: 1},
		S: currPrice,
		K: strike,
		R: 0.01,
		T: t,
	}
	x.Sigma = guant.NewtonRaphson(x, mark)
	fmt.Println("Calculated IV:", x.Sigma)
	y := guant.BlackScholes(x)
	fmt.Println("price W/Calculated IV:", y)

	//Given sigma calculate the value of a Call option
	var sigma float64 = 1.20
	fmt.Println("Given IV", sigma)
	i := guant.Derivative{
		N:     distuv.Normal{Mu: 0, Sigma: 1},
		S:     currPrice,
		K:     strike,
		R:     guant.DefaultRfir(),
		Sigma: sigma,
		T:     t,
	}
	a := guant.BlackScholes(i)
	fmt.Println("price W/Given IV:", a)

}
