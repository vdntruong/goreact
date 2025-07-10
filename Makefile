# Makefile for Go-React Project

# Define directories and executable name
.PHONY: all build-and-run setup build-ui build-be run-be clean

UI_DIR := ui
BE_DIR := be
BUILD_TARGET := goreact # Name of your Go executable (or whatever you prefer)

# --- Main Targets ---

# Default target (runs build-and-run)
all: build-and-run

# Full build and run sequence
build-and-run: setup build-ui build-be run-be

# --- Individual Build Steps ---

# Install Node.js dependencies for the UI
setup:
	@echo "--- Installing UI dependencies (npm install) ---"
	npm install --prefix $(UI_DIR)

# Build the React frontend
# Parcel is now executed from the project root (where Makefile is).
# Paths are therefore relative to the project root.
build-ui:
	@echo "--- Building UI (Parcel build) ---"
	# CORRECTED: Call Parcel executable located inside $(UI_DIR)/node_modules
	$(CURDIR)/$(UI_DIR)/node_modules/.bin/parcel build $(UI_DIR)/index.html --public-url /static/ --dist-dir $(BE_DIR)/static

# Build the Go backend executable
build-be:
	@echo "--- Building Backend (Go build) ---"
	cd $(BE_DIR) && go build -o $(BUILD_TARGET) .

# Run the Go backend server
run-be:
	@echo "--- Running Backend Server ---"
	cd $(BE_DIR) && ./$(BUILD_TARGET)

# --- Clean Target ---

# Clean up generated files and caches
clean:
	@echo "--- Cleaning project artifacts ---"
	rm -rf $(UI_DIR)/node_modules
	rm -rf $(UI_DIR)/.parcel-cache
	rm -rf $(BE_DIR)/static
	rm -f $(BE_DIR)/$(BUILD_TARGET)
	@echo "Clean complete."
