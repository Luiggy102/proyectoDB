-- crud procedures 
-- create (insert)
-- read (select)
-- update (update)
-- delete (drop?)
-- provincia/canton/parroquia
-- primero mostrar (read)
use ecuador;

-- Mostrar (read)
DELIMITER //
create PROCEDURE mostrar_provincias()
BEGIN
SELECT id, nombre FROM provincia;
END //

create PROCEDURE mostrar_cantones()
BEGIN
SELECT id, nombre, id_provincia FROM canton;
END //

create PROCEDURE mostrar_parroquia()
BEGIN
SELECT id, nombre, id_canton FROM parroquia;
END //

-- Borrar (delete)
create PROCEDURE borrar_provincia (PID BIGINT)
BEGIN
    DELETE FROM provincia WHERE id = PID;
END //

create PROCEDURE borrar_canton (CID BIGINT)
BEGIN
    DELETE FROM canton WHERE id = CID;
END //

create PROCEDURE borrar_parroquia (PID BIGINT)
BEGIN
    DELETE FROM parroquia WHERE id = PID;
END //

-- crear (insert)
create PROCEDURE crear_provincia (ID BIGINT, NOMBRE TEXT)
BEGIN
INSERT INTO provincia (
    id , nombre
) VALUES ( ID, NOMBRE );
END //

create PROCEDURE crear_canton (ID BIGINT,NOMBRE text, ID_PROVINCIA BIGINT)
BEGIN
INSERT INTO canton (
    id , nombre, id_provincia
) VALUES ( ID, NOMBRE, ID_PROVINCIA );
END //

create PROCEDURE crear_parroquia (ID BIGINT,NOMBRE text, ID_CANTON BIGINT)
BEGIN
INSERT INTO parroquia (
    id , nombre, id_canton
) VALUES ( ID, NOMBRE, ID_CANTON );
END //

-- actualizar (update)
create PROCEDURE actualizar_provincia (PID BIGINT, NUEVO_NOMBRE TEXT)
BEGIN
UPDATE provincia
    SET nombre = NUEVO_NOMBRE
    WHERE id = PID;
END //

create PROCEDURE actualizar_canton (CID BIGINT, NUEVO_NOMBRE TEXT, ID_PROVINCIA BIGINT)
BEGIN
UPDATE canton
    SET nombre = NUEVO_NOMBRE, id_provincia = ID_PROVINCIA
    WHERE id = CID;
END //

create PROCEDURE actualizar_parroquia (PID BIGINT, NUEVO_NOMBRE TEXT, ID_CANTON BIGINT)
BEGIN
UPDATE parroquia
    SET nombre = NUEVO_NOMBRE, id_canton = ID_CANTON
    WHERE id = PID;
END //
DELIMITER ;
