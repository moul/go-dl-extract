NAME = go-dl-extract
BUILDER = $(NAME)-builder
HOST_BIN = $(NAME)-$(shell uname -s)-$(shell uname -m)
TEST_BIN = $(NAME)-Linux-x86_64
SRCS = go-dl-extract.go

build:	dist/$(HOST_BIN)

.build:	Dockerfile Godeps
	@echo
	@rm -f .build
	docker build -t $(BUILDER) .
	docker inspect -f '{{.Id}}' $(BUILDER) > .build

Godeps:	$(SRCS)
	# godep save
	touch Godeps

dist/$(TEST_BIN):	dist/$(HOST_BIN)
dist/$(HOST_BIN):	.build
	@echo
	@docker rm $(BUILDER) 2>/dev/null || true
	docker run --name=$(BUILDER) $(BUILDER) true
	docker cp $(BUILDER):/go/bin tmp
	docker rm $(BUILDER)
	mkdir -p dist && touch tmp/bin/* && mv tmp/bin/* dist/ && rm -rf tmp && rm -f dist/godep

clean:
	rm -rf Godeps/

fclean:	clean
	rm -rf dist/

test:	dist/$(TEST_BIN)
	@echo
	cp dist/$(TEST_BIN) tests/
	$(MAKE) -C tests/ test BINARY=$(TEST_BIN)
