# use-k8s-secret

## Tasks

### docker-build

```bash
docker build -t use-k8s-secret:latest .
```

### k8s-namespace-create

```bash
kubectl namespace create use-k8s-secret
```

### create-tls-cert

```bash
openssl ecparam -genkey -name secp384r1 -out tls.key
openssl req -new -x509 -sha256 -key tls.key -out tls.crt -days 3650 -subj "/O=a-h/CN=use-k8s-secret"
```

### k8s-create-secret

When mounted into a volume, k8s TLS secrets are always called `tls.crt` and `tls.key` regardless of what filename you use here.

```bash
kubectl create secret tls tls-secret --cert=tls.crt --key=tls.key --namespace=use-k8s-secret
```

### k8s-apply

```bash
kubectl apply -f deployment.yaml -f service.yaml --namespace=use-k8s-secret
```

### k8s-get-services

```bash
kubectl get services --namespace=use-k8s-secret
```

### k8s-get-pods

```bash
kubectl get pods --namespace=use-k8s-secret
```

### k8s-attach

```bash
kubectl exec -it deploy/use-k8s-secret --namespace=use-k8s-secret -- /bin/sh
```
