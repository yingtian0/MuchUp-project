.PHONY: gen lint breaking clean

gen:
	buf generate

lint:
	buf lint

breaking:
	buf breaking --against '.git#branch=main'

clean:
	rm -rf gen/*