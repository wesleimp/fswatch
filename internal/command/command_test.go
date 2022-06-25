package command

import (
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/wesleimp/fswatch/internal/testutils"
)

func TestRunCommand(t *testing.T) {
	tf, teardown := testutils.Mock(t)
	defer teardown()

	Run([]string{"echo", "hello"})

	time.Sleep(10 * time.Millisecond)

	fc, err := ioutil.ReadFile(tf.Name())
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(fc), "hello") {
		t.Fatal("command was not run")
	}
}
