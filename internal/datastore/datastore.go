package datastore

import "errors"

type TreeStore struct {
	root *Node
}

type Index struct {
	Name     string
	L1       string
	L2       string
	LeafName string
}

type Node struct {
	Isleaf   bool
	Name     string
	Children []*Node
	Cid      string
}

func (s *TreeStore) AddIndex(idx Index) error {
	cur := s.root
	//一个index至少要包含一个有效字段
	if idx.Name == "" {
		return errors.New("name is empty")
	}
	if idx.LeafName != "" {
		return errors.New("leafname is not empty")
	}
	for _, v := range []string{idx.Name, idx.L1, idx.L2} {
		if v == "" {
			break
		}
		tmp := s.findNode(cur, v)
		if tmp == nil {
			tmp = &Node{Name: v}
			cur.Children = append(cur.Children, tmp)
		}
		cur = tmp
	}
	return nil
}

func (s *TreeStore) GetNode(idx Index) *Node {
	return s.findNodeByIndex(idx)
}

func (s *TreeStore) AddLeaf(idx Index, cid string) error {
	node := s.findNodeByIndex(idx)
	if node == nil {
		return errors.New("invalid index")
	}
	node.Children = append(node.Children, &Node{Name: idx.LeafName, Cid: cid, Isleaf: true})
	return nil
}

//查找等于指定key的节点
func (s *TreeStore) findNode(cur *Node, key string) *Node {
	for _, v := range cur.Children {
		if v.Name == key {
			return v
		}
	}
	return nil
}

func (s *TreeStore) findNodeByIndex(idx Index) *Node {
	cur := s.root
	for _, v := range []string{idx.Name, idx.L1, idx.L2} {
		if v == "" {
			return cur
		}
		tmp := s.findNode(cur, v)
		if tmp == nil {
			return nil
		}
		cur = tmp
	}
	return cur

}
func NewTreeStore() *TreeStore {
	return &TreeStore{
		root: &Node{Name: "root"},
	}
}

func NewTreeStoreWithData() *TreeStore {
	TreeStore := NewTreeStore()
	TreeStore.AddIndex(Index{Name: "监控", L1: "一号楼", L2: "1号房间", LeafName: ""})
	TreeStore.AddIndex(Index{Name: "监控", L1: "一号楼", L2: "2号房间", LeafName: ""})
	TreeStore.AddIndex(Index{Name: "监控", L1: "二号楼", L2: "1号房间", LeafName: ""})
	TreeStore.AddIndex(Index{Name: "监控", L1: "二号楼", L2: "2号房间", LeafName: ""})
	TreeStore.AddIndex(Index{Name: "心率", L1: "1号房间", L2: "1号床", LeafName: ""})
	TreeStore.AddIndex(Index{Name: "心率", L1: "2号房间", L2: "1号床", LeafName: ""})
	TreeStore.AddIndex(Index{Name: "心率", L1: "1号房间", L2: "2号床", LeafName: ""})
	TreeStore.AddIndex(Index{Name: "心率", L1: "2号房间", L2: "2号床", LeafName: ""})
	TreeStore.AddIndex(Index{Name: "血压", L1: "1号房间", L2: "1号床", LeafName: ""})
	TreeStore.AddIndex(Index{Name: "血压", L1: "2号房间", L2: "1号床", LeafName: ""})
	TreeStore.AddIndex(Index{Name: "血压", L1: "1号房间", L2: "2号床", LeafName: ""})
	TreeStore.AddIndex(Index{Name: "血压", L1: "2号房间", L2: "2号床", LeafName: ""})
	TreeStore.AddIndex(Index{Name: "温度", L1: "一号楼", L2: "1号房间", LeafName: ""})
	TreeStore.AddIndex(Index{Name: "温度", L1: "一号楼", L2: "2号房间", LeafName: ""})
	TreeStore.AddIndex(Index{Name: "温度", L1: "二号楼", L2: "1号房间", LeafName: ""})
	TreeStore.AddIndex(Index{Name: "温度", L1: "二号楼", L2: "2号房间", LeafName: ""})
	return TreeStore
}
