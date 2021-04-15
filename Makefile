
BIN_NAME="chargemywifi"

exe:
	env GOOS=windows GOARCH=amd64 go build -o $(BIN_NAME).exe -v .	

bin:
	go build -o $(BIN_NAME) -v .

clean:
	rm $(BIN_NAME)*
