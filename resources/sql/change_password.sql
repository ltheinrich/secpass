UPDATE
    users
SET
    PASSWORD = $1,
    crypter = $2
WHERE
    name = $3;

