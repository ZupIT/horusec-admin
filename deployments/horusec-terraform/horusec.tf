resource "kubernetes_namespace" "horusec" {
  metadata {
    name = "horusec-system"
  }
}

resource "kubernetes_secret" "horusec_broker" {
  metadata {
    name = "horusec-broker"
    namespace = var.horusec_namespace
  }

  data = {
    "username" = "user"
    "password" = data.kubernetes_secret.rabbit.data.rabbitmq-password
  }
}

resource "kubernetes_secret" "horusec_database" {
  metadata {
    name = "horusec-database"
    namespace = var.horusec_namespace
  }

  data = {
    "username" = "postgres"
    "password" = data.kubernetes_secret.postgres.data.postgresql-password
  }
}

resource "kubernetes_secret" "horusec_jwt" {
  metadata {
    name = "horusec-jwt"
    namespace = var.horusec_namespace
  }

  data = {
    "secret-key" = "74266279-766d-3075-7a2f-36587132a5eb"
  }
}

resource "kubernetes_secret" "horusec_smtp" {
  metadata {
    name = "horusec-smtp"
    namespace = var.horusec_namespace
  }

  data = {
    "username" = "3dcf6374062286"
    "password" = "1a29e895468521"
  }
}
