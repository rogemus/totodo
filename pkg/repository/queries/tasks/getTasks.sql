SELECT
  t.id,
  t.name,
  t.created,
  t.status,
  t.projectId,
  p.name AS projectName
FROM
  tasks AS t LEFT OUTER JOIN projects as p
ON
  t.projectId = p.id
WHERE
  p.id = $1
ORDER BY
  t.created
DESC;
