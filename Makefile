build: go-dl-extract-i386

go-dl-extract-i386: Dockerfile wrapper.go
	docker build -t go-dl-extract-builder .
	docker run --name=go-dl-extract-builder go-dl-extract-builder true
	docker cp go-dl-extract-builder:/gi/bin tmp
	docker rm go-dl-extract-builder
	touch tmp/bin/*
	mv tmp/bin/* .
	rm -rf tmp

clean:
	rm -f go-dl-extract-{i386,armel,armhf,amd64}
