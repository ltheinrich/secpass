SELECT
    name,
    passwordHash,
    secret,
    crypter
FROM
    users
WHERE
    name = $1;

