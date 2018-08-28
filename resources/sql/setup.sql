-- Users Table name, password, secret
CREATE TABLE IF NOT EXISTS users (
        name VARCHAR(255) UNIQUE,
        PASSWORD VARCHAR(255),
        secret VARCHAR(255),
        crypter VARCHAR(255),
        PRIMARY KEY (name)
);

-- Passwords Table name, password, author
CREATE TABLE IF NOT EXISTS passwords (
        title VARCHAR(255),
        name VARCHAR(255),
        PASSWORD VARCHAR(255),
        author VARCHAR(255),
        PRIMARY KEY (title,
            name,
            author),
        FOREIGN KEY (author) REFERENCES users (name) ON DELETE CASCADE
)
