# Simple calculator application written in Go lang.

1. [Introduction](#intro)
2. [Build](#build) <br>
   2.1. [Exec](#build.exe) <br>
   2.2. [Docker](#build.docker)
3. [Deploy](#deploy) <br>
 3.1. [Kubernetes](#deploy.k8s) <br>
4. [Usage](#usage)
5. [Continous Integration](#ci)


## Introduction <a name="intro"></a>

Simple Go Calculator project with some math function:<a name="intro"></a>
- Sum
- Sqrt
- Factorial
- IsPrime
- Log

## BUILD <a name="build"></a>

### Executable <a name="build.exe"></a>
```
go build -o main cmd/calculator/main.go
./main

http://localhost:8888/
```

### Docker container <a name="build.docker"></a>
```
docker build -t go-calc -f build/package/Dockerfile .
docker run -d -p 80:8888 go-calc

http://localhost/
```

## DEPLOY <a name="deploy"></a>

### Kubernetes <a name="deploy.k8s"></a>
```
kubectl create -f deployments/kubernetes/k8s-replicaSet.yml

http://192.168.99.100:30000/
```

## USEAGE <a name="usage"></a>

1. SUM
- (x int64, y int64)
- http://localhost:8888/v1/sum
- JSON: {"num1":"2","num2":"4"}
- curl:
```
curl --location --request POST "http://localhost:8888/v1/sum" \
--header "Content-Type: application/json" \
--data "{\"num1\":\"2\", \"num2\":\"4\"}"
```
2. SQRT
- (x float64)
- http://localhost:8888/v1/sqrt
- JSON: {"number":"144"}
- curl:
```
curl --location --request POST "http://localhost:8888/v1/sqrt" \
--header "Content-Type: application/json" \
--data "{\"number\":\"144\"}"
```
3. FACTORIAL
- (n uint64)
- http://localhost:8888/v1/factorial
- JSON: {"number":"6"}
- curl:
```
curl --location --request POST "http://localhost:8888/v1/factorial" \
--header "Content-Type: application/json" \
--data "{\"number\":\"6\"}"
```
4. ISPRIME
- (n int)
- Function IsPrime returns bool, according to provided number (if it's prime or not).
- http://localhost:8888/v1/isPrime
- JSON: {"number":"6"}
- curl:
```
curl --location --request POST "http://localhost:8888/v1/isPrime" \
--header "Content-Type: application/json" \
--data "{\"number\":\"6\"}"
```
5. LOG
- (n float64)
- http://localhost:8888/v1/log
- JSON: {"number":"6"}
- curl:
```
curl --location --request POST "http://localhost:8888/v1/log" \
--header "Content-Type: application/json" \
--data "{\"number\":\"6\"}"
```
