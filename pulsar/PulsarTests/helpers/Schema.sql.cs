namespace PulsarTests.helpers
{
    public class Schema
    {
        public static string Get()
        {
            // FIXME(RaboliotLeGris): until we have a more centralized wai to initialize the database, it will do
            return @"
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
            ";
        }
    }
}