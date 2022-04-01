run:
	go run main.go
	
models:
	$(call print-target)
	swagger generate model -m models -f ./book.yaml -t ./app/

generate: models