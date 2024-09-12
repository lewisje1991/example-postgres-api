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
    "content" text NOT NULL,
    "status" text NOT NULL,
    "tags" text NOT NULL,
    "created_at" timestamp NOT NULL, 
    "updated_at" timestamp NOT NULL, 
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "task_diary" (
    "task_id" uuid NOT NULL, 
    "diary_id" uuid NOT NULL, 
    PRIMARY KEY ("task_id", "diary_id")
);