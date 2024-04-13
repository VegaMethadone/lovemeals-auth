CREATE TABLE IF NOT EXISTS metadata (
    id SERIAL PRIMARY KEY,
    users INTEGER,
    executors INTEGER
);

INSERT INTO metadata (users, executors) VALUES (0, 0);


CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    phone VARCHAR(11)
);

CREATE TABLE  IF NOT EXISTS  executors (
    id SERIAL PRIMARY KEY,
    executor_id INTEGER,
    login VARCHAR(64),
    password VARCHAR(64),
    phone VARCHAR(11)
);

