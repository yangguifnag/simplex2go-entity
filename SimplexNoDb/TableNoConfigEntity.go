package SimplexNoDbEntity



import (
    time "time"
    )



/*

    - 64cce7c8-b239-490f-9563-2c9c9dfef611 -

    模板文件 自动生成
    实体映射
    库名: simplex_no_db
    表名: no_config
    解释: Simplex2 CodeCreator
    作者: Simplex2
    版本: 1.0
    创建时间: 2025-02-08 01:56:40
    是否虚拟表: 否
    是否存在审计字段: 否
    是否存在逻辑状态字段: 是

    Template File Auto-Generation
    Entity Mapping

    Database Name: simplex_no_db
    Table Name: no_config
    Description: Simplex2 CodeCreator
    Author: Simplex2
    Version: 1.0
    Creation Time: 2025-02-08 01:56:40
    Is Virtual Table: No
    Has Audit Fields: No
    Has Logical Status Field: Yes

*/

type NoConfigEntity struct {




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




    /*  字段名: defaultValue
        类型: VARCHAR
        字段解释: --
        字段长度: 100
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    DefaultValue string `mapstructure:",omitempty" gorm:"column:defaultValue" json:"defaultValue"`




    /*  字段名: id
        类型: INT
        字段解释: --
        字段长度: -
        是否主键: 是
        是否为空: 否
        归宿主键: PRIMARY

    */

    Id int `mapstructure:",omitempty" gorm:"primaryKey;column:id" json:"id"`




    /*  字段名: isTimeliness
        类型: TINYINT
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    IsTimeliness int8 `mapstructure:",omitempty" gorm:"column:isTimeliness" json:"isTimeliness"`




    /*  字段名: maxValue
        类型: VARCHAR
        字段解释: --
        字段长度: 100
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    MaxValue string `mapstructure:",omitempty" gorm:"column:maxValue" json:"maxValue"`




    /*  字段名: minValue
        类型: VARCHAR
        字段解释: --
        字段长度: 100
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    MinValue string `mapstructure:",omitempty" gorm:"column:minValue" json:"minValue"`




    /*  字段名: noCode
        类型: VARCHAR
        字段解释: --
        字段长度: 50
        是否主键: 否
        是否为空: 否
        归宿主键: no_config_pk_2

    */

    NoCode string `mapstructure:",omitempty" gorm:"column:noCode" json:"noCode"`




    /*  字段名: order
        类型: DOUBLE
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: no_config_pk_2

    */

    Order float64 `mapstructure:",omitempty" gorm:"column:order" json:"order"`




    /*  字段名: part
        类型: VARCHAR
        字段解释: --
        字段长度: 50
        是否主键: 否
        是否为空: 否
        归宿主键: no_config_pk_2

    */

    Part string `mapstructure:",omitempty" gorm:"column:part" json:"part"`




    /*  字段名: partLength
        类型: INT
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    PartLength int `mapstructure:",omitempty" gorm:"column:partLength" json:"partLength"`




    /*  字段名: partName
        类型: VARCHAR
        字段解释: --
        字段长度: 100
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    PartName string `mapstructure:",omitempty" gorm:"column:partName" json:"partName"`




    /*  字段名: partType
        类型: VARCHAR
        字段解释: --
        字段长度: 50
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    PartType string `mapstructure:",omitempty" gorm:"column:partType" json:"partType"`




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
        归宿主键: -

    */

    Version string `mapstructure:",omitempty" gorm:"column:version" json:"version"`
}



    /* 方法名: GetTableName
       解释: 获取表名
       返回值: []string

    */

func (t *NoConfigEntity) GetTableName() string{
    return "no_config"
}


    /* 方法名: GetFields
       解释: 获取字段数
       返回值: []string

    */

func (t *NoConfigEntity) GetColumnNum() int{
    return 17
}


    /* 方法名: GetFields
       解释: 获取字段
       返回值: []string

    */

func (t *NoConfigEntity) GetFields() []string{
    return []string{"createBy","createDate","defaultValue","id","isTimeliness","maxValue","minValue","noCode","order","part","partLength","partName","partType","status","updateBy","updateDate","version",}
}

/*** 获取主键/联合主键方法 开始 方法名皆为Get0开头 ***/





    /* 方法名: Get0NoConfigPk2
       返回值: []string

    */

func (t *NoConfigEntity) Get0NoConfigPk2 () []string{
    return []string{"noCode","part","order",}
}



    /* 方法名: Get0PRIMARY
       返回值: []string

    */

func (t *NoConfigEntity) Get0PRIMARY () []string{
    return []string{"id",}
}

/*** 获取主键/联合主键方法 结束 ***/


