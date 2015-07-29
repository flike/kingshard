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

const yyNprod = 206
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 566

var yyAct = []int{

	98, 298, 167, 365, 333, 65, 169, 252, 95, 106,
	243, 208, 125, 290, 245, 184, 179, 96, 67, 192,
	84, 143, 144, 170, 3, 374, 85, 79, 374, 374,
	138, 54, 32, 33, 34, 35, 72, 296, 259, 138,
	89, 128, 69, 138, 206, 74, 206, 344, 76, 240,
	68, 101, 80, 343, 342, 56, 105, 235, 42, 111,
	44, 90, 315, 317, 45, 73, 88, 102, 103, 104,
	324, 75, 49, 376, 124, 93, 375, 373, 323, 109,
	70, 244, 132, 288, 237, 295, 135, 285, 140, 127,
	134, 283, 236, 47, 205, 48, 171, 316, 92, 244,
	172, 277, 107, 108, 86, 50, 51, 52, 142, 112,
	121, 175, 116, 143, 144, 178, 69, 123, 62, 69,
	225, 182, 188, 187, 68, 55, 339, 68, 326, 110,
	291, 66, 238, 166, 168, 189, 255, 186, 118, 135,
	131, 90, 214, 188, 118, 202, 341, 212, 218, 291,
	203, 223, 224, 198, 227, 228, 229, 230, 231, 232,
	233, 234, 226, 215, 219, 213, 340, 266, 267, 268,
	269, 270, 196, 271, 272, 199, 90, 90, 313, 312,
	114, 69, 69, 117, 216, 217, 248, 143, 144, 68,
	250, 309, 307, 256, 239, 241, 310, 308, 251, 311,
	206, 133, 350, 247, 328, 69, 257, 78, 137, 262,
	261, 185, 185, 68, 319, 156, 157, 158, 260, 120,
	212, 360, 276, 263, 279, 280, 254, 247, 17, 349,
	195, 197, 194, 154, 155, 156, 157, 158, 278, 32,
	33, 34, 35, 90, 151, 152, 153, 154, 155, 156,
	157, 158, 138, 287, 264, 118, 359, 211, 211, 297,
	81, 284, 294, 358, 293, 101, 210, 210, 204, 136,
	105, 176, 174, 111, 212, 212, 305, 306, 173, 181,
	88, 102, 103, 104, 322, 289, 83, 180, 113, 93,
	275, 325, 141, 109, 55, 371, 70, 69, 181, 330,
	320, 318, 331, 334, 302, 329, 274, 301, 55, 201,
	200, 183, 92, 372, 335, 63, 107, 108, 86, 129,
	126, 122, 119, 112, 345, 77, 115, 321, 36, 346,
	151, 152, 153, 154, 155, 156, 157, 158, 347, 327,
	82, 135, 17, 110, 354, 356, 348, 38, 39, 40,
	41, 61, 362, 334, 17, 363, 364, 282, 53, 366,
	366, 366, 69, 367, 368, 246, 378, 369, 190, 101,
	68, 130, 59, 57, 105, 379, 299, 111, 355, 380,
	357, 381, 338, 101, 70, 102, 103, 104, 105, 300,
	253, 111, 220, 93, 221, 222, 337, 109, 70, 102,
	103, 104, 304, 185, 64, 377, 17, 93, 17, 361,
	37, 109, 16, 15, 14, 13, 92, 12, 191, 43,
	107, 108, 352, 353, 258, 193, 105, 112, 46, 111,
	92, 71, 249, 370, 107, 108, 70, 102, 103, 104,
	105, 112, 351, 111, 332, 136, 336, 110, 303, 109,
	70, 102, 103, 104, 17, 18, 19, 20, 286, 136,
	177, 110, 242, 109, 151, 152, 153, 154, 155, 156,
	157, 158, 107, 108, 266, 267, 268, 269, 270, 112,
	271, 272, 21, 100, 97, 99, 107, 108, 292, 94,
	145, 91, 314, 112, 209, 265, 207, 87, 273, 110,
	146, 150, 148, 149, 151, 152, 153, 154, 155, 156,
	157, 158, 139, 110, 58, 31, 60, 11, 10, 162,
	163, 164, 165, 9, 159, 160, 161, 8, 7, 6,
	5, 4, 26, 27, 28, 2, 29, 30, 22, 23,
	25, 24, 1, 0, 0, 0, 147, 151, 152, 153,
	154, 155, 156, 157, 158, 281, 0, 0, 151, 152,
	153, 154, 155, 156, 157, 158,
}
var yyPact = []int{

	449, -1000, -1000, 190, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -35, -2, -21, 12, -1000, -1000, -1000, -1000,
	259, 403, 356, -1000, -1000, -1000, 354, -1000, 322, 280,
	395, 45, -62, -29, 259, -1000, -22, 259, -1000, 290,
	-71, 259, -71, 311, 242, -1000, -1000, -1000, 245, -1000,
	249, 280, 293, 36, 280, 91, 287, -1000, 174, -1000,
	34, 286, 50, 259, -1000, 285, -1000, -55, 284, 351,
	76, 259, 280, 415, 199, -1000, -1000, 273, 32, 122,
	479, -1000, 363, 349, -1000, -1000, -1000, 415, 234, 228,
	-1000, 227, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 415, -1000, 254, 261, 276, 393, 261, -1000,
	415, 259, -1000, 348, -81, -1000, 140, -1000, 275, -1000,
	-1000, 274, -1000, 235, -7, 436, 401, 222, 245, -1000,
	-1000, 259, 90, 363, 363, 415, 225, 371, 415, 415,
	95, 415, 415, 415, 415, 415, 415, 415, 415, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 479, -44, -9,
	-17, 479, -1000, 31, 245, -1000, 403, 20, 436, 337,
	261, 261, 202, -1000, 377, 363, -1000, 436, -1000, -1000,
	-1000, 72, 259, -1000, -58, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 337, 261, -1000, 415, 201, 420, 271,
	223, 25, -1000, -1000, -1000, -1000, -1000, -1000, 436, -1000,
	225, 415, 415, 436, 490, -1000, 332, 162, 162, 162,
	142, 142, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -10,
	245, -14, 2, -1000, 363, 66, 225, 190, 85, -16,
	-1000, 377, 361, 375, 122, 272, -1000, -1000, 269, -1000,
	-1000, 91, 436, 391, 222, 222, -1000, -1000, 138, 137,
	145, 125, 124, 0, -1000, 266, 113, 265, -1000, 436,
	262, 415, -1000, -1000, -23, -1000, -12, -1000, 415, 48,
	-1000, 309, 151, -1000, -1000, -1000, 261, 361, -1000, 415,
	415, -1000, -1000, 384, 368, 420, 62, -1000, 112, -1000,
	92, -1000, -1000, -1000, -1000, -40, -41, -47, -1000, -1000,
	-1000, 415, 436, -1000, -1000, 436, 415, 307, 225, -1000,
	-1000, 176, 149, -1000, 396, -1000, 377, 363, 415, 363,
	-1000, -1000, 219, 212, 177, 436, 436, 402, -1000, 415,
	415, -1000, -1000, -1000, 361, 122, 147, 122, 259, 259,
	259, 261, 436, -1000, 279, -24, -1000, -25, -28, 91,
	-1000, 398, 345, -1000, 259, -1000, -1000, -1000, 259, -1000,
	259, -1000,
}
var yyPgo = []int{

	0, 542, 535, 23, 531, 530, 529, 528, 527, 523,
	518, 517, 328, 516, 515, 514, 20, 26, 512, 498,
	497, 496, 11, 495, 494, 118, 492, 3, 15, 40,
	491, 490, 14, 489, 2, 17, 6, 488, 485, 9,
	484, 8, 483, 462, 10, 460, 458, 448, 446, 7,
	444, 4, 442, 1, 433, 16, 432, 13, 5, 18,
	207, 431, 428, 425, 424, 419, 418, 0, 12, 417,
	415, 414, 413, 412, 410,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 3, 3, 3, 4,
	4, 72, 72, 5, 6, 7, 7, 69, 70, 71,
	73, 8, 8, 8, 9, 9, 9, 10, 11, 11,
	11, 74, 12, 13, 13, 14, 14, 14, 14, 14,
	15, 15, 16, 16, 17, 17, 17, 20, 20, 18,
	18, 18, 21, 21, 22, 22, 22, 22, 19, 19,
	19, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	24, 24, 24, 25, 25, 26, 26, 26, 26, 27,
	27, 28, 28, 29, 29, 29, 29, 29, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 31, 31,
	31, 31, 31, 31, 31, 32, 32, 37, 37, 35,
	35, 39, 36, 36, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 38, 38, 40, 40, 40, 42, 45, 45, 43,
	43, 44, 46, 46, 41, 41, 33, 33, 33, 33,
	47, 47, 48, 48, 49, 49, 50, 50, 51, 52,
	52, 52, 53, 53, 53, 54, 54, 54, 55, 55,
	56, 56, 57, 57, 58, 58, 59, 60, 60, 61,
	61, 62, 62, 63, 63, 63, 63, 63, 64, 64,
	65, 65, 66, 66, 67, 68,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 4, 12, 3, 7,
	7, 6, 6, 8, 7, 3, 4, 1, 1, 1,
	5, 5, 8, 4, 6, 7, 4, 5, 4, 5,
	5, 0, 2, 0, 2, 1, 2, 1, 1, 1,
	0, 1, 1, 3, 1, 2, 3, 1, 1, 0,
	1, 2, 1, 3, 3, 3, 3, 5, 0, 1,
	2, 1, 1, 2, 3, 2, 3, 2, 2, 2,
	1, 3, 1, 1, 3, 0, 5, 5, 5, 1,
	3, 0, 2, 1, 3, 3, 2, 3, 3, 3,
	4, 3, 4, 5, 6, 3, 4, 2, 1, 1,
	1, 1, 1, 1, 1, 2, 1, 1, 3, 3,
	1, 3, 1, 3, 1, 1, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 3, 4, 5, 4,
	1, 1, 1, 1, 1, 1, 5, 0, 1, 1,
	2, 4, 0, 2, 1, 3, 1, 1, 1, 1,
	0, 3, 0, 2, 0, 3, 1, 3, 2, 0,
	1, 1, 0, 2, 4, 0, 2, 4, 0, 3,
	1, 3, 0, 5, 1, 3, 3, 0, 2, 0,
	3, 0, 1, 1, 1, 1, 1, 1, 0, 1,
	0, 1, 0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, 5, 6, 7,
	8, 33, 89, 90, 92, 91, 83, 84, 85, 87,
	88, -14, 49, 50, 51, 52, -12, -74, -12, -12,
	-12, -12, 93, -65, 95, 99, -62, 95, 97, 93,
	93, 94, 95, -12, -67, 35, -3, 17, -15, 18,
	-13, 29, -25, 35, 9, -58, 86, -59, -41, -67,
	35, -61, 98, 94, -67, 93, -67, 35, -60, 98,
	-67, -60, 29, 44, -16, -17, 73, -20, 35, -29,
	-34, -30, 67, 44, -33, -41, -35, -40, -67, -38,
	-42, 20, 36, 37, 38, 25, -39, 71, 72, 48,
	98, 28, 78, 39, -25, 33, 76, -25, 53, 35,
	45, 76, 35, 67, -67, -68, 35, -68, 96, 35,
	20, 64, -67, -25, -36, -34, 44, 9, 53, -18,
	-67, 19, 76, 65, 66, -31, 21, 67, 23, 24,
	22, 68, 69, 70, 71, 72, 73, 74, 75, 45,
	46, 47, 40, 41, 42, 43, -29, -34, -29, -36,
	-3, -34, -34, 44, 44, -39, 44, -45, -34, -55,
	33, 44, -58, 35, -28, 10, -59, -34, -67, -68,
	20, -66, 100, -63, 92, 90, 32, 91, 13, 35,
	35, 35, -68, -55, 33, 101, 53, -21, -22, -24,
	44, 35, -39, -17, -67, 73, -29, -29, -34, -35,
	21, 23, 24, -34, -34, 25, 67, -34, -34, -34,
	-34, -34, -34, -34, -34, 101, 101, 101, 101, -16,
	18, -16, -43, -44, 79, -32, 28, -3, -58, -56,
	-41, -28, -49, 13, -29, 64, -67, -68, -64, 96,
	-32, -58, -34, -28, 53, -23, 54, 55, 56, 57,
	58, 60, 61, -19, 35, 19, -22, 76, -35, -34,
	-34, 65, 25, 101, -16, 101, -46, -44, 81, -29,
	-57, 64, -37, -35, -57, 101, 53, -49, -53, 15,
	14, 35, 35, -47, 11, -22, -22, 54, 59, 54,
	59, 54, 54, 54, -26, 62, 97, 63, 35, 101,
	35, 65, -34, 101, 82, -34, 80, 30, 53, -41,
	-53, -34, -50, -51, -34, -68, -48, 12, 14, 64,
	54, 54, 94, 94, 94, -34, -34, 31, -35, 53,
	53, -52, 26, 27, -49, -29, -36, -29, 44, 44,
	44, 7, -34, -51, -53, -27, -67, -27, -27, -58,
	-54, 16, 34, 101, 53, 101, 101, 7, 21, -67,
	-67, -67,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 41, 41, 41,
	41, 41, 200, 191, 0, 0, 27, 28, 29, 41,
	0, 0, 45, 47, 48, 49, 50, 43, 0, 0,
	0, 0, 189, 0, 0, 201, 0, 0, 192, 0,
	187, 0, 187, 0, 0, 204, 18, 46, 0, 51,
	42, 0, 0, 83, 0, 25, 0, 184, 0, 154,
	204, 0, 0, 0, 205, 0, 205, 0, 0, 0,
	0, 0, 0, 0, 16, 52, 54, 59, 204, 57,
	58, 93, 0, 0, 124, 125, 126, 0, 154, 0,
	140, 0, 156, 157, 158, 159, 120, 143, 144, 145,
	141, 142, 147, 44, 178, 0, 0, 91, 0, 26,
	0, 0, 205, 0, 202, 33, 0, 36, 0, 38,
	188, 0, 205, 178, 0, 122, 0, 0, 0, 55,
	60, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
	109, 110, 111, 112, 113, 114, 96, 0, 0, 0,
	0, 122, 135, 0, 0, 107, 0, 0, 148, 0,
	0, 0, 91, 84, 164, 0, 185, 186, 155, 31,
	190, 0, 0, 205, 198, 193, 194, 195, 196, 197,
	37, 39, 40, 0, 0, 30, 0, 91, 62, 68,
	0, 80, 82, 53, 61, 56, 94, 95, 98, 99,
	0, 0, 0, 101, 0, 105, 0, 127, 128, 129,
	130, 131, 132, 133, 134, 97, 119, 121, 136, 0,
	0, 0, 152, 149, 0, 182, 0, 116, 182, 0,
	180, 164, 172, 0, 92, 0, 203, 34, 0, 199,
	21, 22, 123, 160, 0, 0, 71, 72, 0, 0,
	0, 0, 0, 85, 69, 0, 0, 0, 100, 102,
	0, 0, 106, 137, 0, 139, 0, 150, 0, 0,
	19, 0, 115, 117, 20, 179, 0, 172, 24, 0,
	0, 205, 35, 162, 0, 63, 66, 73, 0, 75,
	0, 77, 78, 79, 64, 0, 0, 0, 70, 65,
	81, 0, 103, 138, 146, 153, 0, 0, 0, 181,
	23, 173, 165, 166, 169, 32, 164, 0, 0, 0,
	74, 76, 0, 0, 0, 104, 151, 0, 118, 0,
	0, 168, 170, 171, 172, 163, 161, 67, 0, 0,
	0, 0, 174, 167, 175, 0, 89, 0, 0, 183,
	17, 0, 0, 86, 0, 87, 88, 176, 0, 90,
	0, 177,
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
		//line sql.y:162
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:168
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
		//line sql.y:187
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyS[yypt-2].bytes2), Distinct: yyS[yypt-1].str, SelectExprs: yyS[yypt-0].selectExprs}
		}
	case 17:
		//line sql.y:191
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 18:
		//line sql.y:195
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 19:
		//line sql.y:202
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 20:
		//line sql.y:206
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 21:
		//line sql.y:218
		{
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: yyS[yypt-1].columns, Rows: yyS[yypt-0].insRows}
		}
	case 22:
		//line sql.y:222
		{
			cols := make(Columns, 0, len(yyS[yypt-0].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-0].updateExprs))
			for _, col := range yyS[yypt-0].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 23:
		//line sql.y:235
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 24:
		//line sql.y:241
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 25:
		//line sql.y:247
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 26:
		//line sql.y:251
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyS[yypt-0].bytes)}}}
		}
	case 27:
		//line sql.y:257
		{
			yyVAL.statement = &Begin{}
		}
	case 28:
		//line sql.y:263
		{
			yyVAL.statement = &Commit{}
		}
	case 29:
		//line sql.y:269
		{
			yyVAL.statement = &Rollback{}
		}
	case 30:
		//line sql.y:275
		{
			yyVAL.statement = &Admin{Name: yyS[yypt-3].bytes, Values: yyS[yypt-1].valExprs}
		}
	case 31:
		//line sql.y:281
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 32:
		//line sql.y:285
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 33:
		//line sql.y:290
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 34:
		//line sql.y:296
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 35:
		//line sql.y:300
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 36:
		//line sql.y:305
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 37:
		//line sql.y:311
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 38:
		//line sql.y:317
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 39:
		//line sql.y:321
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 40:
		//line sql.y:326
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 41:
		//line sql.y:331
		{
			SetAllowComments(yylex, true)
		}
	case 42:
		//line sql.y:335
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 43:
		//line sql.y:341
		{
			yyVAL.bytes2 = nil
		}
	case 44:
		//line sql.y:345
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 45:
		//line sql.y:351
		{
			yyVAL.str = AST_UNION
		}
	case 46:
		//line sql.y:355
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 47:
		//line sql.y:359
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 48:
		//line sql.y:363
		{
			yyVAL.str = AST_EXCEPT
		}
	case 49:
		//line sql.y:367
		{
			yyVAL.str = AST_INTERSECT
		}
	case 50:
		//line sql.y:372
		{
			yyVAL.str = ""
		}
	case 51:
		//line sql.y:376
		{
			yyVAL.str = AST_DISTINCT
		}
	case 52:
		//line sql.y:382
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 53:
		//line sql.y:386
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 54:
		//line sql.y:392
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 55:
		//line sql.y:396
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 56:
		//line sql.y:400
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 57:
		//line sql.y:406
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 58:
		//line sql.y:410
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 59:
		//line sql.y:415
		{
			yyVAL.bytes = nil
		}
	case 60:
		//line sql.y:419
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 61:
		//line sql.y:423
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 62:
		//line sql.y:429
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 63:
		//line sql.y:433
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 64:
		//line sql.y:439
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 65:
		//line sql.y:443
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 66:
		//line sql.y:447
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 67:
		//line sql.y:451
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 68:
		//line sql.y:456
		{
			yyVAL.bytes = nil
		}
	case 69:
		//line sql.y:460
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 70:
		//line sql.y:464
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 71:
		//line sql.y:470
		{
			yyVAL.str = AST_JOIN
		}
	case 72:
		//line sql.y:474
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 73:
		//line sql.y:478
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 74:
		//line sql.y:482
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 75:
		//line sql.y:486
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 76:
		//line sql.y:490
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 77:
		//line sql.y:494
		{
			yyVAL.str = AST_JOIN
		}
	case 78:
		//line sql.y:498
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 79:
		//line sql.y:502
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 80:
		//line sql.y:508
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 81:
		//line sql.y:512
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 82:
		//line sql.y:516
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 83:
		//line sql.y:522
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 84:
		//line sql.y:526
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 85:
		//line sql.y:531
		{
			yyVAL.indexHints = nil
		}
	case 86:
		//line sql.y:535
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 87:
		//line sql.y:539
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 88:
		//line sql.y:543
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 89:
		//line sql.y:549
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 90:
		//line sql.y:553
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 91:
		//line sql.y:558
		{
			yyVAL.boolExpr = nil
		}
	case 92:
		//line sql.y:562
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 93:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 94:
		//line sql.y:569
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 95:
		//line sql.y:573
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 96:
		//line sql.y:577
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 97:
		//line sql.y:581
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 98:
		//line sql.y:587
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 99:
		//line sql.y:591
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].tuple}
		}
	case 100:
		//line sql.y:595
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].tuple}
		}
	case 101:
		//line sql.y:599
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 102:
		//line sql.y:603
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 103:
		//line sql.y:607
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 104:
		//line sql.y:611
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 105:
		//line sql.y:615
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 106:
		//line sql.y:619
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 107:
		//line sql.y:623
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 108:
		//line sql.y:629
		{
			yyVAL.str = AST_EQ
		}
	case 109:
		//line sql.y:633
		{
			yyVAL.str = AST_LT
		}
	case 110:
		//line sql.y:637
		{
			yyVAL.str = AST_GT
		}
	case 111:
		//line sql.y:641
		{
			yyVAL.str = AST_LE
		}
	case 112:
		//line sql.y:645
		{
			yyVAL.str = AST_GE
		}
	case 113:
		//line sql.y:649
		{
			yyVAL.str = AST_NE
		}
	case 114:
		//line sql.y:653
		{
			yyVAL.str = AST_NSE
		}
	case 115:
		//line sql.y:659
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 116:
		//line sql.y:663
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 117:
		//line sql.y:669
		{
			yyVAL.values = Values{yyS[yypt-0].tuple}
		}
	case 118:
		//line sql.y:673
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].tuple)
		}
	case 119:
		//line sql.y:679
		{
			yyVAL.tuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 120:
		//line sql.y:683
		{
			yyVAL.tuple = yyS[yypt-0].subquery
		}
	case 121:
		//line sql.y:689
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 122:
		//line sql.y:695
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 123:
		//line sql.y:699
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 124:
		//line sql.y:705
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 125:
		//line sql.y:709
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 126:
		//line sql.y:713
		{
			yyVAL.valExpr = yyS[yypt-0].tuple
		}
	case 127:
		//line sql.y:717
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 128:
		//line sql.y:721
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 129:
		//line sql.y:725
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 130:
		//line sql.y:729
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 131:
		//line sql.y:733
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 132:
		//line sql.y:737
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 133:
		//line sql.y:741
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 134:
		//line sql.y:745
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 135:
		//line sql.y:749
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
	case 136:
		//line sql.y:764
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 137:
		//line sql.y:768
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 138:
		//line sql.y:772
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 139:
		//line sql.y:776
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 140:
		//line sql.y:780
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 141:
		//line sql.y:786
		{
			yyVAL.bytes = IF_BYTES
		}
	case 142:
		//line sql.y:790
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 143:
		//line sql.y:796
		{
			yyVAL.byt = AST_UPLUS
		}
	case 144:
		//line sql.y:800
		{
			yyVAL.byt = AST_UMINUS
		}
	case 145:
		//line sql.y:804
		{
			yyVAL.byt = AST_TILDA
		}
	case 146:
		//line sql.y:810
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 147:
		//line sql.y:815
		{
			yyVAL.valExpr = nil
		}
	case 148:
		//line sql.y:819
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 149:
		//line sql.y:825
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 150:
		//line sql.y:829
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 151:
		//line sql.y:835
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 152:
		//line sql.y:840
		{
			yyVAL.valExpr = nil
		}
	case 153:
		//line sql.y:844
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 154:
		//line sql.y:850
		{
			yyVAL.colName = &ColName{Name: yyS[yypt-0].bytes}
		}
	case 155:
		//line sql.y:854
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 156:
		//line sql.y:860
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 157:
		//line sql.y:864
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 158:
		//line sql.y:868
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 159:
		//line sql.y:872
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 160:
		//line sql.y:877
		{
			yyVAL.valExprs = nil
		}
	case 161:
		//line sql.y:881
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 162:
		//line sql.y:886
		{
			yyVAL.boolExpr = nil
		}
	case 163:
		//line sql.y:890
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 164:
		//line sql.y:895
		{
			yyVAL.orderBy = nil
		}
	case 165:
		//line sql.y:899
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 166:
		//line sql.y:905
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 167:
		//line sql.y:909
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 168:
		//line sql.y:915
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 169:
		//line sql.y:920
		{
			yyVAL.str = AST_ASC
		}
	case 170:
		//line sql.y:924
		{
			yyVAL.str = AST_ASC
		}
	case 171:
		//line sql.y:928
		{
			yyVAL.str = AST_DESC
		}
	case 172:
		//line sql.y:933
		{
			yyVAL.limit = nil
		}
	case 173:
		//line sql.y:937
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 174:
		//line sql.y:941
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 175:
		//line sql.y:946
		{
			yyVAL.str = ""
		}
	case 176:
		//line sql.y:950
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 177:
		//line sql.y:954
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
	case 178:
		//line sql.y:967
		{
			yyVAL.columns = nil
		}
	case 179:
		//line sql.y:971
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 180:
		//line sql.y:977
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 181:
		//line sql.y:981
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 182:
		//line sql.y:986
		{
			yyVAL.updateExprs = nil
		}
	case 183:
		//line sql.y:990
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 184:
		//line sql.y:996
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 185:
		//line sql.y:1000
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 186:
		//line sql.y:1006
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 187:
		//line sql.y:1011
		{
			yyVAL.empty = struct{}{}
		}
	case 188:
		//line sql.y:1013
		{
			yyVAL.empty = struct{}{}
		}
	case 189:
		//line sql.y:1016
		{
			yyVAL.empty = struct{}{}
		}
	case 190:
		//line sql.y:1018
		{
			yyVAL.empty = struct{}{}
		}
	case 191:
		//line sql.y:1021
		{
			yyVAL.empty = struct{}{}
		}
	case 192:
		//line sql.y:1023
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		//line sql.y:1027
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		//line sql.y:1029
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		//line sql.y:1031
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		//line sql.y:1033
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		//line sql.y:1035
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1038
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		//line sql.y:1040
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		//line sql.y:1043
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		//line sql.y:1045
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		//line sql.y:1048
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		//line sql.y:1050
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		//line sql.y:1054
		{
			yyVAL.bytes = bytes.ToLower(yyS[yypt-0].bytes)
		}
	case 205:
		//line sql.y:1059
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
