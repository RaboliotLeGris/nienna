-- SCHEMA

CREATE TABLE meta_info
(
    version INT,

    PRIMARY KEY(version)
);

CREATE TABLE users
(
    id       INT GENERATED ALWAYS AS IDENTITY,
    username TEXT UNIQUE,

    PRIMARY KEY (id)
);

CREATE TABLE videos
(
    id          INT GENERATED ALWAYS AS IDENTITY,
    slug        TEXT UNIQUE,
    uploader    INT,
    title       TEXT,
    description TEXT,
    status      TEXT,

    PRIMARY KEY (id),
    CONSTRAINT fk_users
        FOREIGN KEY (uploader)
            REFERENCES users (id)
);

-- initial data

INSERT INTO meta_info VALUES (1);