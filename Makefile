run:
	@reflex -s -r '\.go$$' -- sh -c 'go run main.go'

templ:
	@reflex -r '\.templ$$' -- sh -c 'templ generate'