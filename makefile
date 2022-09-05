default: auto

run:
	go build -o test.exe
	test.exe 

auto:
	autocmd -v -t ".*\.go" -t example.yml go test -v 

autorun:
	autocmd -v -t ".*\.go" -t example.yml -t makefile make run

