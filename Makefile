build: go-dl-extract-osx

go-dl-extract-osx: Dockerfile go-dl-extract.go
	# godep save
	docker build -t go-dl-extract-builder .
	-docker rm go-dl-extract-builder || true 2>/dev/null
	docker run --name=go-dl-extract-builder go-dl-extract-builder true
	docker cp go-dl-extract-builder:/go/bin tmp
	docker rm go-dl-extract-builder
	touch tmp/bin/*
	mv tmp/bin/* .
	rm -rf tmp

clean:
	rm -f go-dl-extract-{i386,armel,armhf,amd64}
