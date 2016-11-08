package InitTestWorkspace

import "testing"
func TestAdd(t *testing.T) {
	var result int;
	result =Add(10,10);
	if result!=20{
		t.Error("Expect 20, but got ",result);
	}
}

func TestSubstract(t *testing.T) {
	var result int;
	result = Substract(20,10);
	if result!=10{
		t.Error("Expect 10 , but got ",result);
	}
}
