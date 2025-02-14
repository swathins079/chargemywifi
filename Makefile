
BIN_NAME="chargemywifi"
BUILD_FOLDER="build"

create-build-folder:
	mkdir -p build

all: exe bin app

exe: clean-exe create-build-folder
	env \
	GOOS=windows \
	GOARCH=amd64 \
	go build  \
	-ldflags "-X github.com/swathinsankaran/chargemywifi/pkg/model.OperatingSystem=0" -o $(BUILD_FOLDER)/$(BIN_NAME).exe -v .

bin: clean-bin create-build-folder
	env \
	GOOS=linux \
	GOARCH=amd64 \
	go build \
	-ldflags "-X github.com/swathinsankaran/chargemywifi/pkg/model.OperatingSystem=1" -o $(BUILD_FOLDER)/$(BIN_NAME).bin -v .

app: clean-app create-build-folder
	env \
	GOOS=darwin \
	GOARCH=amd64 \
	go build \
	-ldflags "-X github.com/swathinsankaran/chargemywifi/pkg/model.OperatingSystem=2" -o $(BUILD_FOLDER)/$(BIN_NAME).app -v .

clean:
	rm -rf $(BUILD_FOLDER)
	rm -f coverage.out

clean-exe:
	rm -f $(BUILD_FOLDER)/$(BIN_NAME).exe

clean-bin:
	rm -f $(BUILD_FOLDER)/$(BIN_NAME).bin

clean-app:
	rm -f $(BUILD_FOLDER)/$(BIN_NAME).app

test:
	go test -coverprofile=coverage.out -v ./...
	go tool cover -func=coverage.out

coverage-html:
	go tool cover -html=coverage.out
