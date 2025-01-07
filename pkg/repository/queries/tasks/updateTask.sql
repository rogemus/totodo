UPDATE
  tasks AS t
SET
  t.name = $2
WHERE
  t.id = $1;
