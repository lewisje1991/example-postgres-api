CREATE TABLE IF NOT EXISTS "tasks" (
    "id" uuid NOT NULL,
    "title" text NOT NULL,
    "description" text NOT NULL,
    "tags" text NOT NULL,
    "created_at" timestamp NOT NULL, 
    "updated_at" timestamp NOT NULL, 
    PRIMARY KEY ("id")
);
