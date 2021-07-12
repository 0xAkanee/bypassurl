package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	helperURL string = "https://lp.nrmn.top/api/bypass?url="
)

var (
	serr *log.Logger = log.New(os.Stderr, "[ERROR] ",
		log.Lmsgprefix|log.LstdFlags|log.Lshortfile)
	sout *log.Logger = log.New(os.Stdout, "[INFO]  ",
		log.Lmsgprefix|log.LstdFlags)

	reURL = regexp.MustCompile("^https?://")
)

func showError(arg ...interface{}) int {
	serr.Print(arg...)
	return 1
}

func adjustURL(url string) string {
	trimmedURL := strings.TrimSpace(url)
	if !reURL.MatchString(trimmedURL) {
		trimmedURL = "https://" + trimmedURL
	}

	return base64.StdEncoding.EncodeToString([]byte(trimmedURL))
}

func getHelp(url string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet, helperURL+url, nil)

	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	isSuccess := result["success"].(bool)

	if !isSuccess {
		return result["err"].(map[string]interface{})["message"].(string), nil
	}

	return result["url"].(string), nil
}

func bypass(wg *sync.WaitGroup, url string) {
	adjustedURL := adjustURL(url)
	result, err := getHelp(adjustedURL)
	if err != nil {
		showError(err)
	}

	sout.Println(url, "=>", result)
	wg.Done()
}

func run(args ...string) int {
	n := len(args)
	if n == 0 {
		return showError("missing input URL")
	}

	var wg sync.WaitGroup
	wg.Add(n)

	for _, arg := range args {
		go bypass(&wg, arg)
	}

	wg.Wait()

	return 0
}

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println(`######################################
#         Bypass Shortlink           #
# Coded By Rndzx   fb.me/negevian.id #
######################################`)
	os.Exit(run(os.Args[1:]...))
}
