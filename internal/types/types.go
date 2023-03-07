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
// swagger:response SimpleMsg
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

// Basic ID request | 基础ID地址参数请求
// swagger:model IDPathReq
type IDPathReq struct {
	// ID
	// Required: true
	Id uint64 `path:"id"`
}

// Basic UUID request | 基础UUID参数请求
// swagger:model UUIDReq
type UUIDReq struct {
	// ID
	// Required: true
	// Max length: 36
	Id string `json:"id" validate:"len=36"`
}

// Basic UUID array request | 基础UUID数组参数请求
// swagger:model UUIDsReq
type UUIDsReq struct {
	// Ids
	// Required: true
	Ids []string `json:"ids"`
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

// The base UUID response data | 基础信息
// swagger:model BaseUUIDInfo
type BaseUUIDInfo struct {
	// ID
	Id string `json:"id"`
	// Create date | 创建日期
	CreatedAt int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt int64 `json:"updatedAt,optional"`
}

// The response data of teacher information | Teacher信息
// swagger:model TeacherInfo
type TeacherInfo struct {
	BaseUUIDInfo
	// Name
	Name string `json:"name,optional"`
	// Age
	Age int64 `json:"age,optional"`
	// AgeInt32
	AgeInt32 int32 `json:"ageInt32,optional"`
	// AgeInt64
	AgeInt64 int64 `json:"ageInt64,optional"`
	// AgeUint
	AgeUint uint64 `json:"ageUint,optional"`
	// AgeUint32
	AgeUint32 uint32 `json:"ageUint32,optional"`
	// AgeUint64
	AgeUint64 uint64 `json:"ageUint64,optional"`
	// WeightFloat
	WeightFloat float64 `json:"weightFloat,optional"`
	// WeightFloat32
	WeightFloat32 float32 `json:"weightFloat32,optional"`
	// ClassId
	ClassId string `json:"classId,optional"`
	// EnrollAt
	EnrollAt int64 `json:"enrollAt,optional"`
	// StatusBool
	StatusBool bool `json:"statusBool,optional"`
}

// The response data of teacher list | Teacher列表数据
// swagger:model TeacherListResp
type TeacherListResp struct {
	BaseDataInfo
	// Teacher list data | Teacher列表数据
	Data TeacherListInfo `json:"data"`
}

// Teacher list data | Teacher列表数据
// swagger:model TeacherListInfo
type TeacherListInfo struct {
	BaseListInfo
	// The API list data | Teacher列表数据
	Data []TeacherInfo `json:"data"`
}

// Get teacher list request params | Teacher列表请求参数
// swagger:model TeacherListReq
type TeacherListReq struct {
	PageInfo
	// Name
	Name string `json:"name,optional"`
}

// Teacher information response | Teacher信息返回体
// swagger:model TeacherInfoResp
type TeacherInfoResp struct {
	BaseDataInfo
	// Teacher information | Teacher数据
	Data TeacherInfo `json:"data"`
}

// The response data of student information | Student信息
// swagger:model StudentInfo
type StudentInfo struct {
	BaseInfo
	// Name
	Name string `json:"name,optional"`
	// Age
	Age int64 `json:"age,optional"`
	// AgeInt32
	AgeInt32 int32 `json:"ageInt32,optional"`
	// AgeInt64
	AgeInt64 int64 `json:"ageInt64,optional"`
	// AgeUint
	AgeUint uint64 `json:"ageUint,optional"`
	// AgeUint32
	AgeUint32 uint32 `json:"ageUint32,optional"`
	// AgeUint64
	AgeUint64 uint64 `json:"ageUint64,optional"`
	// WeightFloat
	WeightFloat float64 `json:"weightFloat,optional"`
	// WeightFloat32
	WeightFloat32 float32 `json:"weightFloat32,optional"`
	// ClassId
	ClassId string `json:"classId,optional"`
	// EnrollAt
	EnrollAt int64 `json:"enrollAt,optional"`
	// StatusBool
	StatusBool bool `json:"statusBool,optional"`
}

// The response data of student list | Student列表数据
// swagger:model StudentListResp
type StudentListResp struct {
	BaseDataInfo
	// Student list data | Student列表数据
	Data StudentListInfo `json:"data"`
}

// Student list data | Student列表数据
// swagger:model StudentListInfo
type StudentListInfo struct {
	BaseListInfo
	// The API list data | Student列表数据
	Data []StudentInfo `json:"data"`
}

// Get student list request params | Student列表请求参数
// swagger:model StudentListReq
type StudentListReq struct {
	PageInfo
	// Name
	Name string `json:"name,optional"`
}

// Student information response | Student信息返回体
// swagger:model StudentInfoResp
type StudentInfoResp struct {
	BaseDataInfo
	// Student information | Student数据
	Data StudentInfo `json:"data"`
}
