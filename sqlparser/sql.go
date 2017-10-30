//line ./sqlparser/sql.y:19
package sqlparser

import __yyfmt__ "fmt"

//line ./sqlparser/sql.y:21
import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line ./sqlparser/sql.y:45
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	tuple       Tuple
	valExprs    ValExprs
	values      Values
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const ALL = 57359
const DISTINCT = 57360
const AS = 57361
const EXISTS = 57362
const NULL = 57363
const ASC = 57364
const DESC = 57365
const VALUES = 57366
const INTO = 57367
const DUPLICATE = 57368
const KEY = 57369
const DEFAULT = 57370
const SET = 57371
const LOCK = 57372
const ID = 57373
const STRING = 57374
const NUMBER = 57375
const VALUE_ARG = 57376
const COMMENT = 57377
const UNION = 57378
const MINUS = 57379
const EXCEPT = 57380
const INTERSECT = 57381
const JOIN = 57382
const STRAIGHT_JOIN = 57383
const LEFT = 57384
const RIGHT = 57385
const INNER = 57386
const OUTER = 57387
const CROSS = 57388
const NATURAL = 57389
const USE = 57390
const FORCE = 57391
const ON = 57392
const OR = 57393
const AND = 57394
const NOT = 57395
const BETWEEN = 57396
const CASE = 57397
const WHEN = 57398
const THEN = 57399
const ELSE = 57400
const LE = 57401
const GE = 57402
const NE = 57403
const NULL_SAFE_EQUAL = 57404
const IS = 57405
const LIKE = 57406
const IN = 57407
const UNARY = 57408
const UNDERSCORE_BINARY = 57409
const END = 57410
const BEGIN = 57411
const START = 57412
const TRANSACTION = 57413
const COMMIT = 57414
const ROLLBACK = 57415
const NAMES = 57416
const REPLACE = 57417
const ADMIN = 57418
const HELP = 57419
const OFFSET = 57420
const COLLATE = 57421
const CREATE = 57422
const ALTER = 57423
const DROP = 57424
const RENAME = 57425
const TABLE = 57426
const INDEX = 57427
const VIEW = 57428
const TO = 57429
const IGNORE = 57430
const IF = 57431
const UNIQUE = 57432
const USING = 57433
const TRUNCATE = 57434

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"COMMENT",
	"'('",
	"'~'",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	"','",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"OR",
	"AND",
	"NOT",
	"BETWEEN",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"'='",
	"'<'",
	"'>'",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"IS",
	"LIKE",
	"IN",
	"'|'",
	"'&'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'^'",
	"'.'",
	"UNARY",
	"UNDERSCORE_BINARY",
	"END",
	"BEGIN",
	"START",
	"TRANSACTION",
	"COMMIT",
	"ROLLBACK",
	"NAMES",
	"REPLACE",
	"ADMIN",
	"HELP",
	"OFFSET",
	"COLLATE",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
	"TRUNCATE",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 220
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 672

