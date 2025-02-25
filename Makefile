# Makefile for ubit-newsletter project

# Compiler and flags
GO = go
GOFLAGS = -v

# Directories
SRC_DIR = src
BIN_DIR = bin

# Target executable
TARGET = ./$(BIN_DIR)/newsletter

# Default target
all: $(TARGET)

# Link object files to create the executable
$(TARGET):
	@mkdir -p $(BIN_DIR)
	$(GO) build $(GOFLAGS) -o $@ $^

# Clean up build files
clean:
	rm -rf $(BIN_DIR)

# Run tests
test:
	$(GO) test $(SRC_DIR)/...

run:
	./$(BIN_DIR)/newsletter

.PHONY: all clean