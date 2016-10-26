//line ./sqlparser/sql.y:20
package sqlparser

import __yyfmt__ "fmt"

//line ./sqlparser/sql.y:20
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
const END = 57409
const BEGIN = 57410
const START = 57411
const TRANSACTION = 57412
const COMMIT = 57413
const ROLLBACK = 57414
const NAMES = 57415
const REPLACE = 57416
const ADMIN = 57417
const HELP = 57418
const OFFSET = 57419
const COLLATE = 57420
const CREATE = 57421
const ALTER = 57422
const DROP = 57423
const RENAME = 57424
const TABLE = 57425
const INDEX = 57426
const VIEW = 57427
const TO = 57428
const IGNORE = 57429
const IF = 57430
const UNIQUE = 57431
const USING = 57432

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

const yyNprod = 212
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 599

var yyAct = [...]int{

	106, 103, 305, 341, 180, 374, 97, 133, 71, 335,
	261, 114, 192, 178, 220, 181, 3, 142, 93, 256,
	73, 383, 58, 89, 155, 154, 202, 92, 85, 35,
	36, 37, 38, 78, 61, 277, 278, 279, 280, 281,
	383, 282, 283, 51, 104, 75, 74, 269, 80, 383,
	63, 82, 149, 149, 149, 86, 249, 45, 218, 47,
	50, 352, 51, 48, 136, 69, 53, 54, 55, 60,
	351, 350, 323, 325, 79, 81, 247, 332, 52, 98,
	132, 56, 76, 288, 153, 385, 126, 164, 140, 129,
	135, 75, 146, 124, 257, 91, 151, 250, 327, 167,
	168, 169, 164, 388, 384, 144, 62, 177, 179, 128,
	131, 141, 257, 382, 300, 182, 331, 297, 295, 183,
	248, 186, 217, 324, 75, 74, 75, 74, 84, 59,
	198, 237, 191, 154, 189, 347, 234, 336, 199, 72,
	265, 196, 197, 155, 154, 190, 194, 139, 212, 233,
	232, 227, 317, 226, 198, 349, 125, 318, 155, 154,
	224, 228, 229, 98, 334, 213, 238, 336, 225, 348,
	230, 321, 329, 235, 236, 320, 239, 240, 241, 242,
	243, 244, 245, 246, 87, 319, 193, 125, 216, 163,
	162, 165, 166, 167, 168, 169, 164, 249, 98, 98,
	263, 315, 231, 266, 260, 148, 316, 358, 272, 264,
	208, 267, 252, 254, 258, 75, 74, 223, 275, 75,
	273, 368, 222, 271, 49, 206, 367, 366, 209, 144,
	193, 270, 274, 121, 224, 127, 187, 287, 149, 18,
	253, 185, 109, 113, 184, 90, 119, 290, 291, 35,
	36, 37, 38, 96, 110, 111, 112, 286, 101, 117,
	75, 74, 125, 294, 301, 223, 68, 98, 303, 285,
	222, 304, 259, 152, 144, 299, 302, 289, 100, 90,
	120, 296, 214, 62, 369, 62, 76, 224, 224, 90,
	313, 314, 205, 207, 204, 328, 115, 116, 94, 162,
	165, 166, 167, 168, 169, 164, 330, 338, 18, 39,
	326, 309, 380, 337, 333, 308, 343, 310, 211, 210,
	339, 342, 60, 109, 113, 118, 381, 119, 251, 41,
	42, 43, 44, 147, 76, 110, 111, 112, 137, 101,
	117, 57, 134, 353, 130, 83, 123, 355, 354, 122,
	293, 364, 363, 88, 365, 362, 18, 200, 138, 100,
	196, 120, 372, 66, 64, 373, 346, 375, 375, 375,
	370, 371, 342, 376, 377, 143, 307, 115, 116, 75,
	74, 109, 113, 306, 389, 119, 262, 386, 345, 390,
	312, 391, 96, 110, 111, 112, 193, 101, 117, 165,
	166, 167, 168, 169, 164, 70, 118, 387, 378, 18,
	40, 109, 113, 17, 18, 119, 16, 100, 15, 120,
	14, 13, 76, 110, 111, 112, 12, 101, 117, 201,
	113, 46, 268, 119, 203, 115, 116, 94, 77, 145,
	76, 110, 111, 112, 379, 127, 117, 100, 359, 120,
	340, 344, 311, 298, 18, 19, 20, 21, 188, 255,
	113, 108, 105, 119, 118, 115, 116, 120, 292, 107,
	76, 110, 111, 112, 356, 127, 117, 215, 22, 102,
	156, 99, 322, 115, 116, 163, 162, 165, 166, 167,
	168, 169, 164, 221, 118, 276, 219, 120, 95, 284,
	33, 150, 65, 34, 163, 162, 165, 166, 167, 168,
	169, 164, 118, 115, 116, 67, 163, 162, 165, 166,
	167, 168, 169, 164, 357, 360, 361, 11, 10, 9,
	8, 7, 27, 28, 6, 29, 30, 195, 31, 32,
	5, 4, 118, 23, 24, 26, 25, 158, 160, 2,
	1, 0, 0, 170, 171, 172, 173, 174, 175, 176,
	161, 159, 157, 163, 162, 165, 166, 167, 168, 169,
	164, 0, 0, 0, 0, 163, 162, 165, 166, 167,
	168, 169, 164, 163, 162, 165, 166, 167, 168, 169,
	164, 277, 278, 279, 280, 281, 0, 282, 283,
}
var yyPact = [...]int{

	449, -1000, -1000, 211, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -41, -40, -20, -32, -1000, -4, -1000,
	-1000, -1000, 38, 252, 404, 347, -1000, -1000, -1000, 345,
	-1000, -59, 291, 396, 51, -70, -25, 252, -1000, -23,
	252, -1000, 314, -75, 252, -75, -1000, 328, 209, -1000,
	15, -1000, -1000, -1000, -1000, 361, -1000, 198, 324, 317,
	291, 145, 439, -1000, 47, -1000, 9, 313, 54, 252,
	-1000, 311, -1000, -37, 307, 338, 94, 252, 291, 351,
	255, 302, 196, -1000, -1000, 254, 4, 89, 491, -1000,
	391, 303, -1000, -1000, -1000, 439, 208, 205, -1000, 200,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	439, -1000, 291, 255, 386, 255, 444, 409, 439, 252,
	-1000, 337, -79, -1000, 197, -1000, 288, -1000, -1000, 287,
	-1000, 253, -1000, 199, 211, 16, -1000, -1000, 186, 361,
	-1000, -1000, 252, 75, 391, 391, 439, 199, 79, 439,
	439, 110, 439, 439, 439, 439, 439, 439, 439, 439,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 491, -30,
	14, -9, 491, -1000, 222, 361, -1000, 404, 35, 511,
	243, 220, 373, 391, -1000, 439, 511, 511, -1000, -1000,
	-1000, 87, 252, -1000, -54, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 351, 255, 166, -1000, -1000, 255, 176,
	548, 238, 234, 3, -1000, -1000, -1000, -1000, -1000, 78,
	511, -1000, 199, 439, 439, 511, 413, -1000, 329, 325,
	226, -1000, 23, 23, 8, 8, 8, -1000, -1000, 439,
	-1000, -1000, 12, 361, 11, 53, -1000, 391, 351, 255,
	373, 368, 362, 89, 511, 284, -1000, -1000, 280, -1000,
	-1000, 145, 199, -1000, 379, 186, 186, -1000, -1000, 158,
	109, 142, 132, 128, 21, -1000, 279, -8, 264, -1000,
	511, 117, 439, -1000, 511, -1000, 10, -1000, -5, -1000,
	439, 104, 84, 114, 368, -1000, 439, 439, -1000, -1000,
	-1000, 376, 352, 548, 82, -1000, 126, -1000, 112, -1000,
	-1000, -1000, -1000, -28, -29, -38, -1000, -1000, -1000, 439,
	511, -1000, -1000, 511, 439, -1000, 321, -1000, -1000, 432,
	165, -1000, 503, -1000, 373, 391, 439, 391, -1000, -1000,
	191, 190, 185, 511, 511, 257, 439, 439, 439, -1000,
	-1000, -1000, 368, 89, 155, 89, 252, 252, 252, 401,
	511, 511, -1000, 296, 7, -1000, -2, -21, 255, -1000,
	400, 32, -1000, 252, -1000, -1000, 145, -1000, 252, -1000,
	252, -1000,
}
var yyPgo = [...]int{

	0, 550, 549, 15, 541, 540, 534, 531, 530, 529,
	528, 527, 309, 515, 503, 502, 224, 27, 18, 501,
	499, 498, 496, 14, 495, 493, 22, 482, 5, 12,
	6, 481, 480, 17, 479, 13, 44, 4, 477, 469,
	11, 462, 1, 461, 459, 19, 458, 453, 452, 451,
	10, 450, 3, 448, 2, 444, 23, 439, 9, 8,
	20, 128, 438, 434, 432, 431, 429, 0, 7, 426,
	421, 420, 418, 416, 413, 410,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 72, 72, 5, 6, 7, 7, 7, 69,
	69, 70, 71, 73, 73, 74, 8, 8, 8, 9,
	9, 9, 10, 11, 11, 11, 75, 12, 13, 13,
	14, 14, 14, 14, 14, 15, 15, 17, 17, 18,
	18, 18, 21, 21, 19, 19, 19, 22, 22, 23,
	23, 23, 23, 20, 20, 20, 24, 24, 24, 24,
	24, 24, 24, 24, 24, 25, 25, 25, 26, 26,
	27, 27, 27, 27, 28, 28, 29, 29, 30, 30,
	30, 30, 30, 31, 31, 31, 31, 31, 31, 31,
	31, 31, 31, 32, 32, 32, 32, 32, 32, 32,
	33, 33, 38, 38, 36, 36, 40, 37, 37, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 35, 35, 39, 39, 41, 41,
	41, 43, 46, 46, 44, 44, 45, 47, 47, 42,
	42, 34, 34, 34, 34, 48, 48, 49, 49, 50,
	50, 51, 51, 52, 53, 53, 53, 54, 54, 54,
	54, 55, 55, 55, 56, 56, 57, 57, 58, 58,
	59, 59, 60, 61, 61, 62, 62, 16, 16, 63,
	63, 63, 63, 63, 64, 64, 65, 65, 66, 66,
	67, 68,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 12, 3,
	8, 8, 6, 6, 8, 7, 3, 4, 6, 1,
	2, 1, 1, 4, 2, 2, 5, 8, 4, 6,
	7, 4, 5, 4, 5, 5, 0, 2, 0, 2,
	1, 2, 1, 1, 1, 0, 1, 1, 3, 1,
	2, 3, 1, 1, 0, 1, 2, 1, 3, 3,
	3, 3, 5, 0, 1, 2, 1, 1, 2, 3,
	2, 3, 2, 2, 2, 1, 3, 1, 1, 3,
	0, 5, 5, 5, 1, 3, 0, 2, 1, 3,
	3, 2, 3, 3, 3, 4, 3, 4, 5, 6,
	3, 4, 2, 1, 1, 1, 1, 1, 1, 1,
	2, 1, 1, 3, 3, 1, 3, 1, 3, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 3, 4, 5, 4, 1, 1, 1, 1, 1,
	1, 5, 0, 1, 1, 2, 4, 0, 2, 1,
	3, 1, 1, 1, 1, 0, 3, 0, 2, 0,
	3, 1, 3, 2, 0, 1, 1, 0, 2, 4,
	4, 0, 2, 4, 0, 3, 1, 3, 0, 5,
	1, 3, 3, 0, 2, 0, 3, 0, 1, 1,
	1, 1, 1, 1, 0, 1, 0, 1, 0, 2,
	1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 29, 94, 95, 97, 96, 83, 84, 86,
	87, 89, 90, 51, -14, 38, 39, 40, 41, -12,
	-75, -12, -12, -12, -12, 98, -65, 100, 104, -16,
	100, 102, 98, 98, 99, 100, 85, -12, -26, 91,
	31, -67, 31, -3, 17, -15, 18, -13, -16, -26,
	9, -59, 88, -60, -42, -67, 31, -62, 103, 99,
	-67, 98, -67, 31, -61, 103, -67, -61, 25, -56,
	36, 80, -17, -18, 76, -21, 31, -30, -35, -31,
	56, 36, -34, -42, -36, -41, -67, -39, -43, 20,
	32, 33, 34, 21, -40, 74, 75, 37, 103, 24,
	58, 35, 25, 29, -26, 42, -35, 36, 62, 80,
	31, 56, -67, -68, 31, -68, 101, 31, 20, 53,
	-67, -26, -33, 24, -3, -57, -42, 31, 9, 42,
	-19, -67, 19, 80, 55, 54, -32, 71, 56, 70,
	57, 69, 73, 72, 79, 74, 75, 76, 77, 78,
	62, 63, 64, 65, 66, 67, 68, -30, -35, -30,
	-37, -3, -35, -35, 36, 36, -40, 36, -46, -35,
	-26, -59, -29, 10, -60, 93, -35, -35, -67, -68,
	20, -66, 105, -63, 97, 95, 28, 96, 13, 31,
	31, 31, -68, -56, 29, -38, -36, 106, 42, -22,
	-23, -25, 36, 31, -40, -18, -67, 76, -30, -30,
	-35, -36, 71, 70, 57, -35, -35, 21, 56, -35,
	-35, -35, -35, -35, -35, -35, -35, 106, 106, 42,
	106, 106, -17, 18, -17, -44, -45, 59, -56, 29,
	-29, -50, 13, -30, -35, 53, -67, -68, -64, 101,
	-33, -59, 42, -42, -29, 42, -24, 43, 44, 45,
	46, 47, 49, 50, -20, 31, 19, -23, 80, -36,
	-35, -35, 55, 21, -35, 106, -17, 106, -47, -45,
	61, -30, -33, -59, -50, -54, 15, 14, 31, 31,
	-36, -48, 11, -23, -23, 43, 48, 43, 48, 43,
	43, 43, -27, 51, 102, 52, 31, 106, 31, 55,
	-35, 106, 82, -35, 60, -58, 53, -58, -54, -35,
	-51, -52, -35, -68, -49, 12, 14, 53, 43, 43,
	99, 99, 99, -35, -35, 26, 42, 92, 42, -53,
	22, 23, -50, -30, -37, -30, 36, 36, 36, 27,
	-35, -35, -52, -54, -28, -67, -28, -28, 7, -55,
	16, 30, 106, 42, 106, 106, -59, 7, 71, -67,
	-67, -67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 46, 46,
	46, 46, 46, 206, 197, 0, 0, 29, 0, 31,
	32, 46, 0, 0, 0, 50, 52, 53, 54, 55,
	48, 197, 0, 0, 0, 195, 0, 0, 207, 0,
	0, 198, 0, 193, 0, 193, 30, 0, 184, 34,
	88, 35, 210, 19, 51, 0, 56, 47, 0, 0,
	0, 26, 0, 190, 0, 159, 210, 0, 0, 0,
	211, 0, 211, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 17, 57, 59, 64, 210, 62, 63, 98,
	0, 0, 129, 130, 131, 0, 159, 0, 145, 0,
	161, 162, 163, 164, 125, 148, 149, 150, 146, 147,
	152, 49, 0, 0, 96, 0, 27, 0, 0, 0,
	211, 0, 208, 38, 0, 41, 0, 43, 194, 0,
	211, 184, 33, 0, 121, 0, 186, 89, 0, 0,
	60, 65, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	113, 114, 115, 116, 117, 118, 119, 101, 0, 0,
	0, 0, 127, 140, 0, 0, 112, 0, 0, 153,
	184, 96, 169, 0, 191, 0, 127, 192, 160, 36,
	196, 0, 0, 211, 204, 199, 200, 201, 202, 203,
	42, 44, 45, 0, 0, 120, 122, 185, 0, 96,
	67, 73, 0, 85, 87, 58, 66, 61, 99, 100,
	103, 104, 0, 0, 0, 106, 0, 110, 0, 132,
	133, 134, 135, 136, 137, 138, 139, 102, 124, 0,
	126, 141, 0, 0, 0, 157, 154, 0, 0, 0,
	169, 177, 0, 97, 28, 0, 209, 39, 0, 205,
	22, 23, 0, 187, 165, 0, 0, 76, 77, 0,
	0, 0, 0, 0, 90, 74, 0, 0, 0, 105,
	107, 0, 0, 111, 128, 142, 0, 144, 0, 155,
	0, 0, 188, 188, 177, 25, 0, 0, 211, 40,
	123, 167, 0, 68, 71, 78, 0, 80, 0, 82,
	83, 84, 69, 0, 0, 0, 75, 70, 86, 0,
	108, 143, 151, 158, 0, 20, 0, 21, 24, 178,
	170, 171, 174, 37, 169, 0, 0, 0, 79, 81,
	0, 0, 0, 109, 156, 0, 0, 0, 0, 173,
	175, 176, 177, 168, 166, 72, 0, 0, 0, 0,
	179, 180, 172, 181, 0, 94, 0, 0, 0, 18,
	0, 0, 91, 0, 92, 93, 189, 182, 0, 95,
	0, 183,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 78, 73, 3,
	36, 106, 76, 74, 42, 75, 80, 77, 3, 3,
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
	97, 98, 99, 100, 101, 102, 103, 104, 105,
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
		//line ./sqlparser/sql.y:188
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:194
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:214
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs}
		}
	case 18:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line ./sqlparser/sql.y:218
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:222
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 20:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./sqlparser/sql.y:229
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 21:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./sqlparser/sql.y:233
		{
			cols := make(Columns, 0, len(yyDollar[7].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, col := range yyDollar[7].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 22:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:245
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 23:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:249
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./sqlparser/sql.y:262
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./sqlparser/sql.y:268
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:274
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:278
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: yyDollar[4].valExpr}}}
		}
	case 28:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:282
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
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:298
		{
			yyVAL.statement = &Begin{}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:302
		{
			yyVAL.statement = &Begin{}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:309
		{
			yyVAL.statement = &Commit{}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:315
		{
			yyVAL.statement = &Rollback{}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:321
		{
			yyVAL.statement = &Admin{Region: yyDollar[2].tableName, Columns: yyDollar[3].columns, Rows: yyDollar[4].insRows}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:325
		{
			yyVAL.statement = &AdminHelp{}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:331
		{
			yyVAL.statement = &UseDB{DB: string(yyDollar[2].bytes)}
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:337
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 37:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./sqlparser/sql.y:341
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:346
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 39:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:352
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 40:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./sqlparser/sql.y:356
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:361
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:367
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 43:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:373
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 44:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:377
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 45:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:382
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:387
		{
			SetAllowComments(yylex, true)
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:391
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:397
		{
			yyVAL.bytes2 = nil
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:401
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:407
		{
			yyVAL.str = AST_UNION
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:411
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:415
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:419
		{
			yyVAL.str = AST_EXCEPT
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:423
		{
			yyVAL.str = AST_INTERSECT
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:428
		{
			yyVAL.str = ""
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:432
		{
			yyVAL.str = AST_DISTINCT
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:438
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:442
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:448
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:452
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:456
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:462
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:466
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:471
		{
			yyVAL.bytes = nil
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:475
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:479
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:485
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:489
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:495
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:499
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:503
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:507
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 73:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:512
		{
			yyVAL.bytes = nil
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:516
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:520
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:526
		{
			yyVAL.str = AST_JOIN
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:530
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:534
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:538
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:542
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:546
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:550
		{
			yyVAL.str = AST_JOIN
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:554
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:558
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:564
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:568
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:572
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:578
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:582
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 90:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:587
		{
			yyVAL.indexHints = nil
		}
	case 91:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:591
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:595
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 93:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:599
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:605
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:609
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 96:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:614
		{
			yyVAL.boolExpr = nil
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:618
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:625
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:629
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:633
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:637
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:643
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:647
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:651
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:655
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 107:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:659
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 108:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:663
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 109:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:667
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:671
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:675
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:679
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:685
		{
			yyVAL.str = AST_EQ
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:689
		{
			yyVAL.str = AST_LT
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:693
		{
			yyVAL.str = AST_GT
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:697
		{
			yyVAL.str = AST_LE
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:701
		{
			yyVAL.str = AST_GE
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:705
		{
			yyVAL.str = AST_NE
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:709
		{
			yyVAL.str = AST_NSE
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:715
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:719
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:725
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:729
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:735
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:739
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:745
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:751
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:755
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:761
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:765
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:769
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:773
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:777
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:781
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:785
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:789
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:793
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:797
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:801
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:805
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				switch yyDollar[1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
			}
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:820
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 142:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:824
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 143:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:828
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 144:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:832
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:836
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:842
		{
			yyVAL.bytes = IF_BYTES
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:846
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:852
		{
			yyVAL.byt = AST_UPLUS
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:856
		{
			yyVAL.byt = AST_UMINUS
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:860
		{
			yyVAL.byt = AST_TILDA
		}
	case 151:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:866
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:871
		{
			yyVAL.valExpr = nil
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:875
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:881
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 155:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:885
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 156:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:891
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:896
		{
			yyVAL.valExpr = nil
		}
	case 158:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:900
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:906
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:910
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:916
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:920
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:924
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:928
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 165:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:933
		{
			yyVAL.valExprs = nil
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:937
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 167:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:942
		{
			yyVAL.boolExpr = nil
		}
	case 168:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:946
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 169:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:951
		{
			yyVAL.orderBy = nil
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:955
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:961
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:965
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 173:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:971
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 174:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:976
		{
			yyVAL.str = AST_ASC
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:980
		{
			yyVAL.str = AST_ASC
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:984
		{
			yyVAL.str = AST_DESC
		}
	case 177:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:989
		{
			yyVAL.limit = nil
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:993
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 179:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:997
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 180:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:1001
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 181:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1006
		{
			yyVAL.str = ""
		}
	case 182:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:1010
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 183:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:1014
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
	case 184:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1027
		{
			yyVAL.columns = nil
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1031
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1037
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1041
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 188:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1046
		{
			yyVAL.updateExprs = nil
		}
	case 189:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:1050
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1056
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 191:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1060
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1066
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 193:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1071
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:1073
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1076
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1078
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1081
		{
			yyVAL.str = ""
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1083
		{
			yyVAL.str = AST_IGNORE
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1087
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1089
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1091
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1093
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1095
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1098
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1100
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1103
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1105
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1108
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:1110
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1114
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 211:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1119
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