var yyAct = [...]int{

	111, 108, 316, 353, 188, 386, 102, 140, 74, 347,
	271, 119, 76, 186, 230, 189, 3, 149, 97, 266,
	200, 98, 60, 92, 163, 162, 109, 118, 212, 88,
	125, 395, 395, 395, 157, 63, 157, 79, 115, 116,
	117, 157, 134, 122, 81, 259, 53, 78, 77, 228,
	83, 280, 66, 85, 47, 143, 49, 89, 364, 206,
	50, 52, 363, 53, 126, 335, 337, 72, 362, 288,
	289, 290, 291, 292, 82, 293, 294, 96, 257, 84,
	120, 121, 103, 139, 37, 38, 39, 40, 123, 133,
	299, 147, 54, 142, 78, 153, 130, 397, 396, 394,
	343, 159, 308, 55, 56, 57, 62, 306, 151, 58,
	124, 258, 185, 187, 148, 227, 79, 336, 155, 344,
	190, 275, 161, 136, 191, 94, 194, 172, 208, 244,
	78, 77, 78, 77, 339, 400, 135, 207, 199, 138,
	197, 267, 243, 242, 202, 209, 163, 162, 204, 205,
	162, 198, 346, 359, 260, 222, 175, 176, 177, 172,
	87, 236, 207, 368, 267, 131, 311, 61, 234, 238,
	239, 103, 223, 237, 75, 247, 348, 226, 240, 235,
	348, 245, 246, 276, 249, 250, 251, 252, 253, 254,
	255, 256, 241, 171, 170, 173, 174, 175, 176, 177,
	172, 163, 162, 146, 51, 361, 103, 103, 273, 218,
	248, 262, 264, 277, 369, 360, 333, 274, 90, 332,
	270, 278, 268, 329, 216, 78, 77, 219, 330, 78,
	284, 327, 331, 282, 201, 131, 328, 259, 380, 151,
	201, 281, 156, 370, 234, 283, 269, 298, 71, 263,
	285, 114, 118, 93, 233, 125, 379, 301, 302, 232,
	64, 378, 101, 115, 116, 117, 286, 106, 122, 300,
	78, 77, 131, 305, 312, 157, 319, 103, 314, 134,
	195, 315, 307, 19, 151, 310, 313, 105, 224, 126,
	193, 192, 215, 217, 214, 93, 93, 127, 234, 234,
	79, 325, 326, 297, 160, 120, 121, 99, 340, 233,
	322, 338, 392, 123, 232, 296, 64, 342, 350, 37,
	38, 39, 40, 321, 349, 345, 393, 381, 355, 320,
	221, 351, 354, 220, 208, 124, 62, 154, 144, 261,
	141, 19, 20, 21, 22, 137, 171, 170, 173, 174,
	175, 176, 177, 172, 86, 365, 372, 373, 129, 367,
	366, 210, 128, 376, 375, 23, 377, 374, 203, 19,
	91, 304, 204, 145, 384, 69, 67, 385, 317, 387,
	387, 387, 382, 383, 354, 388, 389, 34, 150, 358,
	318, 78, 77, 272, 201, 357, 401, 324, 73, 398,
	399, 402, 390, 403, 19, 42, 171, 170, 173, 174,
	175, 176, 177, 172, 173, 174, 175, 176, 177, 172,
	28, 29, 18, 30, 31, 17, 32, 33, 16, 15,
	14, 24, 25, 27, 26, 114, 118, 341, 13, 125,
	12, 41, 95, 35, 211, 48, 101, 115, 116, 117,
	279, 106, 122, 19, 171, 170, 173, 174, 175, 176,
	177, 172, 43, 44, 45, 46, 213, 80, 114, 118,
	152, 105, 125, 126, 59, 391, 371, 65, 352, 79,
	115, 116, 117, 356, 106, 122, 323, 309, 196, 120,
	121, 99, 265, 113, 110, 112, 225, 123, 107, 114,
	118, 164, 104, 125, 105, 334, 126, 231, 287, 229,
	79, 115, 116, 117, 100, 106, 122, 19, 295, 124,
	158, 68, 120, 121, 36, 288, 289, 290, 291, 292,
	123, 293, 294, 118, 70, 105, 125, 126, 11, 10,
	9, 8, 7, 79, 115, 116, 117, 6, 134, 122,
	5, 4, 124, 120, 121, 2, 118, 303, 1, 125,
	0, 123, 0, 132, 0, 0, 79, 115, 116, 117,
	126, 134, 122, 0, 171, 170, 173, 174, 175, 176,
	177, 172, 0, 124, 0, 0, 120, 121, 118, 0,
	0, 125, 0, 126, 123, 0, 0, 0, 79, 115,
	116, 117, 0, 134, 122, 0, 0, 0, 0, 120,
	121, 0, 0, 0, 0, 0, 124, 123, 170, 173,
	174, 175, 176, 177, 172, 126, 171, 170, 173, 174,
	175, 176, 177, 172, 0, 0, 0, 0, 0, 124,
	0, 120, 121, 0, 166, 168, 0, 0, 0, 123,
	178, 179, 180, 181, 182, 183, 184, 169, 167, 165,
	171, 170, 173, 174, 175, 176, 177, 172, 0, 0,
	0, 124,
}
var yyPact = [...]int{

	336, -1000, -1000, 281, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -45, -40, -7, 4, -1000, 23,
	-1000, -1000, -1000, 75, 229, -1000, 399, 359, -1000, -1000,
	-1000, 357, -1000, -57, 305, 389, 85, -60, -26, 229,
	-1000, -20, 229, -1000, 323, -75, 229, -75, -1000, 345,
	260, -1000, 45, -1000, -1000, -22, -1000, -1000, 415, -1000,
	262, 337, 329, 305, 193, 535, -1000, 74, -1000, 43,
	314, 83, 229, -1000, 309, -1000, -47, 307, 353, 150,
	229, 305, 364, 269, 306, 305, -1000, 233, -1000, -1000,
	285, 42, 147, 588, -1000, 479, 448, -1000, -1000, -1000,
	567, 255, 254, -1000, 244, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 567, -1000, 305, 269,
	384, 269, -1000, 274, 512, 6, 303, -1000, 341, -78,
	-1000, 196, -1000, 302, -1000, -1000, 299, -1000, 259, -1000,
	243, 281, 7, -1000, -1000, -1000, 223, 415, -1000, -1000,
	229, 97, 479, 479, 567, 243, 72, 567, 567, 154,
	567, 567, 567, 567, 567, 567, 567, 567, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 588, -30, 3, 46,
	588, -1000, 231, 415, -1000, 399, 82, 554, 217, 230,
	380, 479, -1000, 567, 554, 554, -1000, -1000, 41, -1000,
	-1000, 130, 229, -1000, -51, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 364, 269, 203, -1000, -1000, 269, 224,
	482, 284, 278, 10, -1000, -1000, -1000, -1000, -1000, 95,
	554, -1000, 243, 567, 567, 554, 502, -1000, 350, 340,
	545, -1000, 80, 80, 48, 48, 48, -1000, -1000, 567,
	-1000, -1000, -1, 415, -6, 105, -1000, 479, 364, 269,
	380, 363, 376, 147, 554, 229, 298, -1000, -1000, 292,
	-1000, -1000, 193, 243, -1000, 386, 223, 223, -1000, -1000,
	188, 180, 189, 176, 173, 14, -1000, 280, 26, 277,
	-1000, 554, 382, 567, -1000, 554, -1000, -8, -1000, 36,
	-1000, 567, 92, 127, 123, 363, -1000, 567, 567, -1000,
	-1000, -1000, -1000, 383, 375, 482, 100, -1000, 172, -1000,
	162, -1000, -1000, -1000, -1000, -32, -38, -42, -1000, -1000,
	-1000, 567, 554, -1000, -1000, 554, 567, -1000, 333, -1000,
	-1000, 121, 201, -1000, 334, -1000, 380, 479, 567, 479,
	-1000, -1000, 225, 220, 202, 554, 554, 300, 567, 567,
	567, -1000, -1000, -1000, 363, 147, 195, 147, 229, 229,
	229, 395, 554, 554, -1000, 296, -9, -1000, -10, -11,
	269, -1000, 393, 64, -1000, 229, -1000, -1000, 193, -1000,
	229, -1000, 229, -1000,
}
var yyPgo = [...]int{

	0, 558, 555, 15, 551, 550, 547, 542, 541, 540,
	539, 538, 441, 534, 524, 521, 204, 18, 21, 520,
	518, 514, 509, 14, 508, 507, 22, 505, 5, 20,
	6, 502, 501, 17, 498, 13, 26, 4, 496, 495,
	11, 494, 1, 493, 492, 19, 488, 487, 486, 483,
	10, 478, 3, 476, 2, 475, 23, 470, 9, 8,
	12, 160, 467, 466, 450, 445, 444, 0, 7, 442,
	440, 438, 430, 429, 428, 425, 422, 405,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 3, 3,
	3, 4, 4, 73, 73, 5, 6, 7, 7, 7,
	7, 70, 70, 71, 72, 74, 74, 75, 76, 8,
	8, 8, 9, 9, 9, 10, 11, 11, 11, 77,
	12, 13, 13, 14, 14, 14, 14, 14, 15, 15,
	17, 17, 18, 18, 18, 21, 21, 19, 19, 19,
	22, 22, 23, 23, 23, 23, 20, 20, 20, 24,
	24, 24, 24, 24, 24, 24, 24, 24, 25, 25,
	25, 26, 26, 27, 27, 27, 27, 28, 28, 29,
	29, 30, 30, 30, 30, 30, 31, 31, 31, 31,
	31, 31, 31, 31, 31, 31, 32, 32, 32, 32,
	32, 32, 32, 33, 33, 38, 38, 36, 36, 40,
	37, 37, 35, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 39,
	39, 41, 41, 41, 41, 43, 46, 46, 44, 44,
	45, 47, 47, 42, 42, 42, 34, 34, 34, 34,
	48, 48, 49, 49, 50, 50, 51, 51, 52, 53,
	53, 53, 54, 54, 54, 54, 55, 55, 55, 56,
	56, 57, 57, 58, 58, 59, 59, 60, 60, 61,
	61, 62, 62, 16, 16, 63, 63, 63, 63, 63,
	64, 64, 65, 65, 66, 66, 67, 68, 69, 69,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 4, 12,
	3, 8, 8, 6, 6, 8, 7, 3, 4, 4,
	6, 1, 2, 1, 1, 4, 2, 2, 4, 5,
	8, 4, 6, 7, 4, 5, 4, 5, 5, 0,
	2, 0, 2, 1, 2, 1, 1, 1, 0, 1,
	1, 3, 1, 2, 3, 1, 1, 0, 1, 2,
	1, 3, 3, 3, 3, 5, 0, 1, 2, 1,
	1, 2, 3, 2, 3, 2, 2, 2, 1, 3,
	1, 1, 3, 0, 5, 5, 5, 1, 3, 0,
	2, 1, 3, 3, 2, 3, 3, 3, 4, 3,
	4, 5, 6, 3, 4, 2, 1, 1, 1, 1,
	1, 1, 1, 2, 1, 1, 3, 3, 1, 3,
	1, 3, 1, 1, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 3, 4, 5, 4, 1, 1,
	1, 1, 1, 1, 1, 5, 0, 1, 1, 2,
	4, 0, 2, 1, 3, 5, 1, 1, 1, 1,
	0, 3, 0, 2, 0, 3, 1, 3, 2, 0,
	1, 1, 0, 2, 4, 4, 0, 2, 4, 0,
	3, 1, 3, 0, 5, 1, 3, 3, 3, 0,
	2, 0, 3, 0, 1, 1, 1, 1, 1, 1,
	0, 1, 0, 1, 0, 2, 1, 0, 0, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -70, -71, -72, -73, -74, -75, -76, 5,
	6, 7, 8, 29, 95, 96, 98, 97, 84, 85,
	87, 88, 90, 91, 51, 107, -14, 38, 39, 40,
	41, -12, -77, -12, -12, -12, -12, 99, -65, 101,
	105, -16, 101, 103, 99, 99, 100, 101, 86, -12,
	-26, 92, 31, -67, 31, -12, -3, 17, -15, 18,
	-13, -16, -26, 9, -59, 89, -60, -42, -67, 31,
	-62, 104, 100, -67, 99, -67, 31, -61, 104, -67,
	-61, 25, -56, 36, 80, -69, 99, -17, -18, 76,
	-21, 31, -30, -35, -31, 56, 36, -34, -42, -36,
	-41, -67, -39, -43, 20, 32, 33, 34, 21, -40,
	74, 75, 37, 82, 104, 24, 58, 35, 25, 29,
	-26, 42, 28, -35, 36, 62, 80, 31, 56, -67,
	-68, 31, -68, 102, 31, 20, 53, -67, -26, -33,
	24, -3, -57, -42, 31, -26, 9, 42, -19, -67,
	19, 80, 55, 54, -32, 71, 56, 70, 57, 69,
	73, 72, 79, 74, 75, 76, 77, 78, 62, 63,
	64, 65, 66, 67, 68, -30, -35, -30, -37, -3,
	-35, -35, 36, 36, -40, 36, -46, -35, -26, -59,
	-29, 10, -60, 94, -35, -35, 53, -67, 31, -68,
	20, -66, 106, -63, 98, 96, 28, 97, 13, 31,
	31, 31, -68, -56, 29, -38, -36, 108, 42, -22,
	-23, -25, 36, 31, -40, -18, -67, 76, -30, -30,
	-35, -36, 71, 70, 57, -35, -35, 21, 56, -35,
	-35, -35, -35, -35, -35, -35, -35, 108, 108, 42,
	108, 108, -17, 18, -17, -44, -45, 59, -56, 29,
	-29, -50, 13, -30, -35, 80, 53, -67, -68, -64,
	102, -33, -59, 42, -42, -29, 42, -24, 43, 44,
	45, 46, 47, 49, 50, -20, 31, 19, -23, 80,
	-36, -35, -35, 55, 21, -35, 108, -17, 108, -47,
	-45, 61, -30, -33, -59, -50, -54, 15, 14, -67,
	31, 31, -36, -48, 11, -23, -23, 43, 48, 43,
	48, 43, 43, 43, -27, 51, 103, 52, 31, 108,
	31, 55, -35, 108, 83, -35, 60, -58, 53, -58,
	-54, -35, -51, -52, -35, -68, -49, 12, 14, 53,
	43, 43, 100, 100, 100, -35, -35, 26, 42, 93,
	42, -53, 22, 23, -50, -30, -37, -30, 36, 36,
	36, 27, -35, -35, -52, -54, -28, -67, -28, -28,
	7, -55, 16, 30, 108, 42, 108, 108, -59, 7,
	71, -67, -67, -67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 49,
	49, 49, 49, 49, 212, 203, 0, 0, 31, 0,
	33, 34, 49, 0, 0, 49, 0, 53, 55, 56,
	57, 58, 51, 203, 0, 0, 0, 201, 0, 0,
	213, 0, 0, 204, 0, 199, 0, 199, 32, 0,
	189, 36, 91, 37, 216, 218, 20, 54, 0, 59,
	50, 0, 0, 0, 27, 0, 195, 0, 163, 216,
	0, 0, 0, 217, 0, 217, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 219, 18, 60, 62,
	67, 216, 65, 66, 101, 0, 0, 132, 133, 134,
	0, 163, 0, 148, 0, 166, 167, 168, 169, 128,
	151, 152, 153, 154, 149, 150, 156, 52, 0, 0,
	99, 0, 28, 29, 0, 0, 0, 217, 0, 214,
	41, 0, 44, 0, 46, 200, 0, 217, 189, 35,
	0, 124, 0, 191, 92, 38, 0, 0, 63, 68,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 116, 117,
	118, 119, 120, 121, 122, 104, 0, 0, 0, 0,
	130, 143, 0, 0, 115, 0, 0, 157, 189, 99,
	174, 0, 196, 0, 130, 197, 198, 164, 216, 39,
	202, 0, 0, 217, 210, 205, 206, 207, 208, 209,
	45, 47, 48, 0, 0, 123, 125, 190, 0, 99,
	70, 76, 0, 88, 90, 61, 69, 64, 102, 103,
	106, 107, 0, 0, 0, 109, 0, 113, 0, 135,
	136, 137, 138, 139, 140, 141, 142, 105, 127, 0,
	129, 144, 0, 0, 0, 161, 158, 0, 0, 0,
	174, 182, 0, 100, 30, 0, 0, 215, 42, 0,
	211, 23, 24, 0, 192, 170, 0, 0, 79, 80,
	0, 0, 0, 0, 0, 93, 77, 0, 0, 0,
	108, 110, 0, 0, 114, 131, 145, 0, 147, 0,
	159, 0, 0, 193, 193, 182, 26, 0, 0, 165,
	217, 43, 126, 172, 0, 71, 74, 81, 0, 83,
	0, 85, 86, 87, 72, 0, 0, 0, 78, 73,
	89, 0, 111, 146, 155, 162, 0, 21, 0, 22,
	25, 183, 175, 176, 179, 40, 174, 0, 0, 0,
	82, 84, 0, 0, 0, 112, 160, 0, 0, 0,
	0, 178, 180, 181, 182, 173, 171, 75, 0, 0,
	0, 0, 184, 185, 177, 186, 0, 97, 0, 0,
	0, 19, 0, 0, 94, 0, 95, 96, 194, 187,
	0, 98, 0, 188,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 78, 73, 3,
	36, 108, 76, 74, 42, 75, 80, 77, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	63, 62, 64, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 79, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 72, 3, 37,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 38, 39, 40, 41, 43, 44,
	45, 46, 47, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 65, 66, 67,
	68, 69, 70, 71, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:194
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:200
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 18:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:221
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs}
		}
	case 19:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line ./sqlparser/sql.y:225
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:229
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 21:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./sqlparser/sql.y:236
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./sqlparser/sql.y:240
		{
			cols := make(Columns, 0, len(yyDollar[7].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, col := range yyDollar[7].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 23:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:252
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:256
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 25:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./sqlparser/sql.y:269
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 26:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./sqlparser/sql.y:275
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:281
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:285
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal("default")}}}
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:289
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: yyDollar[4].valExpr}}}
		}
	case 30:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:293
		{
			yyVAL.statement = &Set{
				Comments: Comments(yyDollar[2].bytes2),
				Exprs: UpdateExprs{
					&UpdateExpr{
						Name: &ColName{Name: []byte("names")}, Expr: yyDollar[4].valExpr,
					},
					&UpdateExpr{
						Name: &ColName{Name: []byte("collate")}, Expr: yyDollar[6].valExpr,
					},
				},
			}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:309
		{
			yyVAL.statement = &Begin{}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:313
		{
			yyVAL.statement = &Begin{}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:320
		{
			yyVAL.statement = &Commit{}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:326
		{
			yyVAL.statement = &Rollback{}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:332
		{
			yyVAL.statement = &Admin{Region: yyDollar[2].tableName, Columns: yyDollar[3].columns, Rows: yyDollar[4].insRows}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:336
		{
			yyVAL.statement = &AdminHelp{}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:342
		{
			yyVAL.statement = &UseDB{DB: string(yyDollar[2].bytes)}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:348
		{
			yyVAL.statement = &Truncate{Comments: Comments(yyDollar[2].bytes2), TableOpt: yyDollar[3].str, Table: yyDollar[4].tableName}
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:354
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 40:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./sqlparser/sql.y:358
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:363
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:369
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 43:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./sqlparser/sql.y:373
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:378
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 45:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:384
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:390
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 47:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:394
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 48:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:399
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:404
		{
			SetAllowComments(yylex, true)
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:408
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:414
		{
			yyVAL.bytes2 = nil
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:418
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:424
		{
			yyVAL.str = AST_UNION
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:428
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:432
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:436
		{
			yyVAL.str = AST_EXCEPT
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:440
		{
			yyVAL.str = AST_INTERSECT
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:445
		{
			yyVAL.str = ""
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:449
		{
			yyVAL.str = AST_DISTINCT
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:455
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:459
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:465
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:469
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:473
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:479
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:483
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:488
		{
			yyVAL.bytes = nil
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:492
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:496
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:502
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:506
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:512
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:516
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:520
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 75:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:524
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 76:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:529
		{
			yyVAL.bytes = nil
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:533
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:537
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:543
		{
			yyVAL.str = AST_JOIN
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:547
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:551
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:555
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:559
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:563
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:567
		{
			yyVAL.str = AST_JOIN
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:571
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:575
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:581
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:585
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:589
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:595
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:599
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 93:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:604
		{
			yyVAL.indexHints = nil
		}
	case 94:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:608
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 95:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:612
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 96:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:616
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:622
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:626
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:631
		{
			yyVAL.boolExpr = nil
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:635
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:642
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:646
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:650
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:654
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:660
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:664
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 108:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:668
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:672
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:676
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 111:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:680
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 112:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:684
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:688
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 114:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:692
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:696
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:702
		{
			yyVAL.str = AST_EQ
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:706
		{
			yyVAL.str = AST_LT
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:710
		{
			yyVAL.str = AST_GT
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:714
		{
			yyVAL.str = AST_LE
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:718
		{
			yyVAL.str = AST_GE
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:722
		{
			yyVAL.str = AST_NE
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:726
		{
			yyVAL.str = AST_NSE
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:732
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:736
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:742
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:746
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:752
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:756
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:762
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:768
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:772
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:778
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:782
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:786
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:790
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:794
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:798
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:802
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:806
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:810
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:814
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:818
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 143:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:822
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				switch yyDollar[1].str {
				case "-":
					yyVAL.valExpr = append(NumVal("-"), num...)
				case "+":
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].str, Expr: yyDollar[2].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].str, Expr: yyDollar[2].valExpr}
			}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:837
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 145:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:841
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 146:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:845
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 147:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:849
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:853
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:859
		{
			yyVAL.bytes = IF_BYTES
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:863
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:869
		{
			yyVAL.str = AST_UPLUS
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:873
		{
			yyVAL.str = AST_UMINUS
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:877
		{
			yyVAL.str = AST_TILDA
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:881
		{
			yyVAL.str = AST_UBinary
		}
	case 155:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:887
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:892
		{
			yyVAL.valExpr = nil
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:896
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:902
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 159:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:906
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 160:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:912
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 161:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:917
		{
			yyVAL.valExpr = nil
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:921
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:927
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:931
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 165:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:935
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[3].bytes, Name: yyDollar[5].bytes}
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:941
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:945
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:949
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:953
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 170:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:958
		{
			yyVAL.valExprs = nil
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:962
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 172:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:967
		{
			yyVAL.boolExpr = nil
		}
	case 173:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:971
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 174:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:976
		{
			yyVAL.orderBy = nil
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:980
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:986
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:990
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:996
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 179:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1001
		{
			yyVAL.str = AST_ASC
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1005
		{
			yyVAL.str = AST_ASC
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1009
		{
			yyVAL.str = AST_DESC
		}
	case 182:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1014
		{
			yyVAL.limit = nil
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:1018
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 184:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:1022
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 185:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:1026
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 186:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1031
		{
			yyVAL.str = ""
		}
	case 187:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:1035
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 188:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:1039
		{
			if !bytes.Equal(yyDollar[3].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyDollar[4].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 189:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1052
		{
			yyVAL.columns = nil
		}
	case 190:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1056
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1062
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1066
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 193:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1071
		{
			yyVAL.updateExprs = nil
		}
	case 194:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:1075
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1081
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1085
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 197:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1091
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 198:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1095
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: StrVal("ON")}
		}
	case 199:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1100
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:1102
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1105
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1107
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1110
		{
			yyVAL.str = ""
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1112
		{
			yyVAL.str = AST_IGNORE
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1116
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1118
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1120
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1122
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1124
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1127
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1129
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1132
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1134
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1137
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:1139
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1143
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1148
		{
			ForceEOF(yylex)
		}
	case 218:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1153
		{
			yyVAL.str = ""
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1157
		{
			yyVAL.str = AST_TABLE
		}
	}
	goto yystack /* stack new state and value */
}
