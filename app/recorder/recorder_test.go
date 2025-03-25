package recorder

import (
	"github.com/zhangpetergo/LiveStreamRecorder/app/resolver/douyin"
	"testing"
)

func TestRecord(t *testing.T) {
	url := "https://live.douyin.com/695496496290"
	data, err := douyin.GetStreamData(url)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	err = Record(data)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestRecordWithValidData(t *testing.T) {
	data := map[string]interface{}{
		"url":  "http://example.com/stream",
		"name": "test_stream",
	}

	err := Record(data)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestRecordWithMissingURL(t *testing.T) {
	data := map[string]interface{}{
		"name": "test_stream",
	}

	err := Record(data)
	if err == nil || err.Error() != "url 不存在" {
		t.Errorf("Expected error 'url 不存在', got %v", err)
	}
}

func TestRecordWithMissingName(t *testing.T) {
	data := map[string]interface{}{
		"url": "http://example.com/stream",
	}

	err := Record(data)
	if err == nil || err.Error() != "name 不存在" {
		t.Errorf("Expected error 'name 不存在', got %v", err)
	}
}

func TestRecordWithInvalidURLType(t *testing.T) {
	data := map[string]interface{}{
		"url":  12345,
		"name": "test_stream",
	}

	err := Record(data)
	if err == nil || err.Error() != "url 不存在" {
		t.Errorf("Expected error 'url 不存在', got %v", err)
	}
}

func TestRecordWithInvalidNameType(t *testing.T) {
	data := map[string]interface{}{
		"url":  "http://example.com/stream",
		"name": 12345,
	}

	err := Record(data)
	if err == nil || err.Error() != "name 不存在" {
		t.Errorf("Expected error 'name 不存在', got %v", err)
	}
}
