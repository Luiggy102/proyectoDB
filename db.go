package main

import "db1final/modelos"

type db interface {
	MostrarProvincias() ([]*modelos.Provincia, error)
	MostrarCatones() ([]*modelos.Canton, error)
	MostrarParroquias() ([]*modelos.Parroquia, error)
	EliminarPorId(tabla string, id string) error
	CrearProvincia(id int, nombre string) error
	CrearCanton(id int, nombre string, IdProvincia string) error
	CrearParroquia(id int, nombre string, IdParroquia string) error
	ActualizarProvincia(id int, nuevoNombre string) error
	ActualizarCantonParroquia(tabla string, id int, idTabla string, nuevoNombre string) error
	Close() error
}
