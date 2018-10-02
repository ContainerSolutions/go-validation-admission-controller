# go-validation-admission-controller

# Dependencies

* Go >= 1.11
* Kubernetes >= 1.11

# Testing

```
go test ./...
```

# Running

```
./gen-cert.sh
./ca-bundle.sh
kubectl apply -f manifest.yaml
```
