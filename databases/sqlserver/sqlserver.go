package sqlserver

import (
	"database/sql"
	"db1final/modelos"
	"errors"
	"fmt"
	_ "github.com/microsoft/go-mssqldb"
	"log"
)

var Url = "sqlserver://sa:yourStrong()Password@localhost:1433?database=ecuador"

type SqlserverDb struct {
	Db *sql.DB
}

func NuevaDBSqlServer(url string) (*SqlserverDb, error) {
	db, err := sql.Open("sqlserver", url)
	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	if err != nil {
		return nil, err
	}
	return &SqlserverDb{Db: db}, nil
}

func (db *SqlserverDb) MostrarProvincias() ([]*modelos.Provincia, error) {
	columnas, err := db.Db.Query("exec mostrar_provincias;")
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
func (db *SqlserverDb) MostrarCatones() ([]*modelos.Canton, error) {
	columnas, err := db.Db.Query("exec mostrar_cantones;")
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

func (db *SqlserverDb) MostrarParroquias() ([]*modelos.Parroquia, error) {
	columnas, err := db.Db.Query("exec mostrar_parroquia;")
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

func (db *SqlserverDb) EliminarPorId(tabla string, id string) error {
	switch tabla {
	case "provincia":
		_, err := db.Db.Exec(
			fmt.Sprintf("exec borrar_provincia %s;", id),
		)
		return err
	case "canton":
		_, err := db.Db.Exec(
			fmt.Sprintf("exec borrar_canton %s;", id),
		)
		return err
	case "parroquia":
		_, err := db.Db.Exec(
			fmt.Sprintf("exec borrar_parroquia %s;", id),
		)
		return err
	default:
		return errors.New("opcion inválida")
	}
}

func (db *SqlserverDb) CrearProvincia(id int, nombre string) error {
	_, err := db.Db.Exec(
		fmt.Sprintf("exec crear_provincia %d,'%s';", id, nombre),
	)
	return err
}

func (db *SqlserverDb) CrearCanton(id int, nombre string, IdProvincia string) error {
	_, err := db.Db.Exec(
		fmt.Sprintf("exec crear_canton %d,'%s',%s;", id, nombre, IdProvincia),
	)
	return err
}
func (db *SqlserverDb) CrearParroquia(id int, nombre string, IdParroquia string) error {
	_, err := db.Db.Exec(
		fmt.Sprintf("exec crear_parroquia %d, '%s', %s;", id, nombre, IdParroquia),
	)
	return err
}

func (db *SqlserverDb) ActualizarProvincia(id int, nuevoNombre string) error {
	_, err := db.Db.Exec(
		fmt.Sprintf("exec actualizar_provincia %d, '%s';", id, nuevoNombre),
	)
	return err
}

func (db *SqlserverDb) ActualizarCantonParroquia(tabla string, id int, idTabla string, nuevoNombre string) error {
	switch tabla {
	case "canton":
		_, err := db.Db.Exec(
			fmt.Sprintf("exec actualizar_canton %d,'%s',%s;",
				id, nuevoNombre, idTabla),
		)
		return err
	case "parroquia":
		_, err := db.Db.Exec(
			fmt.Sprintf("exec actualizar_parroquia %d,'%s',%s;",
				id, nuevoNombre, idTabla),
		)
		return err
	default:
		return errors.New("opcion invalida")
	}
}

func (db *SqlserverDb) Close() error {
	err := db.Db.Close()
	if err != nil {
		return err
	}
	return nil
}
