CREATE TABLE IF NOT EXISTS events(
    id serial PRIMARY KEY,
    user_id int NOT NULL,
    title VARCHAR (200),
    description VARCHAR (500),
    timezone VARCHAR (50),
    duration int,
    event_time timestamp without time zone,
        FOREIGN KEY(user_id)
            REFERENCES users(id)
);