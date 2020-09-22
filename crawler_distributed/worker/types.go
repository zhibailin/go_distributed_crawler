package worker

// 传递函数名和函数参数，序列化出来的是
// {"ParseCityList", nil}, {"ProfileParser", userId, userName}
type SerializedParser struct {
	FunctionName string
	Args         interface{}
}
