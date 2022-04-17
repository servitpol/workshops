CREATE TABLE IF NOT EXISTS events(
    id serial PRIMARY KEY,
    user_id int NOT NULL,
    title VARCHAR (200),
    description VARCHAR (500),
    timestamp_from int,
    timestamp_to int,
        FOREIGN KEY(user_id)
            REFERENCES users(id)
);