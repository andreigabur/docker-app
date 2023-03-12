CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  name varchar(250) NOT NULL
);

INSERT INTO products (name) VALUES
    ('First great product'),
    ('Another amazing product'),
    ('The best product we have');