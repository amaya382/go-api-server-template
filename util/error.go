package util

import "fmt"

type UniqueViolationErr struct {
	TableName string
	FieldName string
	Value     string
}

func NewUniqueViolationErr(tableName string, fieldName string, value string) *UniqueViolationErr {
	return &UniqueViolationErr{tableName, fieldName, value}
}

func (self UniqueViolationErr) Error() string {
	return fmt.Sprintf("%s was submitted. But already exists (%s on %s)\n",
		self.Value, self.FieldName, self.TableName)
}

type RecordNotFoundErr struct {
	Target string
}

func NewRecordNotFoundErr(target string) *RecordNotFoundErr {
	return &RecordNotFoundErr{target}
}

func (self RecordNotFoundErr) Error() string {
	return "Record not found: " + self.Target
}

type ErrParamContents struct {
	IsValid bool
	Reason  string
}

func NewErrParamContents(isValid bool, reason string) *ErrParamContents {
	return &ErrParamContents{isValid, reason}
}

type InvalidRequest struct {
	Message   string
	ErrParams map[string]*ErrParamContents
}

func NewInvalidRequest(message string, errParams map[string]*ErrParamContents) *InvalidRequest {
	return &InvalidRequest{message, errParams}
}

func (self InvalidRequest) Error() string {
	return self.Message
}
