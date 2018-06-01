// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bizday.proto

/*
Package github_com_magaldima_bizday_api is a generated protocol buffer package.

It is generated from these files:
	bizday.proto

It has these top-level messages:
	Date
	HolidayCalendar
	BinaryDateRequest
	NumberOfDaysResponse
	UnaryDateRequest
	UnaryBoolResponse
	UnaryDateResponse
	UnaryTransformRequest
*/
package github_com_magaldima_bizday_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WeekDays int32

const (
	WeekDays_Sunday    WeekDays = 0
	WeekDays_Monday    WeekDays = 1
	WeekDays_Tuesday   WeekDays = 2
	WeekDays_Wednesday WeekDays = 3
	WeekDays_Thursday  WeekDays = 4
	WeekDays_Friday    WeekDays = 5
	WeekDays_Saturday  WeekDays = 6
)

var WeekDays_name = map[int32]string{
	0: "Sunday",
	1: "Monday",
	2: "Tuesday",
	3: "Wednesday",
	4: "Thursday",
	5: "Friday",
	6: "Saturday",
}
var WeekDays_value = map[string]int32{
	"Sunday":    0,
	"Monday":    1,
	"Tuesday":   2,
	"Wednesday": 3,
	"Thursday":  4,
	"Friday":    5,
	"Saturday":  6,
}

func (x WeekDays) String() string {
	return proto.EnumName(WeekDays_name, int32(x))
}
func (WeekDays) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Month int32

const (
	Month_January   Month = 0
	Month_February  Month = 1
	Month_March     Month = 2
	Month_April     Month = 3
	Month_May       Month = 4
	Month_June      Month = 5
	Month_July      Month = 6
	Month_August    Month = 7
	Month_September Month = 8
	Month_October   Month = 9
	Month_November  Month = 10
	Month_December  Month = 11
)

var Month_name = map[int32]string{
	0:  "January",
	1:  "February",
	2:  "March",
	3:  "April",
	4:  "May",
	5:  "June",
	6:  "July",
	7:  "August",
	8:  "September",
	9:  "October",
	10: "November",
	11: "December",
}
var Month_value = map[string]int32{
	"January":   0,
	"February":  1,
	"March":     2,
	"April":     3,
	"May":       4,
	"June":      5,
	"July":      6,
	"August":    7,
	"September": 8,
	"October":   9,
	"November":  10,
	"December":  11,
}

func (x Month) String() string {
	return proto.EnumName(Month_name, int32(x))
}
func (Month) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Date struct {
	Year  int32 `protobuf:"varint,1,opt,name=year" json:"year,omitempty"`
	Month Month `protobuf:"varint,2,opt,name=month,enum=github.com.magaldima.bizday.api.Month" json:"month,omitempty"`
	Day   int32 `protobuf:"varint,3,opt,name=day" json:"day,omitempty"`
}

func (m *Date) Reset()                    { *m = Date{} }
func (m *Date) String() string            { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()               {}
func (*Date) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Date) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Date) GetMonth() Month {
	if m != nil {
		return m.Month
	}
	return Month_January
}

func (m *Date) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

type HolidayCalendar struct {
	Name     string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Holidays []*Date `protobuf:"bytes,2,rep,name=holidays" json:"holidays,omitempty"`
}

func (m *HolidayCalendar) Reset()                    { *m = HolidayCalendar{} }
func (m *HolidayCalendar) String() string            { return proto.CompactTextString(m) }
func (*HolidayCalendar) ProtoMessage()               {}
func (*HolidayCalendar) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HolidayCalendar) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HolidayCalendar) GetHolidays() []*Date {
	if m != nil {
		return m.Holidays
	}
	return nil
}

// Binary Messages
type BinaryDateRequest struct {
	Cal   *HolidayCalendar `protobuf:"bytes,1,opt,name=cal" json:"cal,omitempty"`
	Start *Date            `protobuf:"bytes,2,opt,name=start" json:"start,omitempty"`
	End   *Date            `protobuf:"bytes,3,opt,name=end" json:"end,omitempty"`
}

func (m *BinaryDateRequest) Reset()                    { *m = BinaryDateRequest{} }
func (m *BinaryDateRequest) String() string            { return proto.CompactTextString(m) }
func (*BinaryDateRequest) ProtoMessage()               {}
func (*BinaryDateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BinaryDateRequest) GetCal() *HolidayCalendar {
	if m != nil {
		return m.Cal
	}
	return nil
}

