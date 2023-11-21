variable "token" {
  type    = string
  default = getenv("DB_TOKEN")
}

env "turso" {
  url     = "libsql+wss://code-bookmarks-lewisje1991.turso.io?authToken=${var.token}"
  exclude = ["_litestream*"]
}

env "local" {
  url     = "sqlite://db:8080"
  exclude = ["_litestream*"]
}