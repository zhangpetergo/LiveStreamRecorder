package fileutil

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCheckDir(t *testing.T) {
	// Test case: directory does not exist
	testDir := "D:\\test\\抖音直播\\jack要加油"
	err := CheckDir(testDir)

	assert.Nil(t, err)

	// Verify directory was created
	info, err := os.Stat(testDir)
	if err != nil {
		t.Fatalf("expected directory to be created, got error: %v", err)
	}
	if !info.IsDir() {
		t.Fatalf("expected %s to be a directory", testDir)
	}

	// Clean up
	err = os.RemoveAll(testDir)
	if err != nil {
		t.Fatalf("failed to remove test directory: %v", err)
	}
}
