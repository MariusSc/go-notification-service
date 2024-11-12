package receivers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReceivers(t *testing.T) {
	receivers := NewReceivers()

	assert.Equal(t, 1, len(receivers), "The function does not return the number of receivers that should be registered.")
	assert.IsType(t, &GithubIssueReceiver{}, receivers[0], "The receiver must be of type GithubIssueReceiver.")
}
