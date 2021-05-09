SELECT a.Student, SUM(a.Score) / 3 AS average FROM Score a
LEFT JOIN `Subject` b ON a.`Subject` = b.`name`
WHERE a.Score > 60
GROUP BY a.Student
HAVING average > 80;