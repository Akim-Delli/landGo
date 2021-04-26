package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type result struct {
	path  string
	match bool
	err   error
}

func main() {
	sigs, err := parseSignatureFile("./nasa-logs/md5sum.txt")
	if err != nil {
		log.Fatalf("error parsing signature file: %s", err)
	}

	out := make(chan *result)
	for path, sig := range sigs {
		go md5worker("./nasa-logs/"+path, sig, out)
	}

	ok := true
	for range sigs {
		r := <-out
		switch {
		case r.err != nil:
			log.Printf("%s: error - %s\n", r.path, r.err)
			ok = false
		case !r.match:
			log.Printf("%s: signature mismatch\n", r.path)
			ok = false

		}
	}

	if !ok {
		os.Exit(1)
	}
}

func md5worker(path string, sig string, out chan *result) {
	r := &result{path: path}
	s, err := fileMD5(path)
	if err != nil {
		r.err = err
		out <- r
		return
	}
	r.match = (s == sig)
	out <- r
}

func fileMD5(path string) (string, error) {

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func parseSignatureFile(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	sigs := make(map[string]string)
	scanner := bufio.NewScanner(f)
	for lnum := 1; scanner.Scan(); lnum++ {

		// 6c6427da7893932731901035edbb9214 nasa-00.log
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			return nil, fmt.Errorf("%s:%d bad line", path, lnum)
		}
		sigs[fields[1]] = fields[0]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sigs, nil
}
