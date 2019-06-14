package columns

import (
	"bytes"
	"fmt"
	"net/url"
	"strconv"

	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/request"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin"
	"github.com/idoall/MicroService-UserPowerManager/web/models"
)

// ColumnsController Controller
type ColumnsController struct {
	admin.AdminBaseController
}

// ColumnRow 解析栏目的struct
type ColumnRow struct {
	ID             int64  `json:"ID"`
	Name           string `json:"Name"`
	URL            string `json:"URL"`
	ParentID       int64  `json:"ParentID,omitempty"`
	Sorts          int64  `json:"Sorts,omitempty"`
	IsShowNav      bool   `json:"IsShowNav,omitempty"`
	CssIcon        string `json:"CssIcon,omitempty"`
	CreateTime     int64  `json:"CreateTime"`
	LastUpdateTime int64  `json:"LastUpdateTime"`
	Nodes          []*ColumnRow
}

var (
	TemplageBaseURL = "columns"
	baseTitle       = "栏目"
	pageSizeDefault = 11
)

// GetTreeViewJSON Default Json
func (e *ColumnsController) GetTreeViewJSON() {
	var err error
	var list []*models.TreeView
	var result models.Result

	// 获取首页 TreeView 需要用到的结构化数据
	if list, err = e.getTreeViewBootstrap(); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	e.Data["json"] = list
	e.ServeJSON()

}

// Get 首页
func (e *ColumnsController) Get() {

	// set Data
	versionAdminURL := e.GetVersionAdminBaseURL()
	e.Data["title"] = fmt.Sprintf("%s管理", baseTitle)
	e.Data["AddUrl"] = fmt.Sprintf("%s/%s/add", versionAdminURL, TemplageBaseURL)
	e.Data["UpdateUrl"] = fmt.Sprintf("%s/%s/update?id=", versionAdminURL, TemplageBaseURL)
	e.Data["BatchDelUrl"] = fmt.Sprintf("%s/%s/batchdelete", versionAdminURL, TemplageBaseURL)
	e.Data["JSONTreeViewListUrl"] = fmt.Sprintf("%s/%s/GetTreeViewJSON", versionAdminURL, TemplageBaseURL)

	e.SetMortStype()
	e.SetMortScript()
	e.AppendCustomScripts([]string{
		// Bootstrap table
		"/static/js/hplus/plugins/bootstrap-table/bootstrap-table.min.js",
		"/static/js/hplus/plugins/bootstrap-table/locale/bootstrap-table-zh-CN.min.js",
		// TreeView
		"/static/js/hplus/plugins/treeview/bootstrap-treeview.min.js",
	})
	e.AppendCustomStyles([]string{
		// Bootstrap table
		"/static/css/hplus/bootstrap-table/bootstrap-table.min.css",
		// TreeView
		"/static/css/hplus/treeview/bootstrap-treeview.min.css",
	})

	e.Layout = "admin/layout/layout.html"
	e.LayoutSections = make(map[string]string)
	e.LayoutSections["CustomHeader"] = "admin/layout/layout-customsheader.html"
	e.TplName = fmt.Sprintf("%s/%s/index.html", admin.TemplageAdminBaseURL, TemplageBaseURL)
}

// Add 添加
func (e *ColumnsController) Add() {
	var err error

	// set data
	versionAdminURL := e.GetVersionAdminBaseURL()
	e.Data["title"] = fmt.Sprintf("添加%s", baseTitle)
	if e.Data["ParentList"], err = e.getTreeHtmlSelect(); err != nil {
		e.Data["json"] = err.Error()
		e.ServeJSON()
		return
	}
	e.Data["AddSaveUrl"] = fmt.Sprintf("%s/%s/addsave", versionAdminURL, TemplageBaseURL)

	// 公用设置，样式、脚本、layout
	e.SetMortStype()
	e.SetMortScript()
	e.AppendCustomScripts(nil)
	e.AppendCustomStyles(nil)
	e.Layout = "admin/layout/layout.html"
	e.LayoutSections = make(map[string]string)
	e.LayoutSections["CustomHeader"] = "admin/layout/layout-customsheader.html"
	e.TplName = fmt.Sprintf("%s/%s/add.html", admin.TemplageAdminBaseURL, TemplageBaseURL)
}

// AddSave 保存添加的交易配置
func (e *ColumnsController) AddSave() {

	// 用于 json 返回的数据
	var result models.Result
	var err error

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("Name", e.GetString("name"))
	params.Set("URL", e.GetString("URL"))
	params.Set("CssIcon", e.GetString("cssicon"))
	params.Set("ParentID", e.GetString("parentid"))
	params.Set("Sorts", e.GetString("sorts"))
	params.Set("IsShowNav", e.GetString("isshownav"))

	// 发送请求的路径
	path := fmt.Sprintf("%s%s", inner.MicroServiceHostProt, utils.TConfig.String("MicroServices::ServiceURL_Column_Add"))

	// 临时 Json解析类
	responseJSON := struct {
		NewID int64 `json:"newid"`
	}{}
	// 发送 http 请求
	if err = request.Request.SendPayload("POST", path, nil, bytes.NewBufferString(params.Encode()), &responseJSON, false, true, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	} else {
		e.Ctx.Redirect(302, fmt.Sprintf("%s/%s", e.GetVersionAdminBaseURL(), TemplageBaseURL))
	}
}

