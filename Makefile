.PHONY: deploy list help

help:
	@echo "Protohackers Deployment Helper"
	@echo ""
	@echo "Usage:"
	@echo "  make deploy TASK=0 LANG=go        - Deploy task 0 (Go implementation)"
	@echo "  make deploy TASK=0 LANG=elixir    - Deploy task 0 (Elixir implementation)"
	@echo "  make deploy TASK=1 LANG=rust      - Deploy task 1 (Rust implementation)"
	@echo "  make list                         - List all available tasks"
	@echo ""
	@echo "Default LANG is 'go' if not specified"
	@echo ""
	@echo "Examples:"
	@echo "  make deploy TASK=0              # deploys Go version"
	@echo "  make deploy TASK=0 LANG=elixir  # deploys Elixir version"

list:
	@echo "Available tasks and implementations:"
	@for task_dir in $$(find . -maxdepth 1 -type d -name "*:*" | sort); do \
		task_name=$$(basename "$$task_dir"); \
		echo "$$task_name:"; \
		for lang_dir in $$task_dir/*/; do \
			if [ -d "$$lang_dir" ]; then \
				lang=$$(basename "$$lang_dir"); \
				if [ -f "$$lang_dir/Dockerfile" ]; then \
					echo "  ✓ $$lang"; \
				else \
					echo "  - $$lang (no Dockerfile)"; \
				fi; \
			fi; \
		done; \
	done

deploy:
	@if [ -z "$(TASK)" ]; then \
		echo "Error: TASK parameter is required"; \
		echo "Usage: make deploy TASK=0 LANG=go"; \
		exit 1; \
	fi; \
	LANG=$${LANG:-go}; \
	if [ -n "$(LANG)" ]; then \
		LANG=$(LANG); \
	fi; \
	TASK_DIR=$$(find . -maxdepth 1 -type d -name "$(TASK):*" | head -n 1); \
	if [ -z "$$TASK_DIR" ]; then \
		echo "Error: Task $(TASK) not found"; \
		echo "Available tasks:"; \
		find . -maxdepth 1 -type d -name "*:*" | sort; \
		exit 1; \
	fi; \
	DOCKERFILE="$$TASK_DIR/$$LANG/Dockerfile"; \
	if [ ! -f "$$DOCKERFILE" ]; then \
		echo "Error: Dockerfile not found at $$DOCKERFILE"; \
		echo "Available implementations:"; \
		ls -d $$TASK_DIR/*/ 2>/dev/null | xargs -n1 basename; \
		exit 1; \
	fi; \
	echo "Deploying task $(TASK) [$$LANG]: $$TASK_DIR/$$LANG"; \
	fly deploy --config fly.toml --dockerfile "$$DOCKERFILE" --build-arg GO_VERSION=1.24.10;