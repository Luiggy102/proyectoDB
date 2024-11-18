package main

import (
	"db1final/modelos"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func seleccionarProvincia(provincias []*modelos.Provincia) (modelos.Provincia, error) {
	var provinciaSeleccionada modelos.Provincia
	var opciones []huh.Option[modelos.Provincia]
	for _, p := range provincias {
		var opcion huh.Option[modelos.Provincia] = huh.Option[modelos.Provincia]{
			Key:   p.Nombre,
			Value: *p,
		}
		opciones = append(opciones, opcion)
	}
	g1 := huh.NewGroup(huh.NewSelect[modelos.Provincia]().Options(opciones...).Value(&provinciaSeleccionada).
		Title("Selecciona la provincia"),
	)
	f := huh.NewForm(g1).WithProgramOptions(tea.WithAltScreen())
	err := f.Run()
	if err != nil {
		return modelos.Provincia{}, err
	}
	return provinciaSeleccionada, nil
}

func seleccionarCanton(cantones []*modelos.Canton) (modelos.Canton, error) {
	var cantonSeleccionado modelos.Canton
	var opciones []huh.Option[modelos.Canton]
	for _, c := range cantones {
		var opcion huh.Option[modelos.Canton] = huh.Option[modelos.Canton]{
			Key:   c.Nombre,
			Value: *c,
		}
		opciones = append(opciones, opcion)
	}
	g1 := huh.NewGroup(
		huh.NewSelect[modelos.Canton]().Value(&cantonSeleccionado).
			Options(opciones...),
	)
	f := huh.NewForm(g1).WithProgramOptions(tea.WithAltScreen())
	err := f.Run()
	if err != nil {
		return modelos.Canton{}, err
	}
	return cantonSeleccionado, nil
}
func seleccionarParroquia(parroquias []*modelos.Parroquia) (modelos.Parroquia, error) {
	var parroquiaSeleccionada modelos.Parroquia
	var opciones []huh.Option[modelos.Parroquia]
	for _, p := range parroquias {
		var opcion huh.Option[modelos.Parroquia] = huh.Option[modelos.Parroquia]{
			Key:   p.Nombre,
			Value: *p,
		}
		opciones = append(opciones, opcion)
	}
	g1 := huh.NewGroup(
		huh.NewSelect[modelos.Parroquia]().Value(&parroquiaSeleccionada).
			Options(opciones...),
	)
	f := huh.NewForm(g1).WithProgramOptions(tea.WithAltScreen())
	err := f.Run()
	if err != nil {
		panic(err)
	}
	return parroquiaSeleccionada, nil
}
