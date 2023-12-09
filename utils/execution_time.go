package utils

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func ExecuteAndLogTime(fn interface{}, args ...interface{}) (interface{}, error) {

	// Check if first param is a function
	fnValue := reflect.ValueOf(fn)
	if fnValue.Kind() != reflect.Func {
		return nil, fmt.Errorf("First parameter is not a function")
	}

	// Check if the number of arguments matches the function's input parameters
	if len(args) != fnValue.Type().NumIn() {
		return nil, fmt.Errorf("Number of arguments does not match function's input parameters")
	}

	// Prepare arguments for the function call
	inArgs := make([]reflect.Value, len(args))
	for i := range args {
		inArgs[i] = reflect.ValueOf(args[i])
	}

	// Find function name
	funcName := strings.Split(runtime.FuncForPC(fnValue.Pointer()).Name(), ".")[1]

	// Call function with the provided arguments
	startTime := time.Now()
	result := fnValue.Call(inArgs)
	elapsedTime := time.Since(startTime).String()
	fmt.Printf("Execution time for \033[31;1m%s\033[0m: \033[34m%v\033[0m\n", funcName, elapsedTime)

	if len(result) > 0 {
		return result[0].Interface(), nil
	}

	return nil, nil
}