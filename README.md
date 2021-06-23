# Horusec-Admin

# Running local
For run this project locally is necessary to communicate with an API server of a Kubernetes cluster.

To point to the same cluster as configured for your `kubectl`, just execute:
```bash
NAMESPACE=default KUBECONFIG=~/.kube/config go run ./cmd/app
```