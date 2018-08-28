SELECT
    title,
    name,
    PASSWORD
FROM
    passwords
WHERE
    author = $1
ORDER BY
    title ASC;

