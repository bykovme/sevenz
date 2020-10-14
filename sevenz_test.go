package sevenz

import (
	"os"
	"testing"
)

const testArchive = "./testdata/test_archive.7z"
const testData = "./testdata/test.mp4"
const testExtractPath1 = "./testdata/1"
const testExtractPath2 = "./testdata/2"

func TestCheck7zAvailability(t *testing.T) {
	err := Check7zAvailability()
	if err != nil {
		t.Error(err)
	}
}

func TestCompressInsaneAndExtract(t *testing.T) {
	err := Check7zAvailability()
	if err != nil {
		t.Error(err)
	}
	errCompress := CompressInsane(testArchive, testData)
	if errCompress != nil {
		t.Error(err)
		_ = os.Remove(testArchive)
		return
	}

	errExtract := Extract(testArchive, testExtractPath1)
	if errExtract != nil {
		t.Error(err)
		_ = os.Remove(testArchive)
		_ = os.Remove(testExtractPath1)
		return
	}
	_ = os.Remove(testArchive)
	_ = os.Remove(testExtractPath1)
}

func TestCompressStandardAndExtract(t *testing.T) {
	err := Check7zAvailability()
	if err != nil {
		t.Error(err)
	}
	errCompress := CompressStandard(testArchive, testData)
	if errCompress != nil {
		t.Error(err)
		_ = os.Remove(testArchive)
		return
	}

	errExtract := Extract(testArchive, testExtractPath2)
	if errExtract != nil {
		t.Error(err)
		_ = os.Remove(testArchive)
		_ = os.Remove(testExtractPath2)
		return
	}
	_ = os.Remove(testArchive)
	_ = os.Remove(testExtractPath2)
}
