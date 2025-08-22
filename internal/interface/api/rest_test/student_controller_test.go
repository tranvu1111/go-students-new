package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tranvu1111/go-students-new/internal/application/command"
	"github.com/tranvu1111/go-students-new/internal/application/common"
	"github.com/tranvu1111/go-students-new/internal/interface/api/rest"
	// "github.com/tranvu1111/go-students-new/internal/interface/api/rest/dto/request"
)

func TestCreateStudent(t *testing.T) {
	// setup
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()	

	mockStudentService := new(MockStudentService)
	rest.NewStudentController(r, mockStudentService)


	
	reqBody := map[string]interface{}{
		"FirstName":"tran",
		"LastName":"vu",
		"DateOfBirth":"2003-03-11",
		"Email":"tranvu123312312@gmail.com",
		"Phone":"0931239991",
		"Major":"CNTT",
		"EnrollmentDate":"2023-03-11",
	}

	dob := time.Date(2003, 3, 11, 0, 0, 0, 0, time.UTC)
	phone := "0931239991"
	major := "CNTT"
	enrollment_date := time.Date(2023, 3, 11, 0, 0, 0, 0, time.UTC)
	createStudentCommandResult := &command.CreateStudentCommandResult{
		Result: &common.StudentResult{
			StudentID:uuid.New(),
			FirstName:"tran",
			LastName:"vu",
			DateOfBirth: &dob,
			Email :	"tranvu123312312@gmail.com", 	
			Phone :	&phone ,
			Major :	&major ,
			EnrollmentDate:enrollment_date,
		},
	}


	mockStudentService.On("CreateStudent", mock.Anything).Return(createStudentCommandResult, nil)
	
	reqBodyBytes, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/students", bytes.NewReader(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t,http.StatusCreated, w.Code)

	var responseBody map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	responseBody, ok := responseBody["student"].(map[string]interface{})
	if !ok {
		t.Fatalf("couldn't cast student to map[string]interface{}")
	}
		
	delete(responseBody, "CreatedAt")
	delete(responseBody, "UpdatedAt")
	delete(responseBody, "StudentID")
	delete(responseBody, "EnrollmentDate")
	delete(responseBody, "DateOfBirth")
	delete(reqBody, "DateOfBirth")
	delete(reqBody, "EnrollmentDate")

	assert.NoError(t,err)
	assert.Equal(t, reqBody, responseBody)

	mockStudentService.AssertExpectations(t)
}

func TestUpdateStudent(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()	

	mockStudentService := new(MockStudentService)
	rest.NewStudentController(r, mockStudentService)


	reqBody := map[string]interface{}{
		"StudentID":"6f69799c-1eb2-4266-b28c-9762a4d02129",	
		"DateOfBirth":"2003-03-11",		
		"Phone":"09312399912",
		"Major"	: "deptrai",
	}

	studentID , _ := uuid.Parse("6f69799c-1eb2-4266-b28c-9762a4d02129")
	dob := time.Date(2003, 3, 11, 0, 0, 0, 0, time.UTC)
	phone := "09312399912"
	major := "deptrai"
	enrollment_date := time.Date(2023, 3, 11, 0, 0, 0, 0, time.UTC)
	updateStudentCommandResult := &command.UpdateStudentCommandResult{
		Result: &common.StudentResult{
			StudentID: studentID,
			FirstName:"tran",
			LastName:"vu",
			DateOfBirth: &dob,
			Email :	"tranvu333@gmail.com", 	
			Phone :	&phone ,
			Major :	&major ,
			EnrollmentDate:enrollment_date,
		},
	}

	mockStudentService.On("UpdateStudent", mock.Anything).Return(updateStudentCommandResult, nil)

	reqBodyBytes, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPut, "/api/v1/students", bytes.NewReader(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t,http.StatusOK, w.Code)

	var responseBody map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	// responseBody, ok := responseBody["student"].(map[string]interface{})
	// if !ok {
	// 	t.Fatalf("couldn't cast student to map[string]interface{}")
	// }
		
	delete(responseBody, "CreatedAt")
	delete(responseBody, "UpdatedAt")
	
	delete(responseBody, "Email")
	delete(responseBody, "FirstName")
	delete(responseBody, "LastName")
	delete(responseBody, "EnrollmentDate")
	delete(responseBody, "DateOfBirth")
	delete(reqBody, "DateOfBirth")
	

	assert.NoError(t,err)
	assert.Equal(t, reqBody, responseBody)

	mockStudentService.AssertExpectations(t)

}