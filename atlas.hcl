env "local" {
  src = "file://schema.sql"
  url = "postgres://postgres:postgres@localhost:5432/code-bookmarks?search_path=public&sslmode=disable"
  dev = "docker://postgres/15/dev?search_path=public"
}

env "supa" {
  src = "file://schema.sql"
  url = "postgresql://postgres:[PWD]@db.ptlinblwdgrzohoatgpv.supabase.co:5432/postgres?search_path=public&sslmode=disable"
  dev = "docker://postgres/15/dev?search_path=public"
}
