CREATE TABLE IF NOT EXISTS bookmarks (
  id UUID PRIMARY KEY,
  url VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  tags VARCHAR(255) NOT NULL,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_tags ON bookmarks (tags);