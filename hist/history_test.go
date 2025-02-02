package hist

import (
	"strings"
	"testing"
)

func Test_simpleEntries(t *testing.T) {

	results, err := TestHistory("../test_files/test1.txt")
	if err != nil {
		t.Fatal("Error in test")
	}

	checkIntValue(t, "Total lines count", 4, results.Lines)

	checkIntValue(t, "Unique value", 4, results.Unique)

	checkIntValue(t, "Duplicate value", 0, results.Dups)

	checkIntValue(t, "UniqueVals slice size", 4, len(results.UniqueVals))

	checkIntValue(t, "DupVals slice size", 0, len(results.DupVals))

}

func Test_simpleDupsAndSingleUnique(t *testing.T) {
	results, err := TestHistory("../test_files/test2.txt")
	if err != nil {
		t.Fatal("Error in test")
	}

	checkIntValue(t, "Total Lines Count", 9, results.Lines)

	checkIntValue(t, "Unique value", 1, results.Unique)

	checkIntValue(t, "Duplicate value", 4, results.Dups)

	checkIntValue(t, "Unique Values Slice size", 1, len(results.UniqueVals))

	checkIntValue(t, "DupVals slice size", 4, len(results.DupVals))

}

func Test_multipleTimesSameValueDuplicated(t *testing.T) {
	results, err := TestHistory("../test_files/test3.txt")
	if err != nil {
		t.Fatal("Error in test")
	}

	checkIntValue(t, "Total Lines Count", 17, results.Lines)

	checkIntValue(t, "Unique value", 0, results.Unique)

	checkIntValue(t, "Duplicate value", 3, results.Dups)

	checkIntValue(t, "Unique Values Slice size", 0, len(results.UniqueVals))

	checkIntValue(t, "DupVals slice size", 3, len(results.DupVals))

}

func Test_multipleSameDupAndSeveralUniques(t *testing.T) {
	results, err := TestHistory("../test_files/test4.txt")
	if err != nil {
		t.Fatal("Error in test")
	}

	checkIntValue(t, "Total Lines Count", 20, results.Lines)

	checkIntValue(t, "Unique value", 3, results.Unique)

	checkIntValue(t, "Duplicate value", 3, results.Dups)

	checkIntValue(t, "Unique Values Slice size", 3, len(results.UniqueVals))

	checkIntValue(t, "DupVals slice size", 3, len(results.DupVals))

}

func Test_showSingleMostDuplicatedEntry(t *testing.T) {
	results, err := TestHistory("../test_files/test4.txt")
	if err != nil {
		t.Fatal("Error in test")
	}

	maxKey, max := FindMaxDupValueAndName(results.DupCounts)

	checkIntValue(t, "Single Most Duplicated value", 7, max)

	if !strings.Contains(maxKey, "three") {
		t.Error("Max dup value did not contain 'three'")
	}
}

func checkIntValue(t *testing.T, valueDesc string, expected int, actual int) {
	if expected != actual {
		t.Errorf("%s wrong! Expected %d and got %d!", valueDesc, expected, actual)
	}
}
