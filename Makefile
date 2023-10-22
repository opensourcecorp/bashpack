test:
	@shellcheck ./src/*
	@bats ./src/test.bats

# Requires `act`: https://github.com/nektos/act
ci-local:
	@act
