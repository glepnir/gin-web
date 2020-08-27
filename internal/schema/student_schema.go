// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package schema

type CreateStudentSchema struct {
	StudentName    string `json:"student_name"`
	StudentAge     string `json:"student_age"`
	StudentSchool  string `json:"student_school"`
	StudentGrade   string `json:"student_grade"`
	FatherName     string `json:"father_name"`
	MotherName     string `json:"mother_name"`
	FatherPhone    string `json:"father_phone"`
	MotherPhone    string `json:"mother_phone"`
	FatherJob      string `json:"father_job"`
	MotherJob      string `json:"mother_job"`
	HomeAddress    string `json:"home_address"`
	Creator        string `json:"creator"`
	CreatorCompany string `json:"creator_company"`
}
