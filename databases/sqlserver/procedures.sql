-- sql server
-- crud procedures 
-- create (insert)
-- read (select)
-- update (update)
-- delete (drop?)
-- provincia/canton/parroquia
-- primero mostrar (read)

-- Mostrar (read) 
CREATE PROCEDURE mostrar_provincias AS 
SELECT id, nombre
FROM provincia;
GO

CREATE PROCEDURE mostrar_cantones AS
SELECT id, nombre, id_provincia
FROM canton;
GO

CREATE PROCEDURE mostrar_parroquia AS
SELECT id, nombre, id_canton
FROM parroquia;
GO

-- Borrar (delete) 
CREATE PROCEDURE borrar_provincia @PID BIGINT AS 
DELETE FROM provincia
WHERE id = @PID;
GO

CREATE PROCEDURE borrar_canton @CID BIGINT AS 
DELETE FROM canton
WHERE id = @CID;
GO

CREATE PROCEDURE borrar_parroquia @PID BIGINT AS 
DELETE FROM parroquia
WHERE id = @PID;
GO

-- crear (insert)
CREATE PROCEDURE crear_provincia @ID BIGINT, @NOMBRE TEXT AS 
INSERT INTO provincia (
id , nombre
) VALUES ( @ID, @NOMBRE );
GO

create PROCEDURE crear_canton @ID BIGINT,@NOMBRE text, @ID_PROVINCIA BIGINT AS 
INSERT INTO canton (
id , nombre, id_provincia
) VALUES ( @ID, @NOMBRE, @ID_PROVINCIA )
GO

create PROCEDURE crear_parroquia @ID BIGINT, @NOMBRE text, @ID_CANTON BIGINT AS 
INSERT INTO parroquia (
id , nombre, id_canton
) VALUES ( @ID, @NOMBRE, @ID_CANTON );
GO

-- actualizar (update)
CREATE PROCEDURE actualizar_provincia @PID BIGINT, @NUEVO_NOMBRE TEXT AS 
UPDATE provincia
SET nombre = @NUEVO_NOMBRE
WHERE id = @PID;
GO

CREATE PROCEDURE actualizar_canton @CID BIGINT, @NUEVO_NOMBRE TEXT, @ID_PROVINCIA BIGINT AS 
UPDATE canton
SET nombre = @NUEVO_NOMBRE, id_provincia = @ID_PROVINCIA
WHERE id = @CID;
GO

CREATE PROCEDURE actualizar_parroquia @PID BIGINT, @NUEVO_NOMBRE TEXT, @ID_CANTON BIGINT AS 
UPDATE parroquia
SET nombre = @NUEVO_NOMBRE, id_canton = @ID_CANTON
WHERE id = @PID;
GO
