run:
	@reflex -s -r '\.go$$' -- sh -c 'go run ./cmd/main.go'

templ:
	@reflex -r '\.templ$$' -- sh -c 'templ generate'

scss:
	@reflex -s -r '\.scss$$' -- sh -c 'npm run build --prefix resources'