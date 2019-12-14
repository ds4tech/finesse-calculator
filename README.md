# assignment for interview

Simple Go Calculator project.

run:
---

go build -o main cmd/calculator/handlers.go cmd/calculator/main.go
./main

http://localhost:8888/

---

docker build -t go-calc -f build/package/Dockerfile .
docker run -d -p 80:8888 go-calc

http://localhost/

---

kubectl create -f k8s-replicaSet.yml

http://192.168.99.100:31898/
