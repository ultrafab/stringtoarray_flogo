package stringtoarray_flogo

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
)
var log = logger.GetLogger("activity-helloworld")
// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	input_string := context.GetInput("input_string").(string)
	split_char := context.GetInput("split_char").(string)
	log.Infof("String to parse: [%s] with char: [%s]", input_string, split_char)
	result := strings.Split(input_string, split_char)
	context.SetOutput("result", result)
	return true, nil
}
