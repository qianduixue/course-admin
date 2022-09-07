package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/opisnoeasy/course-service/global"
	"github.com/opisnoeasy/course-service/model/web"
	"github.com/opisnoeasy/course-service/rk"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	maxRetries  = 3
	maxPartSize = int64(5 * 1024 * 1024)
)

// newSession Create S3 session
func newSession() *session.Session {
	s, _ := session.NewSession(&aws.Config{
		Region:           aws.String(global.GVA_CONFIG.AwsS3.Region),
		Endpoint:         aws.String(global.GVA_CONFIG.AwsS3.Endpoint),
		S3ForcePathStyle: aws.Bool(global.GVA_CONFIG.AwsS3.S3ForcePathStyle),
		DisableSSL:       aws.Bool(global.GVA_CONFIG.AwsS3.DisableSSL),
		Credentials: credentials.NewStaticCredentials(
			global.GVA_CONFIG.AwsS3.SecretID,
			global.GVA_CONFIG.AwsS3.SecretKey,
			"",
		),
	})
	return s
}

func CreateMultipartUpload(fileName string, svc *s3.S3) (*s3.CreateMultipartUploadOutput, error) {
	input := &s3.CreateMultipartUploadInput{
		Bucket: aws.String(global.GVA_CONFIG.AwsS3.Bucket),
		Key:    aws.String(fileName),
	}
	result, err := svc.CreateMultipartUpload(input)
	err = AwsReturnErr(err)
	if err != nil {
		return nil, err
	}
	return result, err
}

func GetFileName(filePath string) string {
	fileKey := fmt.Sprintf("%d%d", time.Now().Unix(), GetRangeNum(4))
	suffix := path.Ext(filePath)
	var build strings.Builder
	build.WriteString(global.GVA_CONFIG.AwsS3.PathPrefix)
	build.WriteString("/")
	build.WriteString(fileKey)
	build.WriteString(suffix)
	fileName := build.String()
	return fileName
}

//UploadPart 上传分段
func UploadPart(svc *s3.S3, resp *s3.CreateMultipartUploadOutput, partNumber int, fileBody []byte) (*s3.CompletedPart, error) {
	tryNum := 1
	input := &s3.UploadPartInput{
		Body:          bytes.NewReader(fileBody), // 文件流
		Bucket:        aws.String(global.GVA_CONFIG.AwsS3.Bucket),
		Key:           resp.Key,
		PartNumber:    aws.Int64(int64(partNumber)),
		UploadId:      resp.UploadId,
		ContentLength: aws.Int64(int64(len(fileBody))),
	}
	for tryNum <= maxRetries {
		result, err := svc.UploadPart(input)
		if err != nil {
			if tryNum == maxRetries {
				if er, ok := err.(awserr.Error); ok {
					return nil, er
				}
				return nil, err
			}
			tryNum++
		} else {
			return &s3.CompletedPart{
				ETag:       result.ETag,
				PartNumber: aws.Int64(int64(partNumber)),
			}, nil
		}
	}
	return nil, nil
}

//UploadPart1 分段上传
func UploadPart1(resp *s3.CreateMultipartUploadOutput, partNumber int, fileBody []byte) (*s3.CompletedPart, error) {
	svc := s3.New(newSession())
	input := &s3.UploadPartInput{
		Body:          bytes.NewReader(fileBody), // 文件流
		Bucket:        aws.String(global.GVA_CONFIG.AwsS3.Bucket),
		Key:           resp.Key,
		PartNumber:    aws.Int64(int64(partNumber)),
		UploadId:      resp.UploadId,
		ContentLength: aws.Int64(int64(len(fileBody))),
	}
	result, err := svc.UploadPart(input)
	if err != nil {
		return nil, err
	}
	return &s3.CompletedPart{
		ETag:       result.ETag,
		PartNumber: aws.Int64(int64(partNumber)),
	}, nil
}

