UPDATE
    categories
SET
    name = $1
WHERE
    id = $2
    AND author = $3;

