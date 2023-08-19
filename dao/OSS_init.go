package dao

import (
	"fmt"
	"log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

const (
	OSSAK          = "1"
	OSSSK          = "2"
	OSSEndPoint    = "oss-cn-hangzhou.aliyuncs.com"
	OSSVideoBucket = "douyin-server-hust"
	OSSImageBucket = "douyin-server-hust-image"
)

var (
	VideoBucket           *oss.Bucket
	ImageBucket           *oss.Bucket
	VideoBucketLinkPrefix string
	ImageBucketLinkPrefix string
)

func Oss_init() {
	ossClient, err := oss.New(OSSEndPoint, OSSAK, OSSSK)
	if err != nil {
		log.Fatalf("OSS Init Failed")
		panic(err)
	}

	VideoBucket, err = ossClient.Bucket(OSSVideoBucket)
	if err != nil {
		log.Println("VideoBucket Init Failed")
		panic(err)
	}
	VideoBucketLinkPrefix = fmt.Sprintf(
		"https://%s.%s/", OSSVideoBucket, OSSEndPoint)

	ImageBucket, err = ossClient.Bucket(OSSImageBucket)
	if err != nil {
		hlog.Fatalf("ImageBucket Init Failed")
		panic(err)
	}
	ImageBucketLinkPrefix = fmt.Sprintf(
		"https://%s.%s/", OSSImageBucket, OSSEndPoint)
}
