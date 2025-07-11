package main

import "fmt"

type grade struct {
	name        string
	grades      map[string]float64
	no_subjects int
}

func newgrade(name string, number int) grade {
	g := grade{
		name:        name,
		grades:      map[string]float64{},
		no_subjects: number,
	}
	return g
}

func (g *grade) addItems(name string, gp float64) {
	g.grades[name] = gp

}

func gradeAverage(total float64, subjects int) float64 {
	return (total / float64(subjects))

}

func (g *grade) format() string {
	fs := g.name + "'s " + "Grade Report:\n"
	total := 0.0

	for k, v := range g.grades {
		fs += fmt.Sprintf("%-25v...%v\n", k+":", v)
		total += v
	}
	av := gradeAverage(total, g.no_subjects)
	fs += fmt.Sprintf("%-25v...%0.2f\n", "Average:", av)
	return fs
}
