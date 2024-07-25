run: templ tailwind
	@go build -o ./bin/out && ./bin/out

templ:
	@templ generate


tailwind:
	@npx tailwindcss -i css/input.css -o public/style.css --minify

sqlc:
	@sqlc generate


gooseup:
	@cd models/sql/schema && goose postgres postgres://postgres:amin235711@amin-laptop.local:5432/webshop up

goosedown:
	@cd models/sql/schema && goose postgres postgres://postgres:amin235711@amin-laptop.local:5432/webshop down