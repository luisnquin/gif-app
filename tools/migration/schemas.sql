CREATE TABLE users(
    id INTEGER GENERATED ALWAYS AS IDENTITY NOT NULL,
    username VARCHAR(25) UNIQUE NOT NULL,
    firstname VARCHAR(40) NOT NULL,
    lastname VARCHAR(40) NOT NULL,
    email VARCHAR(40) NOT NULL,
    password VARCHAR(60) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'USER',
    birthday TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT smart_update CHECK(updated_at >= created_at)
);