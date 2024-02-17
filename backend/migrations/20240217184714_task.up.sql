CREATE TYPE priority_enum AS ENUM ('low', 'medium', 'high');
CREATE TYPE status_enum AS ENUM ('not started', 'in progress', 'completed');

CREATE TABLE task (
    id          VARCHAR PRIMARY KEY,
    title       VARCHAR NOT NULL,
    task_desc   VARCHAR NOT NULL,
    label       VARCHAR NOT NULL,
    priority    priority_enum NOT NULL,
    status      status_enum NOT NULL,
    tags        VARCHAR,
    project  VARCHAR REFERENCES project(id),
    created_at  TIMESTAMP NOT NULL,
    due         TIMESTAMP
);