// Update 修改
func (e *ColumnsController) Update() {
	var result models.Result
	var err error
	var ID int64
	var HTMLSelect []*models.HtmlSelect

	// 获取要修改的 ID
	if ID, err = e.GetInt64("id", 0); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("ID", strconv.FormatInt(ID, 10))

	// 发送请求的路径
	path := fmt.Sprintf("%s%s?%s",
		inner.MicroServiceHostProt,
		utils.TConfig.String("MicroServices::ServiceURL_Column_Get"),
		params.Encode(),
	)

	// 临时 Json解析类
	var responseJson map[string]interface{}
	// 发送 http 请求
	if err = request.Request.SendPayload("GET", path, nil, nil, &responseJson, false, true, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// 获取所有的HTMLSelect部分
	if HTMLSelect, err = e.getTreeHtmlSelect(); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}
	fmt.Println(responseJson)
	// 设置所属的上级选中
	parentID := int64(responseJson["ParentID"].(float64))
	// 设置 HTMLSelect 选中
	for _, v := range HTMLSelect {
		if v.ID == parentID {
			v.IsSelected = true
			break
		}
	}
	fmt.Println(responseJson)
	//set Data
	versionAdminURL := e.GetVersionAdminBaseURL()

	e.Data["Model"] = responseJson
	e.Data["ParentList"] = HTMLSelect
	e.Data["title"] = fmt.Sprintf("修改%s", baseTitle)
	e.Data["UpdateSaveUrl"] = fmt.Sprintf("%s/%s/updatesave", versionAdminURL, TemplageBaseURL)

	//公用设置，样式、脚本、layout
	e.SetMortStype()
	e.SetMortScript()
	e.AppendCustomScripts([]string{"/static/js/admin/symbol_add.js"})
	e.AppendCustomStyles(nil)
	e.Layout = "admin/layout/layout.html"
	e.LayoutSections = make(map[string]string)
	e.LayoutSections["CustomHeader"] = "admin/layout/layout-customsheader.html"
	e.TplName = fmt.Sprintf("%s/%s/update.html", admin.TemplageAdminBaseURL, TemplageBaseURL)
}

// UpdateSave 保存修改
func (e *ColumnsController) UpdateSave() {
	// 用于 json 返回的数据
	var result models.Result
	var err error

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("ID", e.GetString("id"))
	params.Set("Name", e.GetString("name"))
	params.Set("URL", e.GetString("URL"))
	params.Set("CssIcon", e.GetString("cssicon"))
	params.Set("ParentID", e.GetString("parentid"))
	params.Set("Sorts", e.GetString("sorts"))
	params.Set("IsShowNav", e.GetString("isshownav"))

	// 发送请求的路径
	path := fmt.Sprintf("%s%s", inner.MicroServiceHostProt, utils.TConfig.String("MicroServices::ServiceURL_Column_Update"))

	// 临时 Json解析类
	responseJSON := struct {
		Updated int64 `json:"updated"`
	}{}
	// 发送 http 请求
	if err = request.Request.SendPayload("POST", path, nil, bytes.NewBufferString(params.Encode()), &responseJSON, false, true, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	} else {
		e.Ctx.Redirect(302, fmt.Sprintf("%s/%s", e.GetVersionAdminBaseURL(), TemplageBaseURL))
	}
}

