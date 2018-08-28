UPDATE
    passwords
SET
    PASSWORD = $1
WHERE
    title = $2
    AND name = $3
    AND author = $4;

