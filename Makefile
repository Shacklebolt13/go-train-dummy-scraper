clean:
	@echo "Cleaning up..."
	@if exist out rmdir /s /q out 2>nul || rm -rf ./out 2>/dev/null || true

dummy: clean
	@echo "Building dummy..."
	@go build -o out/ scraper/cmd/dummy
	@echo "Done."