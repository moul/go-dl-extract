NAME = go-dl-extract-test

test: build
	@echo
	docker run -it --rm $(NAME) /bin/ls

build:
	@echo
	docker build -t $(NAME) .
