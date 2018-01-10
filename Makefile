
JAVA_SRC=`find ./java/ -name "*.java"`
JAVA_TGT=target/classes
TEST_DT=target/tdata

default: prepare
	go build .


prepare:
	mkdir -p ${JAVA_TGT}
	mkdir -p ${TEST_DT}


testdata: prepare
	javac  ${JAVA_SRC} -d ${JAVA_TGT}
	java -cp ${JAVA_TGT} cmd.TestGenerator ${TEST_DT}

test: testdata
	go test ./...
clean:
	rm -rf target