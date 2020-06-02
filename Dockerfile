FROM ubuntu:19.04
COPY go-histogram /
COPY mydata.txt /
#ENTRYPOINT [ "bash" ]