# Define the executable name
EXECUTABLE := fileorganizer

# Define the source directory
SRC_DIR := .

# Define the destination directory
DEST_DIR := /usr/local/bin

# Define the Go build command
GO_BUILD := go build -o $(EXECUTABLE)

# Define Git commit message
COMMIT_MSG := "Automated commit"

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

# Add files to git
git-add:
	@echo "Adding files to Git..."
	@git add .

# Commit changes to git
git-commit:
	@echo "committing changes to Git..."
	@git commit -m "$(COMMIT_MSG)"

# Push changes to remote repository
git-push:
	@echo "Pushing changes to remote repository"
	@git push

# Git operations combined
git: git-add git-commit git-push


# PHONY targets
.PHONY: all build move clean git-add git-commit git-push git
