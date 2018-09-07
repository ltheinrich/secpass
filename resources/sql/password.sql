SELECT
    title,
    name,
    mail,
    passwordValue,
    url,
    backupCode,
    notes
FROM
    passwords
WHERE
    id = $1
    AND author = $2;

