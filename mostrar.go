package main

import (
	"fmt"
	"strconv"

	btable "github.com/charmbracelet/bubbles/table"
)

func MostrarProvincias(db db) {
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
func MostrarCantones(db db) {
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
func MostrarParroquias(db db) {
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
