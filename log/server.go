package logper

import (
	"io"
	"log"
	"net/http"
	"os"
)

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description: 日志服务
 * @Date: 2024-10-11 22:46
 */

var logper *log.Logger

type fileLog string

// 是对io.Write的实现
// New creates a new Logger. The out variable sets the
// destination to which log data will be written.
// The prefix appears at the beginning of each generated log line, or
// after the log header if the Lmsgprefix flag is provided.
// The flag argument defines the logging properties.
//
//	这里通过 io.Writer和prefix(前缀)来创建对应的log
//	func New(out io.Writer, prefix string, flag int) *Logger {
//		l := &Logger{out: out, prefix: prefix, flag: flag}
//		if out == io.Discard {
//			l.isDiscard.Store(true)
//		}
//		return l
//	}
//
// 为后续使用 log.New 方法提供相对应的Write功能
func (ffl fileLog) Write(data []byte) (int, error) {
	// 0000: 无权限。
	// 0400: 只有文件所有者具有读取权限。
	// 0200: 只有文件所有者具有写入权限。
	// 0600: 只有文件所有者具有读写权限（常用于日志文件）。
	// 0070: 只有文件所有者和组具有读、写和执行权限。
	// 0755: 所有者具有读、写和执行权限，组和其他用户具有读和执行权限（常用于可执行文件）。
	// 0777: 所有用户均具有读、写和执行权限（通常不推荐，因为安全风险）
	f, err := os.OpenFile(string(ffl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}
func Run(destination string) {
	// const (
	// 	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	// 	Ltime                         // the time in the local time zone: 01:23:23
	// 	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	// 	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	// 	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	// 	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	// 	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	// 	LstdFlags     = Ldate | Ltime // initial values for the standard logger
	// )
	fileWriter := fileLog(destination)
	consoleWriter := os.Stdout
	multiWriter := io.MultiWriter(fileWriter, consoleWriter) // 这是一个组合Writer的函数，底层是一个实现io.Writer接口的切片
	logper = log.New(multiWriter, "Go logper: ", log.LstdFlags)
}
func write(data string) {
	logper.Printf("%s", data)
}
func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := io.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}
