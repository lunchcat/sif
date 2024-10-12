.POSIX:
.SUFFIXES:

GO ?= go
RM ?= rm
GOFLAGS ?=
PREFIX ?= /usr/local
BINDIR ?= bin

all: check_go_version sif
	@echo "✅ All tasks completed successfully! 🎉"

check_go_version:
	@echo "🔍 Checking Go version..."
	@$(GO) version | grep -q "go1\.23\." || (echo "❌ Error: Please install the latest version of Go" && exit 1)
	@echo "✅ Go version check passed!"

sif: check_go_version
	@echo "🛠️ Building sif..."
	$(GO) build $(GOFLAGS) ./cmd/sif
	@echo "✅ sif built successfully! 🚀"

clean:
	@echo "🧹 Cleaning up..."
	$(RM) -rf sif
	@echo "✨ Cleanup complete!"

install: check_go_version
	@echo "📦 Installing sif..."
	mkdir -p $(DESTDIR)$(PREFIX)/$(BINDIR)
	cp -f sif $(DESTDIR)$(PREFIX)/$(BINDIR)
	@echo "✅ sif installed successfully! 🎊"

uninstall:
	@echo "🗑️ Uninstalling sif..."
	$(RM) $(DESTDIR)$(PREFIX)/$(BINDIR)/sif
	@echo "✅ sif uninstalled successfully!"

.PHONY: all check_go_version sif clean install uninstall