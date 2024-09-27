PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE cars(id integer primary key, make string, model string, year integer);
INSERT INTO cars VALUES(1,'chevy','camaro',1966);
COMMIT;
