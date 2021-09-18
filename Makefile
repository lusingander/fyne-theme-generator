.PHONY: credit
credit: 
	go run github.com/lusingander/fyne-credits-generator/cmd/fyne-credits-generator@latest -package ui > internal/ui/credits.go

.PHONY: run
run:
	go run main.go
