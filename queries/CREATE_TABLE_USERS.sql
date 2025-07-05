DROP TABle IF EXISTS users;
CREATE TABLE users
(
    email      VARCHAR PRIMARY KEY NOT NULL UNIQUE,
    first_name VARCHAR             NOT NULL,
    last_name  VARCHAR,
    avatar_url VARCHAR
);