func (m *BinaryDateRequest) GetStart() *Date {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *BinaryDateRequest) GetEnd() *Date {
	if m != nil {
		return m.End
	}
	return nil
}

type NumberOfDaysResponse struct {
	Days int32 `protobuf:"varint,1,opt,name=days" json:"days,omitempty"`
}

func (m *NumberOfDaysResponse) Reset()                    { *m = NumberOfDaysResponse{} }
func (m *NumberOfDaysResponse) String() string            { return proto.CompactTextString(m) }
func (*NumberOfDaysResponse) ProtoMessage()               {}
func (*NumberOfDaysResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *NumberOfDaysResponse) GetDays() int32 {
	if m != nil {
		return m.Days
	}
	return 0
}

// Unary Date Retrieval Messages
type UnaryDateRequest struct {
	Cal  *HolidayCalendar `protobuf:"bytes,1,opt,name=cal" json:"cal,omitempty"`
	Date *Date            `protobuf:"bytes,2,opt,name=date" json:"date,omitempty"`
}

func (m *UnaryDateRequest) Reset()                    { *m = UnaryDateRequest{} }
func (m *UnaryDateRequest) String() string            { return proto.CompactTextString(m) }
func (*UnaryDateRequest) ProtoMessage()               {}
func (*UnaryDateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UnaryDateRequest) GetCal() *HolidayCalendar {
	if m != nil {
		return m.Cal
	}
	return nil
}

func (m *UnaryDateRequest) GetDate() *Date {
	if m != nil {
		return m.Date
	}
	return nil
}

type UnaryBoolResponse struct {
	Ok bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
}

