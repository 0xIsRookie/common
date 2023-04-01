package file

import (
	"github.com/0xIsRookie/common/file"
	"os"
	"testing"
)

func TestTraverseFiles(t *testing.T) {
	// 创建一个临时目录，并在其中创建一些文件和子目录
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	// 删除临时目录
	defer os.RemoveAll(tempDir)

	tempFile1, err := os.CreateTemp(tempDir, "file1.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer tempFile1.Close()

	tempFile2, err := os.CreateTemp(tempDir, "file2.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer tempFile2.Close()

	tempSubdir, err := os.MkdirTemp(tempDir, "subdir")
	if err != nil {
		t.Fatal(err)
	}

	tempFile3, err := os.CreateTemp(tempSubdir, "file3.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer tempFile3.Close()

	// 调用 TraverseFiles 函数，并验证其输出
	expectedFiles := []string{tempFile1.Name(), tempFile2.Name(), tempFile3.Name()}
	actualFiles, err := file.TraverseFiles(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	if len(expectedFiles) != len(actualFiles) {
		t.Fatalf("期望找到 %d 个文件，实际找到 %d 个", len(expectedFiles), len(actualFiles))
	}

	for i, expectedPath := range expectedFiles {
		actualPath := actualFiles[i]
		if expectedPath != actualPath {
			t.Errorf("期望路径为 %q，但实际路径为 %q", expectedPath, actualPath)
		}
	}
}
