default:
	@echo "\nUSAGE: make [option] \n"
	@echo "Options:"
	@echo "	runc : run the client main.go script"
	@echo "	build : build client"
	@echo "	buildw : build client on windows"
	@echo "	clean : clean all executable file"

build:
	go build -o client.o ./cmd/agent/main.go 

buildw:
	go build -o client.exe ./cmd/agent/main.go 

clean:
	rm  ./*.o

runc:
	go run ./cmd/agent/main.go
