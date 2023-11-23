variable "token" {
  type    = string
  default = getenv("DB_TOKEN")
}

env "dev" {
  url     = "libsql+wss://code-bookmarks-dev-lewisje1991.turso.io?authToken=${var.token}"
  exclude = ["_litestream*"]
}
