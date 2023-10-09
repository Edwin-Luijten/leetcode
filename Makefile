build-parser:
	go build -o benchmark-parser ./cmd/parser/bench_results_parser.go && \
	chmod +x benchmark-parser

bench:
	go test -bench=. -benchmem ./... | gobenchdata --json bench.json \

parse:
	./benchmark-parser

run: bench parse

# Generate a new package for a challenge
safeName = $(shell echo $(challenge) | sed -r 's/ /_/g')
pkg = $(safeName)
method = $(shell echo $(safeName) | sed -r 's/_([a-z])/\U\1/g')
methodPart = $(shell echo $(safeName) | sed -r 's/(^|_)([a-z])/\U\2/g')

new:
	mkdir -p ./$(safeName) && \
	cp -r ./.templates/* ./$(safeName) && \
	sed -i "s/challenge_pkg/${pkg}/g" ./$(safeName)/function.go
	sed -i "s/challenge_pkg/${pkg}/g" ./$(safeName)/function_test.go
	sed -i "s/\(challenge*\)/${method}/g" ./$(safeName)/function.go
	sed -i "s/\(challenge*\)/${method}/g" ./$(safeName)/function_test.go
	sed -i "s/\(Challenge*\)/${methodPart}/g" ./$(safeName)/function_test.go
	mv ./$(safeName)/function.go ./$(safeName)/$(safeName).go
	mv ./$(safeName)/function_test.go ./$(safeName)/$(safeName)_test.go