SELECT
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
    id = $1
    AND author = $2;

