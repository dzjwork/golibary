package yaml

import (
	"fmt"
	"io"
)

// 表示版本信息
type yaml_version_directive_t struct {
	major int8 // 主版本
	minor int8 // 次版本
}

// 表示标签信息
type yaml_tag_directive_t struct {
	handle []byte // 标签处理器
	prefix []byte // 标签前缀
}

// YAML编码格式
type yaml_encoding_t int

// The stream encoding.
const (
	// 解析器选择编码格式
	yaml_ANY_ENCODING yaml_encoding_t = iota

	yaml_UTF8_ENCODING    // UTF-8编码
	yaml_UTF16LE_ENCODING // UTF-16-LE编码
	yaml_UTF16BE_ENCODING // UTF-16-BE编码
)

// YAML中换行符
type yaml_break_t int

const (
	// 解析器选择换行类型
	yaml_ANY_BREAK yaml_break_t = iota

	yaml_CR_BREAK   // MAC的换行类型为CR (Mac style).
	yaml_LN_BREAK   // UNIX的换行类型为LN
	yaml_CRLN_BREAK // DOS的换行类型为CRLN
)

// 解析期间的错误类型
type yaml_error_type_t int

const (
	// 没有产生错误
	yaml_NO_ERROR yaml_error_type_t = iota

	yaml_MEMORY_ERROR   // 内存不足异常
	yaml_READER_ERROR   // 无法读取或解码输入流
	yaml_SCANNER_ERROR  // 无法扫描输入流
	yaml_PARSER_ERROR   // 无法解析输入流
	yaml_COMPOSER_ERROR // 无法编写YAML文档
	yaml_WRITER_ERROR   // 无法写出数据
	yaml_EMITTER_ERROR  // 无法触发YAML流
)

// 解析YAML时指针的位置
type yaml_mark_t struct {
	index  int // 位置索引
	line   int // 解析的行
	column int // 解析的列
}

// 表示YAML样式
type yaml_style_t int8

type yaml_scalar_style_t yaml_style_t

// Scalar styles.
const (
	// Let the emitter choose the style.
	yaml_ANY_SCALAR_STYLE yaml_scalar_style_t = 0

	yaml_PLAIN_SCALAR_STYLE         yaml_scalar_style_t = 1 << iota // The plain scalar style.
	yaml_SINGLE_QUOTED_SCALAR_STYLE                                 // The single-quoted scalar style.
	yaml_DOUBLE_QUOTED_SCALAR_STYLE                                 // The double-quoted scalar style.
	yaml_LITERAL_SCALAR_STYLE                                       // The literal scalar style.
	yaml_FOLDED_SCALAR_STYLE                                        // The folded scalar style.
)

// 表示YAML中序列样式
type yaml_sequence_style_t yaml_style_t

const (
	// Let the emitter choose the style.
	yaml_ANY_SEQUENCE_STYLE yaml_sequence_style_t = iota

	yaml_BLOCK_SEQUENCE_STYLE // The block sequence style.
	yaml_FLOW_SEQUENCE_STYLE  // The flow sequence style.
)

// 表示YAML中Map样式
type yaml_mapping_style_t yaml_style_t

const (
	// Let the emitter choose the style.
	yaml_ANY_MAPPING_STYLE yaml_mapping_style_t = iota

	yaml_BLOCK_MAPPING_STYLE // The block mapping style.
	yaml_FLOW_MAPPING_STYLE  // The flow mapping style.
)

// 表示token类型
type yaml_token_type_t int

