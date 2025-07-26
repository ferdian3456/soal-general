CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name varchar(40) NOT NULL,
    email varchar(80) NOT NULL,
    phone varchar(15) NOT NULL
)