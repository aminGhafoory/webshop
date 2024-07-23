run: templ tailwind
	@go build -o ./bin/out && ./bin/out

templ:
	@templ generate


tailwind:
	@npx tailwindcss -i css/input.css -o public/style.css --minify

sqlc:
	@sqlc generate