// Token types.
const (
	// An empty token.
	yaml_NO_TOKEN yaml_token_type_t = iota

	yaml_STREAM_START_TOKEN // A STREAM-START token.
	yaml_STREAM_END_TOKEN   // A STREAM-END token.

	yaml_VERSION_DIRECTIVE_TOKEN // A VERSION-DIRECTIVE token.
	yaml_TAG_DIRECTIVE_TOKEN     // A TAG-DIRECTIVE token.
	yaml_DOCUMENT_START_TOKEN    // A DOCUMENT-START token.
	yaml_DOCUMENT_END_TOKEN      // A DOCUMENT-END token.

	yaml_BLOCK_SEQUENCE_START_TOKEN // A BLOCK-SEQUENCE-START token.
	yaml_BLOCK_MAPPING_START_TOKEN  // A BLOCK-SEQUENCE-END token.
	yaml_BLOCK_END_TOKEN            // A BLOCK-END token.

	yaml_FLOW_SEQUENCE_START_TOKEN // A FLOW-SEQUENCE-START token.
	yaml_FLOW_SEQUENCE_END_TOKEN   // A FLOW-SEQUENCE-END token.
	yaml_FLOW_MAPPING_START_TOKEN  // A FLOW-MAPPING-START token.
	yaml_FLOW_MAPPING_END_TOKEN    // A FLOW-MAPPING-END token.

	yaml_BLOCK_ENTRY_TOKEN // A BLOCK-ENTRY token.
	yaml_FLOW_ENTRY_TOKEN  // A FLOW-ENTRY token.
	yaml_KEY_TOKEN         // A KEY token.
	yaml_VALUE_TOKEN       // A VALUE token.

	yaml_ALIAS_TOKEN  // An ALIAS token.
	yaml_ANCHOR_TOKEN // An ANCHOR token.
	yaml_TAG_TOKEN    // A TAG token.
	yaml_SCALAR_TOKEN // A SCALAR token.
)

// 根据token码获取对应的名称
func (tt yaml_token_type_t) String() string {
	switch tt {
	case yaml_NO_TOKEN:
		return "yaml_NO_TOKEN"
	case yaml_STREAM_START_TOKEN:
		return "yaml_STREAM_START_TOKEN"
	case yaml_STREAM_END_TOKEN:
		return "yaml_STREAM_END_TOKEN"
	case yaml_VERSION_DIRECTIVE_TOKEN:
		return "yaml_VERSION_DIRECTIVE_TOKEN"
	case yaml_TAG_DIRECTIVE_TOKEN:
		return "yaml_TAG_DIRECTIVE_TOKEN"
	case yaml_DOCUMENT_START_TOKEN:
		return "yaml_DOCUMENT_START_TOKEN"
	case yaml_DOCUMENT_END_TOKEN:
		return "yaml_DOCUMENT_END_TOKEN"
	case yaml_BLOCK_SEQUENCE_START_TOKEN:
		return "yaml_BLOCK_SEQUENCE_START_TOKEN"
	case yaml_BLOCK_MAPPING_START_TOKEN:
		return "yaml_BLOCK_MAPPING_START_TOKEN"
	case yaml_BLOCK_END_TOKEN:
		return "yaml_BLOCK_END_TOKEN"
	case yaml_FLOW_SEQUENCE_START_TOKEN:
		return "yaml_FLOW_SEQUENCE_START_TOKEN"
	case yaml_FLOW_SEQUENCE_END_TOKEN:
		return "yaml_FLOW_SEQUENCE_END_TOKEN"
	case yaml_FLOW_MAPPING_START_TOKEN:
		return "yaml_FLOW_MAPPING_START_TOKEN"
	case yaml_FLOW_MAPPING_END_TOKEN:
		return "yaml_FLOW_MAPPING_END_TOKEN"
	case yaml_BLOCK_ENTRY_TOKEN:
		return "yaml_BLOCK_ENTRY_TOKEN"
	case yaml_FLOW_ENTRY_TOKEN:
		return "yaml_FLOW_ENTRY_TOKEN"
	case yaml_KEY_TOKEN:
		return "yaml_KEY_TOKEN"
	case yaml_VALUE_TOKEN:
		return "yaml_VALUE_TOKEN"
	case yaml_ALIAS_TOKEN:
		return "yaml_ALIAS_TOKEN"
	case yaml_ANCHOR_TOKEN:
		return "yaml_ANCHOR_TOKEN"
	case yaml_TAG_TOKEN:
		return "yaml_TAG_TOKEN"
	case yaml_SCALAR_TOKEN:
		return "yaml_SCALAR_TOKEN"
	}
	return "<unknown token>"
}

