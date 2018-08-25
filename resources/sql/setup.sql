-- Users Table name, password, secret
CREATE TABLE IF NOT EXISTS users (name VARCHAR(255) UNIQUE, password VARCHAR(255), secret VARCHAR(255), PRIMARY KEY (name));

-- Passwords Table name, password, author
CREATE TABLE IF NOT EXISTS passwords (name VARCHAR(255), password VARCHAR(255), author VARCHAR(255),
    PRIMARY KEY (name, author),
    FOREIGN KEY (author) REFERENCES users(name) ON DELETE CASCADE)