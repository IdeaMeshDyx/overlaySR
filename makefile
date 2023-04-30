default:
	@echo "\nUSAGE: make [option] \n"
	@echo "Options:"
	@echo "For project:"
	@echo "	build : build client and server"
	@echo "	clean : clean all executable file"
	@echo "	runs : run the server"
	@echo "	runc : run the client"
	@echo "For git:"
	@echo "	gl : lookup git status and logs before commit"
	@echo "	gc : git add and commit"


gl:
	git status
	git log
gc: 
	git add .
	git commit

build:
	cd ./client && go build -o client.o ./cmd/agent/main.go 
	cd ..
	cd ./server/ && go build -o server.o ./cmd/server/main.go 
	cd ..

clean:
	rm  ./client/*.o
	rm  ./server/*.o

runs:
	sudo ./server/server.o

runc:	
	sudo ./client/client.o