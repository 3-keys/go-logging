# go-logging
A lite go logging component

## Install

    go get -v github.com/vlinx2015/go-logging

## Usage
	import "github.com/vlinx2015/go-logging"
	
	logger := logging.New("info.log", "error.log")
	logger.Debug("it's a debug msg")
	logger.Info("it's an info msg")
	logger.Error("it's an error msg")
	
![vlinx go logging](https://vlinx.io/resources/go-logging.png)




	