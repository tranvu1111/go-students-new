package query

import (
	"github.com/tranvu1111/go-students-new/internal/application/common"

)

type StudentQueryResult struct {

	Result *common.StudentResult
}

type StudentQueryListResult struct {

	Result []*common.StudentResult
}