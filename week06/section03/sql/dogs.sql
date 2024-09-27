PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE dogs(id integer primary key, name text);
INSERT INTO dogs VALUES(1,'frisky');
INSERT INTO dogs VALUES(2,'Goldy');
COMMIT;
