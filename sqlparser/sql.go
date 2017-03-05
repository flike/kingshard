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
const TRUNCATE = 57433

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

const yyNprod = 219
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 647

var yyAct = [...]int{

	111, 108, 315, 352, 187, 385, 102, 139, 74, 346,
	270, 119, 76, 185, 229, 188, 3, 148, 97, 265,
	199, 98, 60, 92, 162, 161, 109, 118, 211, 88,
	124, 81, 394, 53, 394, 63, 279, 79, 115, 116,
	117, 394, 133, 122, 142, 156, 156, 78, 77, 47,
	83, 49, 66, 85, 156, 50, 258, 89, 227, 205,
	363, 362, 114, 118, 125, 361, 124, 72, 82, 37,
	38, 39, 40, 79, 115, 116, 117, 256, 106, 122,
	120, 121, 103, 138, 334, 336, 52, 96, 53, 132,
	343, 146, 84, 141, 78, 152, 129, 396, 105, 395,
	125, 158, 62, 55, 56, 57, 393, 54, 150, 123,
	342, 307, 184, 186, 147, 79, 120, 121, 154, 305,
	189, 257, 58, 226, 190, 298, 193, 274, 160, 78,
	77, 78, 77, 135, 94, 335, 206, 198, 259, 196,
	171, 399, 134, 201, 208, 123, 266, 203, 204, 137,
	197, 243, 162, 161, 221, 174, 175, 176, 171, 207,
	235, 206, 61, 161, 242, 241, 358, 233, 237, 238,
	103, 222, 75, 162, 161, 87, 225, 239, 234, 345,
	244, 245, 246, 248, 249, 250, 251, 252, 253, 254,
	255, 240, 130, 347, 287, 288, 289, 290, 291, 275,
	292, 293, 145, 347, 236, 103, 103, 272, 217, 328,
	261, 263, 276, 266, 329, 310, 273, 247, 360, 269,
	277, 267, 326, 215, 78, 77, 218, 327, 78, 283,
	51, 359, 281, 90, 37, 38, 39, 40, 150, 200,
	280, 332, 331, 233, 330, 130, 297, 200, 262, 284,
	114, 118, 258, 369, 124, 282, 300, 301, 338, 379,
	155, 101, 115, 116, 117, 64, 106, 122, 299, 78,
	77, 285, 304, 311, 71, 318, 103, 313, 378, 130,
	314, 306, 268, 150, 309, 312, 105, 126, 125, 93,
	214, 216, 213, 156, 19, 377, 223, 233, 233, 133,
	324, 325, 232, 93, 120, 121, 99, 231, 194, 321,
	172, 173, 174, 175, 176, 171, 341, 349, 192, 19,
	232, 191, 93, 348, 344, 231, 79, 354, 339, 337,
	350, 353, 296, 123, 114, 118, 320, 260, 124, 391,
	19, 20, 21, 22, 295, 79, 115, 116, 117, 159,
	106, 122, 319, 392, 364, 220, 219, 207, 62, 365,
	153, 64, 375, 374, 23, 376, 373, 143, 140, 136,
	105, 203, 125, 383, 86, 128, 384, 380, 386, 386,
	386, 381, 382, 353, 387, 388, 34, 366, 120, 121,
	78, 77, 127, 91, 19, 400, 303, 209, 397, 144,
	401, 69, 402, 67, 357, 170, 169, 172, 173, 174,
	175, 176, 171, 149, 316, 317, 271, 123, 28, 29,
	356, 30, 31, 367, 32, 33, 202, 323, 200, 24,
	25, 27, 26, 114, 118, 73, 398, 124, 389, 19,
	42, 35, 18, 17, 101, 115, 116, 117, 19, 106,
	122, 16, 15, 170, 169, 172, 173, 174, 175, 176,
	171, 14, 13, 12, 118, 95, 210, 124, 48, 105,
	278, 125, 212, 368, 79, 115, 116, 117, 80, 133,
	122, 151, 390, 118, 370, 351, 124, 120, 121, 99,
	131, 355, 322, 79, 115, 116, 117, 118, 133, 122,
	124, 125, 371, 372, 340, 308, 195, 79, 115, 116,
	117, 264, 133, 122, 113, 110, 123, 120, 121, 112,
	125, 170, 169, 172, 173, 174, 175, 176, 171, 224,
	107, 163, 104, 333, 125, 230, 120, 121, 170, 169,
	172, 173, 174, 175, 176, 171, 123, 286, 228, 302,
	120, 121, 170, 169, 172, 173, 174, 175, 176, 171,
	100, 294, 157, 68, 36, 123, 170, 169, 172, 173,
	174, 175, 176, 171, 165, 167, 70, 11, 10, 123,
	177, 178, 179, 180, 181, 182, 183, 168, 166, 164,
	170, 169, 172, 173, 174, 175, 176, 171, 169, 172,
	173, 174, 175, 176, 171, 287, 288, 289, 290, 291,
	41, 292, 293, 9, 8, 7, 6, 5, 4, 2,
	1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 43, 44, 45, 46, 0, 0, 0, 0, 0,
	0, 0, 0, 59, 0, 0, 65,
}
var yyPact = [...]int{

	335, -1000, -1000, 196, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -49, -14, 9, 5, -1000, 37,
	-1000, -1000, -1000, 71, 234, -1000, 434, 386, -1000, -1000,
	-1000, 383, -1000, -69, 327, 426, 84, -72, -31, 234,
	-1000, -6, 234, -1000, 343, -74, 234, -74, -1000, 368,
	286, -1000, 54, -1000, -1000, -11, -1000, -1000, 413, -1000,
	252, 367, 346, 327, 203, 462, -1000, 80, -1000, 53,
	338, 93, 234, -1000, 337, -1000, -57, 336, 379, 149,
	234, 327, 389, 295, 329, 327, -1000, 251, -1000, -1000,
	330, 48, 98, 518, -1000, 42, 314, -1000, -1000, -1000,
	476, 285, 282, -1000, 272, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 476, -1000, 327, 295, 418,
	295, -1000, 333, 443, 6, 326, -1000, 377, -77, -1000,
	195, -1000, 325, -1000, -1000, 324, -1000, 267, -1000, 263,
	196, 16, -1000, -1000, -1000, 271, 413, -1000, -1000, 234,
	128, 42, 42, 476, 263, 94, 476, 476, 161, 476,
	476, 476, 476, 476, 476, 476, 476, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 518, -30, 14, 31, 518,
	-1000, 230, 413, -1000, 434, 87, 466, 253, 237, 403,
	42, -1000, 476, 466, 466, -1000, -1000, 47, -1000, -1000,
	146, 234, -1000, -65, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 389, 295, 213, -1000, -1000, 295, 229, 562,
	313, 289, 45, -1000, -1000, -1000, -1000, -1000, 108, 466,
	-1000, 263, 476, 476, 466, 494, -1000, 375, 236, 525,
	-1000, 79, 79, 61, 61, 61, -1000, -1000, 476, -1000,
	-1000, 12, 413, 4, 154, -1000, 42, 389, 295, 403,
	399, 401, 98, 466, 234, 321, -1000, -1000, 305, -1000,
	-1000, 203, 263, -1000, 416, 271, 271, -1000, -1000, 179,
	166, 201, 199, 198, 33, -1000, 298, 151, 297, -1000,
	466, 449, 476, -1000, 466, -1000, 3, -1000, 8, -1000,
	476, 119, 140, 150, 399, -1000, 476, 476, -1000, -1000,
	-1000, -1000, 408, 390, 562, 113, -1000, 188, -1000, 175,
	-1000, -1000, -1000, -1000, -34, -38, -39, -1000, -1000, -1000,
	476, 466, -1000, -1000, 466, 476, -1000, 361, -1000, -1000,
	381, 211, -1000, 480, -1000, 403, 42, 476, 42, -1000,
	-1000, 259, 242, 223, 466, 466, 350, 476, 476, 476,
	-1000, -1000, -1000, 399, 98, 210, 98, 234, 234, 234,
	431, 466, 466, -1000, 323, -1, -1000, -8, -10, 295,
	-1000, 429, 70, -1000, 234, -1000, -1000, 203, -1000, 234,
	-1000, 234, -1000,
}
var yyPgo = [...]int{

	0, 620, 619, 15, 618, 617, 616, 615, 614, 613,
	578, 577, 610, 576, 564, 563, 230, 18, 21, 562,
	561, 560, 548, 14, 547, 535, 22, 533, 5, 20,
	6, 532, 531, 17, 530, 13, 26, 4, 529, 519,
	11, 515, 1, 514, 511, 19, 506, 505, 492, 491,
	10, 485, 3, 484, 2, 482, 23, 481, 9, 8,
	12, 175, 478, 472, 470, 468, 466, 0, 7, 465,
	463, 462, 461, 452, 451, 443, 442, 440,
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
	39, 41, 41, 41, 43, 46, 46, 44, 44, 45,
	47, 47, 42, 42, 42, 34, 34, 34, 34, 48,
	48, 49, 49, 50, 50, 51, 51, 52, 53, 53,
	53, 54, 54, 54, 54, 55, 55, 55, 56, 56,
	57, 57, 58, 58, 59, 59, 60, 60, 61, 61,
	62, 62, 16, 16, 63, 63, 63, 63, 63, 64,
	64, 65, 65, 66, 66, 67, 68, 69, 69,
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
	1, 1, 1, 1, 5, 0, 1, 1, 2, 4,
	0, 2, 1, 3, 5, 1, 1, 1, 1, 0,
	3, 0, 2, 0, 3, 1, 3, 2, 0, 1,
	1, 0, 2, 4, 4, 0, 2, 4, 0, 3,
	1, 3, 0, 5, 1, 3, 3, 3, 0, 2,
	0, 3, 0, 1, 1, 1, 1, 1, 1, 0,
	1, 0, 1, 0, 2, 1, 0, 0, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -70, -71, -72, -73, -74, -75, -76, 5,
	6, 7, 8, 29, 94, 95, 97, 96, 83, 84,
	86, 87, 89, 90, 51, 106, -14, 38, 39, 40,
	41, -12, -77, -12, -12, -12, -12, 98, -65, 100,
	104, -16, 100, 102, 98, 98, 99, 100, 85, -12,
	-26, 91, 31, -67, 31, -12, -3, 17, -15, 18,
	-13, -16, -26, 9, -59, 88, -60, -42, -67, 31,
	-62, 103, 99, -67, 98, -67, 31, -61, 103, -67,
	-61, 25, -56, 36, 80, -69, 98, -17, -18, 76,
	-21, 31, -30, -35, -31, 56, 36, -34, -42, -36,
	-41, -67, -39, -43, 20, 32, 33, 34, 21, -40,
	74, 75, 37, 103, 24, 58, 35, 25, 29, -26,
	42, 28, -35, 36, 62, 80, 31, 56, -67, -68,
	31, -68, 101, 31, 20, 53, -67, -26, -33, 24,
	-3, -57, -42, 31, -26, 9, 42, -19, -67, 19,
	80, 55, 54, -32, 71, 56, 70, 57, 69, 73,
	72, 79, 74, 75, 76, 77, 78, 62, 63, 64,
	65, 66, 67, 68, -30, -35, -30, -37, -3, -35,
	-35, 36, 36, -40, 36, -46, -35, -26, -59, -29,
	10, -60, 93, -35, -35, 53, -67, 31, -68, 20,
	-66, 105, -63, 97, 95, 28, 96, 13, 31, 31,
	31, -68, -56, 29, -38, -36, 107, 42, -22, -23,
	-25, 36, 31, -40, -18, -67, 76, -30, -30, -35,
	-36, 71, 70, 57, -35, -35, 21, 56, -35, -35,
	-35, -35, -35, -35, -35, -35, 107, 107, 42, 107,
	107, -17, 18, -17, -44, -45, 59, -56, 29, -29,
	-50, 13, -30, -35, 80, 53, -67, -68, -64, 101,
	-33, -59, 42, -42, -29, 42, -24, 43, 44, 45,
	46, 47, 49, 50, -20, 31, 19, -23, 80, -36,
	-35, -35, 55, 21, -35, 107, -17, 107, -47, -45,
	61, -30, -33, -59, -50, -54, 15, 14, -67, 31,
	31, -36, -48, 11, -23, -23, 43, 48, 43, 48,
	43, 43, 43, -27, 51, 102, 52, 31, 107, 31,
	55, -35, 107, 82, -35, 60, -58, 53, -58, -54,
	-35, -51, -52, -35, -68, -49, 12, 14, 53, 43,
	43, 99, 99, 99, -35, -35, 26, 42, 92, 42,
	-53, 22, 23, -50, -30, -37, -30, 36, 36, 36,
	27, -35, -35, -52, -54, -28, -67, -28, -28, 7,
	-55, 16, 30, 107, 42, 107, 107, -59, 7, 71,
	-67, -67, -67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 49,
	49, 49, 49, 49, 211, 202, 0, 0, 31, 0,
	33, 34, 49, 0, 0, 49, 0, 53, 55, 56,
	57, 58, 51, 202, 0, 0, 0, 200, 0, 0,
	212, 0, 0, 203, 0, 198, 0, 198, 32, 0,
	188, 36, 91, 37, 215, 217, 20, 54, 0, 59,
	50, 0, 0, 0, 27, 0, 194, 0, 162, 215,
	0, 0, 0, 216, 0, 216, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 218, 18, 60, 62,
	67, 215, 65, 66, 101, 0, 0, 132, 133, 134,
	0, 162, 0, 148, 0, 165, 166, 167, 168, 128,
	151, 152, 153, 149, 150, 155, 52, 0, 0, 99,
	0, 28, 29, 0, 0, 0, 216, 0, 213, 41,
	0, 44, 0, 46, 199, 0, 216, 188, 35, 0,
	124, 0, 190, 92, 38, 0, 0, 63, 68, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 116, 117, 118,
	119, 120, 121, 122, 104, 0, 0, 0, 0, 130,
	143, 0, 0, 115, 0, 0, 156, 188, 99, 173,
	0, 195, 0, 130, 196, 197, 163, 215, 39, 201,
	0, 0, 216, 209, 204, 205, 206, 207, 208, 45,
	47, 48, 0, 0, 123, 125, 189, 0, 99, 70,
	76, 0, 88, 90, 61, 69, 64, 102, 103, 106,
	107, 0, 0, 0, 109, 0, 113, 0, 135, 136,
	137, 138, 139, 140, 141, 142, 105, 127, 0, 129,
	144, 0, 0, 0, 160, 157, 0, 0, 0, 173,
	181, 0, 100, 30, 0, 0, 214, 42, 0, 210,
	23, 24, 0, 191, 169, 0, 0, 79, 80, 0,
	0, 0, 0, 0, 93, 77, 0, 0, 0, 108,
	110, 0, 0, 114, 131, 145, 0, 147, 0, 158,
	0, 0, 192, 192, 181, 26, 0, 0, 164, 216,
	43, 126, 171, 0, 71, 74, 81, 0, 83, 0,
	85, 86, 87, 72, 0, 0, 0, 78, 73, 89,
	0, 111, 146, 154, 161, 0, 21, 0, 22, 25,
	182, 174, 175, 178, 40, 173, 0, 0, 0, 82,
	84, 0, 0, 0, 112, 159, 0, 0, 0, 0,
	177, 179, 180, 181, 172, 170, 75, 0, 0, 0,
	0, 183, 184, 176, 185, 0, 97, 0, 0, 0,
	19, 0, 0, 94, 0, 95, 96, 193, 186, 0,
	98, 0, 187,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 78, 73, 3,
	36, 107, 76, 74, 42, 75, 80, 77, 3, 3,
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
		//line ./sqlparser/sql.y:193
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:199
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 18:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:220
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs}
		}
	case 19:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line ./sqlparser/sql.y:224
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:228
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 21:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./sqlparser/sql.y:235
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./sqlparser/sql.y:239
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
		//line ./sqlparser/sql.y:251
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:255
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
		//line ./sqlparser/sql.y:268
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 26:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./sqlparser/sql.y:274
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:280
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:284
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal("default")}}}
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:288
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: yyDollar[4].valExpr}}}
		}
	case 30:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:292
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
		//line ./sqlparser/sql.y:308
		{
			yyVAL.statement = &Begin{}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:312
		{
			yyVAL.statement = &Begin{}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:319
		{
			yyVAL.statement = &Commit{}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:325
		{
			yyVAL.statement = &Rollback{}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:331
		{
			yyVAL.statement = &Admin{Region: yyDollar[2].tableName, Columns: yyDollar[3].columns, Rows: yyDollar[4].insRows}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:335
		{
			yyVAL.statement = &AdminHelp{}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:341
		{
			yyVAL.statement = &UseDB{DB: string(yyDollar[2].bytes)}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:347
		{
			yyVAL.statement = &Truncate{Comments: Comments(yyDollar[2].bytes2), TableOpt: yyDollar[3].str, Table: yyDollar[4].tableName}
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:353
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 40:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./sqlparser/sql.y:357
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:362
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:368
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 43:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./sqlparser/sql.y:372
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:377
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 45:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:383
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:389
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 47:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:393
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 48:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:398
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:403
		{
			SetAllowComments(yylex, true)
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:407
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:413
		{
			yyVAL.bytes2 = nil
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:417
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:423
		{
			yyVAL.str = AST_UNION
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:427
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:431
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:435
		{
			yyVAL.str = AST_EXCEPT
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:439
		{
			yyVAL.str = AST_INTERSECT
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:444
		{
			yyVAL.str = ""
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:448
		{
			yyVAL.str = AST_DISTINCT
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:454
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:458
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:464
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:468
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:472
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:478
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:482
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:487
		{
			yyVAL.bytes = nil
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:491
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:495
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:501
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:505
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:511
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:515
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:519
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 75:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:523
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 76:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:528
		{
			yyVAL.bytes = nil
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:532
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:536
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:542
		{
			yyVAL.str = AST_JOIN
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:546
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:550
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:554
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:558
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:562
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:566
		{
			yyVAL.str = AST_JOIN
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:570
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:574
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:580
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:584
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:588
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:594
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:598
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 93:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:603
		{
			yyVAL.indexHints = nil
		}
	case 94:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:607
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 95:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:611
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 96:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:615
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:621
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:625
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:630
		{
			yyVAL.boolExpr = nil
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:634
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:641
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:645
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:649
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:653
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:659
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:663
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 108:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:667
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:671
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:675
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 111:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:679
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 112:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:683
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:687
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 114:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:691
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:695
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:701
		{
			yyVAL.str = AST_EQ
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:705
		{
			yyVAL.str = AST_LT
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:709
		{
			yyVAL.str = AST_GT
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:713
		{
			yyVAL.str = AST_LE
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:717
		{
			yyVAL.str = AST_GE
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:721
		{
			yyVAL.str = AST_NE
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:725
		{
			yyVAL.str = AST_NSE
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:731
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:735
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:741
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:745
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:751
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:755
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:761
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:767
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:771
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:777
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:781
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:785
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:789
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:793
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:797
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:801
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:805
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:809
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:813
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:817
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 143:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:821
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
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:836
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 145:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:840
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 146:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:844
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 147:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:848
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:852
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:858
		{
			yyVAL.bytes = IF_BYTES
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:862
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:868
		{
			yyVAL.byt = AST_UPLUS
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:872
		{
			yyVAL.byt = AST_UMINUS
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:876
		{
			yyVAL.byt = AST_TILDA
		}
	case 154:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:882
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:887
		{
			yyVAL.valExpr = nil
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:891
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:897
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 158:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:901
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 159:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:907
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:912
		{
			yyVAL.valExpr = nil
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:916
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:922
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:926
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 164:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:930
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[3].bytes, Name: yyDollar[5].bytes}
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:936
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:940
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:944
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:948
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 169:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:953
		{
			yyVAL.valExprs = nil
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:957
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 171:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:962
		{
			yyVAL.boolExpr = nil
		}
	case 172:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:966
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 173:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:971
		{
			yyVAL.orderBy = nil
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:975
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:981
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:985
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 177:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:991
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 178:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:996
		{
			yyVAL.str = AST_ASC
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1000
		{
			yyVAL.str = AST_ASC
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1004
		{
			yyVAL.str = AST_DESC
		}
	case 181:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1009
		{
			yyVAL.limit = nil
		}
	case 182:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:1013
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 183:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:1017
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 184:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:1021
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 185:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1026
		{
			yyVAL.str = ""
		}
	case 186:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:1030
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 187:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:1034
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
	case 188:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1047
		{
			yyVAL.columns = nil
		}
	case 189:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1051
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1057
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 191:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1061
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 192:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1066
		{
			yyVAL.updateExprs = nil
		}
	case 193:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:1070
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1076
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 195:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1080
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1086
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 197:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1090
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: StrVal("ON")}
		}
	case 198:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1095
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:1097
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1100
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1102
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1105
		{
			yyVAL.str = ""
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1107
		{
			yyVAL.str = AST_IGNORE
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1111
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1113
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1115
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1117
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1119
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1122
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1124
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1127
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1129
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1132
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:1134
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1138
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1143
		{
			ForceEOF(yylex)
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1148
		{
			yyVAL.str = ""
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1152
		{
			yyVAL.str = AST_TABLE
		}
	}
	goto yystack /* stack new state and value */
}
