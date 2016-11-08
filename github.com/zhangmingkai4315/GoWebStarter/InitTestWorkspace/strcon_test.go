package InitTestWorkspace

import "testing"

func TestSwapCase(t *testing.T) {
	//var str string;
	str:="Hello World";
	if SwapCase(str)!="hELLO wORLD"{
		t.Error("Expect hELLO wORLD, but got ",SwapCase(str));
	}
}
