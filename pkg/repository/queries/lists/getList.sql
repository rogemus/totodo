SELECT
  l.id,
  l.name,
  l.created
FROM
  lists AS l
WHERE
  l.id = $1;

