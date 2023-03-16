package main

import (
	"github.com/jtfm/ec2-scheduler/internal/integrate"
	"github.com/jtfm/go-cdk-core/pkg/lambda"
)

func main() {
	lambda.SwitchingListenAndServe(integrate.Router(), ":3000")
}
