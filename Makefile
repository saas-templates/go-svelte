NAME:="go-svelte"
DIST_DIR:="dist"

all: clean tidy be

clean:
	@echo "Cleaning up old build..."
	@mkdir -p $(DIST_DIR)

tidy:
	@echo "Reformat and tidy Go..."
	@gofumpt -l -w .
	@go mod tidy -v

be:
	@echo "Running go generate..."
	@go generate ./...
	@echo "Building backend..."
	@go build -o $(DIST_DIR)/$(NAME) ./
