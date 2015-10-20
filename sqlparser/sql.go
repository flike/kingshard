//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
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

//line sql.y:31
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
const IN = 57363
const IS = 57364
const LIKE = 57365
const BETWEEN = 57366
const NULL = 57367
const ASC = 57368
const DESC = 57369
const VALUES = 57370
const INTO = 57371
const DUPLICATE = 57372
const KEY = 57373
const DEFAULT = 57374
const SET = 57375
const LOCK = 57376
const ID = 57377
const STRING = 57378
const NUMBER = 57379
const VALUE_ARG = 57380
const COMMENT = 57381
const LE = 57382
const GE = 57383
const NE = 57384
const NULL_SAFE_EQUAL = 57385
const UNION = 57386
const MINUS = 57387
const EXCEPT = 57388
const INTERSECT = 57389
const JOIN = 57390
const STRAIGHT_JOIN = 57391
const LEFT = 57392
const RIGHT = 57393
const INNER = 57394
const OUTER = 57395
const CROSS = 57396
const NATURAL = 57397
const USE = 57398
const FORCE = 57399
const ON = 57400
const AND = 57401
const OR = 57402
const NOT = 57403
const UNARY = 57404
const CASE = 57405
const WHEN = 57406
const THEN = 57407
const ELSE = 57408
const END = 57409
const BEGIN = 57410
const START = 57411
const TRANSACTION = 57412
const COMMIT = 57413
const ROLLBACK = 57414
const NAMES = 57415
const REPLACE = 57416
const ADMIN = 57417
const CREATE = 57418
const ALTER = 57419
const DROP = 57420
const RENAME = 57421
const TABLE = 57422
const INDEX = 57423
const VIEW = 57424
const TO = 57425
const IGNORE = 57426
const IF = 57427
const UNIQUE = 57428
const USING = 57429

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
	"IN",
	"IS",
	"LIKE",
	"BETWEEN",
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
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"'('",
	"'='",
	"'<'",
	"'>'",
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
	"AND",
	"OR",
	"NOT",
	"'&'",
	"'|'",
	"'^'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'.'",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"BEGIN",
	"START",
	"TRANSACTION",
	"COMMIT",
	"ROLLBACK",
	"NAMES",
	"REPLACE",
	"ADMIN",
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
const yyMaxDepth = 200

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 209
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 572

