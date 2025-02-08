package SimplexCoreDbEntity







/*

    - 64cce7c8-b239-490f-9563-2c9c9dfef611 -

    模板文件 自动生成
    实体映射
    库名: simplex_core_db
    表名: field_type_dic
    解释: Simplex2 CodeCreator
    作者: Simplex2
    版本: 1.0
    创建时间: 2025-02-08 01:56:40
    是否虚拟表: 否
    是否存在审计字段: 否
    是否存在逻辑状态字段: 是

    Template File Auto-Generation
    Entity Mapping

    Database Name: simplex_core_db
    Table Name: field_type_dic
    Description: Simplex2 CodeCreator
    Author: Simplex2
    Version: 1.0
    Creation Time: 2025-02-08 01:56:40
    Is Virtual Table: No
    Has Audit Fields: No
    Has Logical Status Field: Yes

*/

type FieldTypeDicEntity struct {




    /*  字段名: developmentLanguage
        类型: VARCHAR
        字段解释: 开发语言（如 Java、Python 等）
        字段长度: 200
        是否主键: 否
        是否为空: 否
        归宿主键: unique_development_field

    */

    DevelopmentLanguage string `mapstructure:",omitempty" gorm:"column:developmentLanguage" json:"developmentLanguage"`




    /*  字段名: fieldType
        类型: VARCHAR
        字段解释: 字段类型（如 String、Integer 等）
        字段长度: 200
        是否主键: 否
        是否为空: 否
        归宿主键: unique_development_field

    */

    FieldType string `mapstructure:",omitempty" gorm:"column:fieldType" json:"fieldType"`




    /*  字段名: id
        类型: INT
        字段解释: 主键ID
        字段长度: -
        是否主键: 是
        是否为空: 否
        归宿主键: PRIMARY

    */

    Id int `mapstructure:",omitempty" gorm:"primaryKey;column:id" json:"id"`




    /*  字段名: mappedType
        类型: VARCHAR
        字段解释: 映射类型（目标系统中的对应类型，如 VARCHAR、INT 等）
        字段长度: 200
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    MappedType string `mapstructure:",omitempty" gorm:"column:mappedType" json:"mappedType"`




    /*  字段名: remarks
        类型: TEXT
        字段解释: 备注信息，用于记录额外说明
        字段长度: 65535
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    Remarks string `mapstructure:",omitempty" gorm:"column:remarks" json:"remarks"`




    /*  字段名: status
        类型: TINYINT
        字段解释: 状态：1 表示有效，0 表示无效，默认为 1
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    Status int8 `mapstructure:",omitempty" gorm:"column:status" json:"status"`




    /*  字段名: version
        类型: VARCHAR
        字段解释: 版本号（如 v1.0、2023.10.1 等）
        字段长度: 50
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

func (t *FieldTypeDicEntity) GetTableName() string{
    return "field_type_dic"
}


    /* 方法名: GetFields
       解释: 获取字段数
       返回值: []string

    */

func (t *FieldTypeDicEntity) GetColumnNum() int{
    return 7
}


    /* 方法名: GetFields
       解释: 获取字段
       返回值: []string

    */

func (t *FieldTypeDicEntity) GetFields() []string{
    return []string{"developmentLanguage","fieldType","id","mappedType","remarks","status","version",}
}

/*** 获取主键/联合主键方法 开始 方法名皆为Get0开头 ***/





    /* 方法名: Get0PRIMARY
       返回值: []string

    */

func (t *FieldTypeDicEntity) Get0PRIMARY () []string{
    return []string{"id",}
}



    /* 方法名: Get0UniqueDevelopmentField
       返回值: []string

    */

func (t *FieldTypeDicEntity) Get0UniqueDevelopmentField () []string{
    return []string{"developmentLanguage","fieldType",}
}

/*** 获取主键/联合主键方法 结束 ***/


