/*
   Copyright (c) 2016, Percona LLC and/or its affiliates. All rights reserved.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package proto

type ExplainQuery struct {
	UUID            string
	Db              string
	Query           string
	Convert         bool // convert if not SELECT and MySQL <= 5.5 or >= 5.6 but no privs
	WithExplainRows []*ExplainRow
}

type ExplainResult struct {
	Classic []*ExplainRow
	JSON    string // since MySQL 5.6.5
}

type ExplainRow struct {
	Id           NullInt64   `slowlog:"id"`
	SelectType   NullString  `slowlog:"select_type"`
	Table        NullString  `slowlog:"table"`
	Partitions   NullString  `slowlog:"partitions"`   // split by comma; since MySQL 5.1
	CreateTable  NullString  `slowlog:"create_table"` // @todo
	Type         NullString  `slowlog:"type"`
	PossibleKeys NullString  `slowlog:"possible_keys"` // split by comma
	Key          NullString  `slowlog:"key"`
	KeyLen       NullString  `slowlog:"key_len"` // https://jira.percona.com/browse/PCT-863
	Ref          NullString  `slowlog:"ref"`
	Rows         NullInt64   `slowlog:"rows"`
	Filtered     NullFloat64 `slowlog:"filtered"` // as of 5.7.3
	Extra        NullString  `slowlog:"extra"`    // split by semicolon
}
