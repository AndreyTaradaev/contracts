# windows
BINARY_NAME_GW=gwcmd.exe
BINARY_NAME_NEWS=newscmd.exe
BINARY_NAME_COMMENT=comcmd.exe
BINARY_NAME_CENSOR=censorcmd.exe
BINARY_DIR=./bin
SOURCE_GATE=./apigw/
SOURCE_NEWS=./news/cmd/
SOURCE_COMMENT=./comment/cmd/
SOURCE_CENSOR=./censor/

.PHONY: dir uninstall

dir:
	mkdir -p ${BINARY_DIR}/gw/
	mkdir -p ${BINARY_DIR}/news/ 
	mkdir -p ${BINARY_DIR}/comment/ 
	mkdir -p ${BINARY_DIR}/censor/

gate: dir  dir ${SOURCE_GATE}main.go
	swag init -d ${SOURCE_GATE} -o ${SOURCE_GATE}/docs  --parseDependency
	go build  -o ${BINARY_DIR}/gw/${BINARY_NAME_GW}  ${SOURCE_GATE}
	
censors: dir ${SOURCE_CENSOR}main.go	
	swag init -d ${SOURCE_CENSOR} -o ${SOURCE_CENSOR}/docs --parseDependency
	go build  -o ${BINARY_DIR}/censor/${BINARY_NAME_CENSOR}  ${SOURCE_CENSOR}	

new:  dir ${SOURCE_NEWS}newscmd.go
	
	go build  -o ${BINARY_DIR}/news/${BINARY_NAME_NEWS}  ${SOURCE_NEWS}
	
	
comments: dir ${SOURCE_COMMENT}comcmd.go
	
	go build  -o ${BINARY_DIR}/comment/${BINARY_NAME_COMMENT}  ${SOURCE_COMMENT}
		
build: 	gate new comments censors

install: ${SOURCE_GATE}/${BINARY_NAME_GW}  ${SOURCE_NEWS}/${BINARY_NAME_NEWS}  ${SOURCE_COMMENT}/${BINARY_NAME_COMMENT}

	 
	cp ${SOURCE_GATE}config.json ./${BINARY_DIR}/gw
	cp ${SOURCE_NEWS}config.json ./${BINARY_DIR}/news
	cp ${SOURCE_COMMENT}config.json ./${BINARY_DIR}/comment
	cp ${SOURCE_CENSOR}config.json ./${BINARY_DIR}/censor
	cp ${SOURCE_GATE}/${BINARY_NAME_GW}   ./${BINARY_DIR}/gw/
	cp ${SOURCE_NEWS}/${BINARY_NAME_NEWS} ./${BINARY_DIR}/news/
	cp ${SOURCE_COMMENT}/${BINARY_NAME_COMMENT} ./${BINARY_DIR}/comment/
	cp ${SOURCE_CENSOR}/${BINARY_NAME_CENSOR} ./${BINARY_DIR}/censor/
	
run:	./${BINARY_DIR}/censor/${BINARY_NAME_CENSOR}  ./${BINARY_DIR}/comment/${BINARY_NAME_COMMENT} ./${BINARY_DIR}/news/${BINARY_NAME_NEWS}  ./${BINARY_DIR}/gw/${BINARY_NAME_GW}	
	runservice.cmd
	
#debug: install	
#	start /d .\${BINARY_DIR}\comment  comcmd.exe -debug
#	start /d  ./${BINARY_DIR}/news  newscmd.exe -debug
#	start /d   ./${BINARY_DIR}/gw  apicmd.exe -debug

uninstall: 
	rm -f ./${BINARY_DIR}/gw/* 
	rm -f ./${BINARY_DIR}/news/*
	rm -f ./${BINARY_DIR}/comment/* 
	rm -f ./${BINARY_DIR}/censor/
	