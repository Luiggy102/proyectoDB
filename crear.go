package main

import (
	"db1final/modelos"
	"fmt"
	"slices"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func CrearProvincia(db db) {
	// saber la logitud actual de las provincias para usarlo como el nuevo id
	provincias, _ := db.MostrarProvincias()
	// idNuevo := provincias[len(provincias)-1].Id + 1
	// generar id nuevo
	ids := []int{}
	for _, p := range provincias {
		ids = append(ids, p.Id)
	}
	idNuevo := slices.Max(ids) + 1
	var nuevoNombre string

	// crear formulario pidiendo el nuevo nombre
	g1 := huh.NewGroup(
		huh.NewInput().Value(&nuevoNombre).Title("Ingrese nueva provincia"),
	)
	f := huh.NewForm(g1).WithProgramOptions(tea.WithAltScreen())
	err := f.Run()
	if err != nil {
		panic(err)
	}

	var crear bool
	g2 := huh.NewGroup(
		huh.NewConfirm().Title(
			fmt.Sprintf("Desea crear: %s", nuevoNombre),
		).Value(&crear).Affirmative("Sí").Negative("No"),
	)
	f = huh.NewForm(g2).WithProgramOptions(tea.WithAltScreen())
	err = f.Run()
	if err != nil {
		panic(err)
	}

	if crear {
		err = db.CrearProvincia(idNuevo, nuevoNombre)
		if err != nil {
			panic(err)
		}
	}

}
func CrearCanton(db db) {
	cantones, _ := db.MostrarCatones()
	// idNuevo := cantones[len(cantones)-1].Id + 1
	// generar id nuevo
	ids := []int{}
	for _, c := range cantones {
		ids = append(ids, c.Id)
	}
	idNuevo := slices.Max(ids) + 1
	var nuevoNombre string

	// que seleccione la provincia
	// llenar las opciones
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

	// nombre y a que provincia pertenerce(id)
	g1 := huh.NewGroup(
		huh.NewInput().Value(&nuevoNombre).Title("Creación de cantón"),
		huh.NewSelect[modelos.Provincia]().Value(&provinciaSeleccionada).
			Options(opciones...).Title("A que provincia pertenerce"),
	)

	f := huh.NewForm(g1).WithProgramOptions(tea.WithAltScreen())
	err := f.Run()
	if err != nil {
		panic(err)
	}

	var crear bool
	g2 := huh.NewGroup(
		huh.NewConfirm().Title(
			fmt.Sprintf("Desea crear: %s", nuevoNombre),
		).Value(&crear).Affirmative("Sí").Negative("No"),
	)
	f = huh.NewForm(g2).WithProgramOptions(tea.WithAltScreen())
	err = f.Run()
	if err != nil {
		panic(err)
	}
	if crear {
		err = db.CrearCanton(idNuevo, nuevoNombre, strconv.Itoa(provinciaSeleccionada.Id))
		if err != nil {
			panic(err)
		}
	}

}
func CrearParroquia(db db) {
	parroquias, _ := db.MostrarParroquias()
	// idNuevo := parroquias[len(parroquias)-1].Id + 1
	// generar id nuevo
	ids := []int{}
	for _, p := range parroquias {
		ids = append(ids, p.Id)
	}
	idNuevo := slices.Max(ids) + 1
	var nuevoNombre string

	var canton modelos.Canton
	cantones, _ := db.MostrarCatones()
	var opciones []huh.Option[modelos.Canton]
	for _, c := range cantones {
		var opcion huh.Option[modelos.Canton] = huh.Option[modelos.Canton]{
			Key:   c.Nombre,
			Value: *c,
		}
		opciones = append(opciones, opcion)
	}

	// nombre y a que provincia pertenerce(id)
	g1 := huh.NewGroup(
		huh.NewInput().Value(&nuevoNombre).Title("Creación de parroquia"),
		huh.NewSelect[modelos.Canton]().Value(&canton).
			Options(opciones...).Title("A que canton pertenerce"),
	)

	f := huh.NewForm(g1).WithProgramOptions(tea.WithAltScreen())
	err := f.Run()
	if err != nil {
		panic(err)
	}

	var crear bool
	g2 := huh.NewGroup(
		huh.NewConfirm().Title(
			fmt.Sprintf("Desea crear: %s", nuevoNombre),
		).Value(&crear).Affirmative("Sí").Negative("No"),
	)
	f = huh.NewForm(g2).WithProgramOptions(tea.WithAltScreen())
	err = f.Run()
	if err != nil {
		panic(err)
	}

	if crear {
		err = db.CrearParroquia(idNuevo, nuevoNombre, strconv.Itoa(canton.Id))
		if err != nil {
			panic(err)
		}
	}
}
