package main

import (
	"db1final/databases/mariadb"
	"db1final/databases/postgres"
	"db1final/databases/sqlserver"
	"db1final/modelos"
	"fmt"
	"log"
	"slices"
	"strconv"

	btable "github.com/charmbracelet/bubbles/table"
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
		Options(huh.NewOptions("Postgres", "Mariadb", "sqlserver")...).
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
	return &db
}

func InicarApp(opcion string) {
	var db db
	var err error

	if opcion == "Postgres" {
		db, err = postgres.NuevaDbPostgres(postgres.Url)
		if err != nil {
			panic(err)
		}
	}
	if opcion == "Mariadb" {
		db, err = mariadb.NuevaDbMariadb(mariadb.Url)
		if err != nil {
			panic(err)
		}
	}
	if opcion == "sqlserver" {
		db, err = sqlserver.NuevaDBSqlServer(sqlserver.Url)
		if err != nil {
			panic(err)
		}
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
					huh.NewOption("salir", "salir"),
				),
		)
		f := huh.NewForm(g1).WithProgramOptions(tea.WithAltScreen())
		f.Run()

		if opcionCrud == "salir" {
			return
		}

		// preguntar por
		g2 := huh.NewGroup(
			huh.NewSelect[string]().Value(&opcionTerritorio).Title("Seleccione un territorio").
				Options(
					huh.NewOption("provincia", "provincia"),
					huh.NewOption("canton", "canton"),
					huh.NewOption("parroquia", "parroquia"),
				),
		)
		f = huh.NewForm(g2).WithProgramOptions(tea.WithAltScreen())
		f.Run()

		// casos
		if opcionCrud == "mostrar" && opcionTerritorio == "provincia" {
			// https://github.com/charmbracelet/bubbletea/blob/main/examples/table/main.go

			// tabla normal
			// t := table.New().Headers("ID", "Provincia").BorderStyle(
			// 	lipgloss.NewStyle().Foreground(lipgloss.Color("86")),
			// )
			// provincias, _ := db.MostrarProvincias()
			// for _, p := range provincias {
			// 	t.Row(strconv.Itoa(p.Id), p.Nombre)
			// }
			// fmt.Println(t)
			// fmt.Println("Presione Enter para continuar")
			// fmt.Scanln()

			// tabla interactiva
			columna := []btable.Column{
				{Title: "ID", Width: 10}, {Title: "Nombre", Width: 90},
			}
			provincias, _ := db.MostrarProvincias()
			var datos = [][]string{}
			for _, p := range provincias {
				datos = append(datos, []string{
					strconv.Itoa(p.Id), p.Nombre,
				})
			}
			filas := []btable.Row{}
			for i := 0; i < len(datos); i++ {
				filas = append(filas, datos[i])
			}
			fmt.Println("Presione Enter para cerrar")
			imprimirTabla(columna, filas)
		}
		if opcionCrud == "mostrar" && opcionTerritorio == "canton" {
			// t := table.New().Headers("ID", "canton", "provincia").BorderStyle(
			// 	lipgloss.NewStyle().Foreground(lipgloss.Color("86")),
			// )
			// provincias, _ := db.MostrarProvincias()
			// cantones, _ := db.MostrarCatones()
			// for _, c := range cantones {
			// 	var provincia string
			// 	for _, p := range provincias {
			// 		if c.IdProvincia == strconv.Itoa(p.Id) {
			// 			provincia = p.Nombre
			// 		}
			// 	}
			// 	t.Row(strconv.Itoa(c.Id), c.Nombre, provincia)
			// }
			// fmt.Println(t)
			// fmt.Println("Presione Enter para continuar")
			// fmt.Scanln()

			//
			columna := []btable.Column{
				{Title: "ID", Width: 10}, {Title: "Nombre", Width: 40}, {Title: "Provincia", Width: 40},
			}
			cantones, _ := db.MostrarCatones()
			provincias, _ := db.MostrarProvincias()
			var datos = [][]string{}
			for _, c := range cantones {
				var provincia string
				for _, p := range provincias {
					if c.IdProvincia == strconv.Itoa(p.Id) {
						provincia = p.Nombre
					}
				}
				datos = append(datos, []string{
					strconv.Itoa(c.Id), c.Nombre, provincia,
				})
			}
			filas := []btable.Row{}
			for i := 0; i < len(datos); i++ {
				filas = append(filas, datos[i])
			}
			fmt.Println("Presione Enter para cerrar")
			imprimirTabla(columna, filas)
		}
		if opcionCrud == "mostrar" && opcionTerritorio == "parroquia" {
			// t := table.New().Headers("ID", "parroquia", "canton").BorderStyle(
			// 	lipgloss.NewStyle().Foreground(lipgloss.Color("86")),
			// )
			// parroquias, _ := db.MostrarParroquias()
			// cantones, _ := db.MostrarCatones()
			// for _, p := range parroquias {
			// 	var canton string
			// 	for _, c := range cantones {
			// 		if p.IdCanton == c.Id {
			// 			canton = c.Nombre
			// 		}
			// 	}
			// 	t.Row(strconv.Itoa(p.Id), p.Nombre, canton)
			// }
			// fmt.Println(t)
			// fmt.Println("Presione Enter para continuar")
			// fmt.Scanln()

			//
			columna := []btable.Column{
				{Title: "ID", Width: 10}, {Title: "Nombre", Width: 40}, {Title: "Canton", Width: 40},
			}
			parroquias, _ := db.MostrarParroquias()
			cantones, _ := db.MostrarCatones()
			var datos = [][]string{}
			for _, par := range parroquias {
				var canton string
				for _, c := range cantones {
					if strconv.Itoa(par.IdCanton) == strconv.Itoa(c.Id) {
						canton = c.Nombre
					}
				}
				datos = append(datos, []string{
					strconv.Itoa(par.Id), par.Nombre, canton,
				})
			}
			filas := []btable.Row{}
			for i := 0; i < len(datos); i++ {
				filas = append(filas, datos[i])
			}
			fmt.Println("Presione Enter para cerrar")
			imprimirTabla(columna, filas)
		}
		if opcionCrud == "borrar" && opcionTerritorio == "provincia" {
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
			f = huh.NewForm(g2).WithProgramOptions(tea.WithAltScreen())
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
			f := huh.NewForm(g1).WithProgramOptions(tea.WithAltScreen())
			f.Run()

			var borrar bool
			g2 := huh.NewGroup(
				huh.NewConfirm().Title(
					fmt.Sprintf("Desea borrar: %s", cantonSeleccionado.Nombre),
				).Value(&borrar).Affirmative("Sí").Negative("No"),
			)
			f = huh.NewForm(g2).WithProgramOptions(tea.WithAltScreen())
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
			f := huh.NewForm(g1).WithProgramOptions(tea.WithAltScreen())
			f.Run()

			var borrar bool
			g2 := huh.NewGroup(
				huh.NewConfirm().Title(
					fmt.Sprintf("Desea borrar: %s", parroquiaSeleccionada.Nombre),
				).Value(&borrar).Affirmative("Sí").Negative("No"),
			)
			f = huh.NewForm(g2).WithProgramOptions(tea.WithAltScreen())
			f.Run()

			// menu de borrado
			if borrar {
				err = db.EliminarPorId("parroquia", strconv.Itoa(parroquiaSeleccionada.Id))
				if err != nil {
					panic(err)
				}
			}
		}
		if opcionCrud == "crear" && opcionTerritorio == "provincia" {
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
			f.Run()

			err := db.CrearProvincia(idNuevo, nuevoNombre)
			if err != nil {
				panic(err)
			}
		}
		if opcionCrud == "crear" && opcionTerritorio == "canton" {
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
			g1 = huh.NewGroup(
				huh.NewInput().Value(&nuevoNombre).Title("Creación de cantón"),
				huh.NewSelect[modelos.Provincia]().Value(&provinciaSeleccionada).
					Options(opciones...).Title("A que provincia pertenerce"),
			)

			f = huh.NewForm(g1).WithProgramOptions(tea.WithAltScreen())
			f.Run()

			err := db.CrearCanton(idNuevo, nuevoNombre, strconv.Itoa(provinciaSeleccionada.Id))
			if err != nil {
				panic(err)
			}
		}
		if opcionCrud == "crear" && opcionTerritorio == "parroquia" {
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
			g1 = huh.NewGroup(
				huh.NewInput().Value(&nuevoNombre).Title("Creación de parroquia"),
				huh.NewSelect[modelos.Canton]().Value(&canton).
					Options(opciones...).Title("A que canton pertenerce"),
			)

			f = huh.NewForm(g1).WithProgramOptions(tea.WithAltScreen())
			f.Run()

			err = db.CrearParroquia(idNuevo, nuevoNombre, strconv.Itoa(canton.Id))
			if err != nil {
				panic(err)
			}
		}
		if opcionCrud == "actualizar" && opcionTerritorio == "provincia" {
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

			f = huh.NewForm(g1, g2).WithProgramOptions(tea.WithAltScreen())
			f.Run()

			err = db.ActualizarProvincia(provinciaSeleccionada.Id, nuevoNombre)
			if err != nil {
				panic(err)
			}

		}
		if opcionCrud == "actualizar" && opcionTerritorio == "canton" {
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
			f.Run()
			// colocar a que provincia debe pertenercer
			fmt.Println("Ingrese nueva provincia para el canton")
			provincia, _ := seleccionarProvincia(provincias)
			// actualizar en db
			err := db.ActualizarCantonParroquia("canton", cantonParaActualizar.Id, strconv.Itoa(provincia.Id), nuevoNombre)
			if err != nil {
				panic(err)
			}
		}
		if opcionCrud == "actualizar" && opcionTerritorio == "parroquia" {
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
			f.Run()
			// colocar a que canton debe pertenercer
			seleccionCanton, _ = seleccionarCanton(cantones)
			// cambiar en db
			err := db.ActualizarCantonParroquia("parroquia", seleccionParroquia.Id, strconv.Itoa(seleccionCanton.Id), nuevo)
			if err != nil {
				panic(err)
			}
		}
	}

}

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

func main() {
	db := *Intro()
	InicarApp(db)
}
