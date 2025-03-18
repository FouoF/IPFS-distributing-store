package ipfsclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/ipfs/go-cid"
)

const (
	ClusterHost = "http://127.0.0.1:9094" // Cluster API 默认端口 9094
	IPFSHost    = "http://127.0.0.1:5001" // IPFS API 默认端口 5001
)

type IPFSClient struct {
	httpClient *http.Client // 复用 HTTP 客户端
}

// 新版本：直接使用 HTTP API 的客户端
func NewIPFSClient() *IPFSClient {
	return &IPFSClient{
		httpClient: &http.Client{},
	}
}

// 方法一：直接通过 Cluster 的 /add 接口上传内存数据（推荐）
func (c *IPFSClient) PinDirect(data []byte, filename string) (cid.Cid, error) {
	if filename == "" {
		filename = "unnamed" // 可设置默认文件名
	}

	// 1. 构造 multipart 请求（含文件名）
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filename) // 关键修改：使用传入的文件名
	if err != nil {
		return cid.Undef, fmt.Errorf("create form file failed: %v", err)
	}
	if _, err = part.Write(data); err != nil {
		return cid.Undef, fmt.Errorf("write data failed: %v", err)
	}
	writer.Close()

	// 2. 发送到 Cluster /add 接口
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/add?cid-version=1&progress=false&name=%s", ClusterHost, filename), // 将文件名加入查询参数
		body,
	)
	if err != nil {
		return cid.Undef, fmt.Errorf("create request failed: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return cid.Undef, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	// 3. 解析响应并返回 CID
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return cid.Undef, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Cid  string `json:"cid"`
		Name string `json:"name"` // Cluster 会返回记录的文件名
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return cid.Undef, fmt.Errorf("decode response failed: %v", err)
	}

	parsedCid, err := cid.Decode(result.Cid)
	if err != nil {
		return cid.Undef, fmt.Errorf("invalid CID: %v", err)
	}

	return parsedCid, nil
}

// 方法二：通过 IPFS API 上传 + Cluster 钉住（分步操作）
func (c *IPFSClient) PinViaIPFS(data []byte, filename string) (cid.Cid, error) {
	// 1. 上传到 IPFS（含文件名）
	ipfsCID, err := c.addToIPFS(data, filename)
	if err != nil {
		return cid.Undef, err
	}

	// 2. 钉住到 Cluster 并记录文件名
	if err := c.pinOnCluster(ipfsCID, filename); err != nil {
		return cid.Undef, err
	}

	return ipfsCID, nil
}

// 辅助方法：通过 IPFS HTTP API 添加数据
func (c *IPFSClient) addToIPFS(data []byte, filename string) (cid.Cid, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filename) // 使用传入的文件名
	if err != nil {
		return cid.Undef, err
	}
	part.Write(data)
	writer.Close()

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/api/v0/add?cid-version=1&progress=false", IPFSHost),
		body,
	)
	if err != nil {
		return cid.Undef, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return cid.Undef, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return cid.Undef, fmt.Errorf("IPFS add failed: %s", resp.Status)
	}

	var ipfsResult struct {
		Hash string `json:"Hash"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&ipfsResult); err != nil {
		return cid.Undef, err
	}

	return cid.Decode(ipfsResult.Hash)
}

// 辅助方法：通过 Cluster HTTP API 钉住 CID
func (c *IPFSClient) pinOnCluster(cid cid.Cid, filename string) error {
	// 构造包含文件名的元数据
	reqBody := bytes.NewBufferString(
		fmt.Sprintf(`{
			"cid": "%s",
			"replication_factor": -1,
			"name": "%s",          // 将文件名存入 Cluster 元数据
			"metadata": {
				"source": "go-client",
				"filename": "%s"   // 可选：额外元数据字段
			}
		}`, cid.String(), filename, filename),
	)

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/pins/%s", ClusterHost, cid.String()),
		reqBody,
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("cluster pin failed: %s", resp.Status)
	}
	return nil
}