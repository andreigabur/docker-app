CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name varchar(250) NOT NULL,
  email varchar(250) NOT NULL
);

INSERT INTO users (name, email) VALUES
    ('John', 'johndoe@mail.com'),
    ('Bill', 'billmc@mail.com'),
    ('Fred', 'fredl@mail.com');