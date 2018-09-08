UPDATE
    users
SET
    passwordHash = $1,
    crypter = $2
WHERE
    name = $3;

