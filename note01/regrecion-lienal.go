package note01

import "gonum.org/v1/gonum/mat"

func calculateLinearRegression(x, y []float64) (float64, float64) {
	// Crear una matriz X con una columna de unos y una columna de valores x
	xMatrix := mat.NewDense(len(x), 2, nil)
	for i := range x {
		xMatrix.Set(i, 0, 1)
		xMatrix.Set(i, 1, x[i])
	}

	// Crear una matriz Y con los valores y
	yMatrix := mat.NewDense(len(y), 1, y)

	// Resolver la ecuación Xb = Y para encontrar los coeficientes b
	var b mat.Dense
	b.Solve(xMatrix, yMatrix)

	// Devolver los coeficientes (intercepto y pendiente)
	return b.At(0, 0), b.At(1, 0)
}
