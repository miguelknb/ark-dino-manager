CREATE TABLE users (
    id BIGINT NOT NULL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL UNIQUE,
    password_hash BYTEA NOT NULL
);

CREATE TABLE tribes (
    id BIGINT NOT NULL PRIMARY KEY
);

CREATE TABLE tribe_members (
    user_id INT NOT NULL,
    tribe_id INT NOT NULL,
    role VARCHAR(10) NOT NULL DEFAULT 'member',

    PRIMARY KEY (user_id, tribe_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (tribe_id) REFERENCES tribes (id)
);

CREATE TABLE species (
    name TEXT NOT NULL PRIMARY KEY
);

CREATE TABLE dinosaurs (
    id BIGINT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    species TEXT NOT NULL,

    FOREIGN KEY (species) REFERENCES species (name)
);
