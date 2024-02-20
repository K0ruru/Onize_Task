CREATE TABLE "user"
(
    id         VARCHAR PRIMARY KEY,
    name       VARCHAR NOT NULL,
    passphrase VARCHAR NOT NULL,
    email      VARCHAR NOT NULL UNIQUE,
    no_telp    VARCHAR
);

