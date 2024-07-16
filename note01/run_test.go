package note01

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(t *testing.T) {

	// Generar 500 muestras
	df := HousePriceSamples(200)

	// Mostrar las primeras filas y un resumen estadístico
	fmt.Println(df.Describe())

	// Guardar como CSV
	f, err := os.Create("data-synthetic.csv")
	if err != nil {
		t.Fatalf("Error creating CSV file: %v", err)
	}
	defer f.Close()

	if err := df.WriteCSV(f); err != nil {
		t.Fatalf("Error writing CSV file: %v", err)
	}

	// Crear y guardar el gráfico
	err = createScatterPlot(df)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Gráfico guardado como 'house_prices_scatter.png'")

}
