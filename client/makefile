default:
	@echo "\nUSAGE: make [option] \n"
	@echo "Options:"
	@echo "	runc : run the client main.go script"
	@echo "	build : build client"
	@echo "	runt : go run test"
	@echo "	buildt : build and run test"
	@echo "	exec : run the client's executable file"
	@echo "	clean : clean all executable file"

build:
	go build -o client.o ./cmd/agent/main.go 

clean:
	rm  ./*.o

exec:	
	sudo ./client.o

runc:
	go run ./cmd/agent/main.go

runt:
	go run ./test/test.go

buildt:
	go build -o ./test/test.o ./test/test.go 
	sudo ./test/test.o