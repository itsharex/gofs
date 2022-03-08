package conf

import (
	"github.com/no-src/gofs/core"
	"time"
)

// Config store all the flag info
type Config struct {
	// other
	PrintVersion bool   `json:"-" yaml:"-"`
	PrintAbout   bool   `json:"-" yaml:"-"`
	Conf         string `json:"-" yaml:"-"`

	// file sync
	Source                core.VFS `json:"source" yaml:"source"`
	Dest                  core.VFS `json:"dest" yaml:"dest"`
	SyncOnce              bool     `json:"sync_once" yaml:"sync_once"`
	SyncCron              string   `json:"sync_cron" yaml:"sync_cron"`
	EnableLogicallyDelete bool     `json:"logically_delete" yaml:"logically_delete"`
	ClearDeletedPath      bool     `json:"-" yaml:"-"`
	IgnoreConf            string   `json:"ignore_conf" yaml:"ignore_conf"`
	IgnoreDeletedPath     bool     `json:"ignore_deleted" yaml:"ignore_deleted"`
	ChunkSize             int64    `json:"chunk_size" yaml:"chunk_size"`

	// retry
	RetryCount int           `json:"retry_count" yaml:"retry_count"`
	RetryWait  time.Duration `json:"retry_wait" yaml:"retry_wait"`
	RetryAsync bool          `json:"retry_async" yaml:"retry_async"`

	// log
	LogLevel         int           `json:"log_level" yaml:"log_level"`
	EnableFileLogger bool          `json:"log_file" yaml:"log_file"`
	LogDir           string        `json:"log_dir" yaml:"log_dir"`
	LogFlush         bool          `json:"log_flush" yaml:"log_flush"`
	LogFlushInterval time.Duration `json:"log_flush_interval" yaml:"log_flush_interval"`
	EnableEventLog   bool          `json:"log_event" yaml:"log_event"`

	// daemon
	IsDaemon           bool          `json:"daemon" yaml:"daemon"`
	DaemonPid          bool          `json:"daemon_pid" yaml:"daemon_pid"`
	DaemonDelay        time.Duration `json:"daemon_delay" yaml:"daemon_delay"`
	DaemonMonitorDelay time.Duration `json:"daemon_monitor_delay" yaml:"daemon_monitor_delay"`
	KillPPid           bool          `json:"kill_ppid" yaml:"kill_ppid"`
	IsSubprocess       bool          `json:"-" yaml:"-"`

	// file server
	EnableFileServer         bool   `json:"server" yaml:"server"`
	FileServerAddr           string `json:"server_addr" yaml:"server_addr"`
	EnableFileServerCompress bool   `json:"server_compress" yaml:"server_compress"`
	EnableManage             bool   `json:"manage" yaml:"manage"`
	ManagePrivate            bool   `json:"manage_private" yaml:"manage_private"`
	EnablePushServer         bool   `json:"push_server" yaml:"push_server"`

	// tls transfer
	EnableTLS   bool   `json:"tls" yaml:"tls"`
	TLSCertFile string `json:"tls_cert_file" yaml:"tls_cert_file"`
	TLSKeyFile  string `json:"tls_key_file" yaml:"tls_key_file"`

	// login user
	Users             string `json:"users" yaml:"users"`
	RandomUserCount   int    `json:"rand_user_count" yaml:"rand_user_count"`
	RandomUserNameLen int    `json:"rand_user_len" yaml:"rand_user_len"`
	RandomPasswordLen int    `json:"rand_pwd_len" yaml:"rand_pwd_len"`
	RandomDefaultPerm string `json:"rand_perm" yaml:"rand_perm"`
}
