package proto

import "github.com/shatteredsilicon/ssm/proto/query"

// DBObjectType represents the type of database object,
// such as table, procedure, view and etc
type DBObjectType int8

const (
	// TypeDBTable database object - table
	TypeDBTable = iota
	// TypeDBProcedure database object - procedure
	TypeDBProcedure
	// TypeDBView database object - view
	TypeDBView
)

type Table query.Table

type TableInfoQuery struct {
	UUID   string
	Create []Table // SHOW CREATE TABLE Db.Table
	Index  []Table // SHOW INDEXES FROM Db.Table
	Status []Table // SHOW TABLE STATUS FROM Db LIKE 'Table'
}

// ShowIndexRow describes one row from `SHOW INDEX FROM %s` query.
//# mysql -e 'SHOW INDEX FROM mysql.user\G'
//*************************** 1. row ***************************
//        Table: user
//   Non_unique: 0
//     Key_name: PRIMARY
// Seq_in_index: 1
//  Column_name: Host
//    Collation: A
//  Cardinality: 2
//     Sub_part: NULL
//       Packed: NULL
//         Null:
//   Index_type: BTREE
//      Comment:
//Index_comment:
//      Visible: YES
//*************************** 2. row ***************************
//        Table: user
//   Non_unique: 0
//     Key_name: PRIMARY
// Seq_in_index: 2
//  Column_name: User
//    Collation: A
//  Cardinality: 5
//     Sub_part: NULL
//       Packed: NULL
//         Null:
//   Index_type: BTREE
//      Comment:
//Index_comment:
//      Visible: YES
//
//# mysql -e 'DESCRIBE INFORMATION_SCHEMA.STATISTICS'
//+---------------+------------------+------+-----+---------+-------+
//| Field         | Type             | Null | Key | Default | Extra |
//+---------------+------------------+------+-----+---------+-------+
//| TABLE_CATALOG | varchar(64)      | NO   |     | NULL    |       |
//| TABLE_SCHEMA  | varchar(64)      | NO   |     | NULL    |       |
//| TABLE_NAME    | varchar(64)      | NO   |     | NULL    |       |
//| NON_UNIQUE    | int(1)           | NO   |     | 0       |       |
//| INDEX_SCHEMA  | varchar(64)      | NO   |     | NULL    |       |
//| INDEX_NAME    | varchar(64)      | YES  |     | NULL    |       |
//| SEQ_IN_INDEX  | int(10) unsigned | NO   |     | NULL    |       |
//| COLUMN_NAME   | varchar(64)      | YES  |     | NULL    |       |
//| COLLATION     | varchar(1)       | YES  |     | NULL    |       |
//| CARDINALITY   | bigint(21)       | YES  |     | NULL    |       |
//| SUB_PART      | bigint(21)       | YES  |     | NULL    |       |
//| PACKED        | binary(0)        | YES  |     | NULL    |       |
//| NULLABLE      | varchar(3)       | NO   |     |         |       |
//| INDEX_TYPE    | varchar(11)      | NO   |     |         |       |
//| COMMENT       | varchar(8)       | NO   |     |         |       |
//| INDEX_COMMENT | varchar(2048)    | NO   |     | NULL    |       |
//| IS_VISIBLE    | varchar(3)       | NO   |     |         |       |
//+---------------+------------------+------+-----+---------+-------+
type ShowIndexRow struct {
	Table        string
	NonUnique    bool
	KeyName      string
	SeqInIndex   int
	ColumnName   string
	Collation    NullString
	Cardinality  NullInt64
	SubPart      NullInt64
	Packed       NullString
	Null         NullString
	IndexType    string
	Comment      NullString
	IndexComment NullString
	Visible      NullString
}

type ShowTableStatus struct {
	Name          string
	Engine        NullString
	Version       string
	RowFormat     NullString
	Rows          NullInt64
	AvgRowLength  NullInt64
	DataLength    NullInt64
	MaxDataLength NullInt64
	IndexLength   NullInt64
	DataFree      NullInt64
	AutoIncrement NullInt64
	CreateTime    NullString
	UpdateTime    NullString
	CheckTime     NullString
	Collation     NullString
	Checksum      NullString
	CreateOptions NullString
	Comment       NullString
}

// GuessDB structure of guessed database infomation
type GuessDB struct {
	DB          string
	IsAmbiguous bool
}

type TableInfo struct {
	Type    DBObjectType              `json:",omitempty"`
	Create  string                    `json:",omitempty"`
	Index   map[string][]ShowIndexRow `json:",omitempty"`
	Status  *ShowTableStatus          `json:",omitempty"`
	Errors  []string                  `json:",omitempty"`
	GuessDB *GuessDB                  `json:",omitempty"`
}

type TableInfoResult map[string]*TableInfo

// QueryInfoParam param structure of
// QueryInfo cmd api
type QueryInfoParam struct {
	UUID      string
	Table     []Table // SHOW CREATE TABLE Db.Table
	Procedure []query.Procedure
	Index     []Table // SHOW INDEXES FROM Db.Table
	Status    []Table // SHOW TABLE STATUS FROM Db LIKE 'Table'
}

// QueryInfo represents a TABLE/PROCEDURE/VIEW
// structure of QueryInfo cmd api
type QueryInfo struct {
	Type   DBObjectType
	Create string                    `json:",omitempty"`
	Index  map[string][]ShowIndexRow `json:",omitempty"`
	Status *ShowTableStatus          `json:",omitempty"`
	Errors []string                  `json:",omitempty"`
}

// QueryInfoResult represents the response
// of QueryInfo cmd api
type QueryInfoResult struct {
	GuessDB *GuessDB
	Info    map[string]*QueryInfo
}