// The token structure.
type yaml_token_t struct {
	// token类型
	typ yaml_token_type_t

	// token的起始和结束位置
	start_mark, end_mark yaml_mark_t

	// 流编码格式(for yaml_STREAM_START_TOKEN).
	encoding yaml_encoding_t

	// The alias/anchor/scalar value or tag/tag directive handle
	// (for yaml_ALIAS_TOKEN, yaml_ANCHOR_TOKEN, yaml_SCALAR_TOKEN, yaml_TAG_TOKEN, yaml_TAG_DIRECTIVE_TOKEN).
	value []byte

	// The tag suffix (for yaml_TAG_TOKEN).
	suffix []byte

	// The tag directive prefix (for yaml_TAG_DIRECTIVE_TOKEN).
	prefix []byte

	// The scalar style (for yaml_SCALAR_TOKEN).
	style yaml_scalar_style_t

	// The version directive major/minor (for yaml_VERSION_DIRECTIVE_TOKEN).
	major, minor int8
}

// 事件类型
type yaml_event_type_t int8

const (
	// 没有事件
	yaml_NO_EVENT yaml_event_type_t = iota

	yaml_STREAM_START_EVENT   // 流开始事件
	yaml_STREAM_END_EVENT     // 流结束事件
	yaml_DOCUMENT_START_EVENT // 文档开始事件
	yaml_DOCUMENT_END_EVENT   // 文档结束事件
	yaml_ALIAS_EVENT          // 别名事件
	yaml_SCALAR_EVENT         // A SCALAR event.
	yaml_SEQUENCE_START_EVENT // 序列开始事件
	yaml_SEQUENCE_END_EVENT   // 序列结束事件
	yaml_MAPPING_START_EVENT  // Map开始事件
	yaml_MAPPING_END_EVENT    // Map结束事件
	yaml_TAIL_COMMENT_EVENT
)

// 所有事件类型对应的名称
var eventStrings = []string{
	yaml_NO_EVENT:             "none",
	yaml_STREAM_START_EVENT:   "stream start",
	yaml_STREAM_END_EVENT:     "stream end",
	yaml_DOCUMENT_START_EVENT: "document start",
	yaml_DOCUMENT_END_EVENT:   "document end",
	yaml_ALIAS_EVENT:          "alias",
	yaml_SCALAR_EVENT:         "scalar",
	yaml_SEQUENCE_START_EVENT: "sequence start",
	yaml_SEQUENCE_END_EVENT:   "sequence end",
	yaml_MAPPING_START_EVENT:  "mapping start",
	yaml_MAPPING_END_EVENT:    "mapping end",
	yaml_TAIL_COMMENT_EVENT:   "tail comment",
}

// 根据事件类型获取其对应的名称
func (e yaml_event_type_t) String() string {
	if e < 0 || int(e) >= len(eventStrings) {
		return fmt.Sprintf("unknown event %d", e)
	}
	return eventStrings[e]
}

// YAML事件
type yaml_event_t struct {

	// 事件类型
	typ yaml_event_type_t

	// 事件的开始和结束位置
	start_mark, end_mark yaml_mark_t

	// 文档编码格式(for yaml_STREAM_START_EVENT).
	encoding yaml_encoding_t

	// 文档版本信息 (for yaml_DOCUMENT_START_EVENT).
	version_directive *yaml_version_directive_t

	// 标签信息(for yaml_DOCUMENT_START_EVENT).
	tag_directives []yaml_tag_directive_t

	// The comments
	head_comment []byte
	line_comment []byte
	foot_comment []byte
	tail_comment []byte

	// The anchor (for yaml_SCALAR_EVENT, yaml_SEQUENCE_START_EVENT, yaml_MAPPING_START_EVENT, yaml_ALIAS_EVENT).
	anchor []byte

	// The tag (for yaml_SCALAR_EVENT, yaml_SEQUENCE_START_EVENT, yaml_MAPPING_START_EVENT).
	tag []byte

	// The scalar value (for yaml_SCALAR_EVENT).
	value []byte

	// Is the document start/end indicator implicit, or the tag optional?
	// (for yaml_DOCUMENT_START_EVENT, yaml_DOCUMENT_END_EVENT, yaml_SEQUENCE_START_EVENT, yaml_MAPPING_START_EVENT, yaml_SCALAR_EVENT).
	implicit bool

	// Is the tag optional for any non-plain style? (for yaml_SCALAR_EVENT).
	quoted_implicit bool

	// The style (for yaml_SCALAR_EVENT, yaml_SEQUENCE_START_EVENT, yaml_MAPPING_START_EVENT).
	style yaml_style_t
}

