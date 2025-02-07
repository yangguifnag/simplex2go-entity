package SimplexNoDbEntity



import (
    time "time"
    )



/*

    - ce5959e0-25a5-40f3-a5b0-f70b8f7c6934 -

    模板文件 自动生成
    实体映射
    库名: simplex_no_db
    表名: no_customization_usage
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
    Table Name: no_customization_usage
    Description: Simplex2 CodeCreator
    Author: Simplex2
    Version: 1.0
    Creation Time: 2025-02-08 01:17:50
    Is Virtual Table: No
    Has Audit Fields: No
    Has Logical Status Field: Yes

*/

type NoCustomizationUsageEntity struct {




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




    /*  字段名: id
        类型: INT
        字段解释: --
        字段长度: -
        是否主键: 是
        是否为空: 否
        归宿主键: PRIMARY

    */

    Id int `mapstructure:",omitempty" gorm:"primaryKey;column:id" json:"id"`




    /*  字段名: no
        类型: VARCHAR
        字段解释: --
        字段长度: 200
        是否主键: 否
        是否为空: 否
        归宿主键: no_customization_usage_pk_2

    */

    No string `mapstructure:",omitempty" gorm:"column:no" json:"no"`




    /*  字段名: noCode
        类型: VARCHAR
        字段解释: --
        字段长度: 50
        是否主键: 否
        是否为空: 否
        归宿主键: no_customization_usage_pk_2

    */

    NoCode string `mapstructure:",omitempty" gorm:"column:noCode" json:"noCode"`




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
}



    /* 方法名: GetTableName
       解释: 获取表名
       返回值: []string

    */

func (t *NoCustomizationUsageEntity) GetTableName() string{
    return "no_customization_usage"
}


    /* 方法名: GetFields
       解释: 获取字段数
       返回值: []string

    */

func (t *NoCustomizationUsageEntity) GetColumnNum() int{
    return 8
}


    /* 方法名: GetFields
       解释: 获取字段
       返回值: []string

    */

func (t *NoCustomizationUsageEntity) GetFields() []string{
    return []string{"createBy","createDate","id","no","noCode","status","updateBy","updateDate",}
}

/*** 获取主键/联合主键方法 开始 方法名皆为Get0开头 ***/





    /* 方法名: Get0NoCustomizationUsagePk2
       返回值: []string

    */

func (t *NoCustomizationUsageEntity) Get0NoCustomizationUsagePk2 () []string{
    return []string{"noCode","no",}
}



    /* 方法名: Get0PRIMARY
       返回值: []string

    */

func (t *NoCustomizationUsageEntity) Get0PRIMARY () []string{
    return []string{"id",}
}

/*** 获取主键/联合主键方法 结束 ***/


