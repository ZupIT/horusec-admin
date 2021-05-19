resource "helm_release" "rabbit" {
  name = "rabbitmq"
  chart = "https://charts.bitnami.com/bitnami/rabbitmq-8.11.1.tgz"
  namespace = kubernetes_namespace.queue.metadata[0].name

  set {
    name = "auth.password"
    value = "qQAUEGhQ6R"
  }

  set {
    name = "auth.erlangCookie"
    value = "DX5NKjaLajEYC9t6hJujJa25PqpbFXF4"
  }
}

resource "kubernetes_namespace" "queue" {
  metadata {
    name = "queue"
  }
}

data "kubernetes_secret" "rabbit" {
  metadata {
    name = helm_release.rabbit.name
    namespace = helm_release.rabbit.namespace
  }

  depends_on = [
    helm_release.rabbit
  ]
}