func (e *yaml_event_t) scalar_style() yaml_scalar_style_t     { return yaml_scalar_style_t(e.style) }
func (e *yaml_event_t) sequence_style() yaml_sequence_style_t { return yaml_sequence_style_t(e.style) }
func (e *yaml_event_t) mapping_style() yaml_mapping_style_t   { return yaml_mapping_style_t(e.style) }

const (
	yaml_NULL_TAG      = "tag:yaml.org,2002:null"      // !!nnull标签，对应值为null
	yaml_BOOL_TAG      = "tag:yaml.org,2002:bool"      // !!bool标签，对应值为true或false
	yaml_STR_TAG       = "tag:yaml.org,2002:str"       // !!str标签，对应值为字符串
	yaml_INT_TAG       = "tag:yaml.org,2002:int"       // !!int标签，对应值为整数
	yaml_FLOAT_TAG     = "tag:yaml.org,2002:float"     // !!float标签，对应值为浮点数
	yaml_TIMESTAMP_TAG = "tag:yaml.org,2002:timestamp" // !!timestamp标签，对应值为时间戳

	yaml_SEQ_TAG = "tag:yaml.org,2002:seq" // !!seq标签，对应值为数组
	yaml_MAP_TAG = "tag:yaml.org,2002:map" // !!map标签，对应值为Map

	yaml_BINARY_TAG = "tag:yaml.org,2002:binary"
	yaml_MERGE_TAG  = "tag:yaml.org,2002:merge"

	yaml_DEFAULT_SCALAR_TAG   = yaml_STR_TAG // The default scalar tag is !!str.
	yaml_DEFAULT_SEQUENCE_TAG = yaml_SEQ_TAG // The default sequence tag is !!seq.
	yaml_DEFAULT_MAPPING_TAG  = yaml_MAP_TAG // The default mapping tag is !!map.
)

// YAML中节点的类型
type yaml_node_type_t int

const (
	// 空节点
	yaml_NO_NODE yaml_node_type_t = iota

	yaml_SCALAR_NODE   // A scalar node.
	yaml_SEQUENCE_NODE // 序列节点
	yaml_MAPPING_NODE  // Map节点
)

// 表示序列节点的元素
type yaml_node_item_t int

// 表示Map节点的元素
type yaml_node_pair_t struct {
	key   int
	value int
}

// 节点结构
type yaml_node_t struct {
	typ yaml_node_type_t // 节点类型
	tag []byte           // 节点名称

	// The scalar parameters (for yaml_SCALAR_NODE).
	scalar struct {
		value  []byte              // The scalar value.
		length int                 // The length of the scalar value.
		style  yaml_scalar_style_t // The scalar style.
	}

	// The sequence parameters (for YAML_SEQUENCE_NODE).
	sequence struct {
		items_data []yaml_node_item_t    // The stack of sequence items.
		style      yaml_sequence_style_t // The sequence style.
	}

	// The mapping parameters (for yaml_MAPPING_NODE).
	mapping struct {
		pairs_data  []yaml_node_pair_t   // The stack of mapping pairs (key, value).
		pairs_start *yaml_node_pair_t    // The beginning of the stack.
		pairs_end   *yaml_node_pair_t    // The end of the stack.
		pairs_top   *yaml_node_pair_t    // The top of the stack.
		style       yaml_mapping_style_t // The mapping style.
	}

	start_mark yaml_mark_t // 节点的开始位置
	end_mark   yaml_mark_t // 节点的结束位置

}

// 表示YAML文档的结构
type yaml_document_t struct {

	// YAML中所有节点
	nodes []yaml_node_t

	// 版本信息
	version_directive *yaml_version_directive_t

	// 标签列表
	tag_directives_data  []yaml_tag_directive_t
	tag_directives_start int // 标签列表的开头
	tag_directives_end   int // 标签列表的结尾

	start_implicit int // Is the document start indicator implicit?
	end_implicit   int // Is the document end indicator implicit?

	// 文档开始、结束位置
	start_mark, end_mark yaml_mark_t
}

