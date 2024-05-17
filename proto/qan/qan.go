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

package qan

import (
	"time"

	"github.com/shatteredsilicon/ssm/proto/metrics"
	"github.com/shatteredsilicon/ssm/proto/query"
)

// A Class represents all events with the same fingerprint and class ID.
// This is only enforced by convention, so be careful not to mix events from
// different classes.
type Class struct {
	Id            string       // 32-character hex checksum of fingerprint
	Fingerprint   string       // canonical form of query: values replaced with "?"
	Metrics       *Metrics     // statistics for each metric, e.g. max Query_time
	TotalQueries  uint         // total number of queries in class
	UniqueQueries uint         // unique number of queries in class
	Example       *Example     `json:",omitempty"` // sample query with max Query_time
	UserSources   []UserSource // user@host sources parsed from slow log
	StartAt       time.Time    // start time of the earliest query of this class
	EndAt         time.Time    // end time of the latest query of this class
	// --
	Outliers uint   `json:"-"`
	LastDb   string `json:"-"`
	Sample   bool   `json:"-"`
}

type Metrics struct {
	TimeMetrics   map[string]*TimeStats   `json:",omitempty"`
	NumberMetrics map[string]*NumberStats `json:",omitempty"`
	BoolMetrics   map[string]*BoolStats   `json:",omitempty"`
}

// TimeStats are microsecond-based metrics like Query_time and Lock_time.
type TimeStats struct {
	Vals       []float64 `json:"-"`
	Sum        float64
	Min        *float64 `json:",omitempty"`
	Avg        *float64 `json:",omitempty"`
	Med        *float64 `json:",omitempty"` // median
	P95        *float64 `json:",omitempty"` // 95th percentile
	Max        *float64 `json:",omitempty"`
	OutlierSum float64  `json:"-"`
}

// NumberStats are integer-based metrics like Rows_sent and Merge_passes.
type NumberStats struct {
	Vals       []uint64 `json:"-"`
	Sum        uint64
	Min        *uint64 `json:",omitempty"`
	Avg        *uint64 `json:",omitempty"`
	Med        *uint64 `json:",omitempty"` // median
	P95        *uint64 `json:",omitempty"` // 95th percentile
	Max        *uint64 `json:",omitempty"`
	OutlierSum uint64  `json:"-"`
}

// BoolStats are boolean-based metrics like QC_Hit and Filesort.
type BoolStats struct {
	Sum        uint64 // %true = Sum/Cnt
	OutlierSum uint64 `json:"-"`
}

// A Example is a real query and its database, timestamp, and Query_time.
// If the query is larger than MaxExampleBytes, it is truncated and TruncatedExampleSuffix
// is appended.
type Example struct {
	QueryTime float64 // Query_time
	Db        string  // Schema: <db> or USE <db>
	Query     string  // truncated to MaxExampleBytes
	Explain   string  // explain
	Size      int     `json:",omitempty"` // Original size of query.
	Ts        string  `json:",omitempty"` // in MySQL time zone
}

// A UserSource is a user@host source parsed from slow log
type UserSource struct {
	Ts   int64 // unix nano timestamp
	User string
	Host string
}

type Report struct {
	UUID    string    // UUID of MySQL instance
	StartTs time.Time // Start time of interval, UTC
	EndTs   time.Time // Stop time of interval, UTC
	RunTime float64   // Time parsing data, seconds
	Global  *Class    // Metrics for all data
	Class   []*Class  // per-class metrics
	// slow log:
	SlowLogFile     string `json:",omitempty"` // not slow_query_log_file if rotated
	SlowLogFileSize int64  `json:",omitempty"`
	StartOffset     int64  `json:",omitempty"` // parsing starts
	EndOffset       int64  `json:",omitempty"` // parsing stops, but...
	StopOffset      int64  `json:",omitempty"` // ...parsing didn't complete if stop < end
	RateLimit       uint   `json:",omitempty"` // Percona Server rate limit
}

type Profile struct {
	InstanceId   string      // UUID of MySQL instance
	Begin        time.Time   // time range [Begin, End)
	End          time.Time   // time range [Being, End)
	TotalTime    uint        // total seconds in time range minus gaps (missing periods)
	TotalQueries uint        // total unique class queries in time range
	RankBy       RankBy      // criteria for ranking queries compared to global
	Query        []QueryRank // 0=global, 1..N=queries
}

type RankBy struct {
	Metric string // default: Query_time
	Stat   string // default: sum
	Limit  uint   // default: 10
}

// start_ts, query_count, Query_time_sum
type QueryLog struct {
	Point          uint
	Start_ts       time.Time
	Query_count    float32
	Query_load     float32
	Query_time_avg float32
}

type QueryRank struct {
	Rank        uint    // compared to global, same as Profile.Ranks index
	Percentage  float64 // of global value
	Id          string  // hex checksum
	Abstract    string  // e.g. SELECT tbl
	Fingerprint string  // e.g. SELECT tbl
	QPS         float64 // ResponseTime.Cnt / Profile.TotalTime
	Load        float64 // Query_time_sum / (Profile.End - Profile.Begin)
	Log         []QueryLog
	Stats       metrics.Stats // this query's Profile.Metric stats
}

type QueryReport struct {
	InstanceId string                   // UUID of MySQL instance
	Begin      time.Time                // time range [Begin, End)
	End        time.Time                // time range [Being, End)
	Query      query.Query              // id, abstract, fingerprint, etc.
	Metrics    map[string]metrics.Stats // keyed on metric name, e.g. Query_time
	Example    query.Example            // query example
	Sparks     []interface{}            `json:",omitempty"`
	Metrics2   interface{}              `json:",omitempty"`
	Sparks2    interface{}              `json:",omitempty"`
}

type Summary struct {
	InstanceId string                   // UUID of MySQL instance
	Begin      time.Time                // time range [Begin, End)
	End        time.Time                // time range [Being, End)
	Metrics    map[string]metrics.Stats // keyed on metric name, e.g. Query_time
	Sparks     []interface{}            `json:",omitempty"`
	Metrics2   interface{}              `json:",omitempty"`
	Sparks2    interface{}              `json:",omitempty"`
}
