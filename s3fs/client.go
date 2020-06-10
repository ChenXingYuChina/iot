package s3fs

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"log"
	"net/url"
)

const (
	region = "us-east-1"
	bucket = "myteststorge"
)
var (
	Id     = "ASIA2EAE5YGHCUDXM36G"
	Secret = "Tm2g7RF0l8xY5V2YkfVo/9+vZHMFm4maj7pCJy9q"
	Token  = "FwoGZXIvYXdzELr//////////wEaDGUNw+5cUCIZlkRliyLNATThmRlfNXvfj5Cp2J4WAogQ3t7w52mPD4IXf67H8Sc9Bs7KZq6zGXTjwOC1t/GJuEmHvs/MzBMki+61lpqWRRoTUD6P7AlyBYRmrAvJGEdsNqsCdpLRaNmomHV7rpxbdPRomIo3iRTWe50Z3yzHVVlHzdzWPTxA4Oc+/122Pc9+mor2pRSX1bfpdVZoVpKvH1oYhvGshgz2v4lOnhiLVk9il7OkLf6M1nzknTAwFn8x4+A4bYHrhvyG/aumHtEu6PpV8wtqlEHAWAd2m5Yov9KA9wUyLQ0iXQ9PtlW1bd5VTAecQOWUEpFr+KTVlAsqumf17aH3lOeBMterqg6t32kk3Q=="
)

var fs *s3.S3

func Start() {
	s, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			Id,
			Secret,
			Token),
		Region: aws.String(region),
	})
	if err != nil {
		log.Println(err)
		return
	}
	fs = s3.New(s)
}

func ListFiles(prefix string, startAfter string, keyNumber int64) (*FileList, error) {
	o, err := fs.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket:     aws.String(bucket),
		MaxKeys:    aws.Int64(keyNumber),
		Prefix:     aws.String(prefix),
		StartAfter: aws.String(startAfter),
		Delimiter:  aws.String("/"),
	})
	if err != nil {
		return nil, err
	}
	//fmt.Println(o.Contents)
	return getFileList(o), nil
}

type FileList struct {
	Files []string `json:"files"`
	Next  string `json:"next,omitempty"`
}

func getFileList(o *s3.ListObjectsV2Output) *FileList {
	goal := &FileList{}
	for _, v := range o.CommonPrefixes {
		goal.Files = append(goal.Files, *(v.Prefix))
	}
	for _, v := range o.Contents {
		fmt.Println(*(v.Key))
		goal.Files = append(goal.Files, *(v.Key))
	}
	if *o.IsTruncated {
		goal.Next = "/api/ls/"+ *o.Prefix + "?start_after="+url.QueryEscape(goal.Files[len(goal.Files) - 1])
	}
	return goal
}

type Object interface {
	io.Reader
	io.Closer
	io.Seeker
	io.ReaderAt
}

func PutFile(f Object, filename string) error {
	_, err := fs.PutObject(&s3.PutObjectInput{
		Body:                      f,
		Bucket:                    aws.String(bucket),
		Key:                       aws.String(filename),
	})
	if err != nil {
		return err
	}
	return f.Close()
}

func DownloadFile(filename string) (io.ReadCloser, error) {
	o, err := fs.GetObject(&s3.GetObjectInput{
		Bucket:                     aws.String(bucket),
		Key:                        aws.String(filename),
	})
	return o.Body, err
}

func DeleteFile(filename string) error {
	_, err := fs.DeleteObject(&s3.DeleteObjectInput{
		Bucket:                    aws.String(bucket),
		Key:                       aws.String(filename),
	})
	return err
}
