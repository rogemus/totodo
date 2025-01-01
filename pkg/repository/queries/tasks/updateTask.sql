UPDATE
  tasks AS t
SET
  t.description = $2
WHERE
  t.id = $1;
