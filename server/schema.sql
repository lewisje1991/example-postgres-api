CREATE TABLE IF NOT EXISTS "bookmarks" (
    "id" uuid NOT NULL,
    "url" text NOT NULL, 
    "description" text NOT NULL, 
    "tags" text NOT NULL, 
    "created_at" timestamp NOT NULL, 
    "updated_at" timestamp NOT NULL, 
    PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "idx_tags" ON "bookmarks" ("tags");

CREATE TABLE IF NOT EXISTS "notes" (
    "id" uuid NOT NULL, 
    "title" text NOT NULL, 
    "content" text NOT NULL,
    "tags" text NOT NULL,
    "created_at" timestamp NOT NULL, 
    "updated_at" timestamp NOT NULL, 
    PRIMARY KEY ("id")
);

-- CREATE TABLE IF NOT EXISTS "diary" (
--     "id" uuid NOT NULL, 
--     "day" date NOT NULL,
--     "created_at" timestamp NOT NULL, 
--     "updated_at" timestamp NOT NULL, 
--     PRIMARY KEY ("id")
-- );

-- CREATE TABLE IF NOT EXISTS "tasks" (
--     "id" uuid NOT NULL, 
--     "title" text NOT NULL, 
--     "content" text NOT NULL,
--     "tags" text NOT NULL,
--     "created_at" timestamp NOT NULL, 
--     "updated_at" timestamp NOT NULL, 
--     PRIMARY KEY ("id")
-- );

-- CREATE TABLE IF NOT EXISTS "diary_tasks" (
--     "id" uuid NOT NULL, 
--     "diary_id" uuid NOT NULL,
--     "task_id" uuid NOT NULL,
--     "created_at" timestamp NOT NULL,
--     "updated_at" timestamp NOT NULL,
--     PRIMARY KEY ("id"),
--     FOREIGN KEY ("diary_id") REFERENCES "diary" ("id"),
--     FOREIGN KEY ("task_id") REFERENCES "tasks" ("id")
-- );