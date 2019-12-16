# assignment for interview

Simple Go Calculator project.

## RUN

Executable:
```
go build -o main cmd/calculator/handlers.go cmd/calculator/main.go
./main

http://localhost:8888/
```

Docker image
```
docker build -t go-calc -f build/package/Dockerfile .
docker run -d -p 80:8888 go-calc

http://localhost/
```

## DEPLOY

Kubernetes
```
kubectl create -f deployments/kubernetes/k8s-replicaSet.yml

http://192.168.99.100:30000/
```

**AWS ECS:**
```
cd deployments/aws/ecs
terraform init
terraform apply
//copy output: repository url
cd ../../../
$(aws ecr get-login --no-include-email --region eu-west-1)
docker build -t 311744426619.dkr.ecr.eu-west-1.amazonaws.com/go-calc:0.0.2 -f build/package/Dockerfile .
docker push 311744426619.dkr.ecr.eu-west-1.amazonaws.com/go-calc:0.0.2
```
