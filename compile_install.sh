#!/bin/bash

IS_COMPILE=0

if [ "c" = $1 ] ;
then
    IS_COMPILE=1;
fi

if [ $IS_COMPILE = 1 ];
then
    exec go build
else #this will list directories your app.exe is installed on 
    exec go list -f '{{.Target}}'
fi

#to make it public add the bin dir 
#(where you should keep .exe's which can be accessed from anywhere ) to path 
#so that you can run software from anyware
#from commandline --uncomment the 20 
# export PATH=$PATH:my-app/bin

echo "$IS_COMPILE"