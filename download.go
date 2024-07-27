package goutil

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"
)

type progressWriter struct {
	onProgress  func(totalSize int64, processSize int64)
	totalSize   atomic.Int64
	processSize atomic.Int64
	lock        sync.Mutex
}

func (p *progressWriter) Reset() {
	p.totalSize.Store(0)
	p.processSize.Store(0)
}
func (p *progressWriter) SetTotal(totalSize int64) {
	p.totalSize.Store(totalSize)
}
func (p *progressWriter) Add(size int64) {
	p.processSize.Add(size)
}
func (p *progressWriter) Write(dat []byte) (n int, err error) {
	n = len(dat)
	p.processSize.Add(int64(n))
	if p.onProgress != nil {
		p.onProgress(p.totalSize.Load(), p.processSize.Load())
	}
	return
}

type Downloader struct {
	concurrency int
	resume      bool
	progress    *progressWriter
}

func NewDownloader(concurrency int, resume bool) *Downloader {
	return &Downloader{concurrency: concurrency, resume: resume}
}

func (d *Downloader) Download(strURL, filename string, onProgress func(totalSize int64, processSize int64)) error {

	resp, err := http.Head(strURL)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 && resp.StatusCode <= 399 {
		strURL = resp.Header.Get("Location")
		resp, err = http.Head(strURL)
		if err != nil {
			return err
		}
	}
	if filename == "" {
		u, _ := url.Parse(strURL)
		filename = filepath.Base(u.Path)
	}
	if filename == "" {
		contentType := resp.Header.Get("Accept-Ranges")
		panic(errors.New(contentType)) //todo
	}

	d.progress = &progressWriter{onProgress: onProgress}
	if d.concurrency > 1 && resp.StatusCode == http.StatusOK && resp.Header.Get("Accept-Ranges") == "bytes" {
		return d.multiDownload(strURL, filename, int(resp.ContentLength))
	}

	return d.singleDownload(strURL, filename)
}

func (d *Downloader) multiDownload(strURL, filename string, contentLen int) error {
	d.progress.SetTotal(int64(contentLen))
	tempFilename := fmt.Sprintf("%d", time.Now().Unix())
	partSize := contentLen / d.concurrency

	var wg sync.WaitGroup
	wg.Add(d.concurrency)

	rangeStart := 0

	for i := 0; i < d.concurrency; i++ {
		go func(i, rangeStart int) {
			defer wg.Done()

			rangeEnd := rangeStart + partSize
			// 最后一部分，总长度不能超过 ContentLength
			if i == d.concurrency-1 {
				rangeEnd = contentLen
			}

			downloaded := 0
			partFileName := d.getPartFilename(tempFilename, i)
			if d.resume {
				content, err := os.ReadFile(partFileName)
				if err == nil {
					downloaded = len(content)
				}
				d.progress.Add(int64(downloaded))
			}

			d.downloadPartial(strURL, partFileName, rangeStart+downloaded, rangeEnd, i)

		}(i, rangeStart)

		rangeStart += partSize + 1
	}

	wg.Wait()

	d.merge(filename, tempFilename)

	return nil
}

func (d *Downloader) downloadPartial(strURL, partFilename string, rangeStart, rangeEnd, i int) {
	if rangeStart >= rangeEnd {
		return
	}

	req, err := http.NewRequest("GET", strURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", rangeStart, rangeEnd))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	flags := os.O_CREATE | os.O_WRONLY
	if d.resume {
		flags |= os.O_APPEND
	}

	partFile, err := os.OpenFile(partFilename, flags, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer partFile.Close()

	buf := make([]byte, 32*1024)
	_, err = io.CopyBuffer(io.MultiWriter(partFile, d.progress), resp.Body, buf)
	if err != nil {
		if err == io.EOF {
			return
		}
		log.Fatal(err)
	}
}

func (d *Downloader) merge(filename, tempFilename string) error {
	destFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer destFile.Close()

	for i := 0; i < d.concurrency; i++ {
		partFileName := d.getPartFilename(tempFilename, i)
		partFile, err := os.Open(partFileName)
		if err != nil {
			return err
		}
		io.Copy(destFile, partFile)
		partFile.Close()
		os.Remove(partFileName)
	}

	return nil
}

func (d *Downloader) getPartFilename(filename string, partNum int) string {
	return fmt.Sprintf("%s/%s-%d", os.TempDir(), filepath.Base(filename), partNum)
}

func (d *Downloader) singleDownload(strURL, filename string) error {
	resp, err := http.Get(strURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	d.progress.SetTotal(int64(resp.ContentLength))

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, 32*1024)
	_, err = io.CopyBuffer(io.MultiWriter(f, d.progress), resp.Body, buf)
	return err
}
