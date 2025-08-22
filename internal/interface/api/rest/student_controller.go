package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	// "github.com/tranvu1111/go-students-new/internal/interface/api/rest/dto/response"
	"github.com/tranvu1111/go-students-new/internal/application/interfaces"

	// "github.com/tranvu1111/go-students-new/internal/application/services"
	"github.com/tranvu1111/go-students-new/internal/interface/api/rest/dto/mapper"
	"github.com/tranvu1111/go-students-new/internal/interface/api/rest/dto/request"
	// "github.com/tranvu1111/go-students-new/internal/interface/api/rest/dto/response"
)

type StudentController struct {
	service interfaces.StudentService
}

func NewStudentController(r *gin.Engine, service interfaces.StudentService) *StudentController {
	controller := &StudentController{
		service: service,
	}
	// API group with JWT authentication
	r.POST("/api/v1/students", controller.CreateStudentController)
	r.GET("/api/v1/students", controller.GetAllStudentController)
	r.GET("/api/v1/students/:id", controller.GetStudentByIdController)
	r.PUT("/api/v1/students", controller.PutStudentController)
	// r.DELETE("/api/v1/students/:id", controller.DeleteStudentController)

	return controller
}



func (sc *StudentController) CreateStudentController( c *gin.Context )  {
	var createStudentRequest request.CreateStudentRequest

	if err := c.ShouldBindJSON(&createStudentRequest); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"message": "Invalid request","error1":err.Error()} )
		return
	}

	createStudentCommand ,err := createStudentRequest.ToCreateStudentCommand()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seller Id format"})
		return
	}

	commandStudentResult, err := sc.service.CreateStudent(createStudentCommand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student", "content":err.Error()})
		return 
	}
	fmt.Printf("result : %v", commandStudentResult.Result.StudentID)

	response := mapper.ToStudentResponse(commandStudentResult.Result)
	c.JSON(http.StatusCreated, gin.H{"message ": "Create a student successfully", "student" : response } )
}

func (sc *StudentController) GetAllStudentController(c *gin.Context) {
	sellers , err := sc.service.FindAllStudent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to load all students", "content":err.Error()})
		return
	}

	response := mapper.ToStudentListResponse(sellers.Result)
	c.JSON(http.StatusOK, response)
	

}

func (sc *StudentController) GetStudentByIdController(c *gin.Context) {
	idRaw := c.Request.URL.Path[len("/api/v1/students/"):]

	id, err := uuid.Parse(idRaw)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student Id format", "context" :err.Error()})
		return
	}

	student , err := sc.service.FindStudentById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Faild to find the student by their ID", "content" : err.Error()})
		return
	}

	if student == nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "Student not found"})
		return
	}

	response := mapper.ToStudentResponse(student.Result)

	c.JSON(http.StatusOK, response)
	
}

func (sc *StudentController) PutStudentController(c *gin.Context) {
	var updateRequest request.UpdateStudentResquest

	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"message": "Invalid request","error1":err.Error()} )
		return
	}

	updateStudentCommand , err := updateRequest.ToUpdateStudentCommand()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Failed to create a update student command", "content" : err.Error() })
		return
	}

	commandResult , err := sc.service.UpdateStudent(updateStudentCommand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to update student (in service)" , "content" : err.Error()})
	}

	response := mapper.ToStudentResponse(commandResult.Result)
	c.JSON(http.StatusOK, response)
	

}

