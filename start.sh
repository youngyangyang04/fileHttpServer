#########################################################################
# File Name:    start.sh
# Author:       sunxiuyang
# mail:         sunxiuyang04@gmail.com
# Created Time: Tue 20 Aug 2019 05:09:02 PM CST
#########################################################################
#!/bin/bash
nohup go run src/server.go  >logs/file_upload.log 2>>logs/err.log &
