@echo off

REM set GOPATH=%~dp0..\3rdparty;%~dp0
REM set GOBIN=%~dp0bin
REM set GOOS=windows

cd /d %~dp0


REM go build -o .\bin\gencmd.exe .\src\tools\gencmd

REM .\bin\gencmd.exe 

.\proto\protoc.exe -I=.\proto --plugin=protoc-gen-go=.\proto\protoc-gen-go.exe --go_out=.\command command.proto
.\proto\protoc.exe -I=.\proto --plugin=protoc-gen-go=.\proto\protoc-gen-go.exe --go_out=.\command sever.proto

echo compile proto ok

pause
