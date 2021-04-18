
BIN_NAME="chargemywifi"

exe: clean-exe
	env \
	GOOS=windows \
	GOARCH=amd64 \
	go build  \
	-ldflags "-X github.com/swathins079/chargemywifi/pkg/model.OperatingSystem=0" -o $(BIN_NAME).exe -v .

bin: clean-bin
	go build \
	-ldflags "-X github.com/swathins079/chargemywifi/pkg/model.OperatingSystem=1" -o $(BIN_NAME) -v .

clean: clean-exe clean-bin
	rm -f coverage.out

clean-exe:
	rm -f $(BIN_NAME).exe

clean-bin:
	rm -f $(BIN_NAME)

test:
	go test -coverprofile=coverage.out -v ./...
	go tool cover -func=coverage.out

coverage-html:
	go tool cover -html=coverage.out
