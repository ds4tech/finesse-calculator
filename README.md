# assignment for interview

Simple Go Calculator project with some math function:
- Sum
- Sqrt
- Factorial(n uint64)
- IsPrime(n int)
- Log(n float64)

Function IsPrime returns bool, according to provided number (is it prime or not).

**Useage:**

1.SUM 
- (x int64, y int64)
- http://localhost:8888/sum
- JSON: {"num1":"2","num2":"4"}
- curl: 
```
curl --location --request POST "http://localhost:8888/sum" \
--header "Content-Type: application/json" \
--data "{\"num1\":\"2\",
\"num2\":\"4\"}"
```

2.SQRT 
- (x float64)
- http://localhost:8888/sqrt
- JSON: {"number":"144"}
- curl: 
```
curl --location --request POST "http://localhost:8888/sum" \
--header "Content-Type: application/json" \
--data "{\"number\":\"144\"}"
```

## RUN

**Executable**
```
go build -o main cmd/calculator/handlers.go cmd/calculator/main.go
./main

http://localhost:8888/
```

**Docker image**
```
docker build -t go-calc -f build/package/Dockerfile .
docker run -d -p 80:8888 go-calc

http://localhost/
```

## DEPLOY

**Kubernetes**
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
