package main

import (
	"errors"
	"fmt"
)

type StudentMgr struct {
	allStudent []*Student
}

var (
	instance *StudentMgr
)

func GetStudentMgrInstance() *StudentMgr {
	if instance == nil {
		instance = new(StudentMgr)
		instance.init()
	}

	return instance
}

func (p *StudentMgr) AddStudent() error {
	stu := InputStudent()

	for _, v := range p.allStudent {
		if v.Username == stu.Username {
			errInfo := fmt.Sprintf("user %s have exist\n", stu.Username)
			return errors.New(errInfo)
		}
	}
	p.allStudent = append(p.allStudent, stu)
	fmt.Printf("insert usr %s success\n", stu.Username)

	return nil
}

func (p *StudentMgr) ModifyStudent() error {
	stu := InputStudent()

	for index, v := range p.allStudent {
		if v.Username == stu.Username {
			p.allStudent[index] = stu // update
			fmt.Printf("modify usr %s success\n", v.Username)
			return nil
		}
	}
	errInfo := fmt.Sprintf("user %s is not found\n", stu.Username)

	return errors.New(errInfo)
}

func (p *StudentMgr) ShowAllStudent() {
	for _, v := range p.allStudent {
		fmt.Printf("user:%s info:%#v\n", v.Username, v)
	}
}

func (p *StudentMgr) init() {
	p.allStudent = make([]*Student, 0)
}
