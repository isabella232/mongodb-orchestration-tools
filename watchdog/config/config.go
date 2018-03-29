package config

import (
	"time"

	"github.com/percona/dcos-mongo-tools/common"
	"github.com/percona/dcos-mongo-tools/common/api"
)

var (
	DefaultAPIPoll               = "10s"
	DefaultDelayWatcher          = "20s"
	DefaultReplsetPoll           = "5s"
	DefaultReplsetTimeout        = "3s"
	DefaultReplsetConfUpdatePoll = "10s"
)

// Watchdog Configuration
type Config struct {
	Tool                  *common.ToolConfig
	Username              string
	Password              string
	FrameworkName         string
	API                   *api.Config
	APIPoll               time.Duration
	ReplsetPoll           time.Duration
	ReplsetTimeout        time.Duration
	ReplsetConfUpdatePoll time.Duration
	DelayWatcher          time.Duration
}
