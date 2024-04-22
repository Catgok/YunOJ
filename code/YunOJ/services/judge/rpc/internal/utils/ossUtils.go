package utils

import (
	"YunOJ/services/judge/rpc/judge"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"strings"
)

type OSSUtil struct {
	client   *oss.Client
	bucket   *oss.Bucket
	rootPath string
}

func NewOSSUtil(accessKeyID, accessKeySecret, endpoint, bucketName, rootPath string) (*OSSUtil, error) {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		return nil, fmt.Errorf("oss error creating OSS client: %s", err)
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, fmt.Errorf("oss error getting bucket: %s", err)
	}

	return &OSSUtil{
		client:   client,
		bucket:   bucket,
		rootPath: rootPath,
	}, nil
}

func (util *OSSUtil) getFilePath(problemID int64, i int64, fileType string) string {
	return fmt.Sprintf("%s/%d/%d_%d.%s", util.rootPath, problemID, problemID, i, fileType)
}
func (util *OSSUtil) getFilePaths(problemID int64, count int, fileType string) []string {
	var paths []string
	for i := 0; i < count; i++ {
		paths = append(paths, fmt.Sprintf("%s/%d/%d_%d.%s", util.rootPath, problemID, problemID, i, fileType))
	}
	return paths
}
func (util *OSSUtil) getInFilePaths(problemID int64, count int) []string {
	return util.getFilePaths(problemID, count, "in")
}
func (util *OSSUtil) getOutFilePaths(problemID int64, count int) []string {
	return util.getFilePaths(problemID, count, "out")
}

// 单个文件操作

func (util *OSSUtil) UploadFile(filePath, localFilePath string) error {
	err := util.bucket.PutObjectFromFile(filePath, localFilePath)
	if err != nil {
		return fmt.Errorf("oss error uploading file: %s", err)
	}
	return nil
}

func (util *OSSUtil) UploadString(filePath, content string) error {
	err := util.bucket.PutObject(filePath, strings.NewReader(content))
	if err != nil {
		return fmt.Errorf("oss error uploading string: %s", err)
	}
	return nil
}

func (util *OSSUtil) DownloadFile(filePath, downloadFilePath string) error {
	err := util.bucket.GetObjectToFile(filePath, downloadFilePath)
	if err != nil {
		return fmt.Errorf("oss error downloading file: %s", err)
	}
	return nil
}

func (util *OSSUtil) DeleteFile(filePath string) error {
	err := util.bucket.DeleteObject(filePath)
	if err != nil {
		return fmt.Errorf("oss error deleting file: %s", err)
	}
	return nil
}

func (util *OSSUtil) UpdateString(filePath, content string) error {
	return util.UploadString(filePath, content)
}

func (util *OSSUtil) DownloadString(filePath string) (string, error) {
	body, err := util.bucket.GetObject(filePath)
	if err != nil {
		return "", fmt.Errorf("oss error downloading string: %s", err)
	}
	defer body.Close()

	content, err := io.ReadAll(body)
	if err != nil {
		return "", fmt.Errorf("oss error reading downloaded string: %s", err)
	}

	return string(content), nil
}

// 批量文件操作

func (util *OSSUtil) UploadFiles(filePaths []string, localFilePaths []string) error {
	if len(filePaths) != len(localFilePaths) {
		return fmt.Errorf("number of file paths and local file paths do not match")
	}

	for i, filePath := range filePaths {
		err := util.UploadFile(filePath, localFilePaths[i])
		if err != nil {
			return fmt.Errorf("oss error uploading file: %s", err)
		}
	}

	return nil
}

func (util *OSSUtil) uploadStrings(filePaths []string, contents []string) error {
	if len(filePaths) != len(contents) {
		return fmt.Errorf("number of file paths and contents do not match")
	}

	for i, filePath := range filePaths {
		err := util.UploadString(filePath, contents[i])
		if err != nil {
			return fmt.Errorf("oss error uploading string: %s", err)
		}
	}

	return nil
}

func (util *OSSUtil) DownloadFiles(filePaths []string, downloadFilePaths []string) error {
	if len(filePaths) != len(downloadFilePaths) {
		return fmt.Errorf("number of file paths and download file paths do not match")
	}

	for i, filePath := range filePaths {
		err := util.DownloadFile(filePath, downloadFilePaths[i])
		if err != nil {
			return fmt.Errorf("oss error downloading file: %s", err)
		}
	}

	return nil
}

func (util *OSSUtil) deleteFiles(filePaths []string) error {
	for _, filePath := range filePaths {
		err := util.DeleteFile(filePath)
		if err != nil {
			return fmt.Errorf("oss error deleting file: %s", err)
		}
	}

	return nil
}

func (util *OSSUtil) UpdateStrings(filePaths []string, contents []string) error {
	return util.uploadStrings(filePaths, contents)
}

func (util *OSSUtil) DownloadStrings(filePaths []string) ([]string, error) {
	var results []string

	for _, filePath := range filePaths {
		content, err := util.DownloadString(filePath)
		if err != nil {
			return nil, fmt.Errorf("oss error downloading strings: %s", err)
		}
		results = append(results, content)
	}

	return results, nil
}

func (util *OSSUtil) getDirectoryFiles(dirPath string) ([]string, error) {
	var files []string
	marker := ""
	for {
		lsRes, err := util.bucket.ListObjects(oss.Prefix(dirPath), oss.Marker(marker))
		if err != nil {
			return nil, fmt.Errorf("oss error listing objects: %s", err)
		}

		for _, object := range lsRes.Objects {
			files = append(files, object.Key)
		}

		if !lsRes.IsTruncated {
			break
		}
		marker = lsRes.NextMarker
	}

	return files, nil
}

func (util *OSSUtil) GetDirectoryFilesByProblemId(problemId int64) ([]string, error) {
	return util.getDirectoryFiles(fmt.Sprintf("%s/%d", util.rootPath, problemId))
}

func (util *OSSUtil) Upload(cases []*judge.JudgeCase, problemId int64) (err error) {
	var inputs, outputs []string
	for _, caseInfo := range cases {
		inputs = append(inputs, caseInfo.Input)
		outputs = append(outputs, caseInfo.Output)
	}
	err = util.uploadStrings(util.getInFilePaths(problemId, len(inputs)), inputs)
	if err != nil {
		return err
	}
	err = util.uploadStrings(util.getOutFilePaths(problemId, len(outputs)), outputs)
	return err
}

func (util *OSSUtil) Delete(problemId int64) (err error) {
	files, err := util.getDirectoryFiles(fmt.Sprintf("%s/%d", util.rootPath, problemId))
	if err != nil {
		return err
	}
	err = util.deleteFiles(files)
	return err
}
