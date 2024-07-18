# Define the executable name
EXECUTABLE := fileorganizer

# Define the source directory
SRC_DIR := .

# Define the destination directory
DEST_DIR := /usr/local/bin

# Define the Go build command
GO_BUILD := go build -o $(EXECUTABLE)

# Default target
all: build move

# Build the Go program
build:
	@echo "Building the Go executable"
	$(GO_BUILD) $(SRC_DIR)

# Move the executable to /usr/local/bin
move:
	@echo "Moving the executable to $(DEST_DIR)..."
	@sudo mv $(EXECUTABLE) $(DEST_DIR)

# Clean up the executable in the source directory
clean:
	@echo "Cleaning up..."
	@rm -f $(EXECUTABLE)

# PHONY targets
.PHONY: all build move clean
