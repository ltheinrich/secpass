UPDATE
    users
SET
    secret = $1
WHERE
    name = $2;

