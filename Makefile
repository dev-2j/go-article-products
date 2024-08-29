.PHONY: getall
getall: 
	@ go get -u github.com/dev-2j/libaryx
	@ go mod tidy 