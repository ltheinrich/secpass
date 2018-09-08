SELECT
    id,
    name
FROM
    categories
WHERE
    author = $1
ORDER BY
    name ASC;

