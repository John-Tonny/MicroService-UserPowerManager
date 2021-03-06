package handler

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/idoall/TokenExchangeCommon/commonutils"

	"github.com/idoall/MicroService-UserPowerManager/utils"

	"github.com/idoall/MicroService-UserPowerManager/api/columns/v1/models"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"

	srvProto "github.com/idoall/MicroService-UserPowerManager/srv/columns/v1/proto"

	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro/errors"
)

// Columns struct
type Columns struct {
	Client srvProto.ProtoColumnsService
}

// Add route POST /mshk/api/v1/ApiColumns/add
// 添加一个栏目
func (e *Columns) Add(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_Columns_Add_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APICOLUMNS

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [ApiColumns][Add] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var name, URL, cssIcon string
	var ParentID int64
	var sorts, isShowNav int32

	if req.Post["Name"] == nil || req.Post["Name"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "Name 不能为空")
	}
	name = req.Post["Name"].Values[0]

	if req.Post["URL"] != nil && req.Post["URL"].Values[0] != "" {
		URL = req.Post["URL"].Values[0]
	}
	if req.Post["CssIcon"] != nil && req.Post["CssIcon"].Values[0] != "" {
		cssIcon = req.Post["CssIcon"].Values[0]
	}

	if req.Post["ParentID"] != nil && req.Post["ParentID"].Values[0] != "" {
		if ParentID, err = commonutils.Int64FromString(req.Post["ParentID"].Values[0]); err != nil {
			return errors.InternalServerError(namespaceID, "ParentID 的格式不正确:%s", err.Error())
		}
	}

	if req.Post["Sorts"] != nil && req.Post["Sorts"].Values[0] != "" {
		if sorts, err = commonutils.Int32FromString(req.Post["Sorts"].Values[0]); err != nil {
			return errors.InternalServerError(namespaceID, "Sorts 的格式不正确:%s", err.Error())
		}
	}
	if req.Post["IsShowNav"] != nil && req.Post["IsShowNav"].Values[0] != "" && (req.Post["IsShowNav"].Values[0] == "1" || req.Post["IsShowNav"].Values[0] == "true") {
		isShowNav = 1
	}
	// 获取请求参数 - 结束

	// make request
	addModel := &srvProto.Columns{
		Name:      name,
		URL:       URL,
		ParentID:  ParentID,
		Sorts:     sorts,
		IsShowNav: isShowNav,
		CssIcon:   cssIcon,
	}

	// 调用服务端方法
	srvResponse, err := e.Client.Add(ctx, &srvProto.AddRequest{Model: addModel})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 输出的 json
	respnseJSON := struct {
		NewID int64 `json:"newid"`
	}{}
	respnseJSON.NewID = srvResponse.NewID
	b, _ := commonutils.JSONEncode(respnseJSON)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info(srvResponse)
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Api_Columns_Add_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("NewID", srvResponse.NewID)
	}

	return nil
}

// GetList 获取用户列表,默认 id 倒排序
func (e *Columns) GetList(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_Columns_GetList_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APICOLUMNS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [ApiColumns][GetList] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var pageSize, currentPageIndex int64
	var orderBy string
	if req.Get["PageSize"] == nil || req.Get["PageSize"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "PageSize 不能为空")
	} else if pageSize, err = commonutils.Int64FromString(req.Get["PageSize"].Values[0]); err != nil {
		return errors.InternalServerError(namespaceID, "PageSize Format Error:%s", err.Error())
	}

	if req.Get["CurrentPageIndex"] == nil || req.Get["CurrentPageIndex"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "CurrentPageIndex 不能为空")
	} else if currentPageIndex, err = commonutils.Int64FromString(req.Get["CurrentPageIndex"].Values[0]); err != nil {
		return errors.InternalServerError(namespaceID, "CurrentPageIndex Format Error:%s", err.Error())

	}

	if req.Get["OrderBy"] != nil {
		orderBy = req.Get["OrderBy"].Values[0]
	}
	// 获取请求参数 - 结束

	// 调用服务端方法
	srvResponse, err := e.Client.GetList(ctx, &srvProto.GetListRequest{
		CurrentPageIndex: currentPageIndex,
		PageSize:         pageSize,
		OrderBy:          orderBy,
	})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// return json

	responseJSON := struct {
		Rows  []*models.ResponseJSONRow `json:"rows"`
		Total int64                     `json:"total"`
	}{}
	for _, v := range srvResponse.List {
		r := &models.ResponseJSONRow{
			ID:             v.ID,
			Name:           v.Name,
			URL:            v.URL,
			ParentID:       v.ParentID,
			Sorts:          v.Sorts,
			IsShowNav:      v.IsShowNav,
			CssIcon:        v.CssIcon,
			CreateTime:     v.CreateTime,
			LastUpdateTime: v.LastUpdateTime,
		}
		responseJSON.Rows = append(responseJSON.Rows, r)
	}

	responseJSON.Total = srvResponse.TotalCount

	// 对 json 序列化并输出
	b, _ := json.Marshal(responseJSON)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Api_Columns_GetList_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("PageSize", pageSize)
		span.SetTag("CurrentPageIndex", currentPageIndex)
		span.SetTag("orderBy", orderBy)
		span.SetTag("TotalCount", srvResponse.TotalCount)
	}

	return nil
}

