// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package bizv1

import (
	eerrors "github.com/1477921168/ego/core/eerrors"
	codes "google.golang.org/grpc/codes"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the ego package it is being compiled against.
const _ = eerrors.SupportPackageIsVersion1

var errInvalid *eerrors.EgoError
var errUserNotFound *eerrors.EgoError
var errContentMissing *eerrors.EgoError

func init() {
	errInvalid = eerrors.New(int(codes.Unknown), "biz.v1.ERR_INVALID", Err_ERR_INVALID.String())
	eerrors.Register(errInvalid)
	errUserNotFound = eerrors.New(int(codes.NotFound), "biz.v1.ERR_USER_NOT_FOUND", Err_ERR_USER_NOT_FOUND.String())
	eerrors.Register(errUserNotFound)
	errContentMissing = eerrors.New(int(codes.InvalidArgument), "biz.v1.ERR_CONTENT_MISSING", Err_ERR_CONTENT_MISSING.String())
	eerrors.Register(errContentMissing)
}

func ErrInvalid() eerrors.Error {
	return errInvalid
}

func ErrUserNotFound() eerrors.Error {
	return errUserNotFound
}

func ErrContentMissing() eerrors.Error {
	return errContentMissing
}
