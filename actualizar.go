package main

import (
	"db1final/modelos"
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func ActualizarProvincia(db db) {
	var provinciaSeleccionada modelos.Provincia
	provincias, _ := db.MostrarProvincias()
	var opciones []huh.Option[modelos.Provincia]
	for _, p := range provincias {
		var opcion huh.Option[modelos.Provincia] = huh.Option[modelos.Provincia]{
			Key:   p.Nombre,
			Value: *p,
		}
		opciones = append(opciones, opcion)
	}
	// seleccionar la provincia
	g1 := huh.NewGroup(huh.NewSelect[modelos.Provincia]().Options(opciones...).Value(&provinciaSeleccionada).
		Title("Selecciona la provincia"),
	) // cambiarle el nombre

	var nuevoNombre string
	g2 := huh.NewGroup(huh.NewInput().Value(&nuevoNombre).Title("Escriba el nuevo nombre"))

	f := huh.NewForm(g1, g2).WithProgramOptions(tea.WithAltScreen())
	err := f.Run()
	if err != nil {
		panic(err)
	}

	err = db.ActualizarProvincia(provinciaSeleccionada.Id, nuevoNombre)
	if err != nil {
		panic(err)
	}

}
func ActualizarCaton(db db) {
	cantones, _ := db.MostrarCatones()
	provincias, _ := db.MostrarProvincias()
	// seleccionar el canton
	cantonParaActualizar, _ := seleccionarCanton(cantones)
	// cambiarle el nombre
	var nuevoNombre string
	g1 := huh.NewGroup(
		huh.NewInput().Value(&nuevoNombre).Title("Ingrese Nuevo nombre"),
	)
	f := huh.NewForm(g1).WithProgramOptions(tea.WithAltScreen())
	err := f.Run()
	if err != nil {
		panic(err)
	}
	// colocar a que provincia debe pertenercer
	fmt.Println("Ingrese nueva provincia para el canton")
	provincia, _ := seleccionarProvincia(provincias)
	// actualizar en db
	err = db.ActualizarCantonParroquia("canton", cantonParaActualizar.Id, strconv.Itoa(provincia.Id), nuevoNombre)
	if err != nil {
		panic(err)
	}

}
func ActualizarParroquia(db db) {
	parroquias, _ := db.MostrarParroquias()
	cantones, _ := db.MostrarCatones()
	var nuevo string
	var seleccionParroquia modelos.Parroquia
	var seleccionCanton modelos.Canton
	// seleccionar la parroquia
	seleccionParroquia, _ = seleccionarParroquia(parroquias)
	// cambiarle el nombre
	f := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Value(&nuevo).Title("Coloque nuevo nombre para parroquia"),
		),
	).WithProgramOptions(tea.WithAltScreen())
	err := f.Run()
	if err != nil {
		panic(err)
	}
	// colocar a que canton debe pertenercer
	seleccionCanton, _ = seleccionarCanton(cantones)
	// cambiar en db
	err = db.ActualizarCantonParroquia("parroquia", seleccionParroquia.Id, strconv.Itoa(seleccionCanton.Id), nuevo)
	if err != nil {
		panic(err)
	}

}
