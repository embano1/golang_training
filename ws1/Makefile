default: builddocker

buildgo:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ws1 .

builddocker: buildgo
	docker build -t localhost:5000/ws1 .

pushdocker:
	docker push localhost:5000/ws1

cleanup:
	rm ws1

buildrelease: buildgo builddocker pushdocker cleanup
