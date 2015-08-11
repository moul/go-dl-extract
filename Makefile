NAME = go-dl-extract
BUILDER = $(NAME)-builder
HOST_BIN = $(NAME)-$(shell uname -s)-$(shell uname -m)
TEST_BIN = $(NAME)-Linux-x86_64
SRCS = go-dl-extract.go

build:	dist/$(HOST_BIN)

.build:	Dockerfile $(SRCS)
	@echo
	@rm -f .build
	docker build -t $(BUILDER) .
	docker inspect -f '{{.Id}}' $(BUILDER) > .build

dist/$(TEST_BIN):	dist/$(HOST_BIN)
dist/$(HOST_BIN):	.build
	@echo
	@docker rm $(BUILDER) 2>/dev/null || true
	mkdir -p dist
	docker run --name=$(BUILDER) $(BUILDER) tar -cf - /etc/ssl > dist/ssl.tar
	docker cp $(BUILDER):/go/bin tmp
	docker rm $(BUILDER)
	touch tmp/bin/* && mv tmp/bin/* dist/
	rm -rf tmp

clean:

fclean:	clean
	rm -rf dist/

test:	dist/$(TEST_BIN)
	@echo
	cp -f dist/ssl.tar dist/$(TEST_BIN) tests/
	$(MAKE) -C tests/ test BINARY=$(TEST_BIN)

trusted:	trusted/Makefile dist/$(HOST_BIN)
	$(MAKE) -C trusted/ build
