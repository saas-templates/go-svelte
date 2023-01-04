NAME:="go-svelte"
DIST_DIR:="dist"

all: clean fe tidy be

clean:
	@echo "Cleaning up old build..."
	@mkdir -p $(DIST_DIR)

tidy:
	@echo "Reformat and tidy Go..."
	@gofumpt -l -w .
	@go mod tidy -v

fe:
	@echo "Building frontend..."
	@cd ui && yarn && yarn build

be:
	@echo "Building backend..."
	@go build -o $(DIST_DIR)/$(NAME) ./
