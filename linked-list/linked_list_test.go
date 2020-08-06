package linked_list

import "testing"

func TestInsert(t *testing.T) {
	cases := []struct {
		preInsert []string
		toInsert  string
		index     int
		expErr    error
		expPrint  string
	}{
		{
			preInsert: []string{"zero", "one", "two", "three"},
			toInsert:  "four",
			index:     4,
			expErr:    nil,
			expPrint:  "[zero,one,two,three,four]",
		},
		{
			preInsert: []string{"zero", "one", "two", "three"},
			toInsert:  "four",
			index:     5,
			expErr:    InvalidIndexError,
			expPrint:  "[zero,one,two,three]",
		},
		{
			preInsert: []string{"zero", "one", "two", "three"},
			toInsert:  "four",
			index:     0,
			expErr:    nil,
			expPrint:  "[four,zero,one,two,three]",
		},
	}

	for _, tc := range cases {
		ll := buildSLL(tc.preInsert)

		err := ll.Insert(tc.toInsert, tc.index)
		if err != tc.expErr {
			t.Errorf("Unexpected Insert() result. Got: %v, Wanted: %v", err, tc.expErr)
		}

		if tc.expPrint != ll.String() {
			t.Errorf("Print mismatch, Got: %v, Wanted: %v", ll, tc.expPrint)
		}
	}

}

func buildSLL(preInsert []string) *SLinkedList {
	s := &SLinkedList{}
	for i, el := range preInsert {
		s.Insert(el, i)
	}

	return s
}
