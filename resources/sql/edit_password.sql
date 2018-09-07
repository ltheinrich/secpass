UPDATE
    passwords
SET
    title = $1,
    name = $2,
    mail = $3,
    passwordValue = $4,
    url = $5,
    backupCode = $6,
    notes = $7
WHERE
    id = $8
    AND author = $9;

