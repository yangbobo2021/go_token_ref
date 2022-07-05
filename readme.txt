编译：
go mod tidy
go build .

分析：
cp ./analyzer /xx/xx/repo
cd /xx/xx/repo
go mod tidy
go build .
./analyzer . > result.txt

python filter.py /xx/xx/repo/result.txt


注意：
1. filter.py仅仅解析了包引用次数；
2. 被分析项目要能够编译通过；