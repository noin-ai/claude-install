package installer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// TestConnection 测试 API 连接
func TestConnection() error {
	fmt.Println("🧪 正在测试 API 连接...")
	fmt.Println()

	// 读取配置
	config, err := loadConfig()
	if err != nil {
		return fmt.Errorf("无法读取配置: %w\n请先运行 'claude-install config' 进行配置", err)
	}

	// 准备测试请求
	testPayload := map[string]interface{}{
		"model": "claude-3-5-sonnet-20241022",
		"messages": []map[string]string{
			{"role": "user", "content": "Hi"},
		},
		"max_tokens": 10,
	}

	payloadBytes, _ := json.Marshal(testPayload)

	// 发送请求
	req, err := http.NewRequest("POST", config.APIBaseURL+"/v1/messages", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", config.APIKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	fmt.Printf("   正在连接: %s\n", config.APIBaseURL)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("❌ 连接失败: %w\n请检查网络和 API Base URL", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// 检查响应
	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		fmt.Println("   ✅ 连接成功！")
		fmt.Println("   ✅ API Key 有效")
		fmt.Println()
		fmt.Println("🎉 恭喜！Claude API 配置正确，可以正常使用！")
		return nil
	}

	// 处理错误
	var errResp map[string]interface{}
	if err := json.Unmarshal(body, &errResp); err == nil {
		if errData, ok := errResp["error"].(map[string]interface{}); ok {
			errMsg := errData["message"].(string)
			return fmt.Errorf("❌ API 错误 (%d): %s", resp.StatusCode, errMsg)
		}
	}

	return fmt.Errorf("❌ 请求失败 (HTTP %d)\n响应: %s", resp.StatusCode, string(body))
}
