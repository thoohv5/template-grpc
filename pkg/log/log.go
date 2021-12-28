package log

import "context"

type IField interface {
	Key() string
	Value() interface{}
}

type ILogger interface {
	Debug(msg string, fields ...IField)
	Info(msg string, fields ...IField)
	Warn(msg string, fields ...IField)
	Error(msg string, fields ...IField)
	KVContext(ctx context.Context) IField
	KVError(err error) IField
	KVString(key string, value string) IField
	KVInt32(key string, value int32) IField
}

type LoggerType string

const (
	Zap LoggerType = "zap"
)

/** KV **/

type fieldContext struct {
	key string
	ctx context.Context
}

func NewFieldContext(ctx context.Context) IField {
	return &fieldContext{
		key: "ctx",
		ctx: ctx,
	}
}

func (f *fieldContext) Key() string {
	return f.key
}

func (f *fieldContext) Value() interface{} {
	return f.ctx
}

func KVContext(ctx context.Context) IField {
	return NewFieldContext(ctx)
}

type fieldError struct {
	key string
	err error
}

func NewFieldError(err error) IField {
	return &fieldError{
		key: "err",
		err: err,
	}
}

func (f *fieldError) Key() string {
	return f.key
}

func (f *fieldError) Value() interface{} {
	return f.err
}

func KVError(err error) IField {
	return NewFieldError(err)
}

type fieldString struct {
	key   string
	value string
}

func NewFieldString(key string, value string) IField {
	return &fieldString{
		key:   key,
		value: value,
	}
}

func (f *fieldString) Key() string {
	return f.key
}

func (f *fieldString) Value() interface{} {
	return f.value
}

func KVString(key string, value string) IField {
	return NewFieldString(key, value)
}

type fieldInt32 struct {
	key   string
	value int32
}

func NewFieldInt32(key string, value int32) IField {
	return &fieldInt32{
		key:   key,
		value: value,
	}
}

func (f *fieldInt32) Key() string {
	return f.key
}

func (f *fieldInt32) Value() interface{} {
	return f.value
}

func KVInt32(key string, value int32) IField {
	return NewFieldInt32(key, value)
}
