package graphics

import (
	"github.com/DiscoViking/goBrains/testutils"
	"os"
	"testing"
)

func CompareOutput(testname string, t *testing.T) {
	match, err := testutils.FilesAreEqual(
		"test_output/"+testname+"_got.bmp",
		"test_output/"+testname+"_exp.bmp")
	if err != nil {
		t.Errorf(err.Error())
	} else if !match {
		t.Errorf(testname + ": Expected and actual outputs differ. Check files manually.")
	} else {
		//Pass, so remove _got file so we dont clog the output directory.
		os.Remove("test_output/" + testname + "_got.bmp")
	}
}
