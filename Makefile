test: clean kcc
	./test.sh

# pathは都度読み込む必要がある
kcc:
	PATH=/usr/local/go/bin:$(PATH) go build

clean:
	rm -f kcc *.out tmp*

.PHONY: test clean