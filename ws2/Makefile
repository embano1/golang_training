default: builddocker

buildgo:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ws2 .

builddocker: buildgo
	docker build -t localhost:5000/ws2 .

pushdocker:
	docker push localhost:5000/ws2

cleanup:
	rm ws2

buildrelease: buildgo builddocker pushdocker cleanup
