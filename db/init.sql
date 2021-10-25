\connect postgres

CREATE TABLE timestamps (
    id serial PRIMARY KEY,
    runner_id integer NOT NULL,
    ts TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP 
);