# gotime
用golang写的一个time的工具集合(A time util for golang)

#Install

    go get github.com/andyidea/gotime
    
#Function

获取时间戳

    func GetTimeUnix(t time.Time) int64 

获取当前时间的时间戳

    func GetNowTimeUnix() int64
    
获取当日晚上24点（次日0点）的时间

    func Get24time(t time.Time) time.Time 
    
获取当日晚上24点（次日0点）的时间戳

    func Get24timeUnix(t time.Time) int64
    
获取今天晚上24点（次日0点）的时间戳

    func GetToday24timeUnix() int64
    
时间转换成日期字符串 (time.Time to "2006-01-02")

    func TimeToDate(t time.Time) string
    
获取当前的日期字符串

    func GetNowDateStr() string
    
时间转换成日期+时间字符串 (time.Time to "2006-01-02 15:04:05")

    func TimeToDateTime(t time.Time) string
    
获取当前的时期+时间字符串

    func GetNowStr() string
    
日期字符串转换成时间 ("2006-01-02" to time.Time)

    func DateStrToTime(d string) time.Time
    
日期+时间字符串转换成时间 ("2006-01-02 15:04:05" to time.Time)

    func DateTimeStrToTime(dt string) time.Time
