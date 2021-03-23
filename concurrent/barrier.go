package barrier

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var timeoutMilliseconds int = 5000

// Entrance
func CaptureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	barrier(endpoints...)
	// MUST close at here !!!
	writer.Close()

	temp := <-out
	return temp
}

type barrierResp struct {
	Err  error
	Resp string
}

func barrier(endpoints ...string) {
	requestNumber := len(endpoints)
	in := make(chan barrierResp, requestNumber)
	defer close(in)

	responses := make([]barrierResp, requestNumber)
	// make request
	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}

	// read request
	var hasError bool = false
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err)
			hasError = true
			break
		}
		responses[i] = resp
	}

	// print
	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}

}

func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) * time.Millisecond),
	}
	// http Get url
	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	// read byte stream, convert to string, and write to out
	stream, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	res.Resp = string(stream)
	out <- res
}
