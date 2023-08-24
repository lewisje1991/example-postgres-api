variable "token" {
  type    = string
  default = getenv("TURSO_TOKEN")
}

env "turso" {
  url     = "libsql+wss://code-bookmarks-lewisje1991.turso.io?authToken=${var.token}"
  exclude = ["_litestream*"]
}