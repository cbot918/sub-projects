# run: tt
# 	./tt
# 	rm tt

clean:
	docker stop ci-dock-testy
	docker container rm ci-dock-testy

build: main.go
	go build -o tt main.go

tt: main.go
	go build -o tt main.go

.PHONY: build run tt clean
.SILENT: build run tt clean
