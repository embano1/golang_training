REPO?=embano1/go-web
VERSION?=v2
TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

default: builddocker

buildgo:
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-w' -o ws .

builddocker: buildgo
	docker build -t ${REPO}${VERSION}:${TAG} .
	docker tag ${REPO}${VERSION}:${TAG} ${REPO}${VERSION}:latest 

pushdocker:
	docker push ${REPO}${VERSION}:${TAG}
	docker push ${REPO}${VERSION}:latest

cleanup:
	rm ws1

release: buildgo builddocker pushdocker cleanup
