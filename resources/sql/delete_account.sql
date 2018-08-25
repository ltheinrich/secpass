-- Passwords
DELETE FROM passwords WHERE user = $1;

-- Account
DELETE FROM users WHERE name = $2;