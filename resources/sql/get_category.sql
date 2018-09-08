SELECT
    name
FROM
    categories
WHERE
    id = $1
    AND author = $2;

