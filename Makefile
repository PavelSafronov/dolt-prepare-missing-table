test:
	mysql -u root -e 'DROP DATABASE obsidian;'
	mysql -u root -e 'CREATE DATABASE obsidian;'
	go test . -count=1

PHONY: test