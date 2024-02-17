CREATE TABLE project (
    id          VARCHAR PRIMARY KEY,
    title       VARCHAR NOT NULL,
    user_id     VARCHAR REFERENCES "user" (id),
    created_at  TIMESTAMP NOT NULL
);
