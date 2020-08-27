// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package entity

type Student struct {
	Base
	StudentName    string `gorm:"column:student_name;size:64;index;default:'';not null;" `
	StudentAge     string `gorm:"column:student_age;size:20;default:'';not null;" `
	StudentSchool  string `gorm:"column:student_school;size:64;default:'';not null;" `
	StudentGrade   string `gorm:"column:student_grade;size:64;default:'';not null;" `
	FatherName     string `gorm:"column:father_name;size:64;index;default:'';not null;" `
	MotherName     string `gorm:"column:mother_name;size:64;index;default:'';not null;" `
	FatherPhone    string `gorm:"column:father_phone;size:20;index;default:'';not null;" `
	MotherPhone    string `gorm:"column:mother_phone;size:20;index;default:'';not null;" `
	FatherJob      string `gorm:"column:father_job;size:100;index;default:'';not null;" `
	MotherJob      string `gorm:"column:mother_job;size:100;index;default:'';not null;" `
	HomeAddress    string `gorm:"column:home_address;size:255;default:'';not null;" `
	Creator        string `gorm:"column:creator;size:64;default:'';not null;" `
	CreatorCompany string `gorm:"column:creator_company;index;default:'';not null;" `
}