// Get 获取单个条记录，根据ID
func (e *Columns) Get(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_Columns_Get_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APICOLUMNS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [ApiColumns][Get] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var ID int64
	if req.Get["ID"] != nil && req.Get["ID"].Values[0] != "0" {
		if ID, err = commonutils.Int64FromString(req.Get["ID"].Values[0]); err != nil {
			return errors.InternalServerError(namespaceID, "ID Format Error:%s", err.Error())
		}
	}
	// 获取请求参数 - 结束

	// 调用服务端方法
	srvResponse, err := e.Client.Get(ctx, &srvProto.GetRequest{
		ID: ID,
	})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 对 json 序列化并输出
	b, _ := json.Marshal(&models.ResponseJSONRow{
		ID:             srvResponse.Model.ID,
		Name:           srvResponse.Model.Name,
		URL:            srvResponse.Model.URL,
		ParentID:       srvResponse.Model.ParentID,
		Sorts:          srvResponse.Model.Sorts,
		IsShowNav:      srvResponse.Model.IsShowNav,
		CssIcon:        srvResponse.Model.CssIcon,
		CreateTime:     srvResponse.Model.CreateTime,
		LastUpdateTime: srvResponse.Model.LastUpdateTime,
	})
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Api_Columns_Get_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("ID", ID)
	}

	return nil
}

// Update 修改用户信息
func (e *Columns) Update(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_Columns_Update_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APICOLUMNS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Update] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var ID int64
	var name, URL, cssIcon string
	var parentID int64
	var sorts, isShowNav int32

	if req.Post["ID"] == nil || req.Post["ID"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "ID 不能为空")
	} else if ID, err = commonutils.Int64FromString(req.Post["ID"].Values[0]); err != nil {
		return errors.InternalServerError(namespaceID, "ID Format Error:%s", err.Error())
	}
	if req.Post["Name"] == nil || req.Post["Name"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "Name 不能为空")
	}
	name = req.Post["Name"].Values[0]

	if req.Post["URL"] != nil && req.Post["URL"].Values[0] != "" {
		URL = req.Post["URL"].Values[0]
	}
	if req.Post["CssIcon"] != nil && req.Post["CssIcon"].Values[0] != "" {
		cssIcon = req.Post["CssIcon"].Values[0]
	}

	if req.Post["ParentID"] != nil && req.Post["ParentID"].Values[0] != "" {
		if parentID, err = commonutils.Int64FromString(req.Post["ParentID"].Values[0]); err != nil {
			return errors.InternalServerError(namespaceID, "ParentID 的格式不正确:%s", err.Error())
		}
	}

	if req.Post["Sorts"] != nil && req.Post["Sorts"].Values[0] != "" {
		if sorts, err = commonutils.Int32FromString(req.Post["Sorts"].Values[0]); err != nil {
			return errors.InternalServerError(namespaceID, "Sorts 的格式不正确:%s", err.Error())
		}
	}
	if req.Post["IsShowNav"] != nil && req.Post["IsShowNav"].Values[0] != "" && (req.Post["IsShowNav"].Values[0] == "1" || req.Post["IsShowNav"].Values[0] == "true") {
		isShowNav = 1
	}
	// 获取请求参数 - 结束

	// 调用服务端方法获取栏目
	srvResponseGet, err := e.Client.Get(ctx, &srvProto.GetRequest{
		ID: ID,
	})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 调用服务端方法 - 修改栏目
	responseUpdateModel := &srvProto.Columns{
		ID:        srvResponseGet.Model.ID,
		Name:      name,
		URL:       URL,
		ParentID:  parentID,
		Sorts:     sorts,
		IsShowNav: isShowNav,
		CssIcon:   cssIcon,
	}
	response, err := e.Client.Update(ctx, &srvProto.UpdateRequest{Model: responseUpdateModel})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 输出的 json
	responseJSON := struct {
		Updated int64 `json:"updated"`
	}{}
	responseJSON.Updated = response.Updated
	b, _ := commonutils.JSONEncode(responseJSON)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Api_Columns_Update_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("ID", ID)
	}

	return nil
}

// BatchDelete 批量删除用户信息
func (e *Columns) BatchDelete(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_Columns_BatchDelete_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APICOLUMNS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [ApiColumns][BatchDelete] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var IDArray []string
	if req.Post["IDArray"] == nil || req.Post["IDArray"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "IDArray 不能为空")
	}
	IDArray = strings.Split(req.Post["IDArray"].Values[0], ",")

	// 获取请求参数 - 结束

	var IDArray64 []int64
	for _, v := range IDArray {
		i64, _ := commonutils.Int64FromString(v)
		IDArray64 = append(IDArray64, i64)
	}

	// 调用服务端方法获取用户
	response, err := e.Client.BatchDelete(ctx, &srvProto.DeleteRequest{
		IDArray: IDArray64,
	})

	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 输出的 json
	responseJSON := struct {
		Deleted int64 `json:"deleted"`
	}{}
	responseJSON.Deleted = response.Deleted
	b, _ := commonutils.JSONEncode(responseJSON)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Api_Columns_BatchDelete_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("IDArray", strings.Join(IDArray, ","))
	}

	return nil
}
