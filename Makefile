.PHONY: plugin

plugin:
	cd plugin; go build -o ~/.config/octant/plugins/terminator .

