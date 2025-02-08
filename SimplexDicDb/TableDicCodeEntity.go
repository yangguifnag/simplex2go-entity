package SimplexDicDbEntity



import (
    time "time"
    )



/*

    - 64cce7c8-b239-490f-9563-2c9c9dfef611 -

    模板文件 自动生成
    实体映射
    库名: simplex_dic_db
    表名: dic_code
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
    Table Name: dic_code
    Description: Simplex2 CodeCreator
    Author: Simplex2
    Version: 1.0
    Creation Time: 2025-02-08 01:56:40
    Is Virtual Table: No
    Has Audit Fields: No
    Has Logical Status Field: Yes

*/

type DicCodeEntity struct {




    /*  字段名: code
        类型: VARCHAR
        字段解释: 字典码
        字段长度: 100
        是否主键: 否
        是否为空: 否
        归宿主键: dic_code_pk_2

    */

    Code string `mapstructure:",omitempty" gorm:"column:code" json:"code"`




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




    /*  字段名: order
        类型: DOUBLE
        字段解释: 排序
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    Order float64 `mapstructure:",omitempty" gorm:"column:order" json:"order"`




    /*  字段名: parentId
        类型: INT
        字段解释: 父id
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    ParentId int `mapstructure:",omitempty" gorm:"column:parentId" json:"parentId"`




    /*  字段名: status
        类型: INT
        字段解释: 状态
        字段长度: -
        是否主键: 否
        是否为空: 否
        归宿主键: dic_code_pk_2

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




    /*  字段名: value
        类型: VARCHAR
        字段解释: 码值
        字段长度: 200
        是否主键: 否
        是否为空: 否
        归宿主键: dic_code_pk_2

    */

    Value string `mapstructure:",omitempty" gorm:"column:value" json:"value"`




    /*  字段名: valueName
        类型: VARCHAR
        字段解释: 值名称
        字段长度: 200
        是否主键: 否
        是否为空: 否
        归宿主键: -

    */

    ValueName string `mapstructure:",omitempty" gorm:"column:valueName" json:"valueName"`




    /*  字段名: version
        类型: VARCHAR
        字段解释: --
        字段长度: 20
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

func (t *DicCodeEntity) GetTableName() string{
    return "dic_code"
}


    /* 方法名: GetFields
       解释: 获取字段数
       返回值: []string

    */

func (t *DicCodeEntity) GetColumnNum() int{
    return 12
}


    /* 方法名: GetFields
       解释: 获取字段
       返回值: []string

    */

func (t *DicCodeEntity) GetFields() []string{
    return []string{"code","createBy","createDate","id","order","parentId","status","updateBy","updateDate","value","valueName","version",}
}

/*** 获取主键/联合主键方法 开始 方法名皆为Get0开头 ***/





    /* 方法名: Get0DicCodePk2
       返回值: []string

    */

func (t *DicCodeEntity) Get0DicCodePk2 () []string{
    return []string{"code","value","status",}
}



    /* 方法名: Get0PRIMARY
       返回值: []string

    */

func (t *DicCodeEntity) Get0PRIMARY () []string{
    return []string{"id",}
}

/*** 获取主键/联合主键方法 结束 ***/


