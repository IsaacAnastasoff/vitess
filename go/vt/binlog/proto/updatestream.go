// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	mproto "github.com/youtube/vitess/go/mysql/proto"
	"github.com/youtube/vitess/go/vt/key"
)

// UpdateStreamRequest is used to make a request for ServeUpdateStream.
type UpdateStreamRequest struct {
	Position string
}

// KeyRangeRequest is used to make a request for StreamKeyRange.
type KeyRangeRequest struct {
	Position       string
	KeyspaceIdType key.KeyspaceIdType
	KeyRange       key.KeyRange
	Charset        *mproto.Charset
}

// TablesRequest is used to make a request for StreamTables.
type TablesRequest struct {
	Position string
	Tables   []string
	Charset  *mproto.Charset
}

// UpdateStream is the interface for the server
type UpdateStream interface {
	// ServeUpdateStream serves the query and streams the result
	// for the full update stream
	ServeUpdateStream(req *UpdateStreamRequest, sendReply func(reply *StreamEvent) error) error

	// StreamKeyRange streams events related to a KeyRange only
	StreamKeyRange(req *KeyRangeRequest, sendReply func(reply *BinlogTransaction) error) error

	// StreamTables streams events related to a set of Tables only
	StreamTables(req *TablesRequest, sendReply func(reply *BinlogTransaction) error) error

	// HandlePanic should be called in a defer,
	// first thing in the RPC implementation.
	HandlePanic(*error)
}
