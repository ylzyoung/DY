package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {

	InitDB()

	assert.NotNil(t, DB)
	// 尝试执行一个简单的数据库查询，验证数据库连接是否正常
	var result int
	err := DB.Raw("SELECT 1").Scan(&result).Error
	assert.NoError(t, err)
	assert.Equal(t, 1, result)

	// 可以添加更多测试用例，例如验证表是否成功创建
	// 验证 AutoMigrate 是否成功
	// 验证是否能成功插入数据等等
}
