build:
	go build -o bin/mocyt ./main.go 
install:
	make build;
	cp ./bin/mocyt /usr/bin;
uninstall:
	rm /usr/bin/mocyt
	
	