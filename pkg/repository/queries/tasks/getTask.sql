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
  t.id = $1;