func (m *UnaryBoolResponse) Reset()                    { *m = UnaryBoolResponse{} }
func (m *UnaryBoolResponse) String() string            { return proto.CompactTextString(m) }
func (*UnaryBoolResponse) ProtoMessage()               {}
func (*UnaryBoolResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UnaryBoolResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type UnaryDateResponse struct {
	Date *Date `protobuf:"bytes,1,opt,name=date" json:"date,omitempty"`
}

func (m *UnaryDateResponse) Reset()                    { *m = UnaryDateResponse{} }
func (m *UnaryDateResponse) String() string            { return proto.CompactTextString(m) }
func (*UnaryDateResponse) ProtoMessage()               {}
func (*UnaryDateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UnaryDateResponse) GetDate() *Date {
	if m != nil {
		return m.Date
	}
	return nil
}

// Unary Date Transformation
// offset can be positive or negative resulting in an addition or substraction of days
type UnaryTransformRequest struct {
	Cal    *HolidayCalendar `protobuf:"bytes,1,opt,name=cal" json:"cal,omitempty"`
	Date   *Date            `protobuf:"bytes,2,opt,name=date" json:"date,omitempty"`
	Offset int32            `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
}

func (m *UnaryTransformRequest) Reset()                    { *m = UnaryTransformRequest{} }
func (m *UnaryTransformRequest) String() string            { return proto.CompactTextString(m) }
func (*UnaryTransformRequest) ProtoMessage()               {}
func (*UnaryTransformRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UnaryTransformRequest) GetCal() *HolidayCalendar {
	if m != nil {
		return m.Cal
	}
	return nil
}

func (m *UnaryTransformRequest) GetDate() *Date {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *UnaryTransformRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func init() {
	proto.RegisterType((*Date)(nil), "github.com.magaldima.bizday.api.Date")
	proto.RegisterType((*HolidayCalendar)(nil), "github.com.magaldima.bizday.api.HolidayCalendar")
	proto.RegisterType((*BinaryDateRequest)(nil), "github.com.magaldima.bizday.api.BinaryDateRequest")
	proto.RegisterType((*NumberOfDaysResponse)(nil), "github.com.magaldima.bizday.api.NumberOfDaysResponse")
	proto.RegisterType((*UnaryDateRequest)(nil), "github.com.magaldima.bizday.api.UnaryDateRequest")
	proto.RegisterType((*UnaryBoolResponse)(nil), "github.com.magaldima.bizday.api.UnaryBoolResponse")
	proto.RegisterType((*UnaryDateResponse)(nil), "github.com.magaldima.bizday.api.UnaryDateResponse")
	proto.RegisterType((*UnaryTransformRequest)(nil), "github.com.magaldima.bizday.api.UnaryTransformRequest")
	proto.RegisterEnum("github.com.magaldima.bizday.api.WeekDays", WeekDays_name, WeekDays_value)
	proto.RegisterEnum("github.com.magaldima.bizday.api.Month", Month_name, Month_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DateCalc service

type DateCalcClient interface {
	// Calculates the number of business days between two dates.
	// This operates like array operations and is inclusive of the start and exclusive of the end.
	BizDaysBetween(ctx context.Context, in *BinaryDateRequest, opts ...grpc.CallOption) (*NumberOfDaysResponse, error)
	// Calculates the number of weekdays between two dates.
	WeekdaysBetween(ctx context.Context, in *BinaryDateRequest, opts ...grpc.CallOption) (*NumberOfDaysResponse, error)
	// Calculates the number of weekend days between two dates.
	WeekendsBetween(ctx context.Context, in *BinaryDateRequest, opts ...grpc.CallOption) (*NumberOfDaysResponse, error)
	// Calculates the number of holidays between two dates.
	HolidaysBetween(ctx context.Context, in *BinaryDateRequest, opts ...grpc.CallOption) (*NumberOfDaysResponse, error)
	// Calculates the number of business days in the month of the date provided
	BizDaysInMonth(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*NumberOfDaysResponse, error)
	// Calculates the number of business days in the year of the date provided
	BizDaysInYear(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*NumberOfDaysResponse, error)
	// Calculates if the date provided is neither a weekend nor a holiday
	IsBizDay(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryBoolResponse, error)
	// Calculates if the date provided is the first business day of the month
	IsFirstBizDayOfMonth(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryBoolResponse, error)
	// Calculates if the date provided is the last business day of the month
	IsLastBizDayOfMonth(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryBoolResponse, error)
	// Calculates the first business day of the month of the date provided
	FirstBizDayOfMonth(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryDateResponse, error)
	// Calculates the last business day of the month of the date provided
	LastBizDayOfMonth(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryDateResponse, error)
	// Calculates the first business day of the quarter
	FirstBizDayOfQtr(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryDateResponse, error)
	// Calculates the last business day of the quarter
	LastBizDayOfQtr(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryDateResponse, error)
	// Calculates the next business day after the date provided
	NextBizDay(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryDateResponse, error)
	// Calculates the previous business day before the date provided
	PrevBizDay(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryDateResponse, error)
	// Add a number of business days to the date provided
	AddBizDays(ctx context.Context, in *UnaryTransformRequest, opts ...grpc.CallOption) (*UnaryDateResponse, error)
}

type dateCalcClient struct {
	cc *grpc.ClientConn
}

func NewDateCalcClient(cc *grpc.ClientConn) DateCalcClient {
	return &dateCalcClient{cc}
}

func (c *dateCalcClient) BizDaysBetween(ctx context.Context, in *BinaryDateRequest, opts ...grpc.CallOption) (*NumberOfDaysResponse, error) {
	out := new(NumberOfDaysResponse)
	err := grpc.Invoke(ctx, "/github.com.magaldima.bizday.api.DateCalc/BizDaysBetween", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dateCalcClient) WeekdaysBetween(ctx context.Context, in *BinaryDateRequest, opts ...grpc.CallOption) (*NumberOfDaysResponse, error) {
	out := new(NumberOfDaysResponse)
	err := grpc.Invoke(ctx, "/github.com.magaldima.bizday.api.DateCalc/WeekdaysBetween", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dateCalcClient) WeekendsBetween(ctx context.Context, in *BinaryDateRequest, opts ...grpc.CallOption) (*NumberOfDaysResponse, error) {
	out := new(NumberOfDaysResponse)
	err := grpc.Invoke(ctx, "/github.com.magaldima.bizday.api.DateCalc/WeekendsBetween", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dateCalcClient) HolidaysBetween(ctx context.Context, in *BinaryDateRequest, opts ...grpc.CallOption) (*NumberOfDaysResponse, error) {
	out := new(NumberOfDaysResponse)
	err := grpc.Invoke(ctx, "/github.com.magaldima.bizday.api.DateCalc/HolidaysBetween", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dateCalcClient) BizDaysInMonth(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*NumberOfDaysResponse, error) {
	out := new(NumberOfDaysResponse)
	err := grpc.Invoke(ctx, "/github.com.magaldima.bizday.api.DateCalc/BizDaysInMonth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dateCalcClient) BizDaysInYear(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*NumberOfDaysResponse, error) {
	out := new(NumberOfDaysResponse)
	err := grpc.Invoke(ctx, "/github.com.magaldima.bizday.api.DateCalc/BizDaysInYear", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dateCalcClient) IsBizDay(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryBoolResponse, error) {
	out := new(UnaryBoolResponse)
	err := grpc.Invoke(ctx, "/github.com.magaldima.bizday.api.DateCalc/IsBizDay", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dateCalcClient) IsFirstBizDayOfMonth(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryBoolResponse, error) {
	out := new(UnaryBoolResponse)
	err := grpc.Invoke(ctx, "/github.com.magaldima.bizday.api.DateCalc/IsFirstBizDayOfMonth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dateCalcClient) IsLastBizDayOfMonth(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryBoolResponse, error) {
	out := new(UnaryBoolResponse)
	err := grpc.Invoke(ctx, "/github.com.magaldima.bizday.api.DateCalc/IsLastBizDayOfMonth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dateCalcClient) FirstBizDayOfMonth(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryDateResponse, error) {
	out := new(UnaryDateResponse)
	err := grpc.Invoke(ctx, "/github.com.magaldima.bizday.api.DateCalc/FirstBizDayOfMonth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dateCalcClient) LastBizDayOfMonth(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryDateResponse, error) {
	out := new(UnaryDateResponse)
	err := grpc.Invoke(ctx, "/github.com.magaldima.bizday.api.DateCalc/LastBizDayOfMonth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dateCalcClient) FirstBizDayOfQtr(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryDateResponse, error) {
	out := new(UnaryDateResponse)
	err := grpc.Invoke(ctx, "/github.com.magaldima.bizday.api.DateCalc/FirstBizDayOfQtr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dateCalcClient) LastBizDayOfQtr(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryDateResponse, error) {
	out := new(UnaryDateResponse)
	err := grpc.Invoke(ctx, "/github.com.magaldima.bizday.api.DateCalc/LastBizDayOfQtr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dateCalcClient) NextBizDay(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryDateResponse, error) {
	out := new(UnaryDateResponse)
	err := grpc.Invoke(ctx, "/github.com.magaldima.bizday.api.DateCalc/NextBizDay", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dateCalcClient) PrevBizDay(ctx context.Context, in *UnaryDateRequest, opts ...grpc.CallOption) (*UnaryDateResponse, error) {
	out := new(UnaryDateResponse)
	err := grpc.Invoke(ctx, "/github.com.magaldima.bizday.api.DateCalc/PrevBizDay", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dateCalcClient) AddBizDays(ctx context.Context, in *UnaryTransformRequest, opts ...grpc.CallOption) (*UnaryDateResponse, error) {
	out := new(UnaryDateResponse)
	err := grpc.Invoke(ctx, "/github.com.magaldima.bizday.api.DateCalc/AddBizDays", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DateCalc service

type DateCalcServer interface {
	// Calculates the number of business days between two dates.
	// This operates like array operations and is inclusive of the start and exclusive of the end.
	BizDaysBetween(context.Context, *BinaryDateRequest) (*NumberOfDaysResponse, error)
	// Calculates the number of weekdays between two dates.
	WeekdaysBetween(context.Context, *BinaryDateRequest) (*NumberOfDaysResponse, error)
	// Calculates the number of weekend days between two dates.
	WeekendsBetween(context.Context, *BinaryDateRequest) (*NumberOfDaysResponse, error)
	// Calculates the number of holidays between two dates.
	HolidaysBetween(context.Context, *BinaryDateRequest) (*NumberOfDaysResponse, error)
	// Calculates the number of business days in the month of the date provided
	BizDaysInMonth(context.Context, *UnaryDateRequest) (*NumberOfDaysResponse, error)
	// Calculates the number of business days in the year of the date provided
	BizDaysInYear(context.Context, *UnaryDateRequest) (*NumberOfDaysResponse, error)
	// Calculates if the date provided is neither a weekend nor a holiday
	IsBizDay(context.Context, *UnaryDateRequest) (*UnaryBoolResponse, error)
	// Calculates if the date provided is the first business day of the month
	IsFirstBizDayOfMonth(context.Context, *UnaryDateRequest) (*UnaryBoolResponse, error)
	// Calculates if the date provided is the last business day of the month
	IsLastBizDayOfMonth(context.Context, *UnaryDateRequest) (*UnaryBoolResponse, error)
	// Calculates the first business day of the month of the date provided
	FirstBizDayOfMonth(context.Context, *UnaryDateRequest) (*UnaryDateResponse, error)
	// Calculates the last business day of the month of the date provided
	LastBizDayOfMonth(context.Context, *UnaryDateRequest) (*UnaryDateResponse, error)
	// Calculates the first business day of the quarter
	FirstBizDayOfQtr(context.Context, *UnaryDateRequest) (*UnaryDateResponse, error)
	// Calculates the last business day of the quarter
	LastBizDayOfQtr(context.Context, *UnaryDateRequest) (*UnaryDateResponse, error)
	// Calculates the next business day after the date provided
	NextBizDay(context.Context, *UnaryDateRequest) (*UnaryDateResponse, error)
	// Calculates the previous business day before the date provided
	PrevBizDay(context.Context, *UnaryDateRequest) (*UnaryDateResponse, error)
	// Add a number of business days to the date provided
	AddBizDays(context.Context, *UnaryTransformRequest) (*UnaryDateResponse, error)
}

func RegisterDateCalcServer(s *grpc.Server, srv DateCalcServer) {
	s.RegisterService(&_DateCalc_serviceDesc, srv)
}

func _DateCalc_BizDaysBetween_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinaryDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateCalcServer).BizDaysBetween(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.magaldima.bizday.api.DateCalc/BizDaysBetween",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateCalcServer).BizDaysBetween(ctx, req.(*BinaryDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DateCalc_WeekdaysBetween_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinaryDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateCalcServer).WeekdaysBetween(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.magaldima.bizday.api.DateCalc/WeekdaysBetween",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateCalcServer).WeekdaysBetween(ctx, req.(*BinaryDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DateCalc_WeekendsBetween_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinaryDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateCalcServer).WeekendsBetween(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.magaldima.bizday.api.DateCalc/WeekendsBetween",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateCalcServer).WeekendsBetween(ctx, req.(*BinaryDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DateCalc_HolidaysBetween_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinaryDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateCalcServer).HolidaysBetween(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.magaldima.bizday.api.DateCalc/HolidaysBetween",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateCalcServer).HolidaysBetween(ctx, req.(*BinaryDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DateCalc_BizDaysInMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateCalcServer).BizDaysInMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.magaldima.bizday.api.DateCalc/BizDaysInMonth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateCalcServer).BizDaysInMonth(ctx, req.(*UnaryDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DateCalc_BizDaysInYear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateCalcServer).BizDaysInYear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.magaldima.bizday.api.DateCalc/BizDaysInYear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateCalcServer).BizDaysInYear(ctx, req.(*UnaryDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DateCalc_IsBizDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateCalcServer).IsBizDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.magaldima.bizday.api.DateCalc/IsBizDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateCalcServer).IsBizDay(ctx, req.(*UnaryDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DateCalc_IsFirstBizDayOfMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateCalcServer).IsFirstBizDayOfMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.magaldima.bizday.api.DateCalc/IsFirstBizDayOfMonth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateCalcServer).IsFirstBizDayOfMonth(ctx, req.(*UnaryDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DateCalc_IsLastBizDayOfMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateCalcServer).IsLastBizDayOfMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.magaldima.bizday.api.DateCalc/IsLastBizDayOfMonth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateCalcServer).IsLastBizDayOfMonth(ctx, req.(*UnaryDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DateCalc_FirstBizDayOfMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateCalcServer).FirstBizDayOfMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.magaldima.bizday.api.DateCalc/FirstBizDayOfMonth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateCalcServer).FirstBizDayOfMonth(ctx, req.(*UnaryDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DateCalc_LastBizDayOfMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateCalcServer).LastBizDayOfMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.magaldima.bizday.api.DateCalc/LastBizDayOfMonth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateCalcServer).LastBizDayOfMonth(ctx, req.(*UnaryDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DateCalc_FirstBizDayOfQtr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateCalcServer).FirstBizDayOfQtr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.magaldima.bizday.api.DateCalc/FirstBizDayOfQtr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateCalcServer).FirstBizDayOfQtr(ctx, req.(*UnaryDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DateCalc_LastBizDayOfQtr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateCalcServer).LastBizDayOfQtr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.magaldima.bizday.api.DateCalc/LastBizDayOfQtr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateCalcServer).LastBizDayOfQtr(ctx, req.(*UnaryDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DateCalc_NextBizDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateCalcServer).NextBizDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.magaldima.bizday.api.DateCalc/NextBizDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateCalcServer).NextBizDay(ctx, req.(*UnaryDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DateCalc_PrevBizDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateCalcServer).PrevBizDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.magaldima.bizday.api.DateCalc/PrevBizDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateCalcServer).PrevBizDay(ctx, req.(*UnaryDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DateCalc_AddBizDays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryTransformRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateCalcServer).AddBizDays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.magaldima.bizday.api.DateCalc/AddBizDays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateCalcServer).AddBizDays(ctx, req.(*UnaryTransformRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DateCalc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.magaldima.bizday.api.DateCalc",
	HandlerType: (*DateCalcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BizDaysBetween",
			Handler:    _DateCalc_BizDaysBetween_Handler,
		},
		{
			MethodName: "WeekdaysBetween",
			Handler:    _DateCalc_WeekdaysBetween_Handler,
		},
		{
			MethodName: "WeekendsBetween",
			Handler:    _DateCalc_WeekendsBetween_Handler,
		},
		{
			MethodName: "HolidaysBetween",
			Handler:    _DateCalc_HolidaysBetween_Handler,
		},
		{
			MethodName: "BizDaysInMonth",
			Handler:    _DateCalc_BizDaysInMonth_Handler,
		},
		{
			MethodName: "BizDaysInYear",
			Handler:    _DateCalc_BizDaysInYear_Handler,
		},
		{
			MethodName: "IsBizDay",
			Handler:    _DateCalc_IsBizDay_Handler,
		},
		{
			MethodName: "IsFirstBizDayOfMonth",
			Handler:    _DateCalc_IsFirstBizDayOfMonth_Handler,
		},
		{
			MethodName: "IsLastBizDayOfMonth",
			Handler:    _DateCalc_IsLastBizDayOfMonth_Handler,
		},
		{
			MethodName: "FirstBizDayOfMonth",
			Handler:    _DateCalc_FirstBizDayOfMonth_Handler,
		},
		{
			MethodName: "LastBizDayOfMonth",
			Handler:    _DateCalc_LastBizDayOfMonth_Handler,
		},
		{
			MethodName: "FirstBizDayOfQtr",
			Handler:    _DateCalc_FirstBizDayOfQtr_Handler,
		},
		{
			MethodName: "LastBizDayOfQtr",
			Handler:    _DateCalc_LastBizDayOfQtr_Handler,
		},
		{
			MethodName: "NextBizDay",
			Handler:    _DateCalc_NextBizDay_Handler,
		},
		{
			MethodName: "PrevBizDay",
			Handler:    _DateCalc_PrevBizDay_Handler,
		},
		{
			MethodName: "AddBizDays",
			Handler:    _DateCalc_AddBizDays_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bizday.proto",
}

func init() { proto.RegisterFile("bizday.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 715 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x97, 0xdf, 0x6e, 0x12, 0x4f,
	0x14, 0xc7, 0x7f, 0xcb, 0x02, 0x5d, 0x0e, 0x6d, 0xd9, 0xce, 0xaf, 0xbf, 0x5f, 0x88, 0x37, 0x36,
	0x18, 0x4d, 0xd3, 0x0b, 0xa2, 0x18, 0x35, 0x46, 0x6f, 0xa0, 0x4d, 0x23, 0x8d, 0x50, 0xa5, 0x6d,
	0x1a, 0x2f, 0x07, 0x76, 0x28, 0x6b, 0x97, 0x19, 0x3a, 0x33, 0x5b, 0xbb, 0xb4, 0x3e, 0x84, 0x97,
	0xbe, 0x84, 0x0f, 0xe0, 0x3b, 0xf8, 0x4e, 0xe6, 0xcc, 0xae, 0x44, 0xaa, 0x09, 0x94, 0x58, 0xf0,
	0xee, 0x9c, 0xec, 0x7c, 0xcf, 0xe7, 0xfc, 0x99, 0xb3, 0x0b, 0xb0, 0xdc, 0xf6, 0x87, 0x1e, 0x8d,
	0xca, 0x03, 0x29, 0xb4, 0x20, 0x77, 0x4f, 0x7c, 0xdd, 0x0b, 0xdb, 0xe5, 0x8e, 0xe8, 0x97, 0xfb,
	0xf4, 0x84, 0x06, 0x9e, 0xdf, 0xa7, 0xe5, 0xe4, 0x08, 0x1d, 0xf8, 0xa5, 0xf7, 0x90, 0xde, 0xa1,
	0x9a, 0x11, 0x02, 0xe9, 0x88, 0x51, 0x59, 0xb4, 0x36, 0xac, 0xcd, 0x4c, 0xcb, 0xd8, 0xe4, 0x25,
	0x64, 0xfa, 0x82, 0xeb, 0x5e, 0x31, 0xb5, 0x61, 0x6d, 0xae, 0x56, 0x1e, 0x94, 0x27, 0x04, 0x2b,
	0x37, 0xf0, 0x74, 0x2b, 0x16, 0x11, 0x17, 0x6c, 0x8f, 0x46, 0x45, 0xdb, 0x04, 0x44, 0xb3, 0xd4,
	0x83, 0xc2, 0x2b, 0x11, 0xf8, 0x1e, 0x8d, 0xb6, 0x69, 0xc0, 0xb8, 0x47, 0x25, 0x62, 0x39, 0xed,
	0x33, 0x83, 0xcd, 0xb5, 0x8c, 0x4d, 0xaa, 0xe0, 0xf4, 0xe2, 0x63, 0xaa, 0x98, 0xda, 0xb0, 0x37,
	0xf3, 0x95, 0xfb, 0x13, 0xc9, 0x58, 0x43, 0x6b, 0x24, 0x2b, 0x7d, 0xb3, 0x60, 0xad, 0xe6, 0x73,
	0x2a, 0x23, 0xf3, 0x80, 0x9d, 0x85, 0x4c, 0x69, 0x52, 0x03, 0xbb, 0x43, 0x03, 0xc3, 0xca, 0x57,
	0x1e, 0x4e, 0x8c, 0x79, 0x2d, 0xd7, 0x16, 0x8a, 0xc9, 0x0b, 0xc8, 0x28, 0x4d, 0xa5, 0x36, 0x3d,
	0x99, 0x3a, 0xb3, 0x58, 0x43, 0x9e, 0x81, 0xcd, 0xb8, 0x67, 0x5a, 0x32, 0xb5, 0x14, 0x15, 0xa5,
	0x2d, 0x58, 0x6f, 0x86, 0xfd, 0x36, 0x93, 0xfb, 0xdd, 0x1d, 0x1a, 0xa9, 0x16, 0x53, 0x03, 0xc1,
	0x95, 0x99, 0x9a, 0x69, 0x53, 0x32, 0x35, 0x53, 0xfb, 0x27, 0x0b, 0xdc, 0xa3, 0xdb, 0x28, 0xfd,
	0x39, 0xc2, 0x34, 0xbb, 0x59, 0xe5, 0x46, 0x52, 0xba, 0x07, 0x6b, 0x26, 0xa5, 0x9a, 0x10, 0xc1,
	0x28, 0xf9, 0x55, 0x48, 0x89, 0x53, 0x93, 0x92, 0xd3, 0x4a, 0x89, 0xd3, 0x52, 0x33, 0x39, 0x14,
	0xe7, 0x9d, 0x1c, 0xfa, 0x01, 0xb5, 0x6e, 0x0e, 0xfd, 0x62, 0xc1, 0x7f, 0x26, 0xe0, 0xa1, 0xa4,
	0x5c, 0x75, 0x85, 0xec, 0xff, 0x1d, 0xdd, 0x20, 0xff, 0x43, 0x56, 0x74, 0xbb, 0x8a, 0xe9, 0x64,
	0x39, 0x12, 0x6f, 0xab, 0x0b, 0xce, 0x31, 0x63, 0xa7, 0x38, 0x61, 0x02, 0x90, 0x3d, 0x08, 0xb9,
	0x47, 0x23, 0xf7, 0x1f, 0xb4, 0x1b, 0xc2, 0xd8, 0x16, 0xc9, 0xc3, 0xd2, 0x61, 0xc8, 0x14, 0x3a,
	0x29, 0xb2, 0x02, 0xb9, 0x63, 0xe6, 0xf1, 0xd8, 0xb5, 0xc9, 0x32, 0x38, 0x87, 0xbd, 0x50, 0x1a,
	0x2f, 0x8d, 0xaa, 0x5d, 0x89, 0x79, 0xbb, 0x19, 0x7c, 0x72, 0x40, 0x75, 0x28, 0xd1, 0xcb, 0x6e,
	0x7d, 0xb6, 0x20, 0x63, 0x56, 0x15, 0xa3, 0xed, 0x51, 0x1e, 0x52, 0x89, 0x98, 0x65, 0x70, 0x76,
	0x59, 0x5b, 0x1a, 0xcf, 0x22, 0x39, 0xc8, 0x34, 0xa8, 0xec, 0xf4, 0xdc, 0x14, 0x9a, 0xd5, 0x81,
	0xf4, 0x03, 0xd7, 0x26, 0x4b, 0x60, 0x37, 0x4c, 0x74, 0x07, 0xd2, 0x7b, 0x21, 0x67, 0x6e, 0x26,
	0xb6, 0x82, 0xc8, 0xcd, 0x22, 0xb1, 0x1a, 0x9e, 0x84, 0x4a, 0xbb, 0x4b, 0x98, 0xda, 0x01, 0x1b,
	0x68, 0x86, 0x97, 0xd6, 0x75, 0x10, 0xb4, 0xdf, 0xd1, 0x02, 0x9d, 0x1c, 0x82, 0x9a, 0xe2, 0x3c,
	0x7e, 0x04, 0xe8, 0xed, 0xb0, 0x4e, 0xec, 0xe5, 0x2b, 0x5f, 0x0b, 0xe0, 0x60, 0xab, 0xb6, 0x69,
	0xd0, 0x21, 0x97, 0xb0, 0x5a, 0xf3, 0x87, 0xd8, 0x8f, 0x1a, 0xd3, 0x1f, 0x18, 0xe3, 0xa4, 0x32,
	0xb1, 0xcf, 0xbf, 0xac, 0xfd, 0x9d, 0x27, 0x13, 0x35, 0xbf, 0xdd, 0xad, 0x2b, 0x28, 0xe0, 0x34,
	0xbc, 0x85, 0xd2, 0x19, 0xf7, 0x16, 0x44, 0x4f, 0x2e, 0xfd, 0x22, 0xe8, 0xc3, 0xd1, 0xd8, 0xeb,
	0x3c, 0xbe, 0xa7, 0x8f, 0x26, 0x06, 0x3a, 0xfa, 0x43, 0xec, 0x08, 0x56, 0x46, 0xec, 0x77, 0xf8,
	0x11, 0x9c, 0x1f, 0xfa, 0x0c, 0x9c, 0xba, 0x8a, 0xe1, 0xb3, 0x50, 0x2b, 0xd3, 0x49, 0xc6, 0x5e,
	0xc1, 0x1f, 0x61, 0xbd, 0xae, 0x76, 0x7d, 0xa9, 0x74, 0xcc, 0xdd, 0xef, 0xce, 0xdc, 0xef, 0x59,
	0xf0, 0x57, 0xf0, 0x6f, 0x5d, 0xbd, 0xa6, 0x0b, 0xa2, 0x5f, 0x02, 0x99, 0x6b, 0xe9, 0x63, 0xdf,
	0xb5, 0x21, 0xac, 0xcd, 0xb3, 0xf0, 0x31, 0x76, 0x04, 0xee, 0x58, 0xe1, 0x6f, 0xb5, 0x9c, 0x17,
	0xfa, 0x02, 0x0a, 0x3f, 0x97, 0x3d, 0x47, 0xb2, 0x02, 0x68, 0xb2, 0x0b, 0x7d, 0xeb, 0xfb, 0x75,
	0x1d, 0xfa, 0x46, 0xb2, 0xf3, 0xf9, 0x42, 0x2f, 0x00, 0xaa, 0x9e, 0x97, 0xbc, 0xc5, 0xc8, 0xd3,
	0xe9, 0x22, 0x5c, 0xff, 0x8d, 0x34, 0x0b, 0xb9, 0x9d, 0x35, 0x7f, 0x3a, 0x1e, 0x7f, 0x0f, 0x00,
	0x00, 0xff, 0xff, 0xf0, 0x4e, 0x28, 0xab, 0x84, 0x0c, 0x00, 0x00,
}