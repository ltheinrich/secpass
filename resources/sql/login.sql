SELECT
    name,
    PASSWORD,
    secret,
    crypter
FROM
    users
WHERE
    name = $1;

