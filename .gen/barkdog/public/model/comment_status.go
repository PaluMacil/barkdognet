//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type CommentStatus string

const (
	CommentStatus_Approved CommentStatus = "approved"
	CommentStatus_Pending  CommentStatus = "pending"
	CommentStatus_Spam     CommentStatus = "spam"
)

func (e *CommentStatus) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("jet: Invalid scan value for AllTypesEnum enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "approved":
		*e = CommentStatus_Approved
	case "pending":
		*e = CommentStatus_Pending
	case "spam":
		*e = CommentStatus_Spam
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for CommentStatus enum")
	}

	return nil
}

func (e CommentStatus) String() string {
	return string(e)
}