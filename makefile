run: front
	cd backend && go run .

build: front
	cd backend && go build

front:
	cd frontend && npm run build

fmt:
	cd backend && gofmt -l -s -w .

.PHONY: run front build fmt
