SELECT
  t.id,
  t.description,
  t.created,
  t.status,
  t.projectId,
  p.name AS projectName
FROM
  tasks AS t LEFT OUTER JOIN projects as p
ON
  t.projectId = p.id
ORDER BY
  t.created
DESC;
