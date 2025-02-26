package config

import (
	"testing"
)

// 测试配置加载
func TestConfigLoad(t *testing.T) {
	// 这里 init() 函数会自动执行，你不需要调用

	// 验证 Global 变量是否成功加载
	if Global.DB.Host == "" {
		t.Errorf("MySQL 配置加载失败，Host 为空")
	} else {
		println(Global.DB.Host)
	}

	if Global.RDB.IP == "" {
		t.Errorf("Redis 配置加载失败，IP 为空")
	}

	if Global.StaticSourcePath == "" {
		t.Errorf("Path 配置加载失败，StaticSourcePath 为空")
	}

	t.Log("配置文件加载测试通过")
}
