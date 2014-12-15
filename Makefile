build: go-dl-extract-i386

go-dl-extract-i386: Dockerfile go-dl-extract.go
	docker build -t go-dl-extract-builder .
	-docker rm go-dl-extract-builder || true
	docker run --name=go-dl-extract-builder go-dl-extract-builder true
	docker cp go-dl-extract-builder:/go/bin tmp
	docker rm go-dl-extract-builder
	touch tmp/bin/*
	mv tmp/bin/* .
	rm -rf tmp

clean:
	rm -f go-dl-extract-{i386,armel,armhf,amd64}