// The prototype of a read handler.
//
// The read handler is called when the parser needs to read more bytes from the
// source. The handler should write not more than size bytes to the buffer.
// The number of written bytes should be set to the size_read variable.
//
// [in,out]   data        A pointer to an application data specified by
//
//	yaml_parser_set_input().
//
// [out]      buffer      The buffer to write the data from the source.
// [in]       size        The size of the buffer.
// [out]      size_read   The actual number of bytes read from the source.
//
// On success, the handler should return 1.  If the handler failed,
// the returned value should be 0. On EOF, the handler should set the
// size_read to 0 and return 1.
type yaml_read_handler_t func(parser *yaml_parser_t, buffer []byte) (n int, err error)

// This structure holds information about a potential simple key.
type yaml_simple_key_t struct {
	possible     bool        // Is a simple key possible?
	required     bool        // Is a simple key required?
	token_number int         // The number of the token.
	mark         yaml_mark_t // The position mark.
}

// 解析器的状态
type yaml_parser_state_t int

const (
	yaml_PARSE_STREAM_START_STATE yaml_parser_state_t = iota

	yaml_PARSE_IMPLICIT_DOCUMENT_START_STATE           // Expect the beginning of an implicit document.
	yaml_PARSE_DOCUMENT_START_STATE                    // Expect DOCUMENT-START.
	yaml_PARSE_DOCUMENT_CONTENT_STATE                  // Expect the content of a document.
	yaml_PARSE_DOCUMENT_END_STATE                      // Expect DOCUMENT-END.
	yaml_PARSE_BLOCK_NODE_STATE                        // Expect a block node.
	yaml_PARSE_BLOCK_NODE_OR_INDENTLESS_SEQUENCE_STATE // Expect a block node or indentless sequence.
	yaml_PARSE_FLOW_NODE_STATE                         // Expect a flow node.
	yaml_PARSE_BLOCK_SEQUENCE_FIRST_ENTRY_STATE        // Expect the first entry of a block sequence.
	yaml_PARSE_BLOCK_SEQUENCE_ENTRY_STATE              // Expect an entry of a block sequence.
	yaml_PARSE_INDENTLESS_SEQUENCE_ENTRY_STATE         // Expect an entry of an indentless sequence.
	yaml_PARSE_BLOCK_MAPPING_FIRST_KEY_STATE           // Expect the first key of a block mapping.
	yaml_PARSE_BLOCK_MAPPING_KEY_STATE                 // Expect a block mapping key.
	yaml_PARSE_BLOCK_MAPPING_VALUE_STATE               // Expect a block mapping value.
	yaml_PARSE_FLOW_SEQUENCE_FIRST_ENTRY_STATE         // Expect the first entry of a flow sequence.
	yaml_PARSE_FLOW_SEQUENCE_ENTRY_STATE               // Expect an entry of a flow sequence.
	yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_KEY_STATE   // Expect a key of an ordered mapping.
	yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_VALUE_STATE // Expect a value of an ordered mapping.
	yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_END_STATE   // Expect the and of an ordered mapping entry.
	yaml_PARSE_FLOW_MAPPING_FIRST_KEY_STATE            // Expect the first key of a flow mapping.
	yaml_PARSE_FLOW_MAPPING_KEY_STATE                  // Expect a key of a flow mapping.
	yaml_PARSE_FLOW_MAPPING_VALUE_STATE                // Expect a value of a flow mapping.
	yaml_PARSE_FLOW_MAPPING_EMPTY_VALUE_STATE          // Expect an empty value of a flow mapping.
	yaml_PARSE_END_STATE                               // Expect nothing.
)

