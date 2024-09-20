# Helpful Queries

Get all diary entries with tasks and worklogs:
```sql
SELECT * FROM diary d
JOIN diary_tasks dt on dt.diary_id = d.id
JOIN tasks t on t.id = dt.task_id
LEFT JOIN worklog w on w.task_id = t.id;
``