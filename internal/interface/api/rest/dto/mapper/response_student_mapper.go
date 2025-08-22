package mapper

import (
	"github.com/tranvu1111/go-students-new/internal/interface/api/rest/dto/response"
	"github.com/tranvu1111/go-students-new/internal/application/common"
	
)

func ToStudentResponse(studentResult *common.StudentResult) *response.StudentResponse {
	return &response.StudentResponse{
		StudentID:      studentResult.StudentID.String(),
		FirstName:      studentResult.FirstName,
		LastName:       studentResult.LastName,
		DateOfBirth: 	studentResult.DateOfBirth,
		Email:          studentResult.Email,
		Phone: 			studentResult.Phone,
		Major: 			studentResult.Major,			
		CreatedAt: 		studentResult.CreatedAt,
		UpdatedAt: 		studentResult.UpdatedAt,
		EnrollmentDate: studentResult.EnrollmentDate,			
	}
}

func ToStudentListResponse(students []*common.StudentResult) *response.StudentResponseList{
	var studentResponseList []*response.StudentResponse

	for _, v := range students {
		studentResponseList = append(studentResponseList,ToStudentResponse(v))
	}

	return &response.StudentResponseList{Students: studentResponseList}
}