test:
	@printf '>>> Running tests on host...\n'
	@shellcheck ./src/*
	@bats ./src/test.bats
	@printf '>>> Building tester container image...\n'
	@docker build -f Containerfile -t bashpack-tester:latest .
	@printf '>>> Running tests in container, as non-root user...\n'
	@docker run --rm --user nonroot bashpack-tester:latest
	@printf '>>> Running tests in container, as root user...\n'
	@docker run --rm --user root bashpack-tester:latest

ci-local:
	@go run github.com/nektos/act@latest
