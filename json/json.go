import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello from Go!",
	}

	// 将 response 转换为 JSON 格式
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 设置响应头部为 JSON 类型
	w.Header().Set("Content-Type", "application/json")

	// 将 JSON 数据写入响应
	w.Write(jsonData)
}
