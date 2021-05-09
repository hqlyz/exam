SELECT Student FROM Score a
JOIN Subject b ON a.`Subject` = b.`name`
WHERE a.Score > 80
GROUP BY Student
HAVING COUNT(Student) >= 2

-- Result is Lucas