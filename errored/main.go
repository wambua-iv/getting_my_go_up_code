package main

import "fmt"

type CustomErrored struct {
	Resource string
	Code int
}

func (ce CustomErrored) Error () string {
	return fmt.Sprintf("%s : %d", ce.Resource, ce.Code)
}

func (ce CustomErrored) Is (err error) bool {
	if other, ok := err.(CustomErrored); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		matchedResource := other.Resource == ce.Resource
		matchedCode := other.Code == ce.Code
		
		return matchedCode && matchedResource ||
			ignoreCode && matchedResource || 
			ignoreResource && matchedCode
	}
	return false
} 