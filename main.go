package main

import (
	"github.com/jtfm/go-cdk-core/pkg/lambda"
	"github.com/jtfm/ec2-scheduler/integrate"
)

func main() {
	lambda.SwitchingListenAndServe(integrate.Router)
}