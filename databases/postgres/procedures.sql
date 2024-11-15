-- crud procedures 
-- create (insert)
-- read (select)
-- update (update)
-- delete (drop?)
-- provincia/canton/parroquia
-- primero mostrar (read)

------------------------ Mostrar (read) ----------------------
CREATE PROCEDURE mostrar_provincias ()
LANGUAGE SQL
AS $$
SELECT id, nombre
FROM provincia;
$$;

CREATE PROCEDURE mostrar_cantones ()
LANGUAGE SQL
AS $$
SELECT id, nombre, id_provincia
FROM canton;
$$;

CREATE PROCEDURE mostrar_parroquia ()
LANGUAGE SQL
AS $$
SELECT id, nombre, id_canton
FROM parroquia;
$$;

------------------- Borrar (delete) ---------------------
CREATE PROCEDURE borrar_provincia (PID BIGINT)
LANGUAGE SQL
AS $$
    DELETE FROM provincia
        WHERE id = PID;
$$;

CREATE PROCEDURE borrar_canton (CID BIGINT)
LANGUAGE SQL
AS $$
    DELETE FROM canton
        WHERE id = CID;
$$;

CREATE PROCEDURE borrar_parroquia (PID BIGINT)
LANGUAGE SQL
AS $$
    DELETE FROM parroquia
        WHERE id = PID;
$$;

------------------ crear (insert) -----------------------
CREATE PROCEDURE crear_provincia (ID BIGINT, NOMBRE TEXT)
LANGUAGE SQL
AS $$
INSERT INTO provincia (
    id , nombre
) VALUES ( ID, NOMBRE )
$$;

create PROCEDURE crear_canton (ID BIGINT,NOMBRE text, ID_PROVINCIA BIGINT)
LANGUAGE SQL
AS $$
INSERT INTO canton (
    id , nombre, id_provincia
) VALUES ( ID, NOMBRE, ID_PROVINCIA )
$$;

create PROCEDURE crear_parroquia (ID BIGINT,NOMBRE text, ID_CANTON BIGINT)
LANGUAGE SQL
AS $$
INSERT INTO parroquia (
    id , nombre, id_canton
) VALUES ( ID, NOMBRE, ID_CANTON )
$$;

------------------ actualizar (update) ---------------------------
CREATE PROCEDURE actualizar_provincia (PID BIGINT, NUEVO_NOMBRE TEXT)
LANGUAGE SQL
AS $$
UPDATE provincia
    SET nombre = NUEVO_NOMBRE
    WHERE id = PID;
$$;

CREATE PROCEDURE actualizar_canton (CID BIGINT, NUEVO_NOMBRE TEXT, ID_PROVINCIA BIGINT)
LANGUAGE SQL
AS $$
UPDATE canton
    SET nombre = NUEVO_NOMBRE, id_provincia = ID_PROVINCIA
    WHERE id = CID;
$$;

CREATE PROCEDURE actualizar_parroquia (PID BIGINT, NUEVO_NOMBRE TEXT, ID_CANTON BIGINT)
LANGUAGE SQL
AS $$
UPDATE parroquia
    SET nombre = NUEVO_NOMBRE, id_canton = ID_CANTON
    WHERE id = PID;
$$;
