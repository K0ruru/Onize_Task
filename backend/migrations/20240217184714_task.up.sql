-- Create priority_enum only if it does not exist
DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'priority_enum') THEN
    CREATE TYPE priority_enum AS ENUM ('low', 'medium', 'high');
  END IF;
END $$;

-- Create status_enum only if it does not exist
DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_enum') THEN
    CREATE TYPE status_enum AS ENUM ('not started', 'in progress', 'completed');
  END IF;
END $$;

-- Create task table
CREATE TABLE task (
    id          VARCHAR PRIMARY KEY,
    title       VARCHAR NOT NULL,
    task_desc   VARCHAR NOT NULL,
    label       VARCHAR NOT NULL,
    priority    priority_enum NOT NULL,
    status      status_enum DEFAULT 'not started',
    tags        VARCHAR,
    project_id     VARCHAR REFERENCES project(id),
    created_at  TIMESTAMP NOT NULL,
    due         TIMESTAMP
);