//CompleteMultipartUpload 完成分段上传
func CompleteMultipartUpload(svc *s3.S3, resp *s3.CreateMultipartUploadOutput, parts []*s3.CompletedPart) (*s3.CompleteMultipartUploadOutput, error) {
	input := &s3.CompleteMultipartUploadInput{
		Bucket: aws.String(global.GVA_CONFIG.AwsS3.Bucket),
		Key:    resp.Key,
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: parts,
		},
		UploadId: resp.UploadId,
	}
	result, err := svc.CompleteMultipartUpload(input)
	return result, err
}

func ListParts(uploadId, key string) (*s3.ListPartsOutput, error) {
	svc := s3.New(newSession())
	input := &s3.ListPartsInput{
		UploadId: aws.String(uploadId),
		Bucket:   aws.String(global.GVA_CONFIG.AwsS3.Bucket),
		Key:      aws.String(key),
	}
	result, err := svc.ListParts(input)
	err = AwsReturnErr(err)
	if err != nil {
		return nil, err
	}
	return result, err
}

func AbortMultipartUpload(svc *s3.S3, resp *s3.CreateMultipartUploadOutput) error {
	fmt.Println("Aborting multipart upload for UploadId#" + *resp.UploadId)
	abortInput := &s3.AbortMultipartUploadInput{
		Bucket:   resp.Bucket,
		Key:      resp.Key,
		UploadId: resp.UploadId,
	}
	_, err := svc.AbortMultipartUpload(abortInput)
	return err
}

func AwsReturnErr(err error) error {
	if err != nil {
		if arr, ok := err.(awserr.Error); ok {
			switch arr.Code() {
			default:
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

//AwsMultipartUpload 分段上传
func AwsMultipartUpload(filePath string) (path string, err error) {
	svc := s3.New(newSession())
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	fileName := GetFileName(filePath)
	resp, err := CreateMultipartUpload(fileName, svc)
	if err != nil {
		global.GVA_LOG.Error("启动分段上传失败", zap.Error(err))
		return "", err
	}
	readFile, err := ioutil.ReadFile(filePath) // 读取文件大小
	if err != nil {
		global.GVA_LOG.Error("读取文件大小失败", zap.Error(err))
		return "", err
	}
	size := len(readFile)
	partSize := 1024 * 1024 * 5
	times := size / partSize
	if size%partSize > 0 {
		times += 1
	}
	var wg sync.WaitGroup
	var completedParts []*s3.CompletedPart
	dataChan := make(chan []byte)
	go func() {
		for i := 0; i < times; i++ {
			data := make([]byte, partSize)
			n, _ := f.Read(data)
			if n == 0 {
				return
			}
			dataChan <- data
		}
		close(dataChan)
	}()
	partNumber := 1
	for i := 0; i < times; i++ {
		select {
		case data, ok := <-dataChan:
			if !ok {
				return
			}
			wg.Add(1)
			go func(number int) {
				part, _ := UploadPart(svc, resp, partNumber, data)
				defer wg.Done()
				completedParts = append(completedParts, part)
			}(partNumber)
			partNumber += 1
		}
	}
	wg.Wait() // 全部上传完毕
	//处理排序
	sort.SliceStable(completedParts, func(i, j int) bool {
		return *completedParts[i].PartNumber < *completedParts[j].PartNumber
	})
	result, err := CompleteMultipartUpload(svc, resp, completedParts)
	if err != nil {
		global.GVA_LOG.Error("完成上传失败", zap.Error(err))
		return "", err
	}
	return *result.Location, nil
}

//CreateVideoPartRecord 创建分段视频记录
func CreateVideoPartRecord(info []*web.VideoPartInfo) error {
	err := global.GVA_DB.CreateInBatches(&info, len(info)).Error
	if err != nil {
		global.GVA_LOG.Error("插入分段视频记录失败", zap.Error(err))
		return err
	}
	return nil
}

//SetPartToCache 存入缓存中
func SetPartToCache(times int, uploadId string) error {
	var build strings.Builder
	build.WriteString(rk.PartInfoUploadId)
	build.WriteString(uploadId)
	key := build.String()
	err := global.GVA_REDIS.SetEX(context.Background(), key, strconv.Itoa(times), 10800).Err()
	if err != nil {
		global.GVA_LOG.Error("存入缓存失败", zap.Error(err))
		return err
	}
	return nil
}
