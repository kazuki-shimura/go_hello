# go_hello

実行方法
go mod init go_hello
-> 今回は
go_hello % go mod init go_hello としたため go_hello 以下を import できるようになる
ex) import "go_hello/study"

study.Part1()
このとき study. をつけないで呼び出そうとすると下記のエラーが発生

# command-line-arguments

./main.go:6:2: imported and not used: "go_hello/study"
./main.go:11:2: undefined: Part1

// 存在しないモジュールをインポートしようとしたとき
main.go:6:2: package go_hello/studyss is not in GOROOT (/usr/local/Cellar/go/1.18/libexec/src/go_hello/studyss)

GO111MODULE=off コマンドを叩いて、存在しないモジュールをインポートしようとしたとき
main.go:6:2: cannot find package "go_hello/studyss" in any of:
/usr/local/Cellar/go/1.18/libexec/src/go_hello/studyss (from $GOROOT)
/Users/shimurakazuki/go/src/go_hello/studyss (from $GOPATH)

GO111MODULE=on or GO111MODULE=auto コマンドを叩いて、存在しないモジュールをインポートしようとしたとき
main.go:6:2: package go_hello/studyss is not in GOROOT (/usr/local/Cellar/go/1.18/libexec/src/go_hello/studyss)

go mod tidy

go run main.go
