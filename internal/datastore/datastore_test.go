package datastore

import (
	"testing"
)

// TestNewTreeStore 测试 NewTreeStore 函数
func TestNewTreeStore(t *testing.T) {
	store := NewTreeStore()
	if store.root == nil {
		t.Errorf("Expected root node to be initialized, got nil")
	}
	if store.root.Name != "root" {
		t.Errorf("Expected root node name to be 'root', got %s", store.root.Name)
	}
}

// TestAddIndex 测试 AddIndex 方法
func TestAddIndex(t *testing.T) {
	store := NewTreeStore()
	idx := Index{Name: "test", L1: "level1", L2: "level2"}

	err := store.AddIndex(idx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 检查节点是否正确添加
	cur := store.root
	for _, v := range []string{"test", "level1", "level2"} {
		tmp := store.findNode(cur, v)
		if tmp == nil {
			t.Errorf("Expected node %s to exist, got nil", v)
		}
		cur = tmp
	}
}

// TestAddIndex_LeafNameNotEmpty 测试 AddIndex 方法当 LeafName 不为空时
func TestAddIndex_LeafNameNotEmpty(t *testing.T) {
	store := NewTreeStore()
	idx := Index{Name: "test", LeafName: "leaf"}

	err := store.AddIndex(idx)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if err.Error() != "leafname is not empty" {
		t.Errorf("Expected error 'leafname is not empty', got %v", err)
	}
}

// TestGetNode 测试 GetNode 方法
func TestGetNode(t *testing.T) {
	store := NewTreeStore()
	idx := Index{Name: "test", L1: "level1", L2: "level2"}

	err := store.AddIndex(idx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	node := store.GetNode(idx)
	if node == nil {
		t.Errorf("Expected node to exist, got nil")
	}
	if node.Name != "level2" {
		t.Errorf("Expected node name to be 'level2', got %s", node.Name)
	}
}

// TestAddLeaf 测试 AddLeaf 方法
func TestAddLeaf(t *testing.T) {
	store := NewTreeStore()
	idx := Index{Name: "test", L1: "level1", L2: "level2", LeafName: "leaf"}
	cid := "QmHash"

	err := store.AddIndex(Index{Name: "test", L1: "level1", L2: "level2"})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = store.AddLeaf(idx, cid)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	node := store.GetNode(Index{Name: "test", L1: "level1", L2: "level2"})
	if node == nil {
		t.Errorf("Expected node to exist, got nil")
	}

	leaf := store.findNode(node, "leaf")
	if leaf == nil {
		t.Errorf("Expected leaf to exist, got nil")
	}
	if leaf.Cid != cid {
		t.Errorf("Expected leaf CID to be %s, got %s", cid, leaf.Cid)
	}
	if !leaf.Isleaf {
		t.Errorf("Expected leaf Isleaf to be true, got false")
	}
}

// TestAddLeaf_InvalidIndex 测试 AddLeaf 方法当索引无效时
func TestAddLeaf_InvalidIndex(t *testing.T) {
	store := NewTreeStore()
	idx := Index{Name: "test", L1: "level1", L2: "level2", LeafName: "leaf"}
	cid := "QmHash"

	err := store.AddLeaf(idx, cid)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if err.Error() != "invalid index" {
		t.Errorf("Expected error 'invalid index', got %v", err)
	}
}
