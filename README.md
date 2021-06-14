# Horusec-Admin

# Running local
For run this project is necessary some environments and configuration

Setup your environment variable to horusec-admin connect in your cluster, see this example:
```bash
export KUBERNETES_SERVICE_HOST=localhost
export KUBERNETES_SERVICE_PORT=45507
export NAMESPACE=default
```

Setup your local token if not exists in path and with a **strong value**, see this example:
```bash
sudo touch /var/run/secrets/kubernetes.io/serviceaccount/token
echo "123456789" > /var/run/secrets/kubernetes.io/serviceaccount/token
```

Setup your local certificate if not exists in path and with a valid certificate, see this example:
```bash
sudo touch /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
echo "CERTIFICATE" > /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
```

and run application
```bash
make run-dev
```