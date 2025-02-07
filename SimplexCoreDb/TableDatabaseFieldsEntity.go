package SimplexCoreDbEntity



import (
    time "time"
    )



/*
    模板文件 自动生成
    实体映射
    库名: simplex_core_db
    表名: database_fields
    解释: Simplex2 CodeCreator
    作者: Simplex2
    版本: 1.0
    创建时间: 2025-02-07 22:47:55
    是否虚拟表: 否
    是否存在审计字段: 是
    是否存在逻辑状态字段: 否

    Template File Auto-Generation
    Entity Mapping

    Database Name: simplex_core_db
    Table Name: database_fields
    Description: Simplex2 CodeCreator
    Author: Simplex2
    Version: 1.0
    Creation Time: 2025-02-07 22:47:55
    Is Virtual Table: No
    Has Audit Fields: Yes
    Has Logical Status Field: No

*/

type DatabaseFieldsEntity struct {




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
        类型: DATE
        字段解释: 创建日期
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    CreateDate time.Time `mapstructure:",omitempty" gorm:"column:createDate;default:null" json:"createDate"`




    /*  字段名: createTime
        类型: TIME
        字段解释: 创建时间
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    CreateTime time.Time `mapstructure:",omitempty" gorm:"column:createTime;default:null" json:"createTime"`




    /*  字段名: dbName
        类型: VARCHAR
        字段解释: 数据库名
        字段长度: 100
        是否主键: 否
        是否为空: 否
        归宿主键: database_fields_pk

    */

    DbName string `mapstructure:",omitempty" gorm:"column:dbName" json:"dbName"`




    /*  字段名: dbTitle
        类型: VARCHAR
        字段解释: 数据库标题
        字段长度: 255
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    DbTitle string `mapstructure:",omitempty" gorm:"column:dbTitle" json:"dbTitle"`




    /*  字段名: fieldComment
        类型: VARCHAR
        字段解释: 字段备注
        字段长度: 255
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    FieldComment string `mapstructure:",omitempty" gorm:"column:fieldComment" json:"fieldComment"`




    /*  字段名: fieldLength
        类型: INT
        字段解释: 字段长度
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    FieldLength int `mapstructure:",omitempty" gorm:"column:fieldLength" json:"fieldLength"`




    /*  字段名: fieldMappingName
        类型: VARCHAR
        字段解释: 字段映射名称
        字段长度: 255
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    FieldMappingName string `mapstructure:",omitempty" gorm:"column:fieldMappingName" json:"fieldMappingName"`




    /*  字段名: fieldMappingType
        类型: VARCHAR
        字段解释: 字段映射类型
        字段长度: 20
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    FieldMappingType string `mapstructure:",omitempty" gorm:"column:fieldMappingType" json:"fieldMappingType"`




    /*  字段名: fieldName
        类型: VARCHAR
        字段解释: 字段名
        字段长度: 100
        是否主键: 否
        是否为空: 否
        归宿主键: database_fields_pk

    */

    FieldName string `mapstructure:",omitempty" gorm:"column:fieldName" json:"fieldName"`




    /*  字段名: fieldStatus
        类型: TINYINT
        字段解释: 字段状态（1:启用，0:禁用）
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: database_fields_pk

    */

    FieldStatus int8 `mapstructure:",omitempty" gorm:"column:fieldStatus" json:"fieldStatus"`




    /*  字段名: fieldTitle
        类型: VARCHAR
        字段解释: 字段标题
        字段长度: 255
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    FieldTitle string `mapstructure:",omitempty" gorm:"column:fieldTitle" json:"fieldTitle"`




    /*  字段名: fieldValueType
        类型: VARCHAR
        字段解释: 字段值类型
        字段长度: 20
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    FieldValueType string `mapstructure:",omitempty" gorm:"column:fieldValueType" json:"fieldValueType"`




    /*  字段名: generateCount
        类型: INT
        字段解释: 生成次数
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    GenerateCount int `mapstructure:",omitempty" gorm:"column:generateCount" json:"generateCount"`




    /*  字段名: id
        类型: INT
        字段解释: 主键ID
        字段长度: -
        是否主键: 是
        是否为空: 否
        归宿主键: PRIMARY

    */

    Id int `mapstructure:",omitempty" gorm:"primaryKey;column:id" json:"id"`




    /*  字段名: isNullable
        类型: TINYINT
        字段解释: 是否为空（1:允许为空，0:不允许为空）
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    IsNullable int8 `mapstructure:",omitempty" gorm:"column:isNullable" json:"isNullable"`




    /*  字段名: isPrimaryKey
        类型: TINYINT
        字段解释: 是否为主键（1:是，0:否）
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    IsPrimaryKey int8 `mapstructure:",omitempty" gorm:"column:isPrimaryKey" json:"isPrimaryKey"`




    /*  字段名: keys
        类型: VARCHAR
        字段解释: --
        字段长度: 500
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    Keys string `mapstructure:",omitempty" gorm:"column:keys" json:"keys"`




    /*  字段名: tableName
        类型: VARCHAR
        字段解释: 表名
        字段长度: 100
        是否主键: 否
        是否为空: 否
        归宿主键: database_fields_pk

    */

    TableName string `mapstructure:",omitempty" gorm:"column:tableName" json:"tableName"`




    /*  字段名: updateBy
        类型: VARCHAR
        字段解释: 修改人
        字段长度: 50
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    UpdateBy string `mapstructure:",omitempty" gorm:"column:updateBy" json:"updateBy"`




    /*  字段名: updateDate
        类型: DATE
        字段解释: 修改日期
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    UpdateDate time.Time `mapstructure:",omitempty" gorm:"column:updateDate;default:null" json:"updateDate"`




    /*  字段名: updateTime
        类型: TIME
        字段解释: 修改时间
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    UpdateTime time.Time `mapstructure:",omitempty" gorm:"column:updateTime;default:null" json:"updateTime"`
}



    /* 方法名: GetTableName
       解释: 获取表名
       返回值: []string

    */

func (t *DatabaseFieldsEntity) GetTableName() string{
    return "database_fields"
}


    /* 方法名: GetFields
       解释: 获取字段数
       返回值: []string

    */

func (t *DatabaseFieldsEntity) GetColumnNum() int{
    return 22
}


    /* 方法名: GetFields
       解释: 获取字段
       返回值: []string

    */

func (t *DatabaseFieldsEntity) GetFields() []string{
    return []string{"createBy","createDate","createTime","dbName","dbTitle","fieldComment","fieldLength","fieldMappingName","fieldMappingType","fieldName","fieldStatus","fieldTitle","fieldValueType","generateCount","id","isNullable","isPrimaryKey","keys","tableName","updateBy","updateDate","updateTime",}
}

/*** 获取主键/联合主键方法 开始 方法名皆为Get0开头 ***/





    /* 方法名: Get0DatabaseFieldsPk
       返回值: []string

    */

func (t *DatabaseFieldsEntity) Get0DatabaseFieldsPk () []string{
    return []string{"dbName","tableName","fieldName","fieldStatus",}
}



    /* 方法名: Get0PRIMARY
       返回值: []string

    */

func (t *DatabaseFieldsEntity) Get0PRIMARY () []string{
    return []string{"id",}
}

/*** 获取主键/联合主键方法 结束 ***/


