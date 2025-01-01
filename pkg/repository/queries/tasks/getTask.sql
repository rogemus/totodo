SELECT
  t.id,
  t.description,
  t.created,
  t.status,
  t.listId,
  l.name AS listName
FROM
  tasks AS t LEFT OUTER JOIN lists as l
ON
  t.listId = l.id
WHERE
  t.id = $1;

