SELECT
  p.id,
  p.name,
  p.created
FROM
  projects AS p
WHERE
  p.id = $1;

