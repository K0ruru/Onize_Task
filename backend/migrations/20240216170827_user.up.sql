CREATE TABLE "user"
(
    id         VARCHAR PRIMARY KEY,
    name       VARCHAR NOT NULL,
    passphrase       VARCHAR NOT NULL,
    email       VARCHAR NOT NULL,
    created_at TIMESTAMP ,
    updated_at TIMESTAMP 
);
