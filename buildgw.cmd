cd E:\Go\skilfactory\module40\final

swag init  -d ./apigw -o ./apigw/docs --parseDependency
go build  -o ./apigw/gwcmd.exe  ./apigw/
pause
