package SimplexCoreDbEntity



import (
    time "time"
    )



/*

    - ce5959e0-25a5-40f3-a5b0-f70b8f7c6934 -

    模板文件 自动生成
    实体映射
    库名: simplex_core_db
    表名: test_table
    解释: Simplex2 CodeCreator
    作者: Simplex2
    版本: 1.0
    创建时间: 2025-02-08 01:17:50
    是否虚拟表: 否
    是否存在审计字段: 是
    是否存在逻辑状态字段: 是

    Template File Auto-Generation
    Entity Mapping

    Database Name: simplex_core_db
    Table Name: test_table
    Description: Simplex2 CodeCreator
    Author: Simplex2
    Version: 1.0
    Creation Time: 2025-02-08 01:17:50
    Is Virtual Table: No
    Has Audit Fields: Yes
    Has Logical Status Field: Yes

*/

type TestTableEntity struct {




    /*  字段名: column_name
        类型: INT
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    ColumnName int `mapstructure:",omitempty" gorm:"column:column_name" json:"column_name"`




    /*  字段名: column_name_2
        类型: VARCHAR
        字段解释: --
        字段长度: 200
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    ColumnName2 string `mapstructure:",omitempty" gorm:"column:column_name_2" json:"column_name_2"`




    /*  字段名: column_name_3
        类型: INT
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    ColumnName3 int `mapstructure:",omitempty" gorm:"column:column_name_3" json:"column_name_3"`




    /*  字段名: column_name_4
        类型: INT
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    ColumnName4 int `mapstructure:",omitempty" gorm:"column:column_name_4" json:"column_name_4"`




    /*  字段名: column_name_5
        类型: INT
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    ColumnName5 int `mapstructure:",omitempty" gorm:"column:column_name_5" json:"column_name_5"`




    /*  字段名: column_name_6
        类型: INT
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    ColumnName6 int `mapstructure:",omitempty" gorm:"column:column_name_6" json:"column_name_6"`




    /*  字段名: column_name_7
        类型: INT
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    ColumnName7 int `mapstructure:",omitempty" gorm:"column:column_name_7" json:"column_name_7"`




    /*  字段名: column_name_8
        类型: INT
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    ColumnName8 int `mapstructure:",omitempty" gorm:"column:column_name_8" json:"column_name_8"`




    /*  字段名: id
        类型: INT
        字段解释: --
        字段长度: -
        是否主键: 是
        是否为空: 否
        归宿主键: PRIMARY

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




    /*  字段名: updateBy
        类型: VARCHAR
        字段解释: --
        字段长度: 20
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    UpdateBy string `mapstructure:",omitempty" gorm:"column:updateBy" json:"updateBy"`




    /*  字段名: updateDate
        类型: DATE
        字段解释: --
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    UpdateDate time.Time `mapstructure:",omitempty" gorm:"column:updateDate;default:null" json:"updateDate"`




    /*  字段名: updateTime
        类型: TIME
        字段解释: --
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

func (t *TestTableEntity) GetTableName() string{
    return "test_table"
}


    /* 方法名: GetFields
       解释: 获取字段数
       返回值: []string

    */

func (t *TestTableEntity) GetColumnNum() int{
    return 13
}


    /* 方法名: GetFields
       解释: 获取字段
       返回值: []string

    */

func (t *TestTableEntity) GetFields() []string{
    return []string{"column_name","column_name_2","column_name_3","column_name_4","column_name_5","column_name_6","column_name_7","column_name_8","id","status","updateBy","updateDate","updateTime",}
}

/*** 获取主键/联合主键方法 开始 方法名皆为Get0开头 ***/





    /* 方法名: Get0PRIMARY
       返回值: []string

    */

func (t *TestTableEntity) Get0PRIMARY () []string{
    return []string{"id",}
}

/*** 获取主键/联合主键方法 结束 ***/


