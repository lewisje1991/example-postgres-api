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