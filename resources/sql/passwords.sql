SELECT
    id,
    title,
    name,
    mail,
    passwordValue,
    url,
    backupCode,
    notes,
    category
FROM
    passwords
WHERE
    author = $1
ORDER BY
    title ASC;

