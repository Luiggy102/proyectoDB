package main

import (
	"db1final/databases/mariadb"
	"db1final/databases/postgres"
	"db1final/databases/sqlserver"
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/huh"
)

// presentación y selección de db
func Intro() *string {
	var db string
	presentación := huh.NewNote().
		Title("Proyecto final de Bases de Datos 1").
		Description("Integrantes:\n\t*Luiggy Tamayo*\n\t*Foy Barba*\n\t*Andrés Vallejo*")
	selecciónDb := huh.NewSelect[string]().
		Options(huh.NewOptions("Postgres", "Mariadb", "sqlserver", "salir")...).
		Value(&db).
		Title("Elija la base de datos a usar:")
	grupo := huh.NewGroup(
		presentación,
		selecciónDb,
	)
	formulario := huh.NewForm(grupo).WithProgramOptions(tea.WithAltScreen())
	if err := formulario.Run(); err != nil {
		log.Fatal(err)
	}
	// if db != "salir" {
	// 	err := spinner.New().
	// 		Title(fmt.Sprintf("Conectando a base de datos %s...", db)).
	// 		Action(func() { time.Sleep(time.Second * 5) }).
	// 		Run()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	//
	// }
	return &db
}

func InicarApp(opcion string) bool {
	var db db
	var err error

	switch opcion {
	case "Postgres":
		db, err = postgres.NuevaDbPostgres(postgres.Url)
		if err != nil {
			panic(err)
		}
	case "Mariadb":
		db, err = mariadb.NuevaDbMariadb(mariadb.Url)
		if err != nil {
			panic(err)
		}
	case "sqlserver":
		db, err = sqlserver.NuevaDBSqlServer(sqlserver.Url)
		if err != nil {
			panic(err)
		}
	case "salir":
		return true
	}
	defer db.Close()

	var opcionCrud string
	var opcionTerritorio string
	var salir bool

	for !salir {
		// seleccionar operación crud
		g1 := huh.NewGroup(
			huh.NewSelect[string]().Value(&opcionCrud).Title("Seleccione una operación").
				Options(
					huh.NewOption("crear", "crear"),
					huh.NewOption("mostrar", "mostrar"),
					huh.NewOption("actualizar", "actualizar"),
					huh.NewOption("borrar", "borrar"),
					huh.NewOption("volver", "volver"),
				),
		)
		f := huh.NewForm(g1).WithProgramOptions(tea.WithAltScreen())
		err = f.Run()
		if err != nil {
			panic(err)
		}

		if opcionCrud == "volver" {
			return true
		}

		// preguntar por
		g2 := huh.NewGroup(
			huh.NewSelect[string]().Value(&opcionTerritorio).Title("Seleccione un territorio").
				Options(
					huh.NewOption("provincia", "provincia"),
					huh.NewOption("canton", "canton"),
					huh.NewOption("parroquia", "parroquia"),
					huh.NewOption("volver", "volver"),
				),
		)
		f = huh.NewForm(g2).WithProgramOptions(tea.WithAltScreen())
		err = f.Run()
		if err != nil {
			panic(err)
		}

		// casos
		// mostrar
		if opcionCrud == "mostrar" && opcionTerritorio == "provincia" {
			MostrarProvincias(db)
		}
		if opcionCrud == "mostrar" && opcionTerritorio == "canton" {
			MostrarCantones(db)
		}
		if opcionCrud == "mostrar" && opcionTerritorio == "parroquia" {
			MostrarParroquias(db)
		}

		// borrar
		if opcionCrud == "borrar" && opcionTerritorio == "provincia" {
			BorrarProvincia(db)
		}
		if opcionCrud == "borrar" && opcionTerritorio == "canton" {
			BorrarCanton(db)
		}
		if opcionCrud == "borrar" && opcionTerritorio == "parroquia" {
			BorrarParroquia(db)
		}

		// crear
		if opcionCrud == "crear" && opcionTerritorio == "provincia" {
			CrearProvincia(db)
		}
		if opcionCrud == "crear" && opcionTerritorio == "canton" {
			CrearCanton(db)
		}
		if opcionCrud == "crear" && opcionTerritorio == "parroquia" {
			CrearParroquia(db)
		}

		// actualizar
		if opcionCrud == "actualizar" && opcionTerritorio == "provincia" {
			ActualizarProvincia(db)
		}
		if opcionCrud == "actualizar" && opcionTerritorio == "canton" {
			ActualizarCaton(db)
		}
		if opcionCrud == "actualizar" && opcionTerritorio == "parroquia" {
			ActualizarParroquia(db)
		}
	}
	return false
}

func main() {
	var opcion string
	for {
		opcion = *Intro()
		if opcion == "salir" {
			return
		}
		InicarApp(opcion)
	}
}
