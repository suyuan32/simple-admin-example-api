// Code generated by goctl. DO NOT EDIT.
package types

// The basic response with data | 基础带数据信息
// swagger:model BaseDataInfo
type BaseDataInfo struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response with data | 基础带数据信息
// swagger:model BaseListInfo
type BaseListInfo struct {
	// The total number of data | 数据总数
	Total uint64 `json:"total"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response without data | 基础不带数据信息
// swagger:model BaseMsgResp
type BaseMsgResp struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
}

// The simplest message | 最简单的信息
type SimpleMsg struct {
	// Message | 信息
	Msg string `json:"msg"`
}

// The page request parameters | 列表请求参数
// swagger:model PageInfo
type PageInfo struct {
	// Page number | 第几页
	// Required: true
	Page uint64 `json:"page" validate:"number"`
	// Page size | 单页数据行数
	// Required: true
	// Maximum: 100000
	PageSize uint64 `json:"pageSize" validate:"number,max=100000"`
}

// Basic ID request | 基础ID参数请求
// swagger:model IDReq
type IDReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
}

// Basic IDs request | 基础ID数组参数请求
// swagger:model IDsReq
type IDsReq struct {
	// IDs
	// Required: true
	Ids []uint64 `json:"ids"`
}

// Basic ID request in path | 基础ID地址参数请求
// swagger:model IDPathReq
type IDPathReq struct {
	// ID
	// Required: true
	Id uint64 `path:"id"`
}

// Basic UUID request | 基础UUID参数请求
// swagger:model UUIDReq
type UUIDReq struct {
	// UUID
	// Required: true
	// Max length: 36
	UUID string `json:"UUID" validate:"len=36"`
}

// The base response data | 基础信息
// swagger:model BaseInfo
type BaseInfo struct {
	// ID
	Id uint64 `json:"id"`
	// Create date | 创建日期
	CreatedAt int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt int64 `json:"updatedAt,optional"`
}

// The request params of setting boolean status | 设置状态参数
// swagger:model StatusCodeReq
type StatusCodeReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
	// Status code | 状态码
	// Required: true
	Status uint64 `json:"status" validate:"number"`
}

// The response data of Student information | Student信息
// swagger:model StudentInfo
type StudentInfo struct {
	BaseInfo
	// Name
	Name string `json:"name"`
	// Age
	Age int64 `json:"age"`
	// AgeInt32
	AgeInt32 int32 `json:"ageInt32"`
	// AgeInt64
	AgeInt64 int64 `json:"ageInt64"`
	// AgeUint
	AgeUint uint64 `json:"ageUint"`
	// AgeUint32
	AgeUint32 uint32 `json:"ageUint32"`
	// AgeUint64
	AgeUint64 uint64 `json:"ageUint64"`
	// WeightFloat
	WeightFloat float64 `json:"weightFloat"`
	// WeightFloat32
	WeightFloat32 float32 `json:"weightFloat32"`
	// ClassId
	ClassId string `json:"classId"`
	// EnrollAt
	EnrollAt int64 `json:"enrollAt"`
	// StatusBool
	StatusBool bool `json:"statusBool"`
}

// Create or update Student information request | 创建或更新Student信息
// swagger:model CreateOrUpdateStudentReq
type CreateOrUpdateStudentReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id"`
	// Name
	Name string `json:"name"`
	// Age
	Age int64 `json:"age"`
	// AgeInt32
	AgeInt32 int32 `json:"ageInt32"`
	// AgeInt64
	AgeInt64 int64 `json:"ageInt64"`
	// AgeUint
	AgeUint uint64 `json:"ageUint"`
	// AgeUint32
	AgeUint32 uint32 `json:"ageUint32"`
	// AgeUint64
	AgeUint64 uint64 `json:"ageUint64"`
	// WeightFloat
	WeightFloat float64 `json:"weightFloat"`
	// WeightFloat32
	WeightFloat32 float32 `json:"weightFloat32"`
	// ClassId
	ClassId string `json:"classId"`
	// EnrollAt
	EnrollAt int64 `json:"enrollAt"`
	// StatusBool
	StatusBool bool `json:"statusBool"`
}

// The response data of Student list | Student列表数据
// swagger:model StudentListResp
type StudentListResp struct {
	BaseDataInfo
	// Student list data | API 列表数据
	Data StudentListInfo `json:"data"`
}

// Student list data | Student 列表数据
// swagger:model StudentListInfo
type StudentListInfo struct {
	BaseListInfo
	// The API list data | API列表数据
	Data []StudentInfo `json:"data"`
}

// Get Student list request params | Student列表请求参数
// swagger:model StudentListReq
type StudentListReq struct {
	PageInfo
	// Name
	Name string `json:"name"`
	// ClassId
	ClassId string `json:"classId"`
}
