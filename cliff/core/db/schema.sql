-- SCHEMA

CREATE TABLE meta_info
(
    version INT PRIMARY KEY
);

CREATE TABLE users
(
    id       INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    username TEXT UNIQUE,
    hashpass TEXT NOT NULL
);

CREATE TYPE video_status AS ENUM ('UPLOADED', 'PROCESSING', 'PROCESSED', 'READY', 'FAILURE');

CREATE TABLE videos
(
    id          INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    slug        TEXT UNIQUE,
    uploader    INT,
    title       TEXT,
    description TEXT,
    status      video_status,

    CONSTRAINT fk_users
        FOREIGN KEY (uploader)
            REFERENCES users (id)
);

-- initial data

INSERT INTO meta_info VALUES (1);
INSERT INTO users (username, hashpass) VALUES ('admin', '$2y$10$XcWmOIgAuT90XB/7cSwK5e1PTEUeJgXcO47Zgjx6RHh2phZVFqc/C'); -- Default password admin/admin