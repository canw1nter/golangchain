package common

// IOptions 所有具有可选参数的结构体都应实现该接口来设置可选字段
type IOptions interface {
	SetOptions(opts ...Options)
}

// Options 通过该函数可以设置结构体可选字段的值，否则可选字段为默认值
// 你可以直接使用结构体对应包中形如`WithXxx`的函数来设置可选字段`xxx`的值
// 除此以外，你还可以自己定义一个Options来自定义的改变一个或多个可选字段的值
// Options内的结构如下：
//
//		if options, ok := obj.(*ObjectOptions); ok {
//		    options.Xxx = value
//	     // 你需要自定义的可选参数
//		}
//
// 其中obj是传入的参数，实际类型应为*ObjectOptions，即具体结构体的可选参数结构体
type Options func(interface{})
