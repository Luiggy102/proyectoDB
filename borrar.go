package main

import (
	"db1final/modelos"
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func BorrarProvincia(db db) {
	provincias, _ := db.MostrarProvincias()
	provinciaSeleccionada, err := seleccionarProvincia(provincias)
	if err != nil {
		panic(err)
	}
	var borrar bool
	g2 := huh.NewGroup(
		huh.NewConfirm().Title(
			fmt.Sprintf("Desea borrar: %s", provinciaSeleccionada.Nombre),
		).Value(&borrar).Affirmative("Sí").Negative("No"),
	)
	f := huh.NewForm(g2).WithProgramOptions(tea.WithAltScreen())
	err = f.Run()
	if err != nil {
		panic(err)
	}
	// menu de borrado
	if borrar {
		err = db.EliminarPorId("provincia", strconv.Itoa(provinciaSeleccionada.Id))
		if err != nil {
			panic(err)
		}
	}
}
func BorrarCanton(db db) {
	var cantonSeleccionado modelos.Canton
	cantones, _ := db.MostrarCatones()

	// que seleccione la provincia
	// llenar las opciones
	var opciones []huh.Option[modelos.Canton]

	for _, c := range cantones {
		var opcion huh.Option[modelos.Canton] = huh.Option[modelos.Canton]{
			Key:   c.Nombre,
			Value: *c,
		}
		opciones = append(opciones, opcion)
	}

	// sacarle el id
	g1 := huh.NewGroup(
		huh.NewSelect[modelos.Canton]().Value(&cantonSeleccionado).
			Options(opciones...),
	)
	f := huh.NewForm(g1).WithProgramOptions(tea.WithAltScreen())
	err := f.Run()
	if err != nil {
		panic(err)
	}

	var borrar bool
	g2 := huh.NewGroup(
		huh.NewConfirm().Title(
			fmt.Sprintf("Desea borrar: %s", cantonSeleccionado.Nombre),
		).Value(&borrar).Affirmative("Sí").Negative("No"),
	)
	f = huh.NewForm(g2).WithProgramOptions(tea.WithAltScreen())
	err = f.Run()
	if err != nil {
		panic(err)
	}

	// menu de borrado
	if borrar {
		err = db.EliminarPorId("canton", strconv.Itoa(cantonSeleccionado.Id))
		if err != nil {
			panic(err)
		}
	}

}
func BorrarParroquia(db db) {
	var parroquiaSeleccionada modelos.Parroquia
	parroquias, _ := db.MostrarParroquias()

	// que seleccione la provincia
	// llenar las opciones
	var opciones []huh.Option[modelos.Parroquia]

	for _, p := range parroquias {
		var opcion huh.Option[modelos.Parroquia] = huh.Option[modelos.Parroquia]{
			Key:   p.Nombre,
			Value: *p,
		}
		opciones = append(opciones, opcion)
	}

	// sacarle el id
	g1 := huh.NewGroup(
		huh.NewSelect[modelos.Parroquia]().Value(&parroquiaSeleccionada).
			Options(opciones...),
	)
	f := huh.NewForm(g1).WithProgramOptions(tea.WithAltScreen())
	err := f.Run()
	if err != nil {
		panic(err)
	}

	var borrar bool
	g2 := huh.NewGroup(
		huh.NewConfirm().Title(
			fmt.Sprintf("Desea borrar: %s", parroquiaSeleccionada.Nombre),
		).Value(&borrar).Affirmative("Sí").Negative("No"),
	)
	f = huh.NewForm(g2).WithProgramOptions(tea.WithAltScreen())
	err = f.Run()
	if err != nil {
		panic(err)
	}

	// menu de borrado
	if borrar {
		err = db.EliminarPorId("parroquia", strconv.Itoa(parroquiaSeleccionada.Id))
		if err != nil {
			panic(err)
		}
	}

}
