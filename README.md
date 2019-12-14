# assignment for interview

Simple Go Calculator project.

run:
---

go build -o calcWs
./calcWs

http://localhost:8888/

---

docker build -t go-calc .
docker run -d -p 80:8888 go-calc

http://localhost/

---

kubectl create -f k8s-replicaSet.yml

http://192.168.99.100:31898/
