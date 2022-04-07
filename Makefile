NAME=fetch
GO = go
BUILD = $(GO) build
FMT = $(GO) fmt
TEST = $(GO) test

$(NAME): build

.PHONY: build
build:
	$(BUILD) -o $(NAME)

.PHONY: fmt
fmt:
	$(FMT) .

.PHONY: test
test:
	$(TEST)