// 根据解析器的状态获取对应状态名称
func (ps yaml_parser_state_t) String() string {
	switch ps {
	case yaml_PARSE_STREAM_START_STATE:
		return "yaml_PARSE_STREAM_START_STATE"
	case yaml_PARSE_IMPLICIT_DOCUMENT_START_STATE:
		return "yaml_PARSE_IMPLICIT_DOCUMENT_START_STATE"
	case yaml_PARSE_DOCUMENT_START_STATE:
		return "yaml_PARSE_DOCUMENT_START_STATE"
	case yaml_PARSE_DOCUMENT_CONTENT_STATE:
		return "yaml_PARSE_DOCUMENT_CONTENT_STATE"
	case yaml_PARSE_DOCUMENT_END_STATE:
		return "yaml_PARSE_DOCUMENT_END_STATE"
	case yaml_PARSE_BLOCK_NODE_STATE:
		return "yaml_PARSE_BLOCK_NODE_STATE"
	case yaml_PARSE_BLOCK_NODE_OR_INDENTLESS_SEQUENCE_STATE:
		return "yaml_PARSE_BLOCK_NODE_OR_INDENTLESS_SEQUENCE_STATE"
	case yaml_PARSE_FLOW_NODE_STATE:
		return "yaml_PARSE_FLOW_NODE_STATE"
	case yaml_PARSE_BLOCK_SEQUENCE_FIRST_ENTRY_STATE:
		return "yaml_PARSE_BLOCK_SEQUENCE_FIRST_ENTRY_STATE"
	case yaml_PARSE_BLOCK_SEQUENCE_ENTRY_STATE:
		return "yaml_PARSE_BLOCK_SEQUENCE_ENTRY_STATE"
	case yaml_PARSE_INDENTLESS_SEQUENCE_ENTRY_STATE:
		return "yaml_PARSE_INDENTLESS_SEQUENCE_ENTRY_STATE"
	case yaml_PARSE_BLOCK_MAPPING_FIRST_KEY_STATE:
		return "yaml_PARSE_BLOCK_MAPPING_FIRST_KEY_STATE"
	case yaml_PARSE_BLOCK_MAPPING_KEY_STATE:
		return "yaml_PARSE_BLOCK_MAPPING_KEY_STATE"
	case yaml_PARSE_BLOCK_MAPPING_VALUE_STATE:
		return "yaml_PARSE_BLOCK_MAPPING_VALUE_STATE"
	case yaml_PARSE_FLOW_SEQUENCE_FIRST_ENTRY_STATE:
		return "yaml_PARSE_FLOW_SEQUENCE_FIRST_ENTRY_STATE"
	case yaml_PARSE_FLOW_SEQUENCE_ENTRY_STATE:
		return "yaml_PARSE_FLOW_SEQUENCE_ENTRY_STATE"
	case yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_KEY_STATE:
		return "yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_KEY_STATE"
	case yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_VALUE_STATE:
		return "yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_VALUE_STATE"
	case yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_END_STATE:
		return "yaml_PARSE_FLOW_SEQUENCE_ENTRY_MAPPING_END_STATE"
	case yaml_PARSE_FLOW_MAPPING_FIRST_KEY_STATE:
		return "yaml_PARSE_FLOW_MAPPING_FIRST_KEY_STATE"
	case yaml_PARSE_FLOW_MAPPING_KEY_STATE:
		return "yaml_PARSE_FLOW_MAPPING_KEY_STATE"
	case yaml_PARSE_FLOW_MAPPING_VALUE_STATE:
		return "yaml_PARSE_FLOW_MAPPING_VALUE_STATE"
	case yaml_PARSE_FLOW_MAPPING_EMPTY_VALUE_STATE:
		return "yaml_PARSE_FLOW_MAPPING_EMPTY_VALUE_STATE"
	case yaml_PARSE_END_STATE:
		return "yaml_PARSE_END_STATE"
	}
	return "<unknown parser state>"
}

// 表示别名信息
type yaml_alias_data_t struct {
	anchor []byte      // 锚点
	index  int         // The node id.
	mark   yaml_mark_t // The anchor mark.
}

