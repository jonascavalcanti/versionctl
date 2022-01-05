package manipulator

import (
	"testing"
)

func Test_ReadLinesInFile(t *testing.T) {
	var fileNametest string = "../test/test.app.properties"
	result := ReadLinesInFile(fileNametest)
	// fmt.Println(result)
	expected := "version=2021.12.28.34"

	if result[0] != expected {
		t.Error("incorrect result: expected "+expected+", got", result[0])
	}
}

func Test_ReplaceInFile(t *testing.T) {

	fileNametest := "../test/test.app.properties"
	result := ReadLinesInFile(fileNametest)
	// fmt.Println(result)

	expected_begin := "version=2021.12.28.33"
	if result[0] == expected_begin {
		ReplaceInFile(fileNametest, "2021.12.28.33", "2021.12.28.34")
		result2 := ReadLinesInFile(fileNametest)
		// fmt.Println(result2)

		expected_end := "version=2021.12.28.34"
		if result[0] != expected_end {
			t.Error("incorrect result: expected "+expected_end+", got", result2[0])
		}
	}
}
