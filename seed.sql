-- First entry in the database
INSERT INTO diary (id, day, created_at, updated_at) VALUES ('e19704d2-e0c7-43cc-b5e9-e945a5a6f4c5','2024-09-20 10:00:00', '2024-09-20 11:00:00', '2024-09-20 11:00:00');

INSERT INTO tasks (id,title, description, tags, created_at, updated_at) VALUES ('f450d25f-88e9-4425-8964-a49320bc9e76', 'create website', 'create a website for a client', 'web, client', '2024-09-20 11:00:00', '2024-09-20 11:00:00');
INSERT INTO diary_tasks (diary_id, task_id, status) VALUES ('e19704d2-e0c7-43cc-b5e9-e945a5a6f4c5', 'f450d25f-88e9-4425-8964-a49320bc9e76', 'todo');

INSERT INTO worklog (id, task_id, content, created_at, updated_at) VALUES ('2daaf6f5-4fe7-445c-83e9-cda8040577b2', 'f450d25f-88e9-4425-8964-a49320bc9e76', 'created the header', '2024-09-20 11:00:00', '2024-09-20 11:00:00');
INSERT INTO worklog (id, task_id, content, created_at, updated_at) VALUES ('3bf3b26c-7af0-4411-a625-fd1128f27aff', 'f450d25f-88e9-4425-8964-a49320bc9e76', 'created the navbar', '2024-09-20 12:00:00', '2024-09-20 12:00:00');

-- Additional task
INSERT INTO tasks (id,title, description, tags, created_at, updated_at) VALUES ('e904f76c-2527-46ce-9d03-7d81ebad770e', 'create app', 'create a ap for a client', 'web, client', '2024-10-20 11:00:00', '2024-10-20 11:00:00');
INSERT INTO diary_tasks (diary_id, task_id, status) VALUES ('e19704d2-e0c7-43cc-b5e9-e945a5a6f4c5', 'e904f76c-2527-46ce-9d03-7d81ebad770e', 'todo');