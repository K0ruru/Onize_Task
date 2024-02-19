CREATE TABLE kegiatan (
    id VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL,
    deskripsi VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    deadline_at TIMESTAMP NOT NULL
)