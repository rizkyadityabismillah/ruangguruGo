-- TODO: answer here
SELECT 
id, 
nik, 
CONCAT(first_name,' ', last_name) AS fullname, 
date_of_birth, 
weight, 
address
FROM people
WHERE gender = 'laki-laki'
ORDER BY weight DESC
LIMIT 5;