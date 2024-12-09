lint:
	golangci-lint run ./... --verbose --fix


.PHONY: pre-build
pre-build:
	mkdir deployments || true
	mkdir deployments/tmp || true
	mkdir deployments/tmp/libs || true
	cp -r ../../libs deployments/tmp

	mkdir deployments/tmp/services || true
	rsync -av --progress ../../services deployments/tmp --exclude /services/user-ms/deployments

	cp ../../go.work ../../go.work.sum deployments/tmp


.PHONY: build
build:
	docker build -t theg-user-ms -f deployments/Dockerfile deployments/tmp/


.PHONY: post-build
post-build:
	rm -rf deployments/tmp


.PHONY: bui
bui: pre-build build post-build


.PHONY: start
start:
	docker run -it --rm -p 8082:8080 --name user-ms theg-user-ms
