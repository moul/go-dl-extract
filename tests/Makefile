NAME = go-dl-extract-test

test: build

build:
	docker build -t $(NAME) .
	docker run -it --rm $(NAME) /bin/ls