var yyAct = [...]int{

	105, 304, 179, 70, 338, 300, 369, 132, 260, 102,
	96, 113, 92, 177, 219, 255, 192, 180, 3, 141,
	72, 91, 88, 378, 378, 108, 201, 103, 84, 77,
	112, 153, 154, 118, 60, 50, 378, 51, 148, 267,
	75, 109, 110, 111, 135, 74, 148, 112, 79, 100,
	118, 81, 62, 116, 73, 85, 349, 75, 109, 110,
	111, 35, 36, 37, 38, 348, 126, 148, 347, 246,
	116, 248, 99, 380, 379, 78, 114, 115, 97, 131,
	217, 80, 52, 119, 56, 125, 377, 139, 330, 134,
	74, 322, 324, 114, 115, 150, 295, 331, 75, 145,
	119, 45, 256, 47, 298, 117, 143, 48, 286, 256,
	176, 178, 153, 154, 181, 249, 152, 293, 182, 128,
	185, 247, 117, 74, 90, 74, 191, 333, 323, 197,
	216, 58, 73, 188, 73, 153, 154, 198, 130, 344,
	195, 196, 301, 236, 189, 194, 263, 211, 53, 54,
	55, 71, 225, 197, 164, 165, 166, 167, 168, 223,
	124, 224, 97, 212, 227, 228, 166, 167, 168, 229,
	215, 301, 234, 235, 68, 238, 239, 240, 241, 242,
	243, 244, 245, 61, 230, 237, 275, 276, 277, 278,
	279, 74, 280, 281, 258, 138, 83, 97, 97, 121,
	73, 123, 264, 346, 262, 251, 253, 143, 259, 257,
	265, 316, 345, 320, 74, 314, 317, 269, 74, 140,
	315, 226, 319, 73, 252, 318, 108, 271, 193, 124,
	143, 112, 268, 223, 118, 272, 285, 35, 36, 37,
	38, 95, 109, 110, 111, 248, 288, 289, 354, 270,
	100, 147, 86, 207, 116, 127, 193, 364, 363, 287,
	222, 362, 292, 126, 302, 213, 97, 299, 303, 221,
	297, 273, 205, 99, 294, 208, 89, 114, 115, 93,
	356, 357, 186, 184, 119, 223, 223, 183, 312, 313,
	275, 276, 277, 278, 279, 148, 280, 281, 309, 124,
	190, 18, 89, 120, 329, 335, 117, 61, 75, 250,
	327, 89, 332, 325, 284, 340, 308, 375, 307, 336,
	339, 151, 161, 162, 163, 164, 165, 166, 167, 168,
	283, 222, 204, 206, 203, 376, 352, 61, 210, 326,
	221, 209, 350, 146, 59, 136, 360, 351, 133, 129,
	358, 82, 122, 359, 334, 361, 87, 195, 18, 367,
	368, 67, 291, 370, 370, 370, 74, 366, 339, 373,
	371, 372, 108, 63, 199, 73, 382, 112, 18, 383,
	118, 142, 39, 384, 137, 385, 65, 95, 109, 110,
	111, 305, 231, 108, 232, 233, 100, 343, 112, 306,
	116, 118, 41, 42, 43, 44, 261, 342, 75, 109,
	110, 111, 311, 193, 57, 69, 18, 100, 40, 99,
	381, 116, 365, 114, 115, 93, 18, 17, 16, 15,
	119, 14, 13, 12, 200, 46, 112, 266, 202, 118,
	99, 49, 76, 144, 114, 115, 75, 109, 110, 111,
	353, 119, 117, 374, 355, 126, 337, 341, 310, 116,
	18, 19, 20, 21, 296, 161, 162, 163, 164, 165,
	166, 167, 168, 117, 156, 160, 158, 159, 187, 254,
	107, 104, 114, 115, 106, 214, 101, 155, 22, 119,
	98, 321, 220, 172, 173, 174, 175, 274, 169, 170,
	171, 161, 162, 163, 164, 165, 166, 167, 168, 218,
	94, 117, 282, 149, 64, 34, 66, 33, 11, 10,
	157, 161, 162, 163, 164, 165, 166, 167, 168, 9,
	8, 7, 6, 5, 4, 2, 1, 0, 27, 28,
	0, 29, 30, 0, 31, 32, 23, 24, 26, 25,
	328, 0, 0, 161, 162, 163, 164, 165, 166, 167,
	168, 290, 0, 0, 161, 162, 163, 164, 165, 166,
	167, 168,
}
var yyPact = [...]int{

	455, -1000, -1000, 188, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 6, -62, -13, 53, -1000, -1, -1000,
	-1000, -1000, 309, 272, 421, 356, -1000, -1000, -1000, 368,
	-1000, 332, 309, 406, 63, -71, -21, 272, -1000, -14,
	272, -1000, 316, -72, 272, -72, -1000, 327, 258, 48,
	-1000, -1000, -1000, -1000, 352, -1000, 264, 309, 319, 309,
	176, 22, -1000, 210, -1000, 43, 314, 71, 272, -1000,
	313, -1000, -54, 310, 364, 131, 272, 309, 353, 273,
	308, 242, -1000, -1000, 302, 40, 70, 453, -1000, 5,
	373, -1000, -1000, -1000, 22, 243, 239, -1000, 238, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 22,
	-1000, 267, 273, 403, 273, 433, 411, 22, 272, -1000,
	354, -76, -1000, 240, -1000, 306, -1000, -1000, 303, -1000,
	232, -1000, 219, 188, 27, -1000, -1000, 225, 352, -1000,
	-1000, 272, 148, 5, 5, 22, 219, 371, 22, 22,
	118, 22, 22, 22, 22, 22, 22, 22, 22, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 453, -34, 18,
	12, 453, -1000, 206, 352, -1000, 421, 30, 433, 353,
	273, 246, 393, 5, -1000, 433, 433, -1000, -1000, -1000,
	82, 272, -1000, -59, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 353, 273, 196, -1000, -1000, 273, 218, 132,
	295, 296, 32, -1000, -1000, -1000, -1000, -1000, -1000, 433,
	-1000, 219, 22, 22, 433, 496, -1000, 337, 83, 83,
	83, 93, 93, -1000, -1000, -1000, -1000, -1000, 22, -1000,
	-1000, 14, 352, -7, 23, -1000, 5, 78, 107, 393,
	376, 385, 70, 283, -1000, -1000, 281, -1000, -1000, 176,
	219, -1000, 401, 225, 225, -1000, -1000, 161, 157, 171,
	168, 159, 29, -1000, 278, 236, 275, -1000, 433, 485,
	22, -1000, 433, -1000, -15, -1000, 15, -1000, 22, 47,
	-1000, 324, -1000, 376, -1000, 22, 22, -1000, -1000, -1000,
	395, 383, 132, 75, -1000, 158, -1000, 149, -1000, -1000,
	-1000, -1000, -28, -31, -40, -1000, -1000, -1000, 22, 433,
	-1000, -1000, 433, 22, 305, -1000, 397, 195, -1000, 254,
	-1000, 393, 5, 22, 5, -1000, -1000, 217, 214, 213,
	433, 433, 415, 22, 22, -1000, -1000, -1000, 376, 70,
	192, 70, 272, 272, 272, 273, 433, -1000, 301, -17,
	-1000, -29, -30, 176, -1000, 413, 355, -1000, 272, -1000,
	-1000, -1000, 272, -1000, 272, -1000,
}
var yyPgo = [...]int{

	0, 536, 535, 17, 534, 533, 532, 531, 530, 529,
	519, 518, 382, 516, 515, 514, 21, 12, 513, 512,
	510, 509, 14, 497, 492, 131, 491, 6, 16, 10,
	490, 487, 19, 486, 13, 27, 2, 485, 484, 11,
	481, 9, 480, 479, 15, 478, 464, 458, 457, 8,
	456, 4, 454, 1, 453, 22, 443, 5, 3, 20,
	196, 442, 441, 438, 437, 435, 434, 0, 7, 433,
	432, 431, 429, 428, 427, 418,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 72, 72, 5, 6, 7, 7, 69, 69,
	70, 71, 73, 74, 8, 8, 8, 9, 9, 9,
	10, 11, 11, 11, 75, 12, 13, 13, 14, 14,
	14, 14, 14, 15, 15, 16, 16, 17, 17, 17,
	20, 20, 18, 18, 18, 21, 21, 22, 22, 22,
	22, 19, 19, 19, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 24, 24, 24, 25, 25, 26, 26,
	26, 26, 27, 27, 28, 28, 29, 29, 29, 29,
	29, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 31, 31, 31, 31, 31, 31, 31, 32, 32,
	37, 37, 35, 35, 39, 36, 36, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 38, 38, 40, 40, 40, 42,
	45, 45, 43, 43, 44, 46, 46, 41, 41, 33,
	33, 33, 33, 47, 47, 48, 48, 49, 49, 50,
	50, 51, 52, 52, 52, 53, 53, 53, 54, 54,
	54, 55, 55, 56, 56, 57, 57, 58, 58, 59,
	60, 60, 61, 61, 62, 62, 63, 63, 63, 63,
	63, 64, 64, 65, 65, 66, 66, 67, 68,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 12, 3,
	7, 7, 6, 6, 8, 7, 3, 4, 1, 2,
	1, 1, 4, 2, 5, 8, 4, 6, 7, 4,
	5, 4, 5, 5, 0, 2, 0, 2, 1, 2,
	1, 1, 1, 0, 1, 1, 3, 1, 2, 3,
	1, 1, 0, 1, 2, 1, 3, 3, 3, 3,
	5, 0, 1, 2, 1, 1, 2, 3, 2, 3,
	2, 2, 2, 1, 3, 1, 1, 3, 0, 5,
	5, 5, 1, 3, 0, 2, 1, 3, 3, 2,
	3, 3, 3, 4, 3, 4, 5, 6, 3, 4,
	2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 3, 3, 1, 3, 1, 3, 1, 1, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 3,
	4, 5, 4, 1, 1, 1, 1, 1, 1, 5,
	0, 1, 1, 2, 4, 0, 2, 1, 3, 1,
	1, 1, 1, 0, 3, 0, 2, 0, 3, 1,
	3, 2, 0, 1, 1, 0, 2, 4, 0, 2,
	4, 0, 3, 1, 3, 0, 5, 1, 3, 3,
	0, 2, 0, 3, 0, 1, 1, 1, 1, 1,
	1, 0, 1, 0, 1, 0, 2, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 33, 91, 92, 94, 93, 83, 84, 86,
	87, 89, 90, 62, -14, 49, 50, 51, 52, -12,
	-75, -12, -12, -12, -12, 95, -65, 97, 101, -62,
	97, 99, 95, 95, 96, 97, 85, -12, -25, 35,
	-67, 35, -3, 17, -15, 18, -13, 29, -25, 9,
	-58, 88, -59, -41, -67, 35, -61, 100, 96, -67,
	95, -67, 35, -60, 100, -67, -60, 29, -55, 44,
	76, -16, -17, 73, -20, 35, -29, -34, -30, 67,
	44, -33, -41, -35, -40, -67, -38, -42, 20, 36,
	37, 38, 25, -39, 71, 72, 48, 100, 28, 78,
	39, -25, 33, -25, 53, -34, 44, 45, 76, 35,
	67, -67, -68, 35, -68, 98, 35, 20, 64, -67,
	-25, -32, 28, -3, -56, -41, 35, 9, 53, -18,
	-67, 19, 76, 65, 66, -31, 21, 67, 23, 24,
	22, 68, 69, 70, 71, 72, 73, 74, 75, 45,
	46, 47, 40, 41, 42, 43, -29, -34, -29, -36,
	-3, -34, -34, 44, 44, -39, 44, -45, -34, -55,
	33, -58, -28, 10, -59, -34, -34, -67, -68, 20,
	-66, 102, -63, 94, 92, 32, 93, 13, 35, 35,
	35, -68, -55, 33, -37, -35, 103, 53, -21, -22,
	-24, 44, 35, -39, -17, -67, 73, -29, -29, -34,
	-35, 21, 23, 24, -34, -34, 25, 67, -34, -34,
	-34, -34, -34, -34, -34, -34, 103, 103, 53, 103,
	103, -16, 18, -16, -43, -44, 79, -32, -58, -28,
	-49, 13, -29, 64, -67, -68, -64, 98, -32, -58,
	53, -41, -28, 53, -23, 54, 55, 56, 57, 58,
	60, 61, -19, 35, 19, -22, 76, -35, -34, -34,
	65, 25, -34, 103, -16, 103, -46, -44, 81, -29,
	-57, 64, -57, -49, -53, 15, 14, 35, 35, -35,
	-47, 11, -22, -22, 54, 59, 54, 59, 54, 54,
	54, -26, 62, 99, 63, 35, 103, 35, 65, -34,
	103, 82, -34, 80, 30, -53, -34, -50, -51, -34,
	-68, -48, 12, 14, 64, 54, 54, 96, 96, 96,
	-34, -34, 31, 53, 53, -52, 26, 27, -49, -29,
	-36, -29, 44, 44, 44, 7, -34, -51, -53, -27,
	-67, -27, -27, -58, -54, 16, 34, 103, 53, 103,
	103, 7, 21, -67, -67, -67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 44, 44,
	44, 44, 44, 203, 194, 0, 0, 28, 0, 30,
	31, 44, 0, 0, 0, 48, 50, 51, 52, 53,
	46, 0, 0, 0, 0, 192, 0, 0, 204, 0,
	0, 195, 0, 190, 0, 190, 29, 0, 181, 86,
	33, 207, 19, 49, 0, 54, 45, 0, 0, 0,
	26, 0, 187, 0, 157, 207, 0, 0, 0, 208,
	0, 208, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 17, 55, 57, 62, 207, 60, 61, 96, 0,
	0, 127, 128, 129, 0, 157, 0, 143, 0, 159,
	160, 161, 162, 123, 146, 147, 148, 144, 145, 150,
	47, 181, 0, 94, 0, 27, 0, 0, 0, 208,
	0, 205, 36, 0, 39, 0, 41, 191, 0, 208,
	181, 32, 0, 119, 0, 183, 87, 0, 0, 58,
	63, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 111,
	112, 113, 114, 115, 116, 117, 99, 0, 0, 0,
	0, 125, 138, 0, 0, 110, 0, 0, 151, 0,
	0, 94, 167, 0, 188, 125, 189, 158, 34, 193,
	0, 0, 208, 201, 196, 197, 198, 199, 200, 40,
	42, 43, 0, 0, 118, 120, 182, 0, 94, 65,
	71, 0, 83, 85, 56, 64, 59, 97, 98, 101,
	102, 0, 0, 0, 104, 0, 108, 0, 130, 131,
	132, 133, 134, 135, 136, 137, 100, 122, 0, 124,
	139, 0, 0, 0, 155, 152, 0, 185, 185, 167,
	175, 0, 95, 0, 206, 37, 0, 202, 22, 23,
	0, 184, 163, 0, 0, 74, 75, 0, 0, 0,
	0, 0, 88, 72, 0, 0, 0, 103, 105, 0,
	0, 109, 126, 140, 0, 142, 0, 153, 0, 0,
	20, 0, 21, 175, 25, 0, 0, 208, 38, 121,
	165, 0, 66, 69, 76, 0, 78, 0, 80, 81,
	82, 67, 0, 0, 0, 73, 68, 84, 0, 106,
	141, 149, 156, 0, 0, 24, 176, 168, 169, 172,
	35, 167, 0, 0, 0, 77, 79, 0, 0, 0,
	107, 154, 0, 0, 0, 171, 173, 174, 175, 166,
	164, 70, 0, 0, 0, 0, 177, 170, 178, 0,
	92, 0, 0, 186, 18, 0, 0, 89, 0, 90,
	91, 179, 0, 93, 0, 180,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 75, 68, 3,
	44, 103, 73, 71, 53, 72, 76, 74, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	46, 45, 47, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 70, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 69, 3, 48,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 49, 50, 51, 52, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	77, 78, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102,
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
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
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
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
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
	if yychar < 0 {
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
		yyVAL = yylval
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
		if yychar < 0 {
			yychar, yytoken = yylex1(yylex, &yylval)
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
			yychar = -1
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
		//line sql.y:163
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:169
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:189
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs}
		}
	case 18:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:193
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:197
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 20:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:204
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 21:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:208
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 22:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:220
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 23:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:224
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
		//line sql.y:237
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:243
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:249
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:253
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: yyDollar[4].valExpr}}}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:259
		{
			yyVAL.statement = &Begin{}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:263
		{
			yyVAL.statement = &Begin{}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:270
		{
			yyVAL.statement = &Commit{}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:276
		{
			yyVAL.statement = &Rollback{}
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:282
		{
			yyVAL.statement = &Admin{Region: yyDollar[2].tableName, Columns: yyDollar[3].columns, Rows: yyDollar[4].insRows}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:288
		{
			yyVAL.statement = &UseDB{DB: string(yyDollar[2].bytes)}
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:294
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 35:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:298
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:303
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 37:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:309
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 38:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:313
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:318
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 40:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:324
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:330
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:334
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:339
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 44:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:344
		{
			SetAllowComments(yylex, true)
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:348
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:354
		{
			yyVAL.bytes2 = nil
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:358
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:364
		{
			yyVAL.str = AST_UNION
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:368
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:372
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:376
		{
			yyVAL.str = AST_EXCEPT
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:380
		{
			yyVAL.str = AST_INTERSECT
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:385
		{
			yyVAL.str = ""
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:389
		{
			yyVAL.str = AST_DISTINCT
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:395
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:399
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:405
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:409
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:413
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:419
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:423
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 62:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:428
		{
			yyVAL.bytes = nil
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:432
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:436
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:442
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:446
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:452
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:456
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:460
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 70:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:464
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 71:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:469
		{
			yyVAL.bytes = nil
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:473
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:477
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:483
		{
			yyVAL.str = AST_JOIN
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:487
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:491
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:495
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:499
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:503
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:507
		{
			yyVAL.str = AST_JOIN
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:511
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:515
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:521
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:525
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:529
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:535
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:539
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 88:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:544
		{
			yyVAL.indexHints = nil
		}
	case 89:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:548
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 90:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:552
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 91:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:556
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:562
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:566
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 94:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:571
		{
			yyVAL.boolExpr = nil
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:575
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:582
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:586
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:590
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:594
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:600
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:604
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 103:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:608
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:612
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:616
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 106:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:620
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 107:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:624
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:628
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 109:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:632
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:636
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:642
		{
			yyVAL.str = AST_EQ
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:646
		{
			yyVAL.str = AST_LT
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:650
		{
			yyVAL.str = AST_GT
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:654
		{
			yyVAL.str = AST_LE
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:658
		{
			yyVAL.str = AST_GE
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:662
		{
			yyVAL.str = AST_NE
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:666
		{
			yyVAL.str = AST_NSE
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:672
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:676
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:682
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:686
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:692
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:696
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:702
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:708
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:712
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:718
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:722
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:726
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:730
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:734
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:738
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:742
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:746
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:750
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:754
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:758
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 138:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:762
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
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:777
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 140:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:781
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 141:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:785
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 142:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:789
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:793
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:799
		{
			yyVAL.bytes = IF_BYTES
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:803
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:809
		{
			yyVAL.byt = AST_UPLUS
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:813
		{
			yyVAL.byt = AST_UMINUS
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:817
		{
			yyVAL.byt = AST_TILDA
		}
	case 149:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:823
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:828
		{
			yyVAL.valExpr = nil
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:832
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:838
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 153:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:842
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:848
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:853
		{
			yyVAL.valExpr = nil
		}
	case 156:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:857
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:863
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:867
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:873
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:877
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:881
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:885
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 163:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:890
		{
			yyVAL.valExprs = nil
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:894
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 165:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:899
		{
			yyVAL.boolExpr = nil
		}
	case 166:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:903
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 167:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:908
		{
			yyVAL.orderBy = nil
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:912
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:918
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:922
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 171:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:928
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 172:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:933
		{
			yyVAL.str = AST_ASC
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:937
		{
			yyVAL.str = AST_ASC
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:941
		{
			yyVAL.str = AST_DESC
		}
	case 175:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:946
		{
			yyVAL.limit = nil
		}
	case 176:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:950
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 177:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:954
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 178:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:959
		{
			yyVAL.str = ""
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:963
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 180:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:967
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
	case 181:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:980
		{
			yyVAL.columns = nil
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:984
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:990
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:994
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 185:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:999
		{
			yyVAL.updateExprs = nil
		}
	case 186:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1003
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1009
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1013
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 189:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1019
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 190:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1024
		{
			yyVAL.empty = struct{}{}
		}
	case 191:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1026
		{
			yyVAL.empty = struct{}{}
		}
	case 192:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1029
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1031
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1034
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1036
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1040
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1042
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1044
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1046
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1048
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1051
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1053
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1056
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1058
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1061
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1063
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1067
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1072
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
