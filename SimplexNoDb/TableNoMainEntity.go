package SimplexNoDbEntity



import (
    time "time"
    )



/*

    - ce5959e0-25a5-40f3-a5b0-f70b8f7c6934 -

    模板文件 自动生成
    实体映射
    库名: simplex_no_db
    表名: no_main
    解释: Simplex2 CodeCreator
    作者: Simplex2
    版本: 1.0
    创建时间: 2025-02-08 01:17:50
    是否虚拟表: 否
    是否存在审计字段: 否
    是否存在逻辑状态字段: 是

    Template File Auto-Generation
    Entity Mapping

    Database Name: simplex_no_db
    Table Name: no_main
    Description: Simplex2 CodeCreator
    Author: Simplex2
    Version: 1.0
    Creation Time: 2025-02-08 01:17:50
    Is Virtual Table: No
    Has Audit Fields: No
    Has Logical Status Field: Yes

*/

type NoMainEntity struct {




    /*  字段名: beginDate
        类型: DATETIME
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    BeginDate time.Time `mapstructure:",omitempty" gorm:"column:beginDate;default:null" json:"beginDate"`




    /*  字段名: createBy
        类型: VARCHAR
        字段解释: --
        字段长度: 50
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    CreateBy string `mapstructure:",omitempty" gorm:"column:createBy" json:"createBy"`




    /*  字段名: createDate
        类型: DATETIME
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    CreateDate time.Time `mapstructure:",omitempty" gorm:"column:createDate;default:null" json:"createDate"`




    /*  字段名: desc
        类型: VARCHAR
        字段解释: --
        字段长度: 100
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    Desc string `mapstructure:",omitempty" gorm:"column:desc" json:"desc"`




    /*  字段名: endDate
        类型: DATETIME
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    EndDate time.Time `mapstructure:",omitempty" gorm:"column:endDate;default:null" json:"endDate"`




    /*  字段名: id
        类型: INT
        字段解释: --
        字段长度: -
        是否主键: 是
        是否为空: 否
        归宿主键: PRIMARY

    */

    Id int `mapstructure:",omitempty" gorm:"primaryKey;column:id" json:"id"`




    /*  字段名: maxNo
        类型: VARCHAR
        字段解释: --
        字段长度: 255
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    MaxNo string `mapstructure:",omitempty" gorm:"column:maxNo" json:"maxNo"`




    /*  字段名: noCode
        类型: VARCHAR
        字段解释: 编码代号
        字段长度: 50
        是否主键: 否
        是否为空: 否
        归宿主键: no_main_pk_2

    */

    NoCode string `mapstructure:",omitempty" gorm:"column:noCode" json:"noCode"`




    /*  字段名: noName
        类型: VARCHAR
        字段解释: 号码名称
        字段长度: 50
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    NoName string `mapstructure:",omitempty" gorm:"column:noName" json:"noName"`




    /*  字段名: noType
        类型: VARCHAR
        字段解释: --
        字段长度: 50
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    NoType string `mapstructure:",omitempty" gorm:"column:noType" json:"noType"`




    /*  字段名: status
        类型: INT
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    Status int `mapstructure:",omitempty" gorm:"column:status" json:"status"`




    /*  字段名: updateBy
        类型: VARCHAR
        字段解释: --
        字段长度: 50
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    UpdateBy string `mapstructure:",omitempty" gorm:"column:updateBy" json:"updateBy"`




    /*  字段名: updateDate
        类型: DATETIME
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    UpdateDate time.Time `mapstructure:",omitempty" gorm:"column:updateDate;default:null" json:"updateDate"`




    /*  字段名: version
        类型: VARCHAR
        字段解释: --
        字段长度: 10
        是否主键: 否
        是否为空: 否
        归宿主键: no_main_pk_2

    */

    Version string `mapstructure:",omitempty" gorm:"column:version" json:"version"`
}



    /* 方法名: GetTableName
       解释: 获取表名
       返回值: []string

    */

func (t *NoMainEntity) GetTableName() string{
    return "no_main"
}


    /* 方法名: GetFields
       解释: 获取字段数
       返回值: []string

    */

func (t *NoMainEntity) GetColumnNum() int{
    return 14
}


    /* 方法名: GetFields
       解释: 获取字段
       返回值: []string

    */

func (t *NoMainEntity) GetFields() []string{
    return []string{"beginDate","createBy","createDate","desc","endDate","id","maxNo","noCode","noName","noType","status","updateBy","updateDate","version",}
}

/*** 获取主键/联合主键方法 开始 方法名皆为Get0开头 ***/





    /* 方法名: Get0NoMainPk2
       返回值: []string

    */

func (t *NoMainEntity) Get0NoMainPk2 () []string{
    return []string{"noCode","version",}
}



    /* 方法名: Get0PRIMARY
       返回值: []string

    */

func (t *NoMainEntity) Get0PRIMARY () []string{
    return []string{"id",}
}

/*** 获取主键/联合主键方法 结束 ***/