// The parser structure.
//
// All members are internal. Manage the structure using the
// yaml_parser_ family of functions.
type yaml_parser_t struct {

	// Error handling

	error   yaml_error_type_t // 错误类型
	problem string            // 错误信息

	// The byte about which the problem occurred.
	problem_offset int
	problem_value  int
	problem_mark   yaml_mark_t

	// The error context.
	context      string
	context_mark yaml_mark_t

	// Reader stuff

	read_handler yaml_read_handler_t // Read handler.

	input_reader io.Reader // File input data.
	input        []byte    // 要解析的数据
	input_pos    int       // 解析到的位置

	eof bool // 是否解析到了结尾处

	buffer     []byte // 解析好的缓冲区
	buffer_pos int    // 下次buffer可读/写的位置

	unread int // The number of unread characters in the buffer.

	newlines int // The number of line breaks since last non-break/non-blank character

	raw_buffer     []byte // 二进制缓冲区
	raw_buffer_pos int    // 下次raw_buffer可读/写的位置

	encoding yaml_encoding_t // 输入源的编码格式

	offset int         // The offset of the current position (in bytes).
	mark   yaml_mark_t // The mark of the current position.

	// Comments

	head_comment []byte // The current head comments
	line_comment []byte // The current line comments
	foot_comment []byte // The current foot comments
	tail_comment []byte // Foot comment that happens at the end of a block.
	stem_comment []byte // Comment in item preceding a nested structure (list inside list item, etc)

	comments      []yaml_comment_t // The folded comments for all parsed tokens
	comments_head int

	// Scanner stuff

	stream_start_produced bool // Have we started to scan the input stream?
	stream_end_produced   bool // 是否到达输入流的末尾

	flow_level int // The number of unclosed '[' and '{' indicators.

	tokens          []yaml_token_t // The tokens queue.
	tokens_head     int            // The head of the tokens queue.
	tokens_parsed   int            // The number of tokens fetched from the queue.
	token_available bool           // Does the tokens queue contain a token ready for dequeueing.

	indent  int   // The current indentation level.
	indents []int // The indentation levels stack.

	simple_key_allowed bool                // May a simple key occur at the current position?
	simple_keys        []yaml_simple_key_t // The stack of simple keys.
	simple_keys_by_tok map[int]int         // possible simple_key indexes indexed by token_number

	// Parser stuff

	state          yaml_parser_state_t    // 当前解析器的状态
	states         []yaml_parser_state_t  // The parser states stack.
	marks          []yaml_mark_t          // The stack of marks.
	tag_directives []yaml_tag_directive_t // The list of TAG directives.

	// Dumper stuff

	aliases []yaml_alias_data_t // The alias data.

	document *yaml_document_t // The currently parsed document.
}

type yaml_comment_t struct {
	scan_mark  yaml_mark_t // Position where scanning for comments started
	token_mark yaml_mark_t // Position after which tokens will be associated with this comment
	start_mark yaml_mark_t // Position of '#' comment mark
	end_mark   yaml_mark_t // Position where comment terminated

	head []byte
	line []byte
	foot []byte
}

// Emitter Definitions

// The prototype of a write handler.
//
// The write handler is called when the emitter needs to flush the accumulated
// characters to the output.  The handler should write @a size bytes of the
// @a buffer to the output.
//
// @param[in,out]   data        A pointer to an application data specified by
//
//	yaml_emitter_set_output().
//
// @param[in]       buffer      The buffer with bytes to be written.
// @param[in]       size        The size of the buffer.
//
// @returns On success, the handler should return @c 1.  If the handler failed,
// the returned value should be @c 0.
type yaml_write_handler_t func(emitter *yaml_emitter_t, buffer []byte) error

type yaml_emitter_state_t int

// The emitter states.
const (
	// Expect STREAM-START.
	yaml_EMIT_STREAM_START_STATE yaml_emitter_state_t = iota

	yaml_EMIT_FIRST_DOCUMENT_START_STATE       // Expect the first DOCUMENT-START or STREAM-END.
	yaml_EMIT_DOCUMENT_START_STATE             // Expect DOCUMENT-START or STREAM-END.
	yaml_EMIT_DOCUMENT_CONTENT_STATE           // Expect the content of a document.
	yaml_EMIT_DOCUMENT_END_STATE               // Expect DOCUMENT-END.
	yaml_EMIT_FLOW_SEQUENCE_FIRST_ITEM_STATE   // Expect the first item of a flow sequence.
	yaml_EMIT_FLOW_SEQUENCE_TRAIL_ITEM_STATE   // Expect the next item of a flow sequence, with the comma already written out
	yaml_EMIT_FLOW_SEQUENCE_ITEM_STATE         // Expect an item of a flow sequence.
	yaml_EMIT_FLOW_MAPPING_FIRST_KEY_STATE     // Expect the first key of a flow mapping.
	yaml_EMIT_FLOW_MAPPING_TRAIL_KEY_STATE     // Expect the next key of a flow mapping, with the comma already written out
	yaml_EMIT_FLOW_MAPPING_KEY_STATE           // Expect a key of a flow mapping.
	yaml_EMIT_FLOW_MAPPING_SIMPLE_VALUE_STATE  // Expect a value for a simple key of a flow mapping.
	yaml_EMIT_FLOW_MAPPING_VALUE_STATE         // Expect a value of a flow mapping.
	yaml_EMIT_BLOCK_SEQUENCE_FIRST_ITEM_STATE  // Expect the first item of a block sequence.
	yaml_EMIT_BLOCK_SEQUENCE_ITEM_STATE        // Expect an item of a block sequence.
	yaml_EMIT_BLOCK_MAPPING_FIRST_KEY_STATE    // Expect the first key of a block mapping.
	yaml_EMIT_BLOCK_MAPPING_KEY_STATE          // Expect the key of a block mapping.
	yaml_EMIT_BLOCK_MAPPING_SIMPLE_VALUE_STATE // Expect a value for a simple key of a block mapping.
	yaml_EMIT_BLOCK_MAPPING_VALUE_STATE        // Expect a value of a block mapping.
	yaml_EMIT_END_STATE                        // Expect nothing.
)

