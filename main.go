package main

// 题目一：暴力破解密码
// 要求：1、设计一个信息系统，该系统可为学籍管理系统、订餐系统、票务管理系统不限，
// 系统必须通过客户端录入账号口令远程登录；
// 2、系统内至少包含三个以上账号，密码为6 位以上数字组成；

// 题目二：字典破解密码
// 要求：1、设计一个信息系统，该系统可为学籍管理系统、订餐系统、票务管理系统不限，
// 系统必须通过客户端录入账号口令远程登录；
// 2、系统内至少包含三个以上账号，密码为6 位以上任意字符组成；

// 题目三：认证审计系统
// 1、设计一个信息系统，系统必须通过客户端录入账号口令远程登录；
// 2、系统内至少包含三个以上账号；
// 3、某账号登录后服务器端可实时显示该账号登录的时间及IP 信息；
// 4、服务器端可查询账号的历史登录信息。

// 题目四：数据嗅探系统
// 1、设计一个信息系统，系统必须通过客户端录入账号口令远程登录；
// 2、登录后客户端可通过键盘输入向服务器发送数据；
// 3、服务器端设置嗅探关键字，如果客户端发送的数据包含该关键字，即将该数据显示出来。

// 题目五：防火墙系统
// 1、设计一个信息系统，系统必须通过客户端录入账号口令远程登录；
// 2、系统内至少包含三个以上账号；
// 3、系统服务器端可设定禁止登录的IP 地址和账号信息；
// 4、如果客户端从禁止的IP 地址登录或使用禁止的账号登录则显示不允许登录，并断开连
// 接。

// 题目六：加密传输系统
// 1、设计客户端程序向服务器端发送数据；
// 2、客户端从键盘输入的数据在发送之前进行加密，加密方法可选择仿射、移位密码；
// 3、服务器端接收到数据后进行显示，然后在解密后再次显示。

// 题目七：图形验证模拟
// 要求：1、开发一个手机锁屏的图形验证程序，以字符命令行形式来实现要求完成的功能有：
// 2、登录时输入用户名；
// 3、输入4*4 坐标下的图形点位置，用字符方式输入；
// 4、输入完成后实现在服务器端对图形进行验证；
// 5、至少有三个以上的用户验证。

func main() {
	LoadENV()
	initSQL()
	dao := NewUserDaoAdapter(DB, NewUserCoreAdapter())
	userHandler = NewUserHandlerAdapter(dao)
	InitializeRouter()
}
