package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

type MetricLength struct {
	Meters float64
}

type ImperialLength struct {
	Feet float64
}

type LengthAdapter interface {
	GetLength() float64
}

type MetricToImperialAdapter struct {
	MetricLength MetricLength
}

func (m MetricToImperialAdapter) GetLength() float64 {
	return m.MetricLength.Meters * 3.28084
}

func main() {
	metricLength := MetricLength{Meters: 10.0}
	adapter := MetricToImperialAdapter{MetricLength: metricLength}
	lengthInFeet := adapter.GetLength()
	fmt.Printf("Length in feet: %f\n", lengthInFeet)
}
