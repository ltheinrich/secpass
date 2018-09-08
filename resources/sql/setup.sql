-- Users Table name, passwordHash, secret, crypter
CREATE TABLE IF NOT EXISTS users (
                name VARCHAR(255) UNIQUE,
                passwordHash VARCHAR(255),
                secret VARCHAR(255),
                crypter VARCHAR(255),
                PRIMARY KEY (name)
);

-- Categories Table id, name, author
CREATE TABLE IF NOT EXISTS categories (
                id SERIAL UNIQUE,
                name VARCHAR(255),
                author VARCHAR(255),
                PRIMARY KEY (id),
                FOREIGN KEY (author) REFERENCES users (name) ON DELETE CASCADE
);

-- Passwords Table id, title, name, mail, passwordValue, url, backupCode, notes, category, author
CREATE TABLE IF NOT EXISTS passwords (
                id SERIAL UNIQUE,
                title VARCHAR(255),
                name VARCHAR(255),
                mail VARCHAR(255),
                passwordValue VARCHAR(255),
                url VARCHAR(255),
                backupCode VARCHAR(255),
                notes VARCHAR(255),
                category INT,
                author VARCHAR(255),
                PRIMARY KEY (id),
                FOREIGN KEY (category) REFERENCES categories (id) ON DELETE CASCADE,
                FOREIGN KEY (author) REFERENCES users (name) ON DELETE CASCADE
)
