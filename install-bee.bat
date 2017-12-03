set GOPATH=C:\Projects\Go
go build github.com\beego\bee
copy bee.exe %GOPATH%\bin\bee.exe
del bee.exe
pause