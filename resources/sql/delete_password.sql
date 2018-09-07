DELETE FROM passwords
WHERE id = $1
    AND author = $2;

