package note01

import (
	"image/color"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func createScatterPlot(df dataframe.DataFrame) error {
	p := plot.New()

	p.Title.Text = "Precio de casas vs Número de habitaciones"
	p.X.Label.Text = "Número de habitaciones"
	p.Y.Label.Text = "Precio"

	rooms := df.Col("rooms").Float()
	prices := df.Col("price").Float()
	exceptions, _ := df.Col("is_exception").Bool()

	regularPts := make(plotter.XYs, 0)
	exceptionPts := make(plotter.XYs, 0)

	for i := 0; i < df.Nrow(); i++ {
		pt := plotter.XY{X: rooms[i], Y: prices[i]}
		if exceptions[i] {
			exceptionPts = append(exceptionPts, pt)
		} else {
			regularPts = append(regularPts, pt)
		}
	}

	// Puntos regulares
	s, err := plotter.NewScatter(regularPts)
	if err != nil {
		return err
	}
	s.GlyphStyle.Radius = vg.Points(2)
	s.Color = color.RGBA{B: 255, A: 255} // Azul
	p.Add(s)

	// Puntos de excepción
	e, err := plotter.NewScatter(exceptionPts)
	if err != nil {
		return err
	}
	e.GlyphStyle.Radius = vg.Points(3)
	e.Color = color.RGBA{R: 255, A: 255} // Rojo
	p.Add(e)

	// Calcular y añadir la línea de regresión
	intercept, slope := calculateLinearRegression(rooms, prices)
	lineData := plotter.XYs{
		{X: 3, Y: intercept + slope*3},
		{X: 9, Y: intercept + slope*9},
	}

	line, err := plotter.NewLine(lineData)
	if err != nil {
		return err
	}

	// Añadir la línea de regresión
	line.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Rojo
	line.Width = vg.Points(2)
	p.Add(line)

	// Guardar el gráfico
	return p.Save(10*vg.Inch, 8*vg.Inch, "plot-note-01.png")
}
