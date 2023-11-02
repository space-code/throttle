bootstrap: brew hook

hook:
	pre-commit install

brew:
	brew bundle check || brew bundle install

PHONY: brew hook