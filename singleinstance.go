package singleinstance
//用go语言实现单例设计模式

//对象的结构体
type instance struct {

}

var i *instance

//通过此函数创建对象，且返回的对象均为同一个
func GetInstance() *instance {
	if i == nil {
		i = new(instance)
		return i
	}

	return i
}