// The emitter structure.
//
// All members are internal.  Manage the structure using the @c yaml_emitter_
// family of functions.
type yaml_emitter_t struct {
	error   yaml_error_type_t // 错误类型
	problem string            // 错误描述

	write_handler yaml_write_handler_t // Write handler.

	output_buffer *[]byte   // String output data.
	output_writer io.Writer // File output data.

	buffer     []byte // The working buffer.
	buffer_pos int    // The current position of the buffer.

	raw_buffer     []byte // The raw buffer.
	raw_buffer_pos int    // The current position of the buffer.

	encoding yaml_encoding_t // The stream encoding.

	// Emitter stuff

	canonical   bool         // If the output is in the canonical style?
	best_indent int          // The number of indentation spaces.
	best_width  int          // The preferred width of the output lines.
	unicode     bool         // Allow unescaped non-ASCII characters?
	line_break  yaml_break_t // The preferred line break.

	state  yaml_emitter_state_t   // The current emitter state.
	states []yaml_emitter_state_t // The stack of states.

	events      []yaml_event_t // The event queue.
	events_head int            // The head of the event queue.

	indents []int // The stack of indentation levels.

	tag_directives []yaml_tag_directive_t // The list of tag directives.

	indent int // The current indentation level.

	flow_level int // The current flow level.

	root_context       bool // Is it the document root context?
	sequence_context   bool // Is it a sequence context?
	mapping_context    bool // Is it a mapping context?
	simple_key_context bool // Is it a simple mapping key context?

	line       int  // The current line.
	column     int  // The current column.
	whitespace bool // If the last character was a whitespace?
	indention  bool // If the last character was an indentation character (' ', '-', '?', ':')?
	open_ended bool // If an explicit document end is required?

	space_above bool // Is there's an empty line above?
	foot_indent int  // The indent used to write the foot comment above, or -1 if none.

	// Anchor analysis.
	anchor_data struct {
		anchor []byte // The anchor value.
		alias  bool   // Is it an alias?
	}

	// Tag analysis.
	tag_data struct {
		handle []byte // The tag handle.
		suffix []byte // The tag suffix.
	}

	// Scalar analysis.
	scalar_data struct {
		value                 []byte              // The scalar value.
		multiline             bool                // Does the scalar contain line breaks?
		flow_plain_allowed    bool                // Can the scalar be expessed in the flow plain style?
		block_plain_allowed   bool                // Can the scalar be expressed in the block plain style?
		single_quoted_allowed bool                // Can the scalar be expressed in the single quoted style?
		block_allowed         bool                // Can the scalar be expressed in the literal or folded styles?
		style                 yaml_scalar_style_t // The output style.
	}

	// Comments
	head_comment []byte
	line_comment []byte
	foot_comment []byte
	tail_comment []byte

	key_line_comment []byte

	// Dumper stuff

	opened bool // If the stream was already opened?
	closed bool // If the stream was already closed?

	// The information associated with the document nodes.
	anchors *struct {
		references int  // The number of references.
		anchor     int  // The anchor id.
		serialized bool // If the node has been emitted?
	}

	last_anchor_id int // The last assigned anchor id.

	document *yaml_document_t // The currently emitted document.
}
