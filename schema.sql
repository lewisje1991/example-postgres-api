CREATE TABLE IF NOT EXISTS "diary" (
    "id" uuid NOT NULL, 
    "day" date NOT NULL,
    "created_at" timestamp NOT NULL, 
    "updated_at" timestamp NOT NULL, 
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "tasks" (
    "id" uuid NOT NULL,
    "title" text NOT NULL,
    "description" text NOT NULL,
    "tags" text NOT NULL,
    "created_at" timestamp NOT NULL, 
    "updated_at" timestamp NOT NULL, 
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "diary_tasks" (
    "diary_id" uuid NOT NULL,
    "task_id" uuid NOT NULL,
    "status" text NOT NULL,
    PRIMARY KEY ("diary_id", "task_id")
);

CREATE TABLE IF NOT EXISTS "worklog" (
    "id" uuid NOT NULL,
    "task_id" uuid NOT NULL,
    "content" text NOT NULL,
    "created_at" timestamp NOT NULL, 
    "updated_at" timestamp NOT NULL, 
    PRIMARY KEY ("id", "task_id")
);
