package tmp_save_code

// Set 设置链表中index索引处的值为t
//func (l *LinkedList[T]) Set(index int, t T) error {
//	if index < 0 || index >= l.length {
//		return newErrIndexOutOfRange(l.length, index)
//	}
//	rv := l.getNode(index)
//	rv.val = t
//	return nil
//}

//func TestLinkedList_Set(t *testing.T) {
//	testCases := []struct {
//		name           string
//		list           *LinkedList[int]
//		wantLinkedList *LinkedList[int]
//		index          int
//		setVal         int
//		wantErr        error
//	}{
//		{
//			name:    "set num to index -1",
//			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
//			index:   -1,
//			wantErr: fmt.Errorf("ekit: 下标超出范围，长度 %d, 下标 %d", 3, -1),
//		},
//		{
//			name:    "set beyond length index 99",
//			list:    NewLinkedListOf[int]([]int{1, 2, 3}),
//			index:   99,
//			wantErr: fmt.Errorf("ekit: 下标超出范围，长度 %d, 下标 %d", 3, 99),
//		},
//		{
//			name:    "set empty node",
//			list:    NewLinkedListOf[int]([]int{}),
//			index:   3,
//			wantErr: fmt.Errorf("ekit: 下标超出范围，长度 %d, 下标 %d", 0, 3),
//		},
//		{
//			name:           "set num to index 3",
//			list:           NewLinkedListOf[int]([]int{11, 22, 33, 44, 55}),
//			index:          2,
//			setVal:         999,
//			wantLinkedList: NewLinkedListOf([]int{11, 22, 999, 44, 55}),
//		},
//		{
//			name:           "set num to head",
//			list:           NewLinkedListOf[int]([]int{11, 22, 33, 44, 55}),
//			index:          0,
//			setVal:         -200,
//			wantLinkedList: NewLinkedListOf([]int{-200, 22, 33, 44, 55}),
//		},
//		{
//			name:           "set num to tail",
//			list:           NewLinkedListOf[int]([]int{-11, 22, -33, 44, -55, 999, -888}),
//			index:          6,
//			setVal:         888,
//			wantLinkedList: NewLinkedListOf([]int{-11, 22, -33, 44, -55, 999, 888}),
//		},
//		{
//			name:    "index == len(*node)",
//			list:    NewLinkedListOf[int]([]int{-11, 22, -33, 44, -55, 999, -888}),
//			index:   7,
//			setVal:  888,
//			wantErr: fmt.Errorf("ekit: 下标超出范围，长度 %d, 下标 %d", 7, 7),
//		},
//		{
//			name:    "len(*node) == 0",
//			list:    NewLinkedListOf[int]([]int{}),
//			index:   0,
//			setVal:  888,
//			wantErr: fmt.Errorf("ekit: 下标超出范围，长度 %d, 下标 %d", 0, 0),
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			err := tc.list.Set(tc.index, tc.setVal)
//			if err != nil {
//				assert.Equal(t, tc.wantErr, err)
//			} else {
//				assert.Equal(t, tc.wantLinkedList, tc.list)
//			}
//		})
//	}
//}
