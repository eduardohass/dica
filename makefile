.PHONY: default run build test docs clean
# Variables
APP_NAME=dica

# Tasks
default: run-with-docs

run:
	@go run main.go
run-with-docs:
	@swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs

# ALTER TABLE questions DELETE COLUMN question;
# ALTER TABLE questions RENAME COLUMN id_question TO idquestion;
# ALTER TABLE questions RENAME COLUMN option_d TO optiond;
# ALTER TABLE questions RENAME COLUMN option_i TO optioni;
# ALTER TABLE questions RENAME COLUMN option_c TO optionc;
# ALTER TABLE questions RENAME COLUMN option_a TO optiona;