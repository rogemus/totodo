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
    	t.projectId == p.id AND
    	t.status == 'done'
  ) as tasksDoneCount,
  (
    SELECT
    	COUNT(*)
    FROM
    	tasks AS t
    WHERE
    	t.projectId == p.id
  ) as tasksCount
FROM
  projects AS p
ORDER BY
  p.created
DESC;
