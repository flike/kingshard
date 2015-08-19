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
const COMMIT = 57411
const ROLLBACK = 57412
const NAMES = 57413
const REPLACE = 57414
const ADMIN = 57415
const CREATE = 57416
const ALTER = 57417
const DROP = 57418
const RENAME = 57419
const TABLE = 57420
const INDEX = 57421
const VIEW = 57422
const TO = 57423
const IGNORE = 57424
const IF = 57425
const UNIQUE = 57426
const USING = 57427

var yyToknames = []string{
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
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 208
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 563

var yyAct = []int{

	103, 100, 174, 302, 367, 336, 176, 68, 298, 258,
	129, 111, 89, 253, 138, 216, 86, 101, 70, 190,
	198, 150, 151, 90, 376, 376, 177, 3, 34, 35,
	36, 37, 82, 58, 376, 94, 273, 274, 275, 276,
	277, 75, 278, 279, 72, 71, 145, 77, 145, 250,
	79, 106, 145, 49, 83, 50, 110, 243, 245, 116,
	60, 265, 214, 132, 347, 95, 93, 107, 108, 109,
	346, 345, 378, 377, 76, 98, 78, 128, 51, 114,
	246, 127, 375, 324, 329, 136, 254, 284, 72, 142,
	131, 73, 149, 147, 328, 204, 293, 254, 97, 296,
	291, 178, 112, 113, 91, 179, 244, 59, 125, 117,
	213, 320, 322, 140, 202, 150, 151, 205, 183, 88,
	186, 72, 71, 72, 71, 342, 194, 193, 189, 115,
	331, 299, 248, 173, 175, 56, 187, 195, 44, 261,
	46, 192, 69, 135, 47, 223, 321, 208, 95, 222,
	194, 52, 53, 54, 209, 226, 220, 212, 231, 232,
	344, 235, 236, 237, 238, 239, 240, 241, 242, 221,
	343, 227, 201, 203, 200, 233, 318, 66, 161, 162,
	163, 164, 165, 247, 95, 95, 224, 225, 317, 72,
	71, 163, 164, 165, 249, 251, 256, 150, 151, 262,
	314, 119, 255, 121, 122, 315, 316, 191, 122, 257,
	263, 72, 71, 81, 140, 72, 269, 234, 267, 299,
	312, 137, 245, 106, 266, 313, 191, 260, 110, 144,
	220, 116, 286, 287, 283, 270, 140, 352, 93, 107,
	108, 109, 34, 35, 36, 37, 285, 98, 290, 268,
	271, 114, 326, 95, 124, 158, 159, 160, 161, 162,
	163, 164, 165, 292, 118, 300, 295, 301, 84, 122,
	97, 210, 362, 145, 112, 113, 91, 361, 18, 360,
	18, 117, 87, 220, 220, 219, 307, 310, 311, 180,
	297, 327, 184, 188, 218, 106, 182, 181, 87, 330,
	110, 115, 282, 116, 87, 333, 334, 337, 219, 59,
	73, 107, 108, 109, 373, 73, 338, 218, 281, 98,
	148, 325, 323, 114, 273, 274, 275, 276, 277, 348,
	278, 279, 374, 380, 349, 306, 59, 305, 207, 206,
	143, 57, 97, 133, 247, 130, 112, 113, 358, 356,
	126, 123, 80, 117, 364, 337, 120, 350, 365, 332,
	366, 368, 368, 368, 72, 71, 369, 370, 85, 65,
	18, 371, 289, 115, 106, 196, 357, 381, 359, 110,
	18, 382, 116, 383, 38, 228, 134, 229, 230, 73,
	107, 108, 109, 139, 63, 61, 303, 341, 98, 304,
	110, 259, 114, 116, 40, 41, 42, 43, 340, 309,
	73, 107, 108, 109, 191, 55, 67, 110, 379, 180,
	116, 97, 363, 114, 18, 112, 113, 73, 107, 108,
	109, 39, 117, 17, 16, 15, 180, 14, 13, 12,
	114, 18, 19, 20, 21, 197, 112, 113, 45, 264,
	199, 48, 115, 117, 74, 153, 157, 155, 156, 141,
	372, 353, 335, 112, 113, 339, 308, 294, 185, 22,
	117, 252, 105, 115, 169, 170, 171, 172, 102, 166,
	167, 168, 104, 211, 99, 152, 96, 354, 355, 319,
	115, 217, 272, 215, 92, 351, 280, 146, 32, 62,
	33, 154, 158, 159, 160, 161, 162, 163, 164, 165,
	158, 159, 160, 161, 162, 163, 164, 165, 64, 27,
	28, 29, 11, 30, 31, 23, 24, 26, 25, 158,
	159, 160, 161, 162, 163, 164, 165, 288, 10, 9,
	158, 159, 160, 161, 162, 163, 164, 165, 158, 159,
	160, 161, 162, 163, 164, 165, 8, 7, 6, 5,
	4, 2, 1,
}
var yyPact = []int{

	436, -1000, -1000, 193, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 45, -42, -15, 58, -1000, -1000, -1000,
	-1000, 306, 274, 419, 378, -1000, -1000, -1000, 376, -1000,
	340, 306, 407, 56, -57, -20, 274, -1000, -17, 274,
	-1000, 317, -66, 274, -66, 339, 254, 43, -1000, -1000,
	-1000, -1000, 203, -1000, 225, 306, 323, 306, 151, 316,
	-1000, 209, -1000, 32, 315, 14, 274, -1000, 310, -1000,
	-33, 308, 366, 79, 274, 306, 365, 280, 305, 220,
	-1000, -1000, 301, 16, 132, 434, -1000, 354, 275, -1000,
	-1000, -1000, 392, 253, 252, -1000, 248, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 392, -1000, 260,
	280, 404, 280, -1000, 392, 274, -1000, 355, -80, -1000,
	82, -1000, 304, -1000, -1000, 303, -1000, 238, -1000, 245,
	193, 9, -1000, -1000, 250, 203, -1000, -1000, 274, 72,
	354, 354, 392, 245, 364, 392, 392, 150, 392, 392,
	392, 392, 392, 392, 392, 392, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 434, -44, 5, -21, 434, -1000,
	375, 31, 203, -1000, 419, 7, 480, 365, 280, 216,
	388, 354, -1000, 480, -1000, -1000, -1000, 75, 274, -1000,
	-35, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 365,
	280, 196, -1000, -1000, 280, 197, 270, 283, 273, 11,
	-1000, -1000, -1000, -1000, -1000, -1000, 480, -1000, 245, 392,
	392, 480, 472, -1000, 347, 107, 107, 107, 118, 118,
	-1000, -1000, -1000, -1000, -1000, 392, -1000, 480, -1000, -1,
	203, -5, 18, -1000, 354, 67, 155, 388, 381, 385,
	132, 302, -1000, -1000, 300, -1000, -1000, 151, 245, -1000,
	398, 250, 250, -1000, -1000, 166, 146, 152, 134, 122,
	49, -1000, 287, -18, 286, -1000, 480, 187, 392, -1000,
	480, -1000, -7, -1000, 2, -1000, 392, 50, -1000, 329,
	-1000, 381, -1000, 392, 392, -1000, -1000, -1000, 396, 383,
	270, 61, -1000, 116, -1000, 106, -1000, -1000, -1000, -1000,
	-23, -24, -30, -1000, -1000, -1000, 392, 480, -1000, -1000,
	480, 392, 326, -1000, 442, 184, -1000, 461, -1000, 388,
	354, 392, 354, -1000, -1000, 235, 233, 228, 480, 480,
	415, 392, 392, -1000, -1000, -1000, 381, 132, 169, 132,
	274, 274, 274, 280, 480, -1000, 298, -19, -1000, -28,
	-29, 151, -1000, 411, 312, -1000, 274, -1000, -1000, -1000,
	274, -1000, 274, -1000,
}
var yyPgo = []int{

	0, 562, 561, 26, 560, 559, 558, 557, 556, 539,
	538, 522, 384, 518, 500, 499, 12, 23, 497, 496,
	494, 493, 15, 492, 491, 135, 489, 4, 19, 35,
	486, 485, 14, 484, 2, 17, 6, 483, 482, 11,
	478, 1, 472, 471, 13, 468, 467, 466, 465, 9,
	462, 5, 461, 3, 460, 16, 459, 8, 7, 18,
	213, 454, 451, 450, 449, 448, 445, 0, 10, 439,
	438, 437, 435, 434, 433, 431,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 72, 72, 5, 6, 7, 7, 69, 70,
	71, 73, 74, 8, 8, 8, 9, 9, 9, 10,
	11, 11, 11, 75, 12, 13, 13, 14, 14, 14,
	14, 14, 15, 15, 16, 16, 17, 17, 17, 20,
	20, 18, 18, 18, 21, 21, 22, 22, 22, 22,
	19, 19, 19, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 24, 24, 24, 25, 25, 26, 26, 26,
	26, 27, 27, 28, 28, 29, 29, 29, 29, 29,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	31, 31, 31, 31, 31, 31, 31, 32, 32, 37,
	37, 35, 35, 39, 36, 36, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 38, 38, 40, 40, 40, 42, 45,
	45, 43, 43, 44, 46, 46, 41, 41, 33, 33,
	33, 33, 47, 47, 48, 48, 49, 49, 50, 50,
	51, 52, 52, 52, 53, 53, 53, 54, 54, 54,
	55, 55, 56, 56, 57, 57, 58, 58, 59, 60,
	60, 61, 61, 62, 62, 63, 63, 63, 63, 63,
	64, 64, 65, 65, 66, 66, 67, 68,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 12, 3,
	7, 7, 6, 6, 8, 7, 3, 4, 1, 1,
	1, 4, 2, 5, 8, 4, 6, 7, 4, 5,
	4, 5, 5, 0, 2, 0, 2, 1, 2, 1,
	1, 1, 0, 1, 1, 3, 1, 2, 3, 1,
	1, 0, 1, 2, 1, 3, 3, 3, 3, 5,
	0, 1, 2, 1, 1, 2, 3, 2, 3, 2,
	2, 2, 1, 3, 1, 1, 3, 0, 5, 5,
	5, 1, 3, 0, 2, 1, 3, 3, 2, 3,
	3, 3, 4, 3, 4, 5, 6, 3, 4, 2,
	1, 1, 1, 1, 1, 1, 1, 2, 1, 1,
	3, 3, 1, 3, 1, 3, 1, 1, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 3, 4,
	5, 4, 1, 1, 1, 1, 1, 1, 5, 0,
	1, 1, 2, 4, 0, 2, 1, 3, 1, 1,
	1, 1, 0, 3, 0, 2, 0, 3, 1, 3,
	2, 0, 1, 1, 0, 2, 4, 0, 2, 4,
	0, 3, 1, 3, 0, 5, 1, 3, 3, 0,
	2, 0, 3, 0, 1, 1, 1, 1, 1, 1,
	0, 1, 0, 1, 0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 33, 89, 90, 92, 91, 83, 84, 85,
	87, 88, 62, -14, 49, 50, 51, 52, -12, -75,
	-12, -12, -12, -12, 93, -65, 95, 99, -62, 95,
	97, 93, 93, 94, 95, -12, -25, 35, -67, 35,
	-3, 17, -15, 18, -13, 29, -25, 9, -58, 86,
	-59, -41, -67, 35, -61, 98, 94, -67, 93, -67,
	35, -60, 98, -67, -60, 29, -55, 44, 76, -16,
	-17, 73, -20, 35, -29, -34, -30, 67, 44, -33,
	-41, -35, -40, -67, -38, -42, 20, 36, 37, 38,
	25, -39, 71, 72, 48, 98, 28, 78, 39, -25,
	33, -25, 53, 35, 45, 76, 35, 67, -67, -68,
	35, -68, 96, 35, 20, 64, -67, -25, -32, 28,
	-3, -56, -41, 35, 9, 53, -18, -67, 19, 76,
	65, 66, -31, 21, 67, 23, 24, 22, 68, 69,
	70, 71, 72, 73, 74, 75, 45, 46, 47, 40,
	41, 42, 43, -29, -34, -29, -36, -3, -34, -34,
	44, 44, 44, -39, 44, -45, -34, -55, 33, -58,
	-28, 10, -59, -34, -67, -68, 20, -66, 100, -63,
	92, 90, 32, 91, 13, 35, 35, 35, -68, -55,
	33, -37, -35, 101, 53, -21, -22, -24, 44, 35,
	-39, -17, -67, 73, -29, -29, -34, -35, 21, 23,
	24, -34, -34, 25, 67, -34, -34, -34, -34, -34,
	-34, -34, -34, 101, 101, 53, 101, -34, 101, -16,
	18, -16, -43, -44, 79, -32, -58, -28, -49, 13,
	-29, 64, -67, -68, -64, 96, -32, -58, 53, -41,
	-28, 53, -23, 54, 55, 56, 57, 58, 60, 61,
	-19, 35, 19, -22, 76, -35, -34, -34, 65, 25,
	-34, 101, -16, 101, -46, -44, 81, -29, -57, 64,
	-57, -49, -53, 15, 14, 35, 35, -35, -47, 11,
	-22, -22, 54, 59, 54, 59, 54, 54, 54, -26,
	62, 97, 63, 35, 101, 35, 65, -34, 101, 82,
	-34, 80, 30, -53, -34, -50, -51, -34, -68, -48,
	12, 14, 64, 54, 54, 94, 94, 94, -34, -34,
	31, 53, 53, -52, 26, 27, -49, -29, -36, -29,
	44, 44, 44, 7, -34, -51, -53, -27, -67, -27,
	-27, -58, -54, 16, 34, 101, 53, 101, 101, 7,
	21, -67, -67, -67,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 43, 43,
	43, 43, 43, 202, 193, 0, 0, 28, 29, 30,
	43, 0, 0, 0, 47, 49, 50, 51, 52, 45,
	0, 0, 0, 0, 191, 0, 0, 203, 0, 0,
	194, 0, 189, 0, 189, 0, 180, 85, 32, 206,
	19, 48, 0, 53, 44, 0, 0, 0, 26, 0,
	186, 0, 156, 206, 0, 0, 0, 207, 0, 207,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 17,
	54, 56, 61, 206, 59, 60, 95, 0, 0, 126,
	127, 128, 0, 156, 0, 142, 0, 158, 159, 160,
	161, 122, 145, 146, 147, 143, 144, 149, 46, 180,
	0, 93, 0, 27, 0, 0, 207, 0, 204, 35,
	0, 38, 0, 40, 190, 0, 207, 180, 31, 0,
	118, 0, 182, 86, 0, 0, 57, 62, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 110, 111, 112, 113,
	114, 115, 116, 98, 0, 0, 0, 0, 124, 137,
	0, 0, 0, 109, 0, 0, 150, 0, 0, 93,
	166, 0, 187, 188, 157, 33, 192, 0, 0, 207,
	200, 195, 196, 197, 198, 199, 39, 41, 42, 0,
	0, 117, 119, 181, 0, 93, 64, 70, 0, 82,
	84, 55, 63, 58, 96, 97, 100, 101, 0, 0,
	0, 103, 0, 107, 0, 129, 130, 131, 132, 133,
	134, 135, 136, 99, 121, 0, 123, 124, 138, 0,
	0, 0, 154, 151, 0, 184, 184, 166, 174, 0,
	94, 0, 205, 36, 0, 201, 22, 23, 0, 183,
	162, 0, 0, 73, 74, 0, 0, 0, 0, 0,
	87, 71, 0, 0, 0, 102, 104, 0, 0, 108,
	125, 139, 0, 141, 0, 152, 0, 0, 20, 0,
	21, 174, 25, 0, 0, 207, 37, 120, 164, 0,
	65, 68, 75, 0, 77, 0, 79, 80, 81, 66,
	0, 0, 0, 72, 67, 83, 0, 105, 140, 148,
	155, 0, 0, 24, 175, 167, 168, 171, 34, 166,
	0, 0, 0, 76, 78, 0, 0, 0, 106, 153,
	0, 0, 0, 170, 172, 173, 174, 165, 163, 69,
	0, 0, 0, 0, 176, 169, 177, 0, 91, 0,
	0, 185, 18, 0, 0, 88, 0, 89, 90, 178,
	0, 92, 0, 179,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 75, 68, 3,
	44, 101, 73, 71, 53, 72, 76, 74, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	46, 45, 47, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 70, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 69, 3, 48,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 49, 50, 51, 52, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	77, 78, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
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

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
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
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
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
			yychar = yylex1(yylex, &yylval)
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
			if yyn < 0 || yyn == yychar {
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
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
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
		//line sql.y:163
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:169
		{
			yyVAL.statement = yyS[yypt-0].selStmt
		}
	case 3:
		yyVAL.statement = yyS[yypt-0].statement
	case 4:
		yyVAL.statement = yyS[yypt-0].statement
	case 5:
		yyVAL.statement = yyS[yypt-0].statement
	case 6:
		yyVAL.statement = yyS[yypt-0].statement
	case 7:
		yyVAL.statement = yyS[yypt-0].statement
	case 8:
		yyVAL.statement = yyS[yypt-0].statement
	case 9:
		yyVAL.statement = yyS[yypt-0].statement
	case 10:
		yyVAL.statement = yyS[yypt-0].statement
	case 11:
		yyVAL.statement = yyS[yypt-0].statement
	case 12:
		yyVAL.statement = yyS[yypt-0].statement
	case 13:
		yyVAL.statement = yyS[yypt-0].statement
	case 14:
		yyVAL.statement = yyS[yypt-0].statement
	case 15:
		yyVAL.statement = yyS[yypt-0].statement
	case 16:
		yyVAL.statement = yyS[yypt-0].statement
	case 17:
		//line sql.y:189
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyS[yypt-2].bytes2), Distinct: yyS[yypt-1].str, SelectExprs: yyS[yypt-0].selectExprs}
		}
	case 18:
		//line sql.y:193
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 19:
		//line sql.y:197
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 20:
		//line sql.y:204
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 21:
		//line sql.y:208
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 22:
		//line sql.y:220
		{
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: yyS[yypt-1].columns, Rows: yyS[yypt-0].insRows}
		}
	case 23:
		//line sql.y:224
		{
			cols := make(Columns, 0, len(yyS[yypt-0].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-0].updateExprs))
			for _, col := range yyS[yypt-0].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 24:
		//line sql.y:237
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 25:
		//line sql.y:243
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 26:
		//line sql.y:249
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 27:
		//line sql.y:253
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyS[yypt-0].bytes)}}}
		}
	case 28:
		//line sql.y:259
		{
			yyVAL.statement = &Begin{}
		}
	case 29:
		//line sql.y:265
		{
			yyVAL.statement = &Commit{}
		}
	case 30:
		//line sql.y:271
		{
			yyVAL.statement = &Rollback{}
		}
	case 31:
		//line sql.y:277
		{
			yyVAL.statement = &Admin{Region: yyS[yypt-2].tableName, Columns: yyS[yypt-1].columns, Rows: yyS[yypt-0].insRows}
		}
	case 32:
		//line sql.y:283
		{
			yyVAL.statement = &UseDB{DB: string(yyS[yypt-0].bytes)}
		}
	case 33:
		//line sql.y:289
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 34:
		//line sql.y:293
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 35:
		//line sql.y:298
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 36:
		//line sql.y:304
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 37:
		//line sql.y:308
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 38:
		//line sql.y:313
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 39:
		//line sql.y:319
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 40:
		//line sql.y:325
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 41:
		//line sql.y:329
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 42:
		//line sql.y:334
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 43:
		//line sql.y:339
		{
			SetAllowComments(yylex, true)
		}
	case 44:
		//line sql.y:343
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 45:
		//line sql.y:349
		{
			yyVAL.bytes2 = nil
		}
	case 46:
		//line sql.y:353
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 47:
		//line sql.y:359
		{
			yyVAL.str = AST_UNION
		}
	case 48:
		//line sql.y:363
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 49:
		//line sql.y:367
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 50:
		//line sql.y:371
		{
			yyVAL.str = AST_EXCEPT
		}
	case 51:
		//line sql.y:375
		{
			yyVAL.str = AST_INTERSECT
		}
	case 52:
		//line sql.y:380
		{
			yyVAL.str = ""
		}
	case 53:
		//line sql.y:384
		{
			yyVAL.str = AST_DISTINCT
		}
	case 54:
		//line sql.y:390
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 55:
		//line sql.y:394
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 56:
		//line sql.y:400
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 57:
		//line sql.y:404
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 58:
		//line sql.y:408
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 59:
		//line sql.y:414
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 60:
		//line sql.y:418
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 61:
		//line sql.y:423
		{
			yyVAL.bytes = nil
		}
	case 62:
		//line sql.y:427
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 63:
		//line sql.y:431
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 64:
		//line sql.y:437
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 65:
		//line sql.y:441
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 66:
		//line sql.y:447
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 67:
		//line sql.y:451
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 68:
		//line sql.y:455
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 69:
		//line sql.y:459
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 70:
		//line sql.y:464
		{
			yyVAL.bytes = nil
		}
	case 71:
		//line sql.y:468
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 72:
		//line sql.y:472
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 73:
		//line sql.y:478
		{
			yyVAL.str = AST_JOIN
		}
	case 74:
		//line sql.y:482
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 75:
		//line sql.y:486
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 76:
		//line sql.y:490
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 77:
		//line sql.y:494
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 78:
		//line sql.y:498
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 79:
		//line sql.y:502
		{
			yyVAL.str = AST_JOIN
		}
	case 80:
		//line sql.y:506
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 81:
		//line sql.y:510
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 82:
		//line sql.y:516
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 83:
		//line sql.y:520
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 84:
		//line sql.y:524
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 85:
		//line sql.y:530
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 86:
		//line sql.y:534
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 87:
		//line sql.y:539
		{
			yyVAL.indexHints = nil
		}
	case 88:
		//line sql.y:543
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 89:
		//line sql.y:547
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 90:
		//line sql.y:551
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 91:
		//line sql.y:557
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 92:
		//line sql.y:561
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 93:
		//line sql.y:566
		{
			yyVAL.boolExpr = nil
		}
	case 94:
		//line sql.y:570
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 95:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 96:
		//line sql.y:577
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 97:
		//line sql.y:581
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 98:
		//line sql.y:585
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 99:
		//line sql.y:589
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 100:
		//line sql.y:595
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 101:
		//line sql.y:599
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].tuple}
		}
	case 102:
		//line sql.y:603
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].tuple}
		}
	case 103:
		//line sql.y:607
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 104:
		//line sql.y:611
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 105:
		//line sql.y:615
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 106:
		//line sql.y:619
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 107:
		//line sql.y:623
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 108:
		//line sql.y:627
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 109:
		//line sql.y:631
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 110:
		//line sql.y:637
		{
			yyVAL.str = AST_EQ
		}
	case 111:
		//line sql.y:641
		{
			yyVAL.str = AST_LT
		}
	case 112:
		//line sql.y:645
		{
			yyVAL.str = AST_GT
		}
	case 113:
		//line sql.y:649
		{
			yyVAL.str = AST_LE
		}
	case 114:
		//line sql.y:653
		{
			yyVAL.str = AST_GE
		}
	case 115:
		//line sql.y:657
		{
			yyVAL.str = AST_NE
		}
	case 116:
		//line sql.y:661
		{
			yyVAL.str = AST_NSE
		}
	case 117:
		//line sql.y:667
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 118:
		//line sql.y:671
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 119:
		//line sql.y:677
		{
			yyVAL.values = Values{yyS[yypt-0].tuple}
		}
	case 120:
		//line sql.y:681
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].tuple)
		}
	case 121:
		//line sql.y:687
		{
			yyVAL.tuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 122:
		//line sql.y:691
		{
			yyVAL.tuple = yyS[yypt-0].subquery
		}
	case 123:
		//line sql.y:697
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 124:
		//line sql.y:703
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 125:
		//line sql.y:707
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 126:
		//line sql.y:713
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 127:
		//line sql.y:717
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 128:
		//line sql.y:721
		{
			yyVAL.valExpr = yyS[yypt-0].tuple
		}
	case 129:
		//line sql.y:725
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 130:
		//line sql.y:729
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 131:
		//line sql.y:733
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 132:
		//line sql.y:737
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 133:
		//line sql.y:741
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 134:
		//line sql.y:745
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 135:
		//line sql.y:749
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 136:
		//line sql.y:753
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 137:
		//line sql.y:757
		{
			if num, ok := yyS[yypt-0].valExpr.(NumVal); ok {
				switch yyS[yypt-1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
			}
		}
	case 138:
		//line sql.y:772
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 139:
		//line sql.y:776
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 140:
		//line sql.y:780
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 141:
		//line sql.y:784
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 142:
		//line sql.y:788
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 143:
		//line sql.y:794
		{
			yyVAL.bytes = IF_BYTES
		}
	case 144:
		//line sql.y:798
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 145:
		//line sql.y:804
		{
			yyVAL.byt = AST_UPLUS
		}
	case 146:
		//line sql.y:808
		{
			yyVAL.byt = AST_UMINUS
		}
	case 147:
		//line sql.y:812
		{
			yyVAL.byt = AST_TILDA
		}
	case 148:
		//line sql.y:818
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 149:
		//line sql.y:823
		{
			yyVAL.valExpr = nil
		}
	case 150:
		//line sql.y:827
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 151:
		//line sql.y:833
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 152:
		//line sql.y:837
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 153:
		//line sql.y:843
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 154:
		//line sql.y:848
		{
			yyVAL.valExpr = nil
		}
	case 155:
		//line sql.y:852
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 156:
		//line sql.y:858
		{
			yyVAL.colName = &ColName{Name: yyS[yypt-0].bytes}
		}
	case 157:
		//line sql.y:862
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 158:
		//line sql.y:868
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 159:
		//line sql.y:872
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 160:
		//line sql.y:876
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 161:
		//line sql.y:880
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 162:
		//line sql.y:885
		{
			yyVAL.valExprs = nil
		}
	case 163:
		//line sql.y:889
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 164:
		//line sql.y:894
		{
			yyVAL.boolExpr = nil
		}
	case 165:
		//line sql.y:898
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 166:
		//line sql.y:903
		{
			yyVAL.orderBy = nil
		}
	case 167:
		//line sql.y:907
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 168:
		//line sql.y:913
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 169:
		//line sql.y:917
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 170:
		//line sql.y:923
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 171:
		//line sql.y:928
		{
			yyVAL.str = AST_ASC
		}
	case 172:
		//line sql.y:932
		{
			yyVAL.str = AST_ASC
		}
	case 173:
		//line sql.y:936
		{
			yyVAL.str = AST_DESC
		}
	case 174:
		//line sql.y:941
		{
			yyVAL.limit = nil
		}
	case 175:
		//line sql.y:945
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 176:
		//line sql.y:949
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 177:
		//line sql.y:954
		{
			yyVAL.str = ""
		}
	case 178:
		//line sql.y:958
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 179:
		//line sql.y:962
		{
			if !bytes.Equal(yyS[yypt-1].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyS[yypt-0].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 180:
		//line sql.y:975
		{
			yyVAL.columns = nil
		}
	case 181:
		//line sql.y:979
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 182:
		//line sql.y:985
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 183:
		//line sql.y:989
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 184:
		//line sql.y:994
		{
			yyVAL.updateExprs = nil
		}
	case 185:
		//line sql.y:998
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 186:
		//line sql.y:1004
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 187:
		//line sql.y:1008
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 188:
		//line sql.y:1014
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 189:
		//line sql.y:1019
		{
			yyVAL.empty = struct{}{}
		}
	case 190:
		//line sql.y:1021
		{
			yyVAL.empty = struct{}{}
		}
	case 191:
		//line sql.y:1024
		{
			yyVAL.empty = struct{}{}
		}
	case 192:
		//line sql.y:1026
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		//line sql.y:1029
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		//line sql.y:1031
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		//line sql.y:1035
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		//line sql.y:1037
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		//line sql.y:1039
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1041
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		//line sql.y:1043
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		//line sql.y:1046
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		//line sql.y:1048
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		//line sql.y:1051
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		//line sql.y:1053
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		//line sql.y:1056
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		//line sql.y:1058
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		//line sql.y:1062
		{
			yyVAL.bytes = bytes.ToLower(yyS[yypt-0].bytes)
		}
	case 207:
		//line sql.y:1067
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
