package main

//func main(){
//	t := &testing.T{}
//	ctrl := gomock.NewController(t)
//	mockIndex := NewMockFoo(ctrl)
//	var s string
//
//	mockIndex.EXPECT().Bar(gomock.AssignableToTypeOf(s)).DoAndReturn(
//		// signature of anonymous function must have the same number of input and output arguments as the mocked method.
//		func(arg string) interface{} {
//			s = arg
//			return "I'm sleepy"
//		},
//	)
//
//	r := mockIndex.Bar("foo")
//	fmt.Printf("%s %s", r, s)
//
//
//}
