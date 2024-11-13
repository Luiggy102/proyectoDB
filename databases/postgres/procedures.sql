-- crud procedures 
-- create (insert)
-- read (select)
-- update (update)
-- delete (drop?)
-- provincia/canton/parroquia
-- primero mostrar (read)

------------------------- provincia -------------------------------
-- read
CREATE PROCEDURE mostrar_provincias ()
LANGUAGE SQL
AS $$
SELECT id, nombre
FROM provincia;
$$;

-- create
CREATE PROCEDURE insertar_provincia (nbr TEXT)
LANGUAGE SQL
AS $$
INSERT INTO provincia (
    nombre
) VALUES ( nbr )
$$;

-- update
CREATE PROCEDURE actualizar_provincia (pid BIGINT, nuevo_nombre TEXT)
LANGUAGE SQL
AS $$
UPDATE provincia
    SET nombre = nuevo_nombre
    WHERE id = pid;
$$;

-- delete
CREATE PROCEDURE eliminar_provincia (pid BIGINT)
LANGUAGE SQL
AS $$
DELETE FROM provincia
    WHERE id = pid;
$$;


---------------------------------- canton --------------------------------------
-- read
CREATE PROCEDURE mostrar_cantones ()
LANGUAGE SQL
AS $$
SELECT id, nombre, id_provincia
FROM canton;
$$;

-- create
CREATE PROCEDURE insertar_canton (nbr TEXT, idp BIGINT)
LANGUAGE SQL
AS $$
INSERT INTO canton (
    nombre, id_provincia
) VALUES ( nbr, idp )
$$;

-- update
CREATE PROCEDURE actualizar_canton (cid BIGINT, nuevo_nombre TEXT)
LANGUAGE SQL
AS $$
UPDATE canton
    SET nombre = nuevo_nombre
    WHERE id = cid;
$$;

-- delete
CREATE PROCEDURE eliminar_canton (cid BIGINT)
LANGUAGE SQL
AS $$
DELETE FROM canton
    WHERE id = cid;
$$;


----------------------------------- parroquia -------------------------
-- read
CREATE PROCEDURE mostrar_parroquia ()
LANGUAGE SQL
AS $$
SELECT id, nombre, id_canton
FROM parroquia;
$$;

-- create
CREATE PROCEDURE insertar_parroquia (nbr TEXT, idc BIGINT)
LANGUAGE SQL
AS $$
INSERT INTO parroquia (
    nombre, id_canton
) VALUES ( nbr, idc )
$$;

-- update
CREATE PROCEDURE actualizar_parroquia (pid BIGINT, nuevo_nombre TEXT)
LANGUAGE SQL
AS $$
UPDATE parroquia
    SET nombre = nuevo_nombre
    WHERE id = pid;
$$;

-- delete
CREATE PROCEDURE eliminar_parroquia (pid BIGINT)
LANGUAGE SQL
AS $$
DELETE FROM parroquia
    WHERE id = pid;
$$;

-- procedures nomral
-- CREATE PROCEDURE us_customers ()
-- LANGUAGE SQL
-- AS $$
-- SELECT customer_id, first_name
-- FROM Customers
-- WHERE Country = 'USA';
-- $$;

-- procedures con parametro
-- CREATE PROCEDURE ctr_customers (ctr VARCHAR(50))
-- LANGUAGE SQL
-- AS $$
-- SELECT customer_id, first_name
-- FROM Customers
-- WHERE Country = ctr;
-- $$;
