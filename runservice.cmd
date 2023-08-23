
rm -f news/cmd/*log
rm -f comment/cmd/*log
cd "e:\Go\skilfactory\module40\final\apigw\"
rm -f *.log
cd "e:\Go\skilfactory\module40\final\news\cmd\"
rm -f *.log
start /d "e:\Go\skilfactory\module40\final\bin\news\" newscmd.exe -debug
cd "e:\Go\skilfactory\module40\final\comment\"
rm -f *.log
start /d "e:\Go\skilfactory\module40\final\bin\comment\" comcmd.exe -debug
start /d "e:\Go\skilfactory\module40\final\bin\censor\" censorcmd.exe -debug
pause