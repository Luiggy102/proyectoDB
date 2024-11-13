package main

import (
	"db1final/databases/postgres"
	"db1final/modelos"
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

// presentación y selección de db
func Intro() *string {
	var db string
	presentación := huh.NewNote().
		Title("Proyecto final de Bases de Datos 1").
		Description("Integrantes:\n\t*Luiggy Tamayo*\n\t*Foy Barba*\n\t*Andrés Vallejo*")
	selecciónDb := huh.NewSelect[string]().
		Options(huh.NewOptions("Postgres", "Mariadb", "sqlserver")...).
		Value(&db).
		Title("Elija la base de datos a usar:")
	grupo := huh.NewGroup(
		presentación,
		selecciónDb,
	)
	formulario := huh.NewForm(grupo)
	if err := formulario.Run(); err != nil {
		log.Fatal(err)
	}
	return &db
}

func InicarApp(db string) {
	if db == "Postgres" {
		opciónPostgres()
	}
	if db == "Mariadb" {

	}
	if db == "sqlserver" {

	}
}

func opciónPostgres() {
	db, err := postgres.NuevaDbPostgres(postgres.Url)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var opcionCrud string
	var opcionTerritorio string
	var salir bool

	for !salir {
		// seleccionar operación crud
		g1 := huh.NewGroup(
			huh.NewSelect[string]().Value(&opcionCrud).
				Options(
					huh.NewOption("crear", "crear"),
					huh.NewOption("mostrar", "mostrar"),
					huh.NewOption("actualizar", "actualizar"),
					huh.NewOption("borrar", "borrar"),
					huh.NewOption("salir", "salir"),
				),
		)
		f := huh.NewForm(g1)
		f.Run()

		if opcionCrud == "salir" {
			return
		}

		// preguntar por
		g2 := huh.NewGroup(
			huh.NewSelect[string]().Value(&opcionTerritorio).
				Options(
					huh.NewOption("provincia", "provincia"),
					huh.NewOption("canton", "canton"),
					huh.NewOption("parroquia", "parroquia"),
				),
		)
		f = huh.NewForm(g2)
		f.Run()

		// casos
		if opcionCrud == "mostrar" && opcionTerritorio == "provincia" {
			t := table.New().Headers("ID", "Provincia").BorderStyle(
				lipgloss.NewStyle().Foreground(lipgloss.Color("86")),
			)
			provincias, _ := db.MostrarProvincias()
			for _, p := range provincias {
				t.Row(strconv.Itoa(p.Id), p.Nombre)
			}
			fmt.Println(t)
			fmt.Println("Presione Enter para continuar")
			fmt.Scanln()
		}
		if opcionCrud == "mostrar" && opcionTerritorio == "canton" {
			t := table.New().Headers("ID", "canton", "provincia").BorderStyle(
				lipgloss.NewStyle().Foreground(lipgloss.Color("86")),
			)
			provincias, _ := db.MostrarProvincias()
			cantones, _ := db.MostrarCatones()
			for _, c := range cantones {
				var provincia string
				for _, p := range provincias {
					if c.IdProvincia == strconv.Itoa(p.Id) {
						provincia = p.Nombre
					}
				}
				t.Row(strconv.Itoa(c.Id), c.Nombre, provincia)
			}
			fmt.Println(t)
			fmt.Println("Presione Enter para continuar")
			fmt.Scanln()
		}
		if opcionCrud == "mostrar" && opcionTerritorio == "parroquia" {
			t := table.New().Headers("ID", "parroquia", "canton").BorderStyle(
				lipgloss.NewStyle().Foreground(lipgloss.Color("86")),
			)
			parroquias, _ := db.MostrarParroquias()
			cantones, _ := db.MostrarCatones()
			for _, p := range parroquias {
				var canton string
				for _, c := range cantones {
					if p.IdCanton == c.Id {
						canton = c.Nombre
					}
				}
				t.Row(strconv.Itoa(p.Id), p.Nombre, canton)
			}
			fmt.Println(t)
			fmt.Println("Presione Enter para continuar")
			fmt.Scanln()
		}
		if opcionCrud == "borrar" && opcionTerritorio == "provincia" {
			var provinciaSeleccionada modelos.Provincia
			provincias, _ := db.MostrarProvincias()

			// que seleccione la provincia
			// llenar las opciones
			var opciones []huh.Option[modelos.Provincia]

			for _, p := range provincias {
				var opcion huh.Option[modelos.Provincia] = huh.Option[modelos.Provincia]{
					Key:   p.Nombre,
					Value: *p,
				}
				opciones = append(opciones, opcion)
			}

			// sacarle el id
			g1 := huh.NewGroup(
				huh.NewSelect[modelos.Provincia]().Value(&provinciaSeleccionada).
					Options(opciones...),
			)
			f := huh.NewForm(g1)
			f.Run()

			var borrar bool
			g2 := huh.NewGroup(
				huh.NewConfirm().Title(
					fmt.Sprintf("Desea borrar: %s", provinciaSeleccionada.Nombre),
				).Value(&borrar),
			)
			f = huh.NewForm(g2)
			f.Run()

			// menu de borrado
			if borrar {
				err = db.EliminarPorId("provincia", strconv.Itoa(provinciaSeleccionada.Id))
				if err != nil {
					panic(err)
				}
			}
		}
		if opcionCrud == "borrar" && opcionTerritorio == "canton" {
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
			f := huh.NewForm(g1)
			f.Run()

			var borrar bool
			g2 := huh.NewGroup(
				huh.NewConfirm().Title(
					fmt.Sprintf("Desea borrar: %s", cantonSeleccionado.Nombre),
				).Value(&borrar),
			)
			f = huh.NewForm(g2)
			f.Run()

			// menu de borrado
			if borrar {
				err = db.EliminarPorId("canton", strconv.Itoa(cantonSeleccionado.Id))
				if err != nil {
					panic(err)
				}
			}
		}
		if opcionCrud == "borrar" && opcionTerritorio == "parroquia" {
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
			f := huh.NewForm(g1)
			f.Run()

			var borrar bool
			g2 := huh.NewGroup(
				huh.NewConfirm().Title(
					fmt.Sprintf("Desea borrar: %s", parroquiaSeleccionada.Nombre),
				).Value(&borrar),
			)
			f = huh.NewForm(g2)
			f.Run()

			// menu de borrado
			if borrar {
				err = db.EliminarPorId("parroquia", strconv.Itoa(parroquiaSeleccionada.Id))
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

func main() {
	var db string

	db = *Intro()
	InicarApp(db)
}
