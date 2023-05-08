DROP TABLE IF EXISTS notes;

CREATE TABLE notes
(
    id           serial UNIQUE,
    chat_id      varchar(80) NOT NULL,
    service_name varchar(80) NOT NULL,
    login        varchar(80) NOT NULL,
    password     varchar(80) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (chat_id, service_name)
);