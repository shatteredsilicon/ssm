// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qanpb/qan.proto

package qanpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ExampleFormat is format of query example: real or query without values.
type ExampleFormat int32

const (
	ExampleFormat_EXAMPLE_FORMAT_INVALID ExampleFormat = 0
	ExampleFormat_EXAMPLE                ExampleFormat = 1
	ExampleFormat_FINGERPRINT            ExampleFormat = 2
)

var ExampleFormat_name = map[int32]string{
	0: "EXAMPLE_FORMAT_INVALID",
	1: "EXAMPLE",
	2: "FINGERPRINT",
}

var ExampleFormat_value = map[string]int32{
	"EXAMPLE_FORMAT_INVALID": 0,
	"EXAMPLE":                1,
	"FINGERPRINT":            2,
}

func (x ExampleFormat) String() string {
	return proto.EnumName(ExampleFormat_name, int32(x))
}

func (ExampleFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6cc2ea5e264b89be, []int{0}
}

// ExampleType is a type of query example selected for this query class in given period of time.
type ExampleType int32

const (
	ExampleType_EXAMPLE_TYPE_INVALID ExampleType = 0
	ExampleType_RANDOM               ExampleType = 1
	ExampleType_SLOWEST              ExampleType = 2
	ExampleType_FASTEST              ExampleType = 3
	ExampleType_WITH_ERROR           ExampleType = 4
)

var ExampleType_name = map[int32]string{
	0: "EXAMPLE_TYPE_INVALID",
	1: "RANDOM",
	2: "SLOWEST",
	3: "FASTEST",
	4: "WITH_ERROR",
}

var ExampleType_value = map[string]int32{
	"EXAMPLE_TYPE_INVALID": 0,
	"RANDOM":               1,
	"SLOWEST":              2,
	"FASTEST":              3,
	"WITH_ERROR":           4,
}

func (x ExampleType) String() string {
	return proto.EnumName(ExampleType_name, int32(x))
}

func (ExampleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6cc2ea5e264b89be, []int{1}
}

// Point contains values that represents abscissa (time) and ordinate (volume etc.)
// of every point in a coordinate system of Sparklines.
type Point struct {
	// The serial number of the chart point from the largest time in the time interval to the lowest time in the time range.
	Point uint32 `protobuf:"varint,1,opt,name=point,proto3" json:"point,omitempty"`
	// Duration beetween two points.
	TimeFrame uint32 `protobuf:"varint,2,opt,name=time_frame,json=timeFrame,proto3" json:"time_frame,omitempty"`
	// Time of point in format RFC3339.
	Timestamp string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// number of queries in bucket.
	NumQueriesPerSec float32 `protobuf:"fixed32,4,opt,name=num_queries_per_sec,json=numQueriesPerSec,proto3" json:"num_queries_per_sec,omitempty"`
	// The statement execution time in seconds.
	MQueryTimeSumPerSec float32 `protobuf:"fixed32,5,opt,name=m_query_time_sum_per_sec,json=mQueryTimeSumPerSec,proto3" json:"m_query_time_sum_per_sec,omitempty"`
	// The time to acquire locks in seconds.
	MLockTimeSumPerSec float32 `protobuf:"fixed32,6,opt,name=m_lock_time_sum_per_sec,json=mLockTimeSumPerSec,proto3" json:"m_lock_time_sum_per_sec,omitempty"`
	// The number of rows sent to the client.
	MRowsSentSumPerSec float32 `protobuf:"fixed32,7,opt,name=m_rows_sent_sum_per_sec,json=mRowsSentSumPerSec,proto3" json:"m_rows_sent_sum_per_sec,omitempty"`
	// Number of rows scanned - SELECT.
	MRowsExaminedSumPerSec float32 `protobuf:"fixed32,8,opt,name=m_rows_examined_sum_per_sec,json=mRowsExaminedSumPerSec,proto3" json:"m_rows_examined_sum_per_sec,omitempty"`
	// Number of rows changed - UPDATE, DELETE, INSERT.
	MRowsAffectedSumPerSec float32 `protobuf:"fixed32,9,opt,name=m_rows_affected_sum_per_sec,json=mRowsAffectedSumPerSec,proto3" json:"m_rows_affected_sum_per_sec,omitempty"`
	// The number of rows read from tables.
	MRowsReadSumPerSec float32 `protobuf:"fixed32,10,opt,name=m_rows_read_sum_per_sec,json=mRowsReadSumPerSec,proto3" json:"m_rows_read_sum_per_sec,omitempty"`
	// The number of merge passes that the sort algorithm has had to do.
	MMergePassesSumPerSec float32 `protobuf:"fixed32,11,opt,name=m_merge_passes_sum_per_sec,json=mMergePassesSumPerSec,proto3" json:"m_merge_passes_sum_per_sec,omitempty"`
	// Counts the number of page read operations scheduled.
	MInnodbIoROpsSumPerSec float32 `protobuf:"fixed32,12,opt,name=m_innodb_io_r_ops_sum_per_sec,json=mInnodbIoROpsSumPerSec,proto3" json:"m_innodb_io_r_ops_sum_per_sec,omitempty"`
	// Similar to innodb_IO_r_ops, but the unit is bytes.
	MInnodbIoRBytesSumPerSec float32 `protobuf:"fixed32,13,opt,name=m_innodb_io_r_bytes_sum_per_sec,json=mInnodbIoRBytesSumPerSec,proto3" json:"m_innodb_io_r_bytes_sum_per_sec,omitempty"`
	// Shows how long (in seconds) it took InnoDB to actually read the data from storage.
	MInnodbIoRWaitSumPerSec float32 `protobuf:"fixed32,14,opt,name=m_innodb_io_r_wait_sum_per_sec,json=mInnodbIoRWaitSumPerSec,proto3" json:"m_innodb_io_r_wait_sum_per_sec,omitempty"`
	// Shows how long (in seconds) the query waited for row locks.
	MInnodbRecLockWaitSumPerSec float32 `protobuf:"fixed32,15,opt,name=m_innodb_rec_lock_wait_sum_per_sec,json=mInnodbRecLockWaitSumPerSec,proto3" json:"m_innodb_rec_lock_wait_sum_per_sec,omitempty"`
	// Shows how long (in seconds) the query spent either waiting to enter the InnoDB queue or inside that queue waiting for execution.
	MInnodbQueueWaitSumPerSec float32 `protobuf:"fixed32,16,opt,name=m_innodb_queue_wait_sum_per_sec,json=mInnodbQueueWaitSumPerSec,proto3" json:"m_innodb_queue_wait_sum_per_sec,omitempty"`
	// Counts approximately the number of unique pages the query accessed.
	MInnodbPagesDistinctSumPerSec float32 `protobuf:"fixed32,17,opt,name=m_innodb_pages_distinct_sum_per_sec,json=mInnodbPagesDistinctSumPerSec,proto3" json:"m_innodb_pages_distinct_sum_per_sec,omitempty"`
	// Shows how long the query is.
	MQueryLengthSumPerSec float32 `protobuf:"fixed32,18,opt,name=m_query_length_sum_per_sec,json=mQueryLengthSumPerSec,proto3" json:"m_query_length_sum_per_sec,omitempty"`
	// The number of bytes sent to all clients.
	MBytesSentSumPerSec float32 `protobuf:"fixed32,19,opt,name=m_bytes_sent_sum_per_sec,json=mBytesSentSumPerSec,proto3" json:"m_bytes_sent_sum_per_sec,omitempty"`
	// Number of temporary tables created on memory for the query.
	MTmpTablesSumPerSec float32 `protobuf:"fixed32,20,opt,name=m_tmp_tables_sum_per_sec,json=mTmpTablesSumPerSec,proto3" json:"m_tmp_tables_sum_per_sec,omitempty"`
	// Number of temporary tables created on disk for the query.
	MTmpDiskTablesSumPerSec float32 `protobuf:"fixed32,21,opt,name=m_tmp_disk_tables_sum_per_sec,json=mTmpDiskTablesSumPerSec,proto3" json:"m_tmp_disk_tables_sum_per_sec,omitempty"`
	// Total Size in bytes for all temporary tables used in the query.
	MTmpTableSizesSumPerSec float32 `protobuf:"fixed32,22,opt,name=m_tmp_table_sizes_sum_per_sec,json=mTmpTableSizesSumPerSec,proto3" json:"m_tmp_table_sizes_sum_per_sec,omitempty"`
	//
	// Boolean metrics:
	// - *_sum_per_sec - how many times this matric was true.
	//
	// Query Cache hits.
	MQcHitSumPerSec float32 `protobuf:"fixed32,23,opt,name=m_qc_hit_sum_per_sec,json=mQcHitSumPerSec,proto3" json:"m_qc_hit_sum_per_sec,omitempty"`
	// The query performed a full table scan.
	MFullScanSumPerSec float32 `protobuf:"fixed32,24,opt,name=m_full_scan_sum_per_sec,json=mFullScanSumPerSec,proto3" json:"m_full_scan_sum_per_sec,omitempty"`
	// The query performed a full join (a join without indexes).
	MFullJoinSumPerSec float32 `protobuf:"fixed32,25,opt,name=m_full_join_sum_per_sec,json=mFullJoinSumPerSec,proto3" json:"m_full_join_sum_per_sec,omitempty"`
	// The query created an implicit internal temporary table.
	MTmpTableSumPerSec float32 `protobuf:"fixed32,26,opt,name=m_tmp_table_sum_per_sec,json=mTmpTableSumPerSec,proto3" json:"m_tmp_table_sum_per_sec,omitempty"`
	// The querys temporary table was stored on disk.
	MTmpTableOnDiskSumPerSec float32 `protobuf:"fixed32,27,opt,name=m_tmp_table_on_disk_sum_per_sec,json=mTmpTableOnDiskSumPerSec,proto3" json:"m_tmp_table_on_disk_sum_per_sec,omitempty"`
	// The query used a filesort.
	MFilesortSumPerSec float32 `protobuf:"fixed32,28,opt,name=m_filesort_sum_per_sec,json=mFilesortSumPerSec,proto3" json:"m_filesort_sum_per_sec,omitempty"`
	// The filesort was performed on disk.
	MFilesortOnDiskSumPerSec float32 `protobuf:"fixed32,29,opt,name=m_filesort_on_disk_sum_per_sec,json=mFilesortOnDiskSumPerSec,proto3" json:"m_filesort_on_disk_sum_per_sec,omitempty"`
	// The number of joins that used a range search on a reference table.
	MSelectFullRangeJoinSumPerSec float32 `protobuf:"fixed32,30,opt,name=m_select_full_range_join_sum_per_sec,json=mSelectFullRangeJoinSumPerSec,proto3" json:"m_select_full_range_join_sum_per_sec,omitempty"`
	// The number of joins that used ranges on the first table.
	MSelectRangeSumPerSec float32 `protobuf:"fixed32,31,opt,name=m_select_range_sum_per_sec,json=mSelectRangeSumPerSec,proto3" json:"m_select_range_sum_per_sec,omitempty"`
	// The number of joins without keys that check for key usage after each row.
	MSelectRangeCheckSumPerSec float32 `protobuf:"fixed32,32,opt,name=m_select_range_check_sum_per_sec,json=mSelectRangeCheckSumPerSec,proto3" json:"m_select_range_check_sum_per_sec,omitempty"`
	// The number of sorts that were done using ranges.
	MSortRangeSumPerSec float32 `protobuf:"fixed32,33,opt,name=m_sort_range_sum_per_sec,json=mSortRangeSumPerSec,proto3" json:"m_sort_range_sum_per_sec,omitempty"`
	// The number of sorted rows.
	MSortRowsSumPerSec float32 `protobuf:"fixed32,34,opt,name=m_sort_rows_sum_per_sec,json=mSortRowsSumPerSec,proto3" json:"m_sort_rows_sum_per_sec,omitempty"`
	// The number of sorts that were done by scanning the table.
	MSortScanSumPerSec float32 `protobuf:"fixed32,35,opt,name=m_sort_scan_sum_per_sec,json=mSortScanSumPerSec,proto3" json:"m_sort_scan_sum_per_sec,omitempty"`
	// The number of queries without index.
	MNoIndexUsedSumPerSec float32 `protobuf:"fixed32,36,opt,name=m_no_index_used_sum_per_sec,json=mNoIndexUsedSumPerSec,proto3" json:"m_no_index_used_sum_per_sec,omitempty"`
	// The number of queries without good index.
	MNoGoodIndexUsedSumPerSec float32 `protobuf:"fixed32,37,opt,name=m_no_good_index_used_sum_per_sec,json=mNoGoodIndexUsedSumPerSec,proto3" json:"m_no_good_index_used_sum_per_sec,omitempty"`
	//
	// MongoDB metrics.
	//
	// The number of returned documents.
	MDocsReturnedSumPerSec float32 `protobuf:"fixed32,38,opt,name=m_docs_returned_sum_per_sec,json=mDocsReturnedSumPerSec,proto3" json:"m_docs_returned_sum_per_sec,omitempty"`
	// The response length of the query result in bytes.
	MResponseLengthSumPerSec float32 `protobuf:"fixed32,39,opt,name=m_response_length_sum_per_sec,json=mResponseLengthSumPerSec,proto3" json:"m_response_length_sum_per_sec,omitempty"`
	// The number of scanned documents.
	MDocsScannedSumPerSec float32 `protobuf:"fixed32,40,opt,name=m_docs_scanned_sum_per_sec,json=mDocsScannedSumPerSec,proto3" json:"m_docs_scanned_sum_per_sec,omitempty"`
	//
	// PostgreSQL metrics.
	//
	// Total number of shared block cache hits by the statement.
	MSharedBlksHitSumPerSec float32 `protobuf:"fixed32,41,opt,name=m_shared_blks_hit_sum_per_sec,json=mSharedBlksHitSumPerSec,proto3" json:"m_shared_blks_hit_sum_per_sec,omitempty"`
	// Total number of shared blocks read by the statement.
	MSharedBlksReadSumPerSec float32 `protobuf:"fixed32,42,opt,name=m_shared_blks_read_sum_per_sec,json=mSharedBlksReadSumPerSec,proto3" json:"m_shared_blks_read_sum_per_sec,omitempty"`
	// Total number of shared blocks dirtied by the statement.
	MSharedBlksDirtiedSumPerSec float32 `protobuf:"fixed32,43,opt,name=m_shared_blks_dirtied_sum_per_sec,json=mSharedBlksDirtiedSumPerSec,proto3" json:"m_shared_blks_dirtied_sum_per_sec,omitempty"`
	// Total number of shared blocks written by the statement.
	MSharedBlksWrittenSumPerSec float32 `protobuf:"fixed32,44,opt,name=m_shared_blks_written_sum_per_sec,json=mSharedBlksWrittenSumPerSec,proto3" json:"m_shared_blks_written_sum_per_sec,omitempty"`
	// Total number of local block cache hits by the statement.
	MLocalBlksHitSumPerSec float32 `protobuf:"fixed32,45,opt,name=m_local_blks_hit_sum_per_sec,json=mLocalBlksHitSumPerSec,proto3" json:"m_local_blks_hit_sum_per_sec,omitempty"`
	// Total number of local blocks read by the statement.
	MLocalBlksReadSumPerSec float32 `protobuf:"fixed32,46,opt,name=m_local_blks_read_sum_per_sec,json=mLocalBlksReadSumPerSec,proto3" json:"m_local_blks_read_sum_per_sec,omitempty"`
	// Total number of local blocks dirtied by the statement.
	MLocalBlksDirtiedSumPerSec float32 `protobuf:"fixed32,47,opt,name=m_local_blks_dirtied_sum_per_sec,json=mLocalBlksDirtiedSumPerSec,proto3" json:"m_local_blks_dirtied_sum_per_sec,omitempty"`
	// Total number of local blocks written by the statement.
	MLocalBlksWrittenSumPerSec float32 `protobuf:"fixed32,48,opt,name=m_local_blks_written_sum_per_sec,json=mLocalBlksWrittenSumPerSec,proto3" json:"m_local_blks_written_sum_per_sec,omitempty"`
	// Total number of temp blocks read by the statement.
	MTempBlksReadSumPerSec float32 `protobuf:"fixed32,49,opt,name=m_temp_blks_read_sum_per_sec,json=mTempBlksReadSumPerSec,proto3" json:"m_temp_blks_read_sum_per_sec,omitempty"`
	// Total number of temp blocks written by the statement.
	MTempBlksWrittenSumPerSec float32 `protobuf:"fixed32,50,opt,name=m_temp_blks_written_sum_per_sec,json=mTempBlksWrittenSumPerSec,proto3" json:"m_temp_blks_written_sum_per_sec,omitempty"`
	// Total time the statement spent reading blocks, in milliseconds (if track_io_timing is enabled, otherwise zero).
	MBlkReadTimeSumPerSec float32 `protobuf:"fixed32,51,opt,name=m_blk_read_time_sum_per_sec,json=mBlkReadTimeSumPerSec,proto3" json:"m_blk_read_time_sum_per_sec,omitempty"`
	// Total time the statement spent writing blocks, in milliseconds (if track_io_timing is enabled, otherwise zero).
	MBlkWriteTimeSumPerSec float32  `protobuf:"fixed32,52,opt,name=m_blk_write_time_sum_per_sec,json=mBlkWriteTimeSumPerSec,proto3" json:"m_blk_write_time_sum_per_sec,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc2ea5e264b89be, []int{0}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetPoint() uint32 {
	if m != nil {
		return m.Point
	}
	return 0
}

func (m *Point) GetTimeFrame() uint32 {
	if m != nil {
		return m.TimeFrame
	}
	return 0
}

func (m *Point) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Point) GetNumQueriesPerSec() float32 {
	if m != nil {
		return m.NumQueriesPerSec
	}
	return 0
}

func (m *Point) GetMQueryTimeSumPerSec() float32 {
	if m != nil {
		return m.MQueryTimeSumPerSec
	}
	return 0
}

func (m *Point) GetMLockTimeSumPerSec() float32 {
	if m != nil {
		return m.MLockTimeSumPerSec
	}
	return 0
}

func (m *Point) GetMRowsSentSumPerSec() float32 {
	if m != nil {
		return m.MRowsSentSumPerSec
	}
	return 0
}

func (m *Point) GetMRowsExaminedSumPerSec() float32 {
	if m != nil {
		return m.MRowsExaminedSumPerSec
	}
	return 0
}

func (m *Point) GetMRowsAffectedSumPerSec() float32 {
	if m != nil {
		return m.MRowsAffectedSumPerSec
	}
	return 0
}

func (m *Point) GetMRowsReadSumPerSec() float32 {
	if m != nil {
		return m.MRowsReadSumPerSec
	}
	return 0
}

func (m *Point) GetMMergePassesSumPerSec() float32 {
	if m != nil {
		return m.MMergePassesSumPerSec
	}
	return 0
}

func (m *Point) GetMInnodbIoROpsSumPerSec() float32 {
	if m != nil {
		return m.MInnodbIoROpsSumPerSec
	}
	return 0
}

func (m *Point) GetMInnodbIoRBytesSumPerSec() float32 {
	if m != nil {
		return m.MInnodbIoRBytesSumPerSec
	}
	return 0
}

func (m *Point) GetMInnodbIoRWaitSumPerSec() float32 {
	if m != nil {
		return m.MInnodbIoRWaitSumPerSec
	}
	return 0
}

func (m *Point) GetMInnodbRecLockWaitSumPerSec() float32 {
	if m != nil {
		return m.MInnodbRecLockWaitSumPerSec
	}
	return 0
}

func (m *Point) GetMInnodbQueueWaitSumPerSec() float32 {
	if m != nil {
		return m.MInnodbQueueWaitSumPerSec
	}
	return 0
}

func (m *Point) GetMInnodbPagesDistinctSumPerSec() float32 {
	if m != nil {
		return m.MInnodbPagesDistinctSumPerSec
	}
	return 0
}

func (m *Point) GetMQueryLengthSumPerSec() float32 {
	if m != nil {
		return m.MQueryLengthSumPerSec
	}
	return 0
}

func (m *Point) GetMBytesSentSumPerSec() float32 {
	if m != nil {
		return m.MBytesSentSumPerSec
	}
	return 0
}

func (m *Point) GetMTmpTablesSumPerSec() float32 {
	if m != nil {
		return m.MTmpTablesSumPerSec
	}
	return 0
}

func (m *Point) GetMTmpDiskTablesSumPerSec() float32 {
	if m != nil {
		return m.MTmpDiskTablesSumPerSec
	}
	return 0
}

func (m *Point) GetMTmpTableSizesSumPerSec() float32 {
	if m != nil {
		return m.MTmpTableSizesSumPerSec
	}
	return 0
}

func (m *Point) GetMQcHitSumPerSec() float32 {
	if m != nil {
		return m.MQcHitSumPerSec
	}
	return 0
}

func (m *Point) GetMFullScanSumPerSec() float32 {
	if m != nil {
		return m.MFullScanSumPerSec
	}
	return 0
}

func (m *Point) GetMFullJoinSumPerSec() float32 {
	if m != nil {
		return m.MFullJoinSumPerSec
	}
	return 0
}

func (m *Point) GetMTmpTableSumPerSec() float32 {
	if m != nil {
		return m.MTmpTableSumPerSec
	}
	return 0
}

func (m *Point) GetMTmpTableOnDiskSumPerSec() float32 {
	if m != nil {
		return m.MTmpTableOnDiskSumPerSec
	}
	return 0
}

func (m *Point) GetMFilesortSumPerSec() float32 {
	if m != nil {
		return m.MFilesortSumPerSec
	}
	return 0
}

func (m *Point) GetMFilesortOnDiskSumPerSec() float32 {
	if m != nil {
		return m.MFilesortOnDiskSumPerSec
	}
	return 0
}

func (m *Point) GetMSelectFullRangeJoinSumPerSec() float32 {
	if m != nil {
		return m.MSelectFullRangeJoinSumPerSec
	}
	return 0
}

func (m *Point) GetMSelectRangeSumPerSec() float32 {
	if m != nil {
		return m.MSelectRangeSumPerSec
	}
	return 0
}

func (m *Point) GetMSelectRangeCheckSumPerSec() float32 {
	if m != nil {
		return m.MSelectRangeCheckSumPerSec
	}
	return 0
}

func (m *Point) GetMSortRangeSumPerSec() float32 {
	if m != nil {
		return m.MSortRangeSumPerSec
	}
	return 0
}

func (m *Point) GetMSortRowsSumPerSec() float32 {
	if m != nil {
		return m.MSortRowsSumPerSec
	}
	return 0
}

func (m *Point) GetMSortScanSumPerSec() float32 {
	if m != nil {
		return m.MSortScanSumPerSec
	}
	return 0
}

func (m *Point) GetMNoIndexUsedSumPerSec() float32 {
	if m != nil {
		return m.MNoIndexUsedSumPerSec
	}
	return 0
}

func (m *Point) GetMNoGoodIndexUsedSumPerSec() float32 {
	if m != nil {
		return m.MNoGoodIndexUsedSumPerSec
	}
	return 0
}

func (m *Point) GetMDocsReturnedSumPerSec() float32 {
	if m != nil {
		return m.MDocsReturnedSumPerSec
	}
	return 0
}

func (m *Point) GetMResponseLengthSumPerSec() float32 {
	if m != nil {
		return m.MResponseLengthSumPerSec
	}
	return 0
}

func (m *Point) GetMDocsScannedSumPerSec() float32 {
	if m != nil {
		return m.MDocsScannedSumPerSec
	}
	return 0
}

func (m *Point) GetMSharedBlksHitSumPerSec() float32 {
	if m != nil {
		return m.MSharedBlksHitSumPerSec
	}
	return 0
}

func (m *Point) GetMSharedBlksReadSumPerSec() float32 {
	if m != nil {
		return m.MSharedBlksReadSumPerSec
	}
	return 0
}

func (m *Point) GetMSharedBlksDirtiedSumPerSec() float32 {
	if m != nil {
		return m.MSharedBlksDirtiedSumPerSec
	}
	return 0
}

func (m *Point) GetMSharedBlksWrittenSumPerSec() float32 {
	if m != nil {
		return m.MSharedBlksWrittenSumPerSec
	}
	return 0
}

func (m *Point) GetMLocalBlksHitSumPerSec() float32 {
	if m != nil {
		return m.MLocalBlksHitSumPerSec
	}
	return 0
}

func (m *Point) GetMLocalBlksReadSumPerSec() float32 {
	if m != nil {
		return m.MLocalBlksReadSumPerSec
	}
	return 0
}

func (m *Point) GetMLocalBlksDirtiedSumPerSec() float32 {
	if m != nil {
		return m.MLocalBlksDirtiedSumPerSec
	}
	return 0
}

func (m *Point) GetMLocalBlksWrittenSumPerSec() float32 {
	if m != nil {
		return m.MLocalBlksWrittenSumPerSec
	}
	return 0
}

func (m *Point) GetMTempBlksReadSumPerSec() float32 {
	if m != nil {
		return m.MTempBlksReadSumPerSec
	}
	return 0
}

func (m *Point) GetMTempBlksWrittenSumPerSec() float32 {
	if m != nil {
		return m.MTempBlksWrittenSumPerSec
	}
	return 0
}

func (m *Point) GetMBlkReadTimeSumPerSec() float32 {
	if m != nil {
		return m.MBlkReadTimeSumPerSec
	}
	return 0
}

func (m *Point) GetMBlkWriteTimeSumPerSec() float32 {
	if m != nil {
		return m.MBlkWriteTimeSumPerSec
	}
	return 0
}

func init() {
	proto.RegisterEnum("qan.ExampleFormat", ExampleFormat_name, ExampleFormat_value)
	proto.RegisterEnum("qan.ExampleType", ExampleType_name, ExampleType_value)
	proto.RegisterType((*Point)(nil), "qan.Point")
}

func init() { proto.RegisterFile("qanpb/qan.proto", fileDescriptor_6cc2ea5e264b89be) }

var fileDescriptor_6cc2ea5e264b89be = []byte{
	// 1173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x97, 0x6d, 0x57, 0xdb, 0x36,
	0x14, 0xc7, 0x17, 0x5a, 0x60, 0x88, 0x51, 0x32, 0x43, 0x21, 0xe5, 0xa1, 0xa5, 0xb4, 0xdb, 0x18,
	0x1b, 0xed, 0x56, 0xb6, 0x17, 0x7b, 0x5e, 0xd2, 0x24, 0x34, 0x5d, 0x48, 0x82, 0xed, 0x8d, 0x6d,
	0xe7, 0xec, 0xe8, 0x38, 0x8e, 0x00, 0x2f, 0x96, 0xe5, 0x58, 0xce, 0xa1, 0xec, 0xdb, 0xec, 0x9b,
	0xee, 0xe8, 0xfa, 0x12, 0x5b, 0xb6, 0xd9, 0xbb, 0xd8, 0xba, 0xbf, 0xff, 0x95, 0xee, 0x83, 0xae,
	0x43, 0x56, 0x27, 0x4e, 0x10, 0x0e, 0x5f, 0x4e, 0x9c, 0xe0, 0x45, 0x18, 0x89, 0x58, 0x18, 0xf7,
	0x26, 0x4e, 0xb0, 0xff, 0x6f, 0x8d, 0xcc, 0x0f, 0x84, 0x17, 0xc4, 0xc6, 0x3a, 0x99, 0x0f, 0xd5,
	0x8f, 0x5a, 0x65, 0xaf, 0x72, 0xb0, 0x62, 0x26, 0x0f, 0xc6, 0x2e, 0x21, 0xb1, 0xc7, 0x19, 0xbd,
	0x88, 0x1c, 0xce, 0x6a, 0x73, 0xb0, 0xb4, 0xa4, 0xde, 0xb4, 0xd5, 0x0b, 0x63, 0x87, 0xc0, 0x83,
	0x8c, 0x1d, 0x1e, 0xd6, 0xee, 0xed, 0x55, 0x0e, 0x96, 0xcc, 0xf4, 0x85, 0x71, 0x44, 0xd6, 0x82,
	0x29, 0xa7, 0x93, 0x29, 0x8b, 0x3c, 0x26, 0x69, 0xc8, 0x22, 0x2a, 0x99, 0x5b, 0xbb, 0xbf, 0x57,
	0x39, 0x98, 0x33, 0xab, 0xc1, 0x94, 0x9f, 0x25, 0x2b, 0x03, 0x16, 0x59, 0xcc, 0x35, 0xbe, 0x26,
	0xb5, 0xc4, 0xf8, 0x86, 0x82, 0x4f, 0x39, 0xe5, 0x33, 0x66, 0x1e, 0x98, 0x35, 0x20, 0x6e, 0x6c,
	0x8f, 0x33, 0x6b, 0xca, 0x11, 0x3b, 0x26, 0x9b, 0x9c, 0xfa, 0xc2, 0x1d, 0x17, 0xa9, 0x05, 0xa0,
	0x0c, 0xde, 0x15, 0xee, 0xb8, 0x04, 0x8a, 0xc4, 0xb5, 0xa4, 0x92, 0x05, 0xb1, 0x06, 0x2d, 0x22,
	0x64, 0x8a, 0x6b, 0x69, 0xb1, 0x20, 0x4e, 0xa1, 0xef, 0xc8, 0x36, 0x42, 0xec, 0x9d, 0xc3, 0xbd,
	0x80, 0x8d, 0x34, 0xf0, 0x7d, 0x00, 0x37, 0x00, 0x6c, 0xa1, 0x41, 0x19, 0xec, 0x5c, 0x5c, 0x30,
	0x37, 0xce, 0xc1, 0x4b, 0x19, 0xb8, 0x8e, 0x06, 0x65, 0xdb, 0x8d, 0x98, 0xa3, 0x83, 0x24, 0xb3,
	0x5d, 0x93, 0x39, 0x19, 0xe8, 0x1b, 0xb2, 0xc5, 0x29, 0x67, 0xd1, 0x25, 0xa3, 0xa1, 0x23, 0x25,
	0x93, 0x1a, 0xb7, 0x0c, 0xdc, 0x43, 0x7e, 0xaa, 0x0c, 0x06, 0xb0, 0x9e, 0xa2, 0x3f, 0x90, 0x5d,
	0x4e, 0xbd, 0x20, 0x10, 0xa3, 0x21, 0xf5, 0x04, 0x8d, 0xa8, 0x08, 0x75, 0xfa, 0x03, 0xdc, 0x6e,
	0x07, 0x6c, 0x3a, 0xc2, 0xec, 0x87, 0x19, 0xbc, 0x4e, 0x9e, 0xe8, 0xf8, 0xf0, 0x26, 0xce, 0xb9,
	0x5f, 0x01, 0x81, 0x5a, 0x2a, 0xd0, 0x50, 0x26, 0xa9, 0xc4, 0x4f, 0xe4, 0xb1, 0x2e, 0x71, 0xed,
	0x78, 0x7a, 0x9e, 0x1e, 0x80, 0xc2, 0x66, 0xaa, 0x70, 0xee, 0x78, 0x99, 0x64, 0x9d, 0x90, 0xfd,
	0x99, 0x40, 0xc4, 0xdc, 0xa4, 0x42, 0x0a, 0x22, 0xab, 0x20, 0xb2, 0x8d, 0x22, 0x26, 0x73, 0x55,
	0xa9, 0xe8, 0x42, 0x8d, 0xcc, 0x61, 0x26, 0x53, 0x36, 0x65, 0x45, 0x95, 0x2a, 0xa8, 0x3c, 0x42,
	0x95, 0x33, 0x65, 0xa4, 0x6b, 0xbc, 0x25, 0xcf, 0x66, 0x1a, 0xa1, 0x73, 0xc9, 0x24, 0x1d, 0x79,
	0x32, 0xf6, 0x02, 0x57, 0xd7, 0xf9, 0x10, 0x74, 0x76, 0x51, 0x67, 0xa0, 0x0c, 0x9b, 0x68, 0x97,
	0x4b, 0x6b, 0xd2, 0x26, 0x3e, 0x0b, 0x2e, 0xe3, 0x2b, 0x4d, 0xc2, 0xc0, 0xb4, 0x42, 0xa3, 0x74,
	0x61, 0x3d, 0x45, 0xa1, 0xc3, 0x30, 0x17, 0xf9, 0xb2, 0x5f, 0xc3, 0x0e, 0x4b, 0xf2, 0xa0, 0xd5,
	0x3d, 0x60, 0x31, 0x0f, 0x69, 0xec, 0x0c, 0xfd, 0x5c, 0x1e, 0xd7, 0x11, 0xb3, 0x79, 0x68, 0xc3,
	0x6a, 0x8a, 0xfd, 0xa8, 0x8a, 0x48, 0x61, 0x23, 0x4f, 0x8e, 0xcb, 0xd8, 0x87, 0x98, 0x41, 0x9b,
	0x87, 0x4d, 0x4f, 0x8e, 0xef, 0xe4, 0x01, 0xa5, 0xd2, 0xfb, 0x27, 0xc7, 0x6f, 0xa4, 0x3c, 0xb0,
	0x96, 0xb2, 0x48, 0xf9, 0x23, 0xb2, 0xce, 0xe9, 0xc4, 0xa5, 0x57, 0xb9, 0x6c, 0x6d, 0x02, 0xb6,
	0xca, 0xcf, 0xdc, 0x37, 0xd9, 0x1c, 0x41, 0x8f, 0x5d, 0x4c, 0x7d, 0x9f, 0x4a, 0xd7, 0x09, 0x34,
	0xa2, 0x86, 0x3d, 0xd6, 0x9e, 0xfa, 0xbe, 0xe5, 0x3a, 0x41, 0x19, 0xf4, 0xb7, 0xf0, 0x74, 0xe8,
	0x51, 0x06, 0x7a, 0x2b, 0xbc, 0x3c, 0x94, 0x39, 0x58, 0x06, 0xda, 0x42, 0x68, 0x76, 0x24, 0xbd,
	0xa7, 0x52, 0x48, 0x04, 0x49, 0x58, 0xb3, 0xf0, 0x36, 0xf6, 0xd4, 0x2d, 0xdc, 0x0f, 0x54, 0x58,
	0x53, 0x89, 0x57, 0x64, 0x83, 0xd3, 0x0b, 0xcf, 0x67, 0x52, 0x44, 0x7a, 0x48, 0x76, 0x6e, 0xf7,
	0x8a, 0x8b, 0x29, 0xf3, 0xb3, 0xea, 0xc3, 0x19, 0x53, 0xe6, 0x75, 0x17, 0xbd, 0xde, 0xb2, 0x79,
	0xaf, 0xbf, 0x90, 0xe7, 0x9c, 0x4a, 0xe6, 0x33, 0x37, 0x4e, 0x22, 0x15, 0x39, 0xc1, 0x25, 0x2b,
	0xc6, 0xeb, 0x31, 0x16, 0xbf, 0x05, 0xa6, 0x2a, 0x6a, 0xa6, 0x32, 0xd4, 0x43, 0x07, 0xc5, 0x8f,
	0x62, 0x89, 0x4e, 0x56, 0xe2, 0x09, 0x16, 0x7f, 0x22, 0x01, 0x78, 0x8a, 0x36, 0xc9, 0x5e, 0x0e,
	0x75, 0xaf, 0x98, 0xab, 0x9f, 0x65, 0x0f, 0x04, 0xb6, 0xb2, 0x02, 0xaf, 0x95, 0x51, 0xae, 0x17,
	0x20, 0x16, 0x45, 0xf7, 0x4f, 0xb1, 0x17, 0x2c, 0x11, 0xe5, 0x9d, 0x43, 0xca, 0x13, 0x0c, 0x86,
	0x4e, 0x86, 0xda, 0xc7, 0xd8, 0x03, 0xa5, 0x66, 0x4e, 0x09, 0x54, 0xa8, 0xc8, 0x67, 0x19, 0x48,
	0xaf, 0xc8, 0x6f, 0xd5, 0x9c, 0x09, 0x04, 0xf5, 0x82, 0x11, 0x7b, 0x47, 0xa7, 0x32, 0x37, 0x67,
	0x9e, 0x63, 0x88, 0x7a, 0xa2, 0xa3, 0x0c, 0x7e, 0x95, 0xd9, 0x31, 0xf3, 0x5a, 0x85, 0x28, 0x10,
	0xf4, 0x52, 0x88, 0xd1, 0x5d, 0x02, 0x1f, 0xe1, 0x5d, 0xd7, 0x13, 0x27, 0x42, 0x8c, 0x4a, 0x44,
	0x60, 0xd0, 0x8d, 0x84, 0xab, 0x66, 0x55, 0x3c, 0x8d, 0xf2, 0x53, 0xf2, 0x63, 0x9c, 0x1c, 0x4d,
	0xe1, 0x4a, 0x13, 0x0d, 0xb2, 0xd7, 0xfe, 0x2e, 0xa7, 0x11, 0x93, 0xa1, 0x08, 0x24, 0x2b, 0xbb,
	0xdf, 0x3e, 0xc1, 0x6a, 0x33, 0xd1, 0x26, 0x7f, 0xc5, 0x41, 0x81, 0x80, 0x77, 0x15, 0xb3, 0xbc,
	0xf3, 0x03, 0x3c, 0xbd, 0x72, 0x6e, 0x25, 0xeb, 0xb9, 0xfb, 0x46, 0x5e, 0x39, 0x11, 0x1b, 0xd1,
	0xa1, 0x3f, 0x96, 0x85, 0x8b, 0xe3, 0x53, 0xbc, 0x6f, 0x2c, 0xb0, 0x69, 0xf8, 0x63, 0xa9, 0x5d,
	0x20, 0xd0, 0x2a, 0x59, 0xbe, 0x30, 0xab, 0x0f, 0x71, 0xf3, 0xa9, 0x80, 0x3e, 0xb1, 0xdb, 0xe4,
	0xa9, 0xae, 0x30, 0xf2, 0xa2, 0xd8, 0xcb, 0x9d, 0xe1, 0x33, 0x1c, 0x59, 0xa9, 0x48, 0x33, 0xb1,
	0xfa, 0x1f, 0x9d, 0xeb, 0xc8, 0x8b, 0x63, 0xa6, 0x97, 0xd0, 0xe7, 0x05, 0x9d, 0xf3, 0xc4, 0x2a,
	0xd5, 0xf9, 0x9e, 0xec, 0xc0, 0xa7, 0x95, 0xe3, 0x97, 0x07, 0xe4, 0x08, 0x73, 0xd9, 0x55, 0x26,
	0x85, 0x78, 0x40, 0x3c, 0x33, 0x74, 0x21, 0x1c, 0x2f, 0x30, 0x9e, 0x33, 0x5c, 0x8f, 0x06, 0x34,
	0x6c, 0x86, 0x2f, 0x0b, 0xc6, 0x4b, 0x6c, 0xd8, 0x99, 0x44, 0x21, 0x16, 0x79, 0x95, 0xb2, 0x50,
	0x7c, 0x91, 0x57, 0x29, 0x8f, 0x44, 0xcc, 0x78, 0x78, 0xc7, 0x51, 0xbe, 0xc4, 0x48, 0xd8, 0x8c,
	0x87, 0xc5, 0x93, 0xc0, 0x27, 0x44, 0x4a, 0x97, 0x6d, 0xe1, 0x15, 0xb6, 0xd5, 0xad, 0x40, 0x61,
	0x07, 0xd0, 0xd7, 0x43, 0x7f, 0x9c, 0xf8, 0x2e, 0x7c, 0xea, 0x1e, 0x63, 0x65, 0x37, 0xfc, 0xb1,
	0xf2, 0xad, 0x7f, 0xed, 0xc2, 0xee, 0x15, 0xab, 0x3c, 0xb3, 0x22, 0xfc, 0x15, 0xee, 0xbe, 0xe1,
	0x8f, 0x95, 0x5f, 0xa6, 0xd1, 0x87, 0x1d, 0xb2, 0xa2, 0x3e, 0x67, 0x43, 0x9f, 0xb5, 0x45, 0xc4,
	0x9d, 0xd8, 0xd8, 0x22, 0x1b, 0xad, 0xdf, 0xeb, 0xa7, 0x83, 0x6e, 0x8b, 0xb6, 0xfb, 0xe6, 0x69,
	0xdd, 0xa6, 0x9d, 0xde, 0x6f, 0xf5, 0x6e, 0xa7, 0x59, 0x7d, 0xcf, 0x58, 0x26, 0x8b, 0xb8, 0x56,
	0xad, 0x18, 0xab, 0x64, 0xb9, 0xdd, 0xe9, 0x9d, 0xb4, 0xcc, 0x81, 0xd9, 0xe9, 0xd9, 0xd5, 0xb9,
	0xc3, 0xbf, 0xc8, 0x32, 0x4a, 0xd9, 0x37, 0x21, 0x33, 0x6a, 0x64, 0xfd, 0x56, 0xc8, 0xfe, 0x63,
	0xd0, 0xca, 0xc8, 0x10, 0xb2, 0x60, 0xd6, 0x7b, 0xcd, 0xfe, 0x69, 0xb5, 0xa2, 0x24, 0xad, 0x6e,
	0xff, 0xbc, 0x65, 0xd9, 0xd5, 0x39, 0xf5, 0xd0, 0xae, 0x5b, 0xb6, 0x7a, 0xb8, 0x67, 0x3c, 0x20,
	0xe4, 0xbc, 0x63, 0xbf, 0xa1, 0x2d, 0xd3, 0xec, 0x9b, 0xd5, 0xfb, 0x8d, 0xc5, 0x3f, 0xe7, 0xe1,
	0x5f, 0xce, 0x70, 0x01, 0xfe, 0xe2, 0x1c, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0xed, 0x1e, 0xd6,
	0x0e, 0xf5, 0x0c, 0x00, 0x00,
}
