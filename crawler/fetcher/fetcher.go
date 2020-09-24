package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler_distributed/config"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36"

var rateLimiter = time.Tick(time.Second / config.Qps)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	log.Printf("Fetching url %s ", url)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("User-Agent", UserAgent)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
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
