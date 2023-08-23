cd E:\Go\skilfactory\module40\final

swag init  -d ./apigw -o ./apigw/docs --parseDependency
go build  -o ./apigw/gwcmd.exe  ./apigw/
go build -o ./news/cmd/newscmd.exe  ./news/cmd/newscmd.go 
go build  -o ./comment/cmd/comcmd.exe  ./comment/cmd/ 

swag init  -d ./censor -o ./censor/docs --parseDependency
go build  -o ./censor/censorcmd.exe  ./censor
pause
