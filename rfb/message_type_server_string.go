// Code generated by "stringer -type=ServerMessageType"; DO NOT EDIT.

package rfb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FramebufferUpdate-0]
	_ = x[SetColorMapEntries-1]
	_ = x[Bell-2]
	_ = x[ServerCutText-3]
}

const _ServerMessageType_name = "FramebufferUpdateSetColorMapEntriesBellServerCutText"

var _ServerMessageType_index = [...]uint8{0, 17, 35, 39, 52}

func (i ServerMessageType) String() string {
	if i >= ServerMessageType(len(_ServerMessageType_index)-1) {
		return "ServerMessageType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ServerMessageType_name[_ServerMessageType_index[i]:_ServerMessageType_index[i+1]]
}
