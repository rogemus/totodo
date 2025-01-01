UPDATE
  lists AS l
SET
  l.name = $2
WHERE
  l.id = $1;
