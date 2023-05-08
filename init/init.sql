DROP TABLE IF EXISTS notes;

CREATE TABLE notes
(
    id       serial UNIQUE,
    user_id  varchar(80) NOT NULL,
    service  varchar(80) NOT NULL,
    login    varchar(80) NOT NULL,
    password varchar(80) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (user_id, service)
);