package mariadb

import (
	"database/sql"
	"db1final/modelos"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Url = "root:F]anS=}G}z+3]Xc@/ecuador"

type Mariadb struct {
	Db *sql.DB
}

func NuevaDbMariadb(url string) (*Mariadb, error) { // constructor
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	return &Mariadb{Db: db}, nil
}

func (db *Mariadb) MostrarProvincias() ([]*modelos.Provincia, error) {
	columnas, err := db.Db.Query("call mostrar_provincias();")
	if err != nil {
		return nil, err
	}
	defer func() {
		err = columnas.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	provincias := []*modelos.Provincia{}
	for columnas.Next() {
		p := modelos.Provincia{}
		err = columnas.Scan(&p.Id, &p.Nombre)
		if err != nil {
			return nil, err
		}
		provincias = append(provincias, &p)
	}
	if err = columnas.Err(); err != nil {
		return nil, err
	}
	return provincias, nil
}
func (db *Mariadb) MostrarCatones() ([]*modelos.Canton, error) {
	columnas, err := db.Db.Query("call mostrar_cantones();")
	if err != nil {
		return nil, err
	}
	defer func() {
		err = columnas.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	cantones := []*modelos.Canton{}
	for columnas.Next() {
		c := modelos.Canton{}
		err = columnas.Scan(&c.Id, &c.Nombre, &c.IdProvincia)
		if err != nil {
			return nil, err
		}
		cantones = append(cantones, &c)
	}
	if err = columnas.Err(); err != nil {
		return nil, err
	}
	return cantones, nil
}

func (db *Mariadb) MostrarParroquias() ([]*modelos.Parroquia, error) {
	columnas, err := db.Db.Query("call mostrar_parroquia();")
	if err != nil {
		return nil, err
	}
	defer func() {
		err = columnas.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	parroquias := []*modelos.Parroquia{}
	for columnas.Next() {
		p := modelos.Parroquia{}
		err = columnas.Scan(&p.Id, &p.Nombre, &p.IdCanton)
		if err != nil {
			return nil, err
		}
		parroquias = append(parroquias, &p)
	}
	if err = columnas.Err(); err != nil {
		return nil, err
	}
	return parroquias, nil
}

func (db *Mariadb) EliminarPorId(tabla string, id string) error {
	switch tabla {
	case "provincia":
		_, err := db.Db.Exec(
			fmt.Sprintf("call borrar_provincia(%s);", id),
		)
		return err
	case "canton":
		_, err := db.Db.Exec(
			fmt.Sprintf("call borrar_canton(%s);", id),
		)
		return err
	case "parroquia":
		_, err := db.Db.Exec(
			fmt.Sprintf("call borrar_parroquia(%s);", id),
		)
		return err
	default:
		return errors.New("opcion inválida")
	}
}

func (db *Mariadb) CrearProvincia(id int, nombre string) error {
	_, err := db.Db.Exec(
		fmt.Sprintf("call crear_provincia(%d,'%s')", id, nombre),
	)
	return err
}

func (db *Mariadb) CrearCanton(id int, nombre string, IdProvincia string) error {
	_, err := db.Db.Exec(
		fmt.Sprintf("call crear_canton(%d,'%s',%s)", id, nombre, IdProvincia),
	)
	return err
}
func (db *Mariadb) CrearParroquia(id int, nombre string, IdParroquia string) error {
	_, err := db.Db.Exec(
		fmt.Sprintf("call crear_parroquia(%d, '%s', %s)", id, nombre, IdParroquia),
	)
	return err
}

func (db *Mariadb) ActualizarProvincia(id int, nuevoNombre string) error {
	_, err := db.Db.Exec(
		fmt.Sprintf("call actualizar_provincia(%d, '%s')", id, nuevoNombre),
	)
	return err
}

func (db *Mariadb) ActualizarCantonParroquia(tabla string, id int, idTabla string, nuevoNombre string) error {
	switch tabla {
	case "canton":
		_, err := db.Db.Exec(
			fmt.Sprintf("call actualizar_canton(%d,'%s',%s);",
				id, nuevoNombre, idTabla),
		)
		return err
	case "parroquia":
		_, err := db.Db.Exec(
			fmt.Sprintf("call actualizar_parroquia(%d,'%s',%s);",
				id, nuevoNombre, idTabla),
		)
		return err
	default:
		return errors.New("opcion invalida")
	}
}

func (db *Mariadb) Close() error {
	err := db.Db.Close()
	if err != nil {
		return err
	}
	return nil
}
