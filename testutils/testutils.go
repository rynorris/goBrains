package testutils

import "math"
import "io/ioutil"
import "crypto/md5"

const Delta = 0.0000001

// FloatsAreEqual returns true if the two given float64
// values are within Delta from each other.
func FloatsAreEqual(x, y float64) bool {
	if dist := math.Abs(x - y); dist > Delta {
		return false
	}

	return true
}

// FileAreEqual returns true if the two given files
// are the same.
// It does this by calculating the md5 checksum of each file
// and comparing.
// If either of the files does not exist, this returns false, and an error.
func FilesAreEqual(a, b string) (bool, error) {
	f1, err1 := ioutil.ReadFile(a)
	if err1 != nil {
		return false, err1
	}

	f2, err2 := ioutil.ReadFile(b)
	if err2 != nil {
		return false, err2
	}

	sum1 := md5.Sum(f1)
	sum2 := md5.Sum(f2)

	return sum1 == sum2, nil
}
