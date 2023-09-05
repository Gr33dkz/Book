Simple Book CRUD

DB:
CREATE TABLE Book (
id VARCHAR ( 50 ) UNIQUE NOT NULL,
author VARCHAR ( 50 ) NOT NULL,
quantity    INTEGER NOT NULL,
price       double precision NOT NULL,
releaseDate DATE NOT NULL,
description varchar ( 500 ) NOT NULL,
createdDate timestamp
);