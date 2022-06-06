.PHONY: vendor
vendor:
	go mod tidy -compat=1.17
	go mod vendor
	git add go.sum go.mod
	git status
