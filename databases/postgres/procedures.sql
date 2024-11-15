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
CREATE PROCEDURE borrar_por_id (tabla TEXT, ID BIGINT)
LANGUAGE SQL
AS $$
DELETE FROM tabla
    WHERE id = ID;
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
CREATE PROCEDURE actualizar_provincia (ID BIGINT, NUEVO_NOMBRE TEXT)
LANGUAGE SQL
AS $$
UPDATE provincia
    SET nombre = NUEVO_NOMBRE
    WHERE id = ID;
$$;

CREATE PROCEDURE actualizar_canton (ID BIGINT, NUEVO_NOMBRE TEXT, ID_PROVINCIA BIGINT)
LANGUAGE SQL
AS $$
UPDATE canton
    SET nombre = NUEVO_NOMBRE, id_provincia = ID_PROVINCIA
    WHERE id = ID;
$$;

CREATE PROCEDURE actualizar_parroquia (ID BIGINT, NUEVO_NOMBRE TEXT, ID_CANTON BIGINT)
LANGUAGE SQL
AS $$
UPDATE parroquia
    SET nombre = NUEVO_NOMBRE, id_canton = ID_CANTON
    WHERE id = ID;
$$;
