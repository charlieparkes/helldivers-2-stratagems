.PHONY: build
build:
	go build -o ./bin/stratagems ./cmd/stratagems
	GOOS=windows go build -o ./bin/stratagems.exe ./cmd/stratagems
