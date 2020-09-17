package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong: status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)

	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// Peek 看的是 bofio.NewReader，对其中的 r 读 1024 个字节并存起来；
	// 但这里 r 已经被读了 1024，上面的 utf8Reader 是从第 1025 字节开始读
	// bytes, err := bufio.NewReader(r).Peek(1024)
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Fatal(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