// BatchDelete 批量删除
func (e *ColumnsController) BatchDelete() {

	// 用于 json 返回的数据
	var result models.Result
	var err error

	userIds := e.GetString("ids")
	if userIds == "" {
		result.Code = -1
		result.Msg = "ids 不能为空"
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("IDArray", userIds)
	fmt.Println(params.Encode())

	// 发送请求的路径
	path := fmt.Sprintf("%s%s", inner.MicroServiceHostProt, utils.TConfig.String("MicroServices::ServiceURL_Column_BatchDelete"))

	// 临时 Json解析类
	responseJson := struct {
		Deleted int64
	}{}
	// 发送 http 请求
	if err = request.Request.SendPayload("POST", path, nil, bytes.NewBufferString(params.Encode()), &responseJson, false, true, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
	} else {
		result.Code = 0
		e.Data["json"] = result
	}

	e.ServeJSON()
}

// 从 go-micro 的 apicolumns 获取所有列的数据
func (e *ColumnsController) getApiServiceColumnRow() ([]*ColumnRow, error) {
	var err error

	// return json
	jsonList := struct {
		Rows       []*ColumnRow `json:"rows"`
		Total      int64        `json:"total"`
		ErrMessage string       `json:"errmsg"`
	}{}
	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("PageSize", fmt.Sprintf("%d", 1000))
	params.Set("CurrentPageIndex", fmt.Sprintf("%d", 1))

	// 发送请求的路径
	path := fmt.Sprintf("%s%s?%s",
		inner.MicroServiceHostProt,
		utils.TConfig.String("MicroServices::ServiceURL_Column_GetList"),
		params.Encode(),
	)

	// 发送 http 请求
	if err = request.Request.SendPayload("GET", path, nil, nil, &jsonList, false, false, false); err != nil {
		return nil, err
	} else {
		return jsonList.Rows, nil
	}
}

//-----------------------HTML Select

// GetTreeHtmlSelect getTreeHtmlSelect
func (e *ColumnsController) getTreeHtmlSelect() ([]*models.HtmlSelect, error) {
	list, err := e.getTreeStruct()
	if err != nil {
		return nil, err
	}

	var columnPowerHTMLSelect []*models.HtmlSelect

	for _, v := range list {

		columnPowerHTMLSelect = append(columnPowerHTMLSelect, &models.HtmlSelect{
			ID:    int64(v.ID),
			Value: v.Name,
		})

		if v.Nodes != nil {
			columnPowerHTMLSelect = e.generateRecursiveHTMLSelect(columnPowerHTMLSelect, v.Nodes)
		}
	}

	return columnPowerHTMLSelect, nil
}

// 递归生成HTML Select
func (e *ColumnsController) generateRecursiveHTMLSelect(columnPowerHTMLSelect []*models.HtmlSelect, list []*ColumnRow) []*models.HtmlSelect {
	for _, v := range list {
		columnPowerHTMLSelect = append(columnPowerHTMLSelect, &models.HtmlSelect{
			ID:    int64(v.ID),
			Value: v.Name,
		})

		if v.Nodes != nil {
			columnPowerHTMLSelect = e.generateRecursiveHTMLSelect(columnPowerHTMLSelect, v.Nodes)
		}
	}

	return columnPowerHTMLSelect
}

//-------------Tree Struct

// getTreeStruct TreeList 获取原始组装好的 Struct 节点
func (e *ColumnsController) getTreeStruct() ([]*ColumnRow, error) {

	var err error
	var apiColumns []*ColumnRow
	// 发送 http 请求
	if apiColumns, err = e.getApiServiceColumnRow(); err != nil {
		return nil, err
	} else {
		// 递归计算第一层
		var returnlist []*ColumnRow
		for _, v := range apiColumns {
			if v.ParentID == 0 {
				v.Name = "✚" + v.Name
				e.getTreeStructRecursive(v, apiColumns, 1)
				returnlist = append(returnlist, v)
			}
		}
		return returnlist, nil
	}
}

// getTreeStructRecursive 递归获取下一级
func (e *ColumnsController) getTreeStructRecursive(m *ColumnRow, list []*ColumnRow, depth int) {

	title := ""
	for i := 0; i < depth; i++ {
		title += "┊"
	}

	for _, v := range list {
		if v.ParentID == m.ID {
			v.Name = title + "├" + v.Name
			m.Nodes = append(m.Nodes, v)
			e.getTreeStructRecursive(v, list, depth+1)
		}
	}
}

//--------------Tree View

// GetTreeViewBootstrap GetTreeViewBootstrap
func (e *ColumnsController) getTreeViewBootstrap() ([]*models.TreeView, error) {

	var err error
	var apiColumns []*ColumnRow
	// 发送 http 请求
	if apiColumns, err = e.getApiServiceColumnRow(); err != nil {
		return nil, err
	} else {
		// 递归计算第一层
		var returnlist []*models.TreeView
		for _, v := range apiColumns {
			if v.ParentID == 0 {
				t := &models.TreeView{
					ID:   v.ID,
					Text: v.Name,
					Href: v.URL,
					Tags: []string{fmt.Sprintf("%d", v.ID), v.URL, fmt.Sprintf("Sort:%d", v.Sorts), fmt.Sprintf("IsShowNav:%t", v.IsShowNav)},
				}
				e.generateRecursiveTreeViewBootstrap(v, t, apiColumns)
				returnlist = append(returnlist, t)
			}
		}
		return returnlist, nil
	}

}

func (e *ColumnsController) generateRecursiveTreeViewBootstrap(m *ColumnRow, t *models.TreeView, list []*ColumnRow) {

	for _, v := range list {
		if v.ParentID == m.ID {
			tr := &models.TreeView{
				ID:   v.ID,
				Text: v.Name,
				Href: v.URL,
				Tags: []string{fmt.Sprintf("%d", v.ID), v.URL, fmt.Sprintf("Sort:%d", v.Sorts), fmt.Sprintf("IsShowNav:%t", v.IsShowNav)},
			}
			t.Nodes = append(t.Nodes, tr)
			e.generateRecursiveTreeViewBootstrap(v, tr, list)
		}
	}
}
