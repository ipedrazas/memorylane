
CREATE TABLE tasks
(
    id SERIAL,
    name TEXT NOT NULL,
    completed BOOLEAN,
    CONSTRAINT tasks_pkey PRIMARY KEY (id)
)