package ipfsclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"ipfs-store/internal/utils"

	"github.com/ipfs/go-cid"
)

const (
	ClusterHost = "http://ipfs.default.svc.cluster.local:9094"
)

type IPFSClient struct {
	httpClient *http.Client
}

func NewIPFSClient() *IPFSClient {
	return &IPFSClient{
		httpClient: &http.Client{
			Timeout: 30 * time.Minute, // 大文件超时设置
		},
	}
}

func (c *IPFSClient) PinDirect(buf *utils.Buffer, filename string) (cid.Cid, error) {
	var data []byte
	buf.Mux.RLock()
	data = buf.Data[:buf.Length()]
	buf.Mux.RUnlock()
	// 严格对齐 CLI 默认参数
	params := url.Values{
		"chunker":             []string{"size-262144"},
		"cid-version":         []string{"1"},
		"format":              []string{"unixfs"},
		"hash":                []string{"sha2-256"},
		"raw-leaves":          []string{"false"},
		"wrap-with-directory": []string{"false"},
		"local":               []string{"false"},
		"nocopy":              []string{"false"},
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 必须使用 "file" 作为字段名
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return cid.Undef, fmt.Errorf("create form file failed: %v", err)
	}
	if _, err = part.Write(data); err != nil {
		return cid.Undef, fmt.Errorf("write data failed: %v", err)
	}
	writer.Close()

	// 精确构造请求 URL
	reqUrl := fmt.Sprintf("%s/add?%s&name=%s",
		ClusterHost,
		params.Encode(), // 先编码基础参数
		url.QueryEscape(filename),
	)

	req, err := http.NewRequest("POST", reqUrl, body)
	if err != nil {
		return cid.Undef, fmt.Errorf("create request failed: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 启用详细日志（调试时）
	fmt.Printf("DEBUG: Sending request to %s\n", req.URL.Redacted())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return cid.Undef, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	// 详细错误处理
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return cid.Undef, fmt.Errorf("HTTP %d error: %s", resp.StatusCode, string(bodyBytes))
	}

	var result struct {
		Cid  string `json:"cid"`
		Name string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return cid.Undef, fmt.Errorf("decode response failed: %v", err)
	}

	time.Sleep(1 * time.Second)

	return cid.Decode(result.Cid)
}
