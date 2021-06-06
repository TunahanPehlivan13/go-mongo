build:
	docker build . -t go-mongo

run: build
	docker run -p 5000:5000 go-mongo