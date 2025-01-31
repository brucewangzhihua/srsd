package discovery

import (
	"time"

	"github.com/brucewangzhihua/srsd/selector"
)

var (
	defaultPrefix    = "/srsd/services/"
	defaultAddresses = []string{"127.0.0.1:2379"}
	defaultTimeout   = 5 * time.Second
	defaultSelectors = []selector.Selector{selector.NewRound()} // 默认使用循环选择器
)

// Option 设置服务注册参数
type Option func(*Options)

// Watch watch回调函数
type Watch func(event int32, host string)

// Options 服务注册参数
type Options struct {
	Addresses []string            // etcd地址
	Username  string              // etcd用户名
	Password  string              // etcd密码
	Prefix    string              //服务注册前缀
	Timeout   time.Duration       // etcd超时时间
	Watch     func(event *Event)  // 服务发生变化时回调函数
	Selectors []selector.Selector // 服务发现
}

// newOptions 创建服务注册参数对象
func newOptions(opts ...Option) *Options {
	opt := &Options{
		Addresses: defaultAddresses,
		Prefix:    defaultPrefix,
		Timeout:   defaultTimeout,
		Selectors: defaultSelectors,
	}

	for _, one := range opts {
		one(opt)
	}
	return opt
}

// Addresses 设置etcd地址
func Addresses(addresses []string) Option {
	return func(opt *Options) {
		opt.Addresses = addresses
	}
}

// Username 设置etcd用户名
func Username(userName string) Option {
	return func(opt *Options) {
		opt.Username = userName
	}
}

// Password 设置etcd密码
func Password(password string) Option {
	return func(opt *Options) {
		opt.Password = password
	}
}

// Prefix 设置服务发现前缀
func Prefix(prefix string) Option {
	if prefix != "" && prefix[len(prefix)-1] != '/' {
		prefix += "/"
	}

	return func(opt *Options) {
		opt.Prefix = prefix
	}
}

// Timeout 设置etcd超时时间
func Timeout(timeout time.Duration) Option {
	return func(opt *Options) {
		opt.Timeout = timeout
	}
}

// Selectors 设置服务发选择器
func Selectors(selectors ...selector.Selector) Option {
	return func(opt *Options) {
		opt.Selectors = selectors
	}
}
