package SimplexCoreDbEntity



import (
    time "time"
    )



/*
    模板文件 自动生成
    实体映射
    库名: simplex_core_db
    表名: core_table_track
    解释: Simplex2 CodeCreator
    作者: Simplex2
    版本: 1.0
    创建时间: 2025-02-07 22:47:55
    是否虚拟表: 否
    是否存在审计字段: 否
    是否存在逻辑状态字段: 是

    Template File Auto-Generation
    Entity Mapping

    Database Name: simplex_core_db
    Table Name: core_table_track
    Description: Simplex2 CodeCreator
    Author: Simplex2
    Version: 1.0
    Creation Time: 2025-02-07 22:47:55
    Is Virtual Table: No
    Has Audit Fields: No
    Has Logical Status Field: Yes

*/

type CoreTableTrackEntity struct {




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




    /*  字段名: dbName
        类型: VARCHAR
        字段解释: --
        字段长度: 200
        是否主键: 否
        是否为空: 否
        归宿主键: core_table_track_pk_2

    */

    DbName string `mapstructure:",omitempty" gorm:"column:dbName" json:"dbName"`




    /*  字段名: id
        类型: INT
        字段解释: --
        字段长度: -
        是否主键: 是
        是否为空: 否
        归宿主键: core_table_track_pk_2,PRIMARY

    */

    Id int `mapstructure:",omitempty" gorm:"primaryKey;column:id" json:"id"`




    /*  字段名: status
        类型: INT
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    Status int `mapstructure:",omitempty" gorm:"column:status" json:"status"`




    /*  字段名: tableName
        类型: VARCHAR
        字段解释: --
        字段长度: 200
        是否主键: 否
        是否为空: 否
        归宿主键: core_table_track_pk_2

    */

    TableName string `mapstructure:",omitempty" gorm:"column:tableName" json:"tableName"`




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

func (t *CoreTableTrackEntity) GetTableName() string{
    return "core_table_track"
}


    /* 方法名: GetFields
       解释: 获取字段数
       返回值: []string

    */

func (t *CoreTableTrackEntity) GetColumnNum() int{
    return 8
}


    /* 方法名: GetFields
       解释: 获取字段
       返回值: []string

    */

func (t *CoreTableTrackEntity) GetFields() []string{
    return []string{"createBy","createDate","dbName","id","status","tableName","updateBy","updateDate",}
}

/*** 获取主键/联合主键方法 开始 方法名皆为Get0开头 ***/





    /* 方法名: Get0CoreTableTrackPk2
       返回值: []string

    */

func (t *CoreTableTrackEntity) Get0CoreTableTrackPk2 () []string{
    return []string{"id","dbName","tableName",}
}



    /* 方法名: Get0PRIMARY
       返回值: []string

    */

func (t *CoreTableTrackEntity) Get0PRIMARY () []string{
    return []string{"id",}
}

/*** 获取主键/联合主键方法 结束 ***/


