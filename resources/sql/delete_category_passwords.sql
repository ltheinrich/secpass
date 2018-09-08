DELETE FROM passwords
WHERE category = $1
    AND author = $2;

