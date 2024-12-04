run:
	@reflex -s -r '\.go$$' -- sh -c 'go run ./cmd/main.go'

templ:
	@reflex -r '\.templ$$' -- sh -c 'templ generate'

scss:
	@reflex -s -r 'resources/.*\.(scss|js)$$' -- sh -c 'npm install --prefix resources && npm run dev --prefix resources'