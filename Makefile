# build will build the main application
build:

	@ echo 😆 Bundling application...
	go build -o ./dist/app ./app/main.go
	@ echo App bundled and app ready to run! 😎

start: build 
	@ echo starting app 
	./dist/app  

test: 
	@ echo Starting tests
	go test -v 