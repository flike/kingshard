//line sql.y:20
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:20
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

//line sql.y:45
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
const HELP = 57418
const LAST_INSERT_ID = 57419
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
	"HELP",
	"LAST_INSERT_ID",
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
const yyMaxDepth = 200

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 213
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 621

var yyAct = [...]int{

	107, 104, 308, 342, 181, 105, 115, 71, 222, 304,
	375, 263, 258, 98, 194, 93, 179, 134, 182, 3,
	92, 143, 155, 156, 73, 204, 85, 384, 89, 35,
	36, 37, 38, 78, 61, 271, 18, 279, 280, 281,
	282, 283, 137, 284, 285, 75, 74, 384, 80, 353,
	384, 82, 150, 63, 352, 86, 114, 150, 150, 120,
	251, 50, 351, 51, 249, 79, 76, 111, 112, 113,
	220, 45, 81, 47, 52, 128, 76, 48, 60, 118,
	133, 386, 99, 53, 54, 55, 56, 252, 141, 127,
	330, 75, 147, 335, 259, 259, 152, 302, 326, 328,
	136, 385, 116, 117, 383, 290, 334, 132, 145, 121,
	154, 299, 297, 130, 250, 178, 180, 187, 91, 183,
	62, 155, 156, 184, 219, 75, 74, 75, 74, 72,
	58, 200, 193, 239, 59, 119, 337, 348, 190, 327,
	166, 167, 168, 169, 170, 198, 199, 155, 156, 201,
	218, 196, 191, 305, 228, 200, 226, 126, 229, 214,
	168, 169, 170, 267, 233, 140, 227, 99, 305, 230,
	231, 215, 350, 69, 232, 240, 210, 237, 238, 349,
	241, 242, 243, 244, 245, 246, 247, 248, 84, 324,
	323, 320, 318, 75, 74, 208, 321, 319, 211, 123,
	261, 125, 99, 99, 322, 268, 254, 256, 262, 265,
	145, 126, 195, 260, 266, 251, 216, 75, 74, 142,
	359, 75, 275, 269, 273, 274, 195, 90, 255, 129,
	110, 226, 225, 289, 145, 114, 276, 272, 120, 192,
	291, 224, 369, 368, 87, 96, 111, 112, 113, 367,
	90, 128, 292, 293, 102, 277, 149, 122, 118, 207,
	209, 206, 35, 36, 37, 38, 62, 188, 296, 126,
	301, 306, 99, 303, 307, 357, 298, 101, 186, 18,
	313, 116, 117, 94, 226, 226, 316, 317, 121, 185,
	163, 164, 165, 166, 167, 168, 169, 170, 361, 362,
	150, 90, 97, 76, 288, 153, 381, 356, 331, 225,
	339, 333, 329, 312, 119, 358, 311, 253, 224, 336,
	287, 62, 213, 212, 382, 124, 340, 343, 148, 344,
	163, 164, 165, 166, 167, 168, 169, 170, 60, 138,
	163, 164, 165, 166, 167, 168, 169, 170, 135, 354,
	131, 83, 365, 338, 355, 88, 197, 363, 18, 68,
	364, 295, 366, 373, 198, 388, 374, 64, 376, 376,
	376, 75, 74, 66, 371, 372, 343, 202, 379, 377,
	378, 144, 110, 139, 309, 389, 347, 114, 310, 390,
	120, 391, 234, 264, 235, 236, 346, 96, 111, 112,
	113, 279, 280, 281, 282, 283, 102, 284, 285, 315,
	118, 195, 70, 387, 332, 370, 18, 163, 164, 165,
	166, 167, 168, 169, 170, 18, 40, 17, 16, 101,
	15, 110, 14, 116, 117, 94, 114, 13, 12, 120,
	121, 203, 46, 270, 205, 110, 76, 111, 112, 113,
	114, 49, 77, 120, 97, 102, 146, 380, 360, 118,
	76, 111, 112, 113, 341, 345, 119, 314, 300, 102,
	189, 257, 109, 118, 106, 108, 217, 103, 101, 157,
	100, 325, 116, 117, 223, 278, 221, 95, 114, 121,
	286, 120, 101, 151, 65, 34, 116, 117, 76, 111,
	112, 113, 67, 121, 11, 10, 9, 128, 8, 7,
	6, 118, 5, 4, 2, 119, 18, 19, 20, 21,
	163, 164, 165, 166, 167, 168, 169, 170, 1, 119,
	158, 162, 160, 161, 116, 117, 0, 39, 0, 0,
	0, 121, 0, 0, 22, 0, 0, 0, 0, 174,
	175, 176, 177, 0, 171, 172, 173, 41, 42, 43,
	44, 0, 0, 0, 0, 0, 0, 119, 0, 57,
	0, 0, 0, 33, 0, 0, 159, 163, 164, 165,
	166, 167, 168, 169, 170, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 27, 28, 0, 29, 30, 0,
	31, 32, 0, 0, 0, 0, 23, 24, 26, 25,
	294, 0, 0, 163, 164, 165, 166, 167, 168, 169,
	170,
}
var yyPact = [...]int{

	511, -1000, -1000, 213, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -28, -40, -25, -16, -1000, 1, -1000,
	-1000, -1000, 43, 231, 420, 350, -1000, -1000, -1000, 355,
	-1000, 330, 303, 403, 41, -71, -35, 231, -1000, -27,
	231, -1000, 316, -78, 231, -78, -1000, 326, 257, -1000,
	42, -1000, -1000, -1000, -1000, 362, -1000, 218, 303, 292,
	303, 158, 463, -1000, 184, -1000, 37, 315, 40, 231,
	-1000, 313, -1000, -60, 304, 363, 101, 231, 303, 353,
	268, 293, 247, -1000, -1000, 286, 34, -1000, 82, 509,
	-1000, 425, 411, -1000, -1000, -1000, 463, 245, 234, -1000,
	223, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 463, -1000, 206, 268, 401, 268, 262, 31, 463,
	231, -1000, 357, -81, -1000, 163, -1000, 288, -1000, -1000,
	287, -1000, 183, -1000, 207, 213, 17, -1000, -1000, 197,
	362, -1000, -1000, 231, 85, 425, 425, 463, 207, 371,
	463, 463, 108, 463, 463, 463, 463, 463, 463, 463,
	463, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 509,
	-43, 7, -20, 509, -1000, 210, 362, -1000, 420, 15,
	452, 353, 268, 216, 380, 425, -1000, 463, 452, 452,
	-1000, -1000, -1000, 99, 231, -1000, -67, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 353, 268, 172, -1000, -1000,
	268, 202, 347, 285, 274, 29, -1000, -1000, -1000, -1000,
	-1000, -1000, 452, -1000, 207, 463, 463, 452, 545, -1000,
	336, 69, 69, 69, 87, 87, -1000, -1000, -1000, -1000,
	-1000, 463, -1000, -1000, 5, 362, 4, 16, -1000, 425,
	89, 104, 380, 369, 374, 82, 452, 281, -1000, -1000,
	278, -1000, -1000, 158, 207, -1000, 398, 197, 197, -1000,
	-1000, 138, 137, 150, 136, 135, 36, -1000, 277, -17,
	273, -1000, 452, 349, 463, -1000, 452, -1000, -1, -1000,
	11, -1000, 463, 56, -1000, 323, -1000, 369, -1000, 463,
	463, -1000, -1000, -1000, 384, 372, 347, 73, -1000, 125,
	-1000, 118, -1000, -1000, -1000, -1000, -38, -46, -51, -1000,
	-1000, -1000, 463, 452, -1000, -1000, 452, 463, 276, -1000,
	222, 167, -1000, 272, -1000, 380, 425, 463, 425, -1000,
	-1000, 205, 199, 198, 452, 452, 408, 463, 463, 463,
	-1000, -1000, -1000, 369, 82, 162, 82, 231, 231, 231,
	268, 452, 452, -1000, 290, -3, -1000, -6, -26, 158,
	-1000, 406, 344, -1000, 231, -1000, -1000, -1000, 231, -1000,
	231, -1000,
}
var yyPgo = [...]int{

	0, 528, 514, 18, 513, 512, 510, 509, 508, 506,
	505, 504, 537, 502, 495, 494, 20, 15, 493, 490,
	487, 486, 8, 485, 484, 130, 481, 10, 14, 13,
	480, 479, 21, 477, 16, 5, 4, 476, 475, 6,
	474, 1, 472, 471, 12, 470, 468, 467, 465, 11,
	464, 3, 458, 2, 457, 28, 456, 9, 7, 24,
	188, 452, 451, 444, 443, 442, 441, 0, 17, 438,
	437, 432, 430, 428, 427, 426,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 72, 72, 5, 6, 7, 7, 7, 69,
	69, 70, 71, 73, 73, 74, 8, 8, 8, 9,
	9, 9, 10, 11, 11, 11, 75, 12, 13, 13,
	14, 14, 14, 14, 14, 15, 15, 16, 16, 17,
	17, 17, 17, 20, 20, 18, 18, 18, 21, 21,
	22, 22, 22, 22, 19, 19, 19, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 24, 24, 24, 25,
	25, 26, 26, 26, 26, 27, 27, 28, 28, 29,
	29, 29, 29, 29, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 31, 31, 31, 31, 31, 31,
	31, 32, 32, 37, 37, 35, 35, 39, 36, 36,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 38, 38, 40,
	40, 40, 42, 45, 45, 43, 43, 44, 46, 46,
	41, 41, 33, 33, 33, 33, 47, 47, 48, 48,
	49, 49, 50, 50, 51, 52, 52, 52, 53, 53,
	53, 53, 54, 54, 54, 55, 55, 56, 56, 57,
	57, 58, 58, 59, 60, 60, 61, 61, 62, 62,
	63, 63, 63, 63, 63, 64, 64, 65, 65, 66,
	66, 67, 68,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 12, 3,
	7, 7, 6, 6, 8, 7, 3, 4, 6, 1,
	2, 1, 1, 4, 2, 2, 5, 8, 4, 6,
	7, 4, 5, 4, 5, 5, 0, 2, 0, 2,
	1, 2, 1, 1, 1, 0, 1, 1, 3, 1,
	2, 3, 1, 1, 1, 0, 1, 2, 1, 3,
	3, 3, 3, 5, 0, 1, 2, 1, 1, 2,
	3, 2, 3, 2, 2, 2, 1, 3, 1, 1,
	3, 0, 5, 5, 5, 1, 3, 0, 2, 1,
	3, 3, 2, 3, 3, 3, 4, 3, 4, 5,
	6, 3, 4, 2, 1, 1, 1, 1, 1, 1,
	1, 2, 1, 1, 3, 3, 1, 3, 1, 3,
	1, 1, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 3, 4, 5, 4, 1, 1, 1, 1,
	1, 1, 5, 0, 1, 1, 2, 4, 0, 2,
	1, 3, 1, 1, 1, 1, 0, 3, 0, 2,
	0, 3, 1, 3, 2, 0, 1, 1, 0, 2,
	4, 4, 0, 2, 4, 0, 3, 1, 3, 0,
	5, 1, 3, 3, 0, 2, 0, 3, 0, 1,
	1, 1, 1, 1, 1, 0, 1, 0, 1, 0,
	2, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 33, 95, 96, 98, 97, 83, 84, 86,
	87, 89, 90, 62, -14, 49, 50, 51, 52, -12,
	-75, -12, -12, -12, -12, 99, -65, 101, 105, -62,
	101, 103, 99, 99, 100, 101, 85, -12, -25, 91,
	35, -67, 35, -3, 17, -15, 18, -13, 29, -25,
	9, -58, 88, -59, -41, -67, 35, -61, 104, 100,
	-67, 99, -67, 35, -60, 104, -67, -60, 29, -55,
	44, 76, -16, -17, 73, -20, 35, 92, -29, -34,
	-30, 67, 44, -33, -41, -35, -40, -67, -38, -42,
	20, 36, 37, 38, 25, -39, 71, 72, 48, 104,
	28, 78, 39, -25, 33, -25, 53, -34, 44, 45,
	76, 35, 67, -67, -68, 35, -68, 102, 35, 20,
	64, -67, -25, -32, 28, -3, -56, -41, 35, 9,
	53, -18, -67, 19, 76, 65, 66, -31, 21, 67,
	23, 24, 22, 68, 69, 70, 71, 72, 73, 74,
	75, 45, 46, 47, 40, 41, 42, 43, -29, -34,
	-29, -36, -3, -34, -34, 44, 44, -39, 44, -45,
	-34, -55, 33, -58, -28, 10, -59, 94, -34, -34,
	-67, -68, 20, -66, 106, -63, 98, 96, 32, 97,
	13, 35, 35, 35, -68, -55, 33, -37, -35, 107,
	53, -21, -22, -24, 44, 35, -39, -17, -67, 73,
	-29, -29, -34, -35, 21, 23, 24, -34, -34, 25,
	67, -34, -34, -34, -34, -34, -34, -34, -34, 107,
	107, 53, 107, 107, -16, 18, -16, -43, -44, 79,
	-32, -58, -28, -49, 13, -29, -34, 64, -67, -68,
	-64, 102, -32, -58, 53, -41, -28, 53, -23, 54,
	55, 56, 57, 58, 60, 61, -19, 35, 19, -22,
	76, -35, -34, -34, 65, 25, -34, 107, -16, 107,
	-46, -44, 81, -29, -57, 64, -57, -49, -53, 15,
	14, 35, 35, -35, -47, 11, -22, -22, 54, 59,
	54, 59, 54, 54, 54, -26, 62, 103, 63, 35,
	107, 35, 65, -34, 107, 82, -34, 80, 30, -53,
	-34, -50, -51, -34, -68, -48, 12, 14, 64, 54,
	54, 100, 100, 100, -34, -34, 31, 53, 93, 53,
	-52, 26, 27, -49, -29, -36, -29, 44, 44, 44,
	7, -34, -34, -51, -53, -27, -67, -27, -27, -58,
	-54, 16, 34, 107, 53, 107, 107, 7, 21, -67,
	-67, -67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 46, 46,
	46, 46, 46, 207, 198, 0, 0, 29, 0, 31,
	32, 46, 0, 0, 0, 50, 52, 53, 54, 55,
	48, 0, 0, 0, 0, 196, 0, 0, 208, 0,
	0, 199, 0, 194, 0, 194, 30, 0, 185, 34,
	89, 35, 211, 19, 51, 0, 56, 47, 0, 0,
	0, 26, 0, 191, 0, 160, 211, 0, 0, 0,
	212, 0, 212, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 17, 57, 59, 65, 211, 62, 63, 64,
	99, 0, 0, 130, 131, 132, 0, 160, 0, 146,
	0, 162, 163, 164, 165, 126, 149, 150, 151, 147,
	148, 153, 49, 185, 0, 97, 0, 27, 0, 0,
	0, 212, 0, 209, 38, 0, 41, 0, 43, 195,
	0, 212, 185, 33, 0, 122, 0, 187, 90, 0,
	0, 60, 66, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 114, 115, 116, 117, 118, 119, 120, 102, 0,
	0, 0, 0, 128, 141, 0, 0, 113, 0, 0,
	154, 0, 0, 97, 170, 0, 192, 0, 128, 193,
	161, 36, 197, 0, 0, 212, 205, 200, 201, 202,
	203, 204, 42, 44, 45, 0, 0, 121, 123, 186,
	0, 97, 68, 74, 0, 86, 88, 58, 67, 61,
	100, 101, 104, 105, 0, 0, 0, 107, 0, 111,
	0, 133, 134, 135, 136, 137, 138, 139, 140, 103,
	125, 0, 127, 142, 0, 0, 0, 158, 155, 0,
	189, 189, 170, 178, 0, 98, 28, 0, 210, 39,
	0, 206, 22, 23, 0, 188, 166, 0, 0, 77,
	78, 0, 0, 0, 0, 0, 91, 75, 0, 0,
	0, 106, 108, 0, 0, 112, 129, 143, 0, 145,
	0, 156, 0, 0, 20, 0, 21, 178, 25, 0,
	0, 212, 40, 124, 168, 0, 69, 72, 79, 0,
	81, 0, 83, 84, 85, 70, 0, 0, 0, 76,
	71, 87, 0, 109, 144, 152, 159, 0, 0, 24,
	179, 171, 172, 175, 37, 170, 0, 0, 0, 80,
	82, 0, 0, 0, 110, 157, 0, 0, 0, 0,
	174, 176, 177, 178, 169, 167, 73, 0, 0, 0,
	0, 180, 181, 173, 182, 0, 95, 0, 0, 190,
	18, 0, 0, 92, 0, 93, 94, 183, 0, 96,
	0, 184,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 75, 68, 3,
	44, 107, 73, 71, 53, 72, 76, 74, 3, 3,
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
		//line sql.y:183
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:189
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:209
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs}
		}
	case 18:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:213
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:217
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 20:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:224
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 21:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:228
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
		//line sql.y:240
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 23:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:244
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
		//line sql.y:257
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:263
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:269
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:273
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: yyDollar[4].valExpr}}}
		}
	case 28:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:277
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
		//line sql.y:293
		{
			yyVAL.statement = &Begin{}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:297
		{
			yyVAL.statement = &Begin{}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:304
		{
			yyVAL.statement = &Commit{}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:310
		{
			yyVAL.statement = &Rollback{}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:316
		{
			yyVAL.statement = &Admin{Region: yyDollar[2].tableName, Columns: yyDollar[3].columns, Rows: yyDollar[4].insRows}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:320
		{
			yyVAL.statement = &AdminHelp{}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:326
		{
			yyVAL.statement = &UseDB{DB: string(yyDollar[2].bytes)}
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:332
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 37:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:336
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:341
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 39:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:347
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 40:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:351
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:356
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:362
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 43:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:368
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 44:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:372
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 45:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:377
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:382
		{
			SetAllowComments(yylex, true)
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:386
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:392
		{
			yyVAL.bytes2 = nil
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:396
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:402
		{
			yyVAL.str = AST_UNION
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:406
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:410
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:414
		{
			yyVAL.str = AST_EXCEPT
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:418
		{
			yyVAL.str = AST_INTERSECT
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:423
		{
			yyVAL.str = ""
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:427
		{
			yyVAL.str = AST_DISTINCT
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:433
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:437
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:443
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:447
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:451
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:455
		{
			yyVAL.selectExpr = &LastInsertId{}
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:461
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:465
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 65:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:470
		{
			yyVAL.bytes = nil
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:474
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:478
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:484
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:488
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:494
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:498
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:502
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 73:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:506
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 74:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:511
		{
			yyVAL.bytes = nil
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:515
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:519
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:525
		{
			yyVAL.str = AST_JOIN
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:529
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:533
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:537
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:541
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:545
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:549
		{
			yyVAL.str = AST_JOIN
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:553
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:557
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:563
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:567
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:571
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:577
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:581
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:586
		{
			yyVAL.indexHints = nil
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:590
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 93:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:594
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 94:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:598
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:604
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:608
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 97:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:613
		{
			yyVAL.boolExpr = nil
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:617
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:624
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:628
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:632
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:636
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:642
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:646
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 106:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:650
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:654
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 108:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:658
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 109:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:662
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 110:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:666
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:670
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:674
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:678
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:684
		{
			yyVAL.str = AST_EQ
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:688
		{
			yyVAL.str = AST_LT
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:692
		{
			yyVAL.str = AST_GT
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:696
		{
			yyVAL.str = AST_LE
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:700
		{
			yyVAL.str = AST_GE
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:704
		{
			yyVAL.str = AST_NE
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:708
		{
			yyVAL.str = AST_NSE
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:714
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:718
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:724
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:728
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:734
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:738
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:744
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:750
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:754
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:760
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:764
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:768
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:772
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:776
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:780
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:784
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:788
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:792
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:796
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:800
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:804
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
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:819
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 143:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:823
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 144:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:827
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 145:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:831
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:835
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:841
		{
			yyVAL.bytes = IF_BYTES
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:845
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:851
		{
			yyVAL.byt = AST_UPLUS
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:855
		{
			yyVAL.byt = AST_UMINUS
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:859
		{
			yyVAL.byt = AST_TILDA
		}
	case 152:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:865
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:870
		{
			yyVAL.valExpr = nil
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:874
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:880
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 156:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:884
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 157:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:890
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:895
		{
			yyVAL.valExpr = nil
		}
	case 159:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:899
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:905
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:909
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:915
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:919
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:923
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:927
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 166:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:932
		{
			yyVAL.valExprs = nil
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:936
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 168:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:941
		{
			yyVAL.boolExpr = nil
		}
	case 169:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:945
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 170:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:950
		{
			yyVAL.orderBy = nil
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:954
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:960
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:964
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:970
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 175:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:975
		{
			yyVAL.str = AST_ASC
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:979
		{
			yyVAL.str = AST_ASC
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:983
		{
			yyVAL.str = AST_DESC
		}
	case 178:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:988
		{
			yyVAL.limit = nil
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:992
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 180:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:996
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 181:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1000
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 182:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1005
		{
			yyVAL.str = ""
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1009
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 184:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1013
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
	case 185:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1026
		{
			yyVAL.columns = nil
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1030
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1036
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1040
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 189:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1045
		{
			yyVAL.updateExprs = nil
		}
	case 190:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1049
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1055
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1059
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 193:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1065
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 194:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1070
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1072
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1075
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1077
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1080
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1082
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1086
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1088
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1090
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1092
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1094
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1097
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1099
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1102
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1104
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1107
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1109
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1113
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 212:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1118
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
