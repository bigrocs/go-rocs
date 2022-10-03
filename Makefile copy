.PHONY: git
git:
	git add .
	git commit -m"自动提交 git 代码"
	git push
.PHONY: tag
tag:
	git push --tags
.PHONY: sql2pb
sql2pb:
	sql2pb -go_package ./pb -host localhost -package pb -password 123456 -port 3306 -schema user -service_name user -user root > ./user.proto
.PHONY: rpc
rpc:
	goctl rpc protoc ./*.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
.PHONY: model
model:
	goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/user" -table="users" -c -dir ./types/model
# .PHONY: api
# api:
# 	goctl api go -api ./order/api/order.api -dir ./order/api
.PHONY: docker
docker:
	docker build -f Dockerfile  -t user-rpc .
.PHONY: up
up:
	docker-compose up
.PHONY: dockerfile
dockerfile:
	goctl docker -go ./user.go
.PHONY: test
test:
	go test main_test.go -test.v
.PHONY: run
run:
	go run user.go