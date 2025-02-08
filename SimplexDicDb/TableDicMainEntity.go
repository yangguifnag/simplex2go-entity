package SimplexDicDbEntity



import (
    time "time"
    )



/*

    - 64cce7c8-b239-490f-9563-2c9c9dfef611 -

    模板文件 自动生成
    实体映射
    库名: simplex_dic_db
    表名: dic_main
    解释: Simplex2 CodeCreator
    作者: Simplex2
    版本: 1.0
    创建时间: 2025-02-08 01:56:40
    是否虚拟表: 否
    是否存在审计字段: 否
    是否存在逻辑状态字段: 是

    Template File Auto-Generation
    Entity Mapping

    Database Name: simplex_dic_db
    Table Name: dic_main
    Description: Simplex2 CodeCreator
    Author: Simplex2
    Version: 1.0
    Creation Time: 2025-02-08 01:56:40
    Is Virtual Table: No
    Has Audit Fields: No
    Has Logical Status Field: Yes

*/

type DicMainEntity struct {




    /*  字段名: code
        类型: VARCHAR
        字段解释: 字典Code
        字段长度: 100
        是否主键: 否
        是否为空: 否
        归宿主键: dic_main_pk_2

    */

    Code string `mapstructure:",omitempty" gorm:"column:code" json:"code"`




    /*  字段名: createBy
        类型: VARCHAR
        字段解释: 创建人
        字段长度: 50
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    CreateBy string `mapstructure:",omitempty" gorm:"column:createBy" json:"createBy"`




    /*  字段名: createDate
        类型: DATETIME
        字段解释: 创建日期
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    CreateDate time.Time `mapstructure:",omitempty" gorm:"column:createDate;default:null" json:"createDate"`




    /*  字段名: desc
        类型: VARCHAR
        字段解释: 备注
        字段长度: 200
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    Desc string `mapstructure:",omitempty" gorm:"column:desc" json:"desc"`




    /*  字段名: dicName
        类型: VARCHAR
        字段解释: 字典名
        字段长度: 100
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    DicName string `mapstructure:",omitempty" gorm:"column:dicName" json:"dicName"`




    /*  字段名: dicTable
        类型: VARCHAR
        字段解释: 字典表名
        字段长度: 100
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    DicTable string `mapstructure:",omitempty" gorm:"column:dicTable" json:"dicTable"`




    /*  字段名: dicType
        类型: VARCHAR
        字段解释: 字典类型
        字段长度: 100
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    DicType string `mapstructure:",omitempty" gorm:"column:dicType" json:"dicType"`




    /*  字段名: id
        类型: INT
        字段解释: --
        字段长度: -
        是否主键: 是
        是否为空: 否
        归宿主键: PRIMARY

    */

    Id int `mapstructure:",omitempty" gorm:"primaryKey;column:id" json:"id"`




    /*  字段名: labelKey
        类型: VARCHAR
        字段解释: 字典键名
        字段长度: 50
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    LabelKey string `mapstructure:",omitempty" gorm:"column:labelKey" json:"labelKey"`




    /*  字段名: parentKey
        类型: VARCHAR
        字段解释: 父值键名
        字段长度: 50
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    ParentKey string `mapstructure:",omitempty" gorm:"column:parentKey" json:"parentKey"`




    /*  字段名: status
        类型: INT
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: dic_main_pk_2

    */

    Status int `mapstructure:",omitempty" gorm:"column:status" json:"status"`




    /*  字段名: updateBy
        类型: VARCHAR
        字段解释: 更新人
        字段长度: 50
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    UpdateBy string `mapstructure:",omitempty" gorm:"column:updateBy" json:"updateBy"`




    /*  字段名: updateDate
        类型: DATETIME
        字段解释: 更新时间
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    UpdateDate time.Time `mapstructure:",omitempty" gorm:"column:updateDate;default:null" json:"updateDate"`




    /*  字段名: valueKey
        类型: VARCHAR
        字段解释: 字典值键
        字段长度: 50
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    ValueKey string `mapstructure:",omitempty" gorm:"column:valueKey" json:"valueKey"`




    /*  字段名: version
        类型: VARCHAR
        字段解释: --
        字段长度: 20
        是否主键: 否
        是否为空: 否
        归宿主键: dic_main_pk_2

    */

    Version string `mapstructure:",omitempty" gorm:"column:version" json:"version"`
}



    /* 方法名: GetTableName
       解释: 获取表名
       返回值: []string

    */

func (t *DicMainEntity) GetTableName() string{
    return "dic_main"
}


    /* 方法名: GetFields
       解释: 获取字段数
       返回值: []string

    */

func (t *DicMainEntity) GetColumnNum() int{
    return 15
}


    /* 方法名: GetFields
       解释: 获取字段
       返回值: []string

    */

func (t *DicMainEntity) GetFields() []string{
    return []string{"code","createBy","createDate","desc","dicName","dicTable","dicType","id","labelKey","parentKey","status","updateBy","updateDate","valueKey","version",}
}

/*** 获取主键/联合主键方法 开始 方法名皆为Get0开头 ***/





    /* 方法名: Get0DicMainPk2
       返回值: []string

    */

func (t *DicMainEntity) Get0DicMainPk2 () []string{
    return []string{"code","version","status",}
}



    /* 方法名: Get0PRIMARY
       返回值: []string

    */

func (t *DicMainEntity) Get0PRIMARY () []string{
    return []string{"id",}
}

/*** 获取主键/联合主键方法 结束 ***/


