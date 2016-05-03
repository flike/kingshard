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

const yyLast = 588

var yyAct = [...]int{

	106, 103, 307, 341, 180, 104, 114, 71, 221, 303,
	374, 262, 257, 97, 193, 93, 178, 133, 181, 3,
	92, 142, 154, 155, 73, 383, 383, 203, 89, 278,
	279, 280, 281, 282, 61, 283, 284, 35, 36, 37,
	38, 45, 383, 47, 149, 75, 74, 48, 80, 149,
	149, 82, 250, 63, 219, 86, 85, 78, 50, 109,
	51, 270, 136, 248, 113, 352, 351, 119, 350, 325,
	327, 131, 60, 79, 96, 110, 111, 112, 385, 384,
	132, 329, 98, 101, 53, 54, 55, 117, 140, 126,
	81, 75, 146, 52, 251, 382, 151, 333, 56, 76,
	135, 334, 298, 296, 258, 249, 100, 218, 144, 326,
	115, 116, 94, 209, 177, 179, 186, 120, 182, 154,
	155, 258, 183, 301, 75, 74, 75, 74, 59, 58,
	199, 192, 207, 289, 336, 210, 153, 189, 167, 168,
	169, 347, 118, 129, 197, 198, 91, 304, 200, 217,
	195, 190, 72, 227, 199, 225, 62, 266, 213, 154,
	155, 238, 139, 232, 125, 226, 98, 349, 229, 230,
	214, 348, 69, 231, 323, 304, 236, 237, 322, 240,
	241, 242, 243, 244, 245, 246, 247, 165, 166, 167,
	168, 169, 75, 74, 228, 206, 208, 205, 122, 260,
	124, 98, 98, 239, 267, 253, 255, 261, 264, 144,
	319, 321, 259, 265, 317, 320, 75, 74, 141, 318,
	75, 274, 268, 272, 125, 194, 250, 254, 358, 109,
	225, 273, 288, 144, 113, 275, 271, 119, 215, 290,
	84, 224, 128, 368, 96, 110, 111, 112, 367, 90,
	223, 291, 292, 101, 148, 331, 194, 117, 162, 163,
	164, 165, 166, 167, 168, 169, 191, 295, 276, 300,
	305, 98, 302, 306, 18, 297, 100, 90, 366, 312,
	115, 116, 94, 225, 225, 315, 316, 120, 127, 109,
	35, 36, 37, 38, 113, 187, 87, 119, 149, 125,
	18, 185, 184, 90, 76, 110, 111, 112, 287, 338,
	332, 121, 118, 101, 62, 252, 76, 117, 335, 330,
	328, 311, 310, 212, 286, 339, 342, 152, 343, 211,
	224, 278, 279, 280, 281, 282, 100, 283, 284, 223,
	115, 116, 380, 62, 147, 60, 137, 120, 353, 134,
	130, 364, 83, 354, 123, 355, 362, 337, 88, 363,
	381, 365, 372, 197, 18, 373, 68, 375, 375, 375,
	75, 74, 118, 370, 371, 342, 294, 378, 376, 377,
	18, 109, 66, 387, 388, 201, 113, 143, 389, 119,
	390, 233, 138, 234, 235, 64, 76, 110, 111, 112,
	113, 308, 346, 119, 39, 101, 309, 263, 345, 117,
	76, 110, 111, 112, 18, 19, 20, 21, 314, 127,
	194, 70, 386, 117, 41, 42, 43, 44, 100, 369,
	18, 40, 115, 116, 17, 16, 57, 15, 113, 120,
	14, 119, 22, 13, 12, 202, 115, 116, 76, 110,
	111, 112, 46, 120, 269, 204, 49, 127, 77, 145,
	379, 117, 359, 340, 118, 344, 313, 299, 188, 256,
	108, 33, 105, 107, 216, 102, 156, 99, 118, 324,
	222, 277, 220, 95, 115, 116, 285, 150, 65, 34,
	67, 120, 27, 28, 11, 29, 30, 10, 31, 32,
	9, 8, 356, 23, 24, 26, 25, 162, 163, 164,
	165, 166, 167, 168, 169, 7, 118, 162, 163, 164,
	165, 166, 167, 168, 169, 157, 161, 159, 160, 6,
	5, 4, 196, 2, 1, 0, 0, 0, 360, 361,
	0, 357, 0, 0, 173, 174, 175, 176, 0, 170,
	171, 172, 293, 0, 0, 162, 163, 164, 165, 166,
	167, 168, 169, 162, 163, 164, 165, 166, 167, 168,
	169, 158, 162, 163, 164, 165, 166, 167, 168, 169,
	162, 163, 164, 165, 166, 167, 168, 169,
}
var yyPact = [...]int{

	409, -1000, -1000, 241, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -57, -42, -5, -14, -1000, 13, -1000,
	-1000, -1000, 37, 279, 425, 378, -1000, -1000, -1000, 364,
	-1000, 337, 310, 412, 64, -46, -26, 279, -1000, -8,
	279, -1000, 317, -47, 279, -47, -1000, 329, 259, -1000,
	70, -1000, -1000, -1000, -1000, 39, -1000, 272, 310, 321,
	310, 171, 413, -1000, 197, -1000, 67, 315, 4, 279,
	-1000, 314, -1000, -39, 311, 372, 98, 279, 310, 359,
	281, 309, 245, -1000, -1000, 308, 60, 94, 504, -1000,
	361, 269, -1000, -1000, -1000, 413, 258, 257, -1000, 251,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	413, -1000, 233, 281, 410, 281, 439, 375, 413, 279,
	-1000, 365, -78, -1000, 100, -1000, 294, -1000, -1000, 288,
	-1000, 205, -1000, 244, 241, 1, -1000, -1000, 206, 39,
	-1000, -1000, 279, 121, 361, 361, 413, 244, 370, 413,
	413, 136, 413, 413, 413, 413, 413, 413, 413, 413,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 504, -43,
	-1, -12, 504, -1000, 209, 39, -1000, 425, 25, 495,
	359, 281, 246, 394, 361, -1000, 413, 495, 495, -1000,
	-1000, -1000, 93, 279, -1000, -40, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 359, 281, 178, -1000, -1000, 281,
	215, 277, 289, 295, 57, -1000, -1000, -1000, -1000, -1000,
	-1000, 495, -1000, 244, 413, 413, 495, 487, -1000, 351,
	116, 116, 116, 65, 65, -1000, -1000, -1000, -1000, -1000,
	413, -1000, -1000, -3, 39, -4, 42, -1000, 361, 83,
	111, 394, 386, 392, 94, 495, 287, -1000, -1000, 286,
	-1000, -1000, 171, 244, -1000, 407, 206, 206, -1000, -1000,
	160, 156, 157, 124, 120, 7, -1000, 285, -25, 284,
	-1000, 495, 190, 413, -1000, 495, -1000, -9, -1000, 19,
	-1000, 413, 54, -1000, 327, -1000, 386, -1000, 413, 413,
	-1000, -1000, -1000, 396, 388, 277, 77, -1000, 117, -1000,
	113, -1000, -1000, -1000, -1000, -31, -33, -34, -1000, -1000,
	-1000, 413, 495, -1000, -1000, 495, 413, 324, -1000, 449,
	175, -1000, 512, -1000, 394, 361, 413, 361, -1000, -1000,
	234, 204, 199, 495, 495, 422, 413, 413, 413, -1000,
	-1000, -1000, 386, 94, 173, 94, 279, 279, 279, 281,
	495, 495, -1000, 326, -11, -1000, -27, -28, 171, -1000,
	415, 362, -1000, 279, -1000, -1000, -1000, 279, -1000, 279,
	-1000,
}
var yyPgo = [...]int{

	0, 534, 533, 18, 531, 530, 529, 515, 501, 500,
	497, 494, 404, 490, 489, 488, 20, 15, 487, 486,
	483, 482, 8, 481, 480, 129, 479, 10, 14, 13,
	477, 476, 21, 475, 16, 5, 4, 474, 473, 6,
	472, 1, 470, 469, 12, 468, 467, 466, 465, 11,
	463, 3, 462, 2, 460, 28, 459, 9, 7, 24,
	240, 458, 456, 455, 454, 452, 445, 0, 17, 444,
	443, 440, 437, 435, 434, 431,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 72, 72, 5, 6, 7, 7, 7, 69,
	69, 70, 71, 73, 73, 74, 8, 8, 8, 9,
	9, 9, 10, 11, 11, 11, 75, 12, 13, 13,
	14, 14, 14, 14, 14, 15, 15, 16, 16, 17,
	17, 17, 20, 20, 18, 18, 18, 21, 21, 22,
	22, 22, 22, 19, 19, 19, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 24, 24, 24, 25, 25,
	26, 26, 26, 26, 27, 27, 28, 28, 29, 29,
	29, 29, 29, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 31, 31, 31, 31, 31, 31, 31,
	32, 32, 37, 37, 35, 35, 39, 36, 36, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 38, 38, 40, 40,
	40, 42, 45, 45, 43, 43, 44, 46, 46, 41,
	41, 33, 33, 33, 33, 47, 47, 48, 48, 49,
	49, 50, 50, 51, 52, 52, 52, 53, 53, 53,
	53, 54, 54, 54, 55, 55, 56, 56, 57, 57,
	58, 58, 59, 60, 60, 61, 61, 62, 62, 63,
	63, 63, 63, 63, 64, 64, 65, 65, 66, 66,
	67, 68,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 12, 3,
	7, 7, 6, 6, 8, 7, 3, 4, 6, 1,
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
	7, 8, 33, 94, 95, 97, 96, 83, 84, 86,
	87, 89, 90, 62, -14, 49, 50, 51, 52, -12,
	-75, -12, -12, -12, -12, 98, -65, 100, 104, -62,
	100, 102, 98, 98, 99, 100, 85, -12, -25, 91,
	35, -67, 35, -3, 17, -15, 18, -13, 29, -25,
	9, -58, 88, -59, -41, -67, 35, -61, 103, 99,
	-67, 98, -67, 35, -60, 103, -67, -60, 29, -55,
	44, 76, -16, -17, 73, -20, 35, -29, -34, -30,
	67, 44, -33, -41, -35, -40, -67, -38, -42, 20,
	36, 37, 38, 25, -39, 71, 72, 48, 103, 28,
	78, 39, -25, 33, -25, 53, -34, 44, 45, 76,
	35, 67, -67, -68, 35, -68, 101, 35, 20, 64,
	-67, -25, -32, 28, -3, -56, -41, 35, 9, 53,
	-18, -67, 19, 76, 65, 66, -31, 21, 67, 23,
	24, 22, 68, 69, 70, 71, 72, 73, 74, 75,
	45, 46, 47, 40, 41, 42, 43, -29, -34, -29,
	-36, -3, -34, -34, 44, 44, -39, 44, -45, -34,
	-55, 33, -58, -28, 10, -59, 93, -34, -34, -67,
	-68, 20, -66, 105, -63, 97, 95, 32, 96, 13,
	35, 35, 35, -68, -55, 33, -37, -35, 106, 53,
	-21, -22, -24, 44, 35, -39, -17, -67, 73, -29,
	-29, -34, -35, 21, 23, 24, -34, -34, 25, 67,
	-34, -34, -34, -34, -34, -34, -34, -34, 106, 106,
	53, 106, 106, -16, 18, -16, -43, -44, 79, -32,
	-58, -28, -49, 13, -29, -34, 64, -67, -68, -64,
	101, -32, -58, 53, -41, -28, 53, -23, 54, 55,
	56, 57, 58, 60, 61, -19, 35, 19, -22, 76,
	-35, -34, -34, 65, 25, -34, 106, -16, 106, -46,
	-44, 81, -29, -57, 64, -57, -49, -53, 15, 14,
	35, 35, -35, -47, 11, -22, -22, 54, 59, 54,
	59, 54, 54, 54, -26, 62, 102, 63, 35, 106,
	35, 65, -34, 106, 82, -34, 80, 30, -53, -34,
	-50, -51, -34, -68, -48, 12, 14, 64, 54, 54,
	99, 99, 99, -34, -34, 31, 53, 92, 53, -52,
	26, 27, -49, -29, -36, -29, 44, 44, 44, 7,
	-34, -34, -51, -53, -27, -67, -27, -27, -58, -54,
	16, 34, 106, 53, 106, 106, 7, 21, -67, -67,
	-67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 46, 46,
	46, 46, 46, 206, 197, 0, 0, 29, 0, 31,
	32, 46, 0, 0, 0, 50, 52, 53, 54, 55,
	48, 0, 0, 0, 0, 195, 0, 0, 207, 0,
	0, 198, 0, 193, 0, 193, 30, 0, 184, 34,
	88, 35, 210, 19, 51, 0, 56, 47, 0, 0,
	0, 26, 0, 190, 0, 159, 210, 0, 0, 0,
	211, 0, 211, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 17, 57, 59, 64, 210, 62, 63, 98,
	0, 0, 129, 130, 131, 0, 159, 0, 145, 0,
	161, 162, 163, 164, 125, 148, 149, 150, 146, 147,
	152, 49, 184, 0, 96, 0, 27, 0, 0, 0,
	211, 0, 208, 38, 0, 41, 0, 43, 194, 0,
	211, 184, 33, 0, 121, 0, 186, 89, 0, 0,
	60, 65, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	113, 114, 115, 116, 117, 118, 119, 101, 0, 0,
	0, 0, 127, 140, 0, 0, 112, 0, 0, 153,
	0, 0, 96, 169, 0, 191, 0, 127, 192, 160,
	36, 196, 0, 0, 211, 204, 199, 200, 201, 202,
	203, 42, 44, 45, 0, 0, 120, 122, 185, 0,
	96, 67, 73, 0, 85, 87, 58, 66, 61, 99,
	100, 103, 104, 0, 0, 0, 106, 0, 110, 0,
	132, 133, 134, 135, 136, 137, 138, 139, 102, 124,
	0, 126, 141, 0, 0, 0, 157, 154, 0, 188,
	188, 169, 177, 0, 97, 28, 0, 209, 39, 0,
	205, 22, 23, 0, 187, 165, 0, 0, 76, 77,
	0, 0, 0, 0, 0, 90, 74, 0, 0, 0,
	105, 107, 0, 0, 111, 128, 142, 0, 144, 0,
	155, 0, 0, 20, 0, 21, 177, 25, 0, 0,
	211, 40, 123, 167, 0, 68, 71, 78, 0, 80,
	0, 82, 83, 84, 69, 0, 0, 0, 75, 70,
	86, 0, 108, 143, 151, 158, 0, 0, 24, 178,
	170, 171, 174, 37, 169, 0, 0, 0, 79, 81,
	0, 0, 0, 109, 156, 0, 0, 0, 0, 173,
	175, 176, 177, 168, 166, 72, 0, 0, 0, 0,
	179, 180, 172, 181, 0, 94, 0, 0, 189, 18,
	0, 0, 91, 0, 92, 93, 182, 0, 95, 0,
	183,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 75, 68, 3,
	44, 106, 73, 71, 53, 72, 76, 74, 3, 3,
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
		//line ./sqlparser/sql.y:181
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:187
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:207
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs}
		}
	case 18:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line ./sqlparser/sql.y:211
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:215
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 20:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./sqlparser/sql.y:222
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 21:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./sqlparser/sql.y:226
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
		//line ./sqlparser/sql.y:238
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 23:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:242
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
		//line ./sqlparser/sql.y:255
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./sqlparser/sql.y:261
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:267
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:271
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: yyDollar[4].valExpr}}}
		}
	case 28:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:275
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
		//line ./sqlparser/sql.y:291
		{
			yyVAL.statement = &Begin{}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:295
		{
			yyVAL.statement = &Begin{}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:302
		{
			yyVAL.statement = &Commit{}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:308
		{
			yyVAL.statement = &Rollback{}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:314
		{
			yyVAL.statement = &Admin{Region: yyDollar[2].tableName, Columns: yyDollar[3].columns, Rows: yyDollar[4].insRows}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:318
		{
			yyVAL.statement = &AdminHelp{}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:324
		{
			yyVAL.statement = &UseDB{DB: string(yyDollar[2].bytes)}
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:330
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 37:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./sqlparser/sql.y:334
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:339
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 39:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:345
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 40:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./sqlparser/sql.y:349
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:354
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:360
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 43:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:366
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 44:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:370
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 45:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:375
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:380
		{
			SetAllowComments(yylex, true)
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:384
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:390
		{
			yyVAL.bytes2 = nil
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:394
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:400
		{
			yyVAL.str = AST_UNION
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:404
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:408
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:412
		{
			yyVAL.str = AST_EXCEPT
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:416
		{
			yyVAL.str = AST_INTERSECT
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:421
		{
			yyVAL.str = ""
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:425
		{
			yyVAL.str = AST_DISTINCT
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:431
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:435
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:441
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:445
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:449
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:455
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:459
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:464
		{
			yyVAL.bytes = nil
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:468
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:472
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:478
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:482
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:488
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:492
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:496
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:500
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 73:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:505
		{
			yyVAL.bytes = nil
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:509
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:513
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:519
		{
			yyVAL.str = AST_JOIN
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:523
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:527
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:531
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:535
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:539
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:543
		{
			yyVAL.str = AST_JOIN
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:547
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:551
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:557
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:561
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:565
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:571
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:575
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 90:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:580
		{
			yyVAL.indexHints = nil
		}
	case 91:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:584
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:588
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 93:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:592
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:598
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:602
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 96:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:607
		{
			yyVAL.boolExpr = nil
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:611
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:618
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:622
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:626
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:630
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:636
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:640
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:644
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:648
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 107:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:652
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 108:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:656
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 109:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./sqlparser/sql.y:660
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:664
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:668
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:672
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:678
		{
			yyVAL.str = AST_EQ
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:682
		{
			yyVAL.str = AST_LT
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:686
		{
			yyVAL.str = AST_GT
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:690
		{
			yyVAL.str = AST_LE
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:694
		{
			yyVAL.str = AST_GE
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:698
		{
			yyVAL.str = AST_NE
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:702
		{
			yyVAL.str = AST_NSE
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:708
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:712
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:718
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:722
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:728
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:732
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:738
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:744
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:748
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:754
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:758
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:762
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:766
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:770
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:774
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:778
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:782
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:786
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:790
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:794
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:798
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
		//line ./sqlparser/sql.y:813
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 142:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:817
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 143:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:821
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 144:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:825
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:829
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:835
		{
			yyVAL.bytes = IF_BYTES
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:839
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:845
		{
			yyVAL.byt = AST_UPLUS
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:849
		{
			yyVAL.byt = AST_UMINUS
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:853
		{
			yyVAL.byt = AST_TILDA
		}
	case 151:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:859
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:864
		{
			yyVAL.valExpr = nil
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:868
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:874
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 155:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:878
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 156:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:884
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:889
		{
			yyVAL.valExpr = nil
		}
	case 158:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:893
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:899
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:903
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:909
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:913
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:917
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:921
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 165:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:926
		{
			yyVAL.valExprs = nil
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:930
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 167:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:935
		{
			yyVAL.boolExpr = nil
		}
	case 168:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:939
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 169:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:944
		{
			yyVAL.orderBy = nil
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:948
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:954
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:958
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 173:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:964
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 174:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:969
		{
			yyVAL.str = AST_ASC
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:973
		{
			yyVAL.str = AST_ASC
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:977
		{
			yyVAL.str = AST_DESC
		}
	case 177:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:982
		{
			yyVAL.limit = nil
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:986
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 179:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:990
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 180:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:994
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 181:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:999
		{
			yyVAL.str = ""
		}
	case 182:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:1003
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 183:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./sqlparser/sql.y:1007
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
		//line ./sqlparser/sql.y:1020
		{
			yyVAL.columns = nil
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1024
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1030
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1034
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 188:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1039
		{
			yyVAL.updateExprs = nil
		}
	case 189:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./sqlparser/sql.y:1043
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1049
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 191:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1053
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1059
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 193:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1064
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:1066
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1069
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./sqlparser/sql.y:1071
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1074
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1076
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1080
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1082
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1084
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1086
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1088
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1091
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1093
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1096
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1098
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1101
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./sqlparser/sql.y:1103
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./sqlparser/sql.y:1107
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 211:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./sqlparser/sql.y:1112
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
