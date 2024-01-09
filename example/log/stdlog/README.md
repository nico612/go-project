
go 日志标准库

log默认输出到标准错误（stderr），每条日志前会自动加上日期和时间。如果日志不是以换行符结尾的，那么log会自动加上换行符。即每条日志会在新行中输出。

log提供了三组函数：

- Print/Printf/Println：正常输出日志；
- Panic/Panicf/Panicln：输出日志后，以拼装好的字符串为参数调用panic；
- Fatal/Fatalf/Fatalln：输出日志后，调用os.Exit(1)退出程序。
