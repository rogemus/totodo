UPDATE
  projects AS p
SET
  p.name = $2
WHERE
  p.id = $1;
