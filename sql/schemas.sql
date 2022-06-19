CREATE TABLE users(
    id INTEGER GENERATED ALWAYS AS IDENTITY NOT NULL UNIQUE,
    username VARCHAR(25) UNIQUE NOT NULL,
    firstname VARCHAR(40) NOT NULL,
    lastname VARCHAR(40) NOT NULL,
    email VARCHAR(40) UNIQUE NOT NULL,
    password VARCHAR(60) NOT NULL,
    roles VARCHAR(20)[] NOT NULL DEFAULT '{ USER }',
    birthday TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT smart_update CHECK(updated_at >= created_at)
);

CREATE TABLE profiles(
    id INTEGER GENERATED ALWAYS AS IDENTITY NOT NULL UNIQUE,
    last_connection TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT profile_users_fk FOREIGN KEY(id) REFERENCES users(id),
    CONSTRAINT smart_update CHECK(updated_at >= created_at)
);

CREATE TABLE posts(
    id INTEGER GENERATED ALWAYS AS IDENTITY NOT NULL UNIQUE,
    profile_id INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    external_source TEXT NOT NULL,
    description TEXT,
    tags VARCHAR(50)[],
    is_read_only BOOLEAN DEFAULT 'f',
    CONSTRAINT profile_post_fk FOREIGN KEY(profile_id) REFERENCES profiles(id),
    CONSTRAINT smart_update CHECK(updated_at >= created_at)
);

CREATE TABLE mentions(
    id INTEGER GENERATED ALWAYS AS IDENTITY NOT NULL UNIQUE,
    source INTEGER NOT NULL,
    target VARCHAR(25) NOT NULL,
    CONSTRAINT source_post_fk FOREIGN KEY(source) REFERENCES posts(id),
    CONSTRAINT target_profile_fk FOREIGN KEY(target) REFERENCES users(username)
);