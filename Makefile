.PHONY: all generate-grpc-task-server

all: generate-grpc-task-server 

generate-grpc-task-server:
	@echo "Generating grpc server for scraper tasks"
	@./scripts/task-proto.sh