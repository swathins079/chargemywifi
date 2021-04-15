
BIN_NAME="chargemywifi"

exe: clean
	env \
	GOOS=windows \
	GOARCH=amd64 \
	go build  \
	-ldflags "-X github.com/swathins079/chargemywifi/pkg/model.OperatingSystem=0" -o $(BIN_NAME).exe -v .

bin: clean
	go build \
	-ldflags "-X github.com/swathins079/chargemywifi/pkg/model.OperatingSystem=1" -o $(BIN_NAME) -v .

clean:
	rm -f $(BIN_NAME)*
