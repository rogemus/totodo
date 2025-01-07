SELECT
  p.id,
  p.name,
  p.created
FROM
  projects AS p
ORDER BY
  p.created
DESC;
