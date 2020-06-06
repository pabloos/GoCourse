# The current makefile contains all the mainstream tasks for a go project -> run, build, test and test coverage

# conts
TEST_TIMEOUT = 30s
COVER_PROFILE = coverage.out

# cmds
TEST = go test ./... -v -race -timeout ${TEST_TIMEOUT}
COVER = go tool cover

run:
	go run .

build:
	go build -v ./...

test:
	${TEST}

clean-test-cache:
	echo "cleaning test cache..."
	go clean -testcache

coverage: clean-test-cache
	${TEST} -covermode=atomic -coverprofile=${COVER_PROFILE}
	${COVER} -func ${COVER_PROFILE}
	${COVER} -html=${COVER_PROFILE}
	rm ${COVER_PROFILE}