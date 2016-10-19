package board

import (
	"fmt"
	"strconv"
)

type BoardMgr struct {
	data [] *Board
}

func NewBoardMgr() *BoardMgr {
	return &BoardMgr{data: make([]*Board, 0)}
}

func (mgr *BoardMgr) Inputboard() {

	var age string
	fmt.Println("insert age")
	fmt.Scan(&age)

	var name string
	fmt.Println("insert name")
	fmt.Scan(&name)

	mgr.data = append(mgr.data, &Board{data: map[string]string{"age": age, "name": name}})
}

func (manager *BoardMgr) Deleteboard() {
	manager.PrintAllData(false)
	var index int
	fmt.Scan(&index)
	if index == 0 {
        manager.data = manager.data[1:len(manager.data)]
	} else if index == len(manager.data) {
        manager.data = manager.data[1 : len(manager.data)-1]
	} else {
        arr1 := manager.data[0:index]
		arr2 := manager.data[index+1 : len(manager.data)]

		manager.data = make([]*Board, 0)
		manager.data = append(manager.data, arr1...)
		manager.data = append(manager.data, arr2...)
	}
}

func (manager *BoardMgr) Readboard() {
	manager.PrintAllData(false)
	var index int
	fmt.Scan(&index)
	manager.data[index].Print()
}

func (boardManager *BoardMgr) PrintAllData(isData bool) {
	for index, data := range boardManager.data {
		fmt.Println("index :" + strconv.Itoa(index))
		if isData == false {
			continue
		}
		data.Print()
	}
}
