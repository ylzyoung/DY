package util

import (
	"douyin/config"
	models2 "douyin/models"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetFileUrl(t *testing.T) {

	fileName := "test.mp4"
	expectedUrl := "http://127.0.0.1:8080/static/test.mp4"
	actualUrl := GetFileUrl(fileName)
	assert.Equal(t, expectedUrl, actualUrl)
}

func TestNewFileName(t *testing.T) {
	userId := int64(123)
	expectedName := "123-0"
	actualName := NewFileName(userId)
	assert.Equal(t, expectedName, actualName)
}

func TestFillVideoListFields(t *testing.T) {
	videos := []*models2.Video{
		{Id: 1, UserInfoId: 101, CreatedAt: time.Now()},
		{Id: 2, UserInfoId: 102, CreatedAt: time.Now().Add(-time.Hour)},
	}
	userId := int64(1001)

	latestTime, err := FillVideoListFields(userId, &videos)
	assert.NoError(t, err)
	assert.NotNil(t, latestTime)
	assert.Equal(t, videos[0].CreatedAt, *latestTime)
	assert.NotNil(t, videos[0].Author)
	assert.NotNil(t, videos[1].Author)
}

func TestSaveImageFromVideo(t *testing.T) {

	videoFileName := "testvideo.mp4"
	videoFilePath := filepath.Join(config.Global.StaticSourcePath, videoFileName)
	file, err := os.Create(videoFilePath)
	assert.NoError(t, err)
	defer os.Remove(videoFilePath)
	defer file.Close()

	err = SaveImageFromVideo("testvideo", true)
	assert.NoError(t, err)

	imageFilePath := filepath.Join(config.Global.StaticSourcePath, "testvideo.jpg")
	_, err = os.Stat(imageFilePath)
	assert.NoError(t, err)
	defer os.Remove(imageFilePath)
}
