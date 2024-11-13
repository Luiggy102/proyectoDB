package postgres

import (
	"database/sql"
	"db1final/modelos"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Url = "postgres://ludwig@localhost:5432/ecuador?sslmode=disable"

type PostgresDb struct {
	Db *sql.DB
}

func NuevaDbPostgres(url string) (*PostgresDb, error) { // constructor
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresDb{Db: db}, nil
}

func (db *PostgresDb) MostrarProvincias() ([]*modelos.Provincia, error) {
	columnas, err := db.Db.Query("select id, nombre from provincia;")
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
func (db *PostgresDb) MostrarCatones() ([]*modelos.Canton, error) {
	columnas, err := db.Db.Query("select id, nombre, id_provincia from canton")
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

func (db *PostgresDb) MostrarParroquias() ([]*modelos.Parroquia, error) {
	columnas, err := db.Db.Query("select id, nombre, id_canton from parroquia")
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

func (db *PostgresDb) EliminarPorId(tabla string, id string) error {
	_, err := db.Db.Exec(
		fmt.Sprintf("delete from %s where id = %s", tabla, id),
	)
	return err
}

func (db *PostgresDb) Close() error {
	err := db.Db.Close()
	if err != nil {
		return err
	}
	return nil
}
