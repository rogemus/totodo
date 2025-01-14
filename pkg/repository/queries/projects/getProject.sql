SELECT
  p.id,
  p.name,
  p.created,
  (
    SELECT
    	COUNT(*)
    FROM
    	tasks AS t
    WHERE
    	t.projectId == $1 AND
    	t.status == 'done'
  ) as tasksDoneCount,
  (
    SELECT
    	COUNT(*)
    FROM
    	tasks AS t
    WHERE
    	t.projectId == $1
  ) as tasksCount
FROM
  projects AS p
WHERE
  p.id = $1;

