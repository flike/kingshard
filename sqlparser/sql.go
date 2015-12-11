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
const OFFSET = 57418
const COLLATE = 57419
const CREATE = 57420
const ALTER = 57421
const DROP = 57422
const RENAME = 57423
const TABLE = 57424
const INDEX = 57425
const VIEW = 57426
const TO = 57427
const IGNORE = 57428
const IF = 57429
const UNIQUE = 57430
const USING = 57431

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

const yyNprod = 211
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 579

var yyAct = [...]int{

	105, 102, 306, 340, 179, 103, 113, 70, 220, 302,
	373, 261, 256, 96, 192, 92, 177, 132, 180, 3,
	91, 141, 153, 154, 72, 382, 382, 202, 88, 277,
	278, 279, 280, 281, 60, 282, 283, 35, 36, 37,
	38, 45, 382, 47, 148, 74, 73, 48, 79, 148,
	148, 81, 249, 62, 218, 85, 84, 77, 108, 50,
	351, 51, 247, 112, 324, 326, 118, 53, 54, 55,
	56, 269, 350, 95, 109, 110, 111, 384, 383, 131,
	328, 97, 100, 135, 349, 78, 116, 139, 125, 80,
	74, 145, 52, 250, 381, 150, 332, 333, 75, 134,
	257, 297, 295, 325, 248, 99, 217, 143, 288, 114,
	115, 93, 208, 176, 178, 185, 119, 181, 153, 154,
	257, 182, 300, 74, 73, 74, 73, 152, 128, 198,
	191, 206, 90, 335, 209, 61, 188, 166, 167, 168,
	117, 153, 154, 196, 197, 130, 346, 199, 216, 194,
	189, 71, 226, 198, 224, 237, 303, 212, 265, 124,
	138, 318, 231, 348, 225, 97, 319, 228, 229, 213,
	303, 58, 230, 227, 83, 235, 236, 347, 239, 240,
	241, 242, 243, 244, 245, 246, 164, 165, 166, 167,
	168, 74, 73, 205, 207, 204, 322, 238, 259, 321,
	97, 97, 320, 266, 252, 254, 260, 263, 143, 124,
	193, 258, 264, 316, 68, 74, 73, 249, 317, 74,
	273, 267, 271, 357, 193, 272, 253, 127, 108, 224,
	86, 287, 143, 112, 274, 270, 118, 214, 289, 121,
	223, 123, 367, 95, 109, 110, 111, 366, 89, 222,
	290, 291, 100, 275, 190, 120, 116, 365, 61, 140,
	35, 36, 37, 38, 75, 89, 294, 124, 299, 304,
	97, 301, 305, 355, 296, 99, 18, 147, 311, 114,
	115, 93, 224, 224, 314, 315, 119, 126, 161, 162,
	163, 164, 165, 166, 167, 168, 161, 162, 163, 164,
	165, 166, 167, 168, 186, 379, 223, 184, 337, 331,
	117, 356, 286, 251, 183, 222, 89, 334, 329, 327,
	195, 148, 310, 380, 338, 341, 330, 342, 285, 161,
	162, 163, 164, 165, 166, 167, 168, 161, 162, 163,
	164, 165, 166, 167, 168, 309, 211, 352, 210, 146,
	363, 59, 353, 136, 133, 361, 129, 82, 362, 151,
	364, 371, 196, 122, 372, 18, 374, 374, 374, 74,
	73, 336, 369, 370, 341, 61, 377, 375, 376, 354,
	108, 87, 18, 387, 67, 112, 293, 388, 118, 389,
	386, 200, 137, 39, 108, 75, 109, 110, 111, 112,
	65, 63, 118, 307, 100, 142, 345, 308, 116, 75,
	109, 110, 111, 41, 42, 43, 44, 18, 100, 262,
	344, 232, 116, 233, 234, 57, 313, 99, 193, 69,
	385, 114, 115, 368, 40, 18, 17, 112, 119, 16,
	118, 99, 15, 14, 13, 114, 115, 75, 109, 110,
	111, 12, 119, 201, 112, 46, 126, 118, 268, 203,
	116, 49, 117, 76, 75, 109, 110, 111, 18, 19,
	20, 21, 144, 126, 378, 358, 117, 116, 339, 343,
	312, 298, 187, 114, 115, 255, 107, 104, 106, 215,
	119, 156, 160, 158, 159, 101, 22, 155, 98, 323,
	114, 115, 277, 278, 279, 280, 281, 119, 282, 283,
	172, 173, 174, 175, 117, 169, 170, 171, 359, 360,
	221, 276, 219, 94, 284, 33, 149, 64, 34, 66,
	11, 117, 10, 9, 8, 7, 6, 157, 161, 162,
	163, 164, 165, 166, 167, 168, 27, 28, 5, 29,
	30, 4, 31, 32, 2, 1, 23, 24, 26, 25,
	161, 162, 163, 164, 165, 166, 167, 168, 292, 0,
	0, 161, 162, 163, 164, 165, 166, 167, 168,
}
var yyPact = [...]int{

	463, -1000, -1000, 211, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -56, -40, -5, -30, -1000, -15, -1000,
	-1000, -1000, 316, 223, 430, 384, -1000, -1000, -1000, 382,
	-1000, 355, 316, 420, 63, -45, -13, 223, -1000, -8,
	223, -1000, 322, -46, 223, -46, -1000, 352, 272, 56,
	-1000, -1000, -1000, -1000, 38, -1000, 216, 316, 330, 316,
	156, 429, -1000, 182, -1000, 52, 321, 78, 223, -1000,
	319, -1000, -17, 318, 372, 96, 223, 316, 377, 229,
	314, 268, -1000, -1000, 340, 51, 76, 470, -1000, 374,
	360, -1000, -1000, -1000, 429, 270, 263, -1000, 260, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 429,
	-1000, 221, 229, 418, 229, 228, 412, 429, 223, -1000,
	371, -77, -1000, 99, -1000, 313, -1000, -1000, 311, -1000,
	204, -1000, 243, 211, 1, -1000, -1000, 205, 38, -1000,
	-1000, 223, 100, 374, 374, 429, 243, 400, 429, 429,
	130, 429, 429, 429, 429, 429, 429, 429, 429, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 470, -43, -1,
	-12, 470, -1000, 208, 38, -1000, 430, 21, 269, 377,
	229, 214, 406, 374, -1000, 429, 269, 269, -1000, -1000,
	-1000, 94, 223, -1000, -29, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 377, 229, 172, -1000, -1000, 229, 200,
	448, 293, 271, 32, -1000, -1000, -1000, -1000, -1000, -1000,
	269, -1000, 243, 429, 429, 269, 503, -1000, 361, 115,
	115, 115, 64, 64, -1000, -1000, -1000, -1000, -1000, 429,
	-1000, -1000, -3, 38, -4, 41, -1000, 374, 92, 106,
	406, 388, 393, 76, 269, 310, -1000, -1000, 287, -1000,
	-1000, 156, 243, -1000, 415, 205, 205, -1000, -1000, 159,
	107, 148, 145, 142, 2, -1000, 284, -25, 283, -1000,
	269, 261, 429, -1000, 269, -1000, -9, -1000, 15, -1000,
	429, 53, -1000, 341, -1000, 388, -1000, 429, 429, -1000,
	-1000, -1000, 408, 392, 448, 82, -1000, 123, -1000, 109,
	-1000, -1000, -1000, -1000, -14, -26, -38, -1000, -1000, -1000,
	429, 269, -1000, -1000, 269, 429, 348, -1000, 220, 170,
	-1000, 492, -1000, 406, 374, 429, 374, -1000, -1000, 213,
	203, 198, 269, 269, 426, 429, 429, 429, -1000, -1000,
	-1000, 388, 76, 164, 76, 223, 223, 223, 229, 269,
	269, -1000, 289, -11, -1000, -27, -28, 156, -1000, 423,
	369, -1000, 223, -1000, -1000, -1000, 223, -1000, 223, -1000,
}
var yyPgo = [...]int{

	0, 555, 554, 18, 551, 548, 536, 535, 534, 533,
	532, 530, 393, 529, 528, 527, 20, 15, 526, 524,
	523, 522, 8, 521, 520, 171, 499, 10, 14, 13,
	498, 497, 21, 495, 16, 5, 4, 489, 488, 6,
	487, 1, 486, 485, 12, 482, 481, 480, 479, 11,
	478, 3, 475, 2, 474, 28, 472, 9, 7, 24,
	174, 463, 461, 459, 458, 455, 453, 0, 17, 451,
	444, 443, 442, 439, 436, 434,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 72, 72, 5, 6, 7, 7, 7, 69,
	69, 70, 71, 73, 74, 8, 8, 8, 9, 9,
	9, 10, 11, 11, 11, 75, 12, 13, 13, 14,
	14, 14, 14, 14, 15, 15, 16, 16, 17, 17,
	17, 20, 20, 18, 18, 18, 21, 21, 22, 22,
	22, 22, 19, 19, 19, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 24, 24, 24, 25, 25, 26,
	26, 26, 26, 27, 27, 28, 28, 29, 29, 29,
	29, 29, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 31, 31, 31, 31, 31, 31, 31, 32,
	32, 37, 37, 35, 35, 39, 36, 36, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 38, 38, 40, 40, 40,
	42, 45, 45, 43, 43, 44, 46, 46, 41, 41,
	33, 33, 33, 33, 47, 47, 48, 48, 49, 49,
	50, 50, 51, 52, 52, 52, 53, 53, 53, 53,
	54, 54, 54, 55, 55, 56, 56, 57, 57, 58,
	58, 59, 60, 60, 61, 61, 62, 62, 63, 63,
	63, 63, 63, 64, 64, 65, 65, 66, 66, 67,
	68,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 12, 3,
	7, 7, 6, 6, 8, 7, 3, 4, 6, 1,
	2, 1, 1, 4, 2, 5, 8, 4, 6, 7,
	4, 5, 4, 5, 5, 0, 2, 0, 2, 1,
	2, 1, 1, 1, 0, 1, 1, 3, 1, 2,
	3, 1, 1, 0, 1, 2, 1, 3, 3, 3,
	3, 5, 0, 1, 2, 1, 1, 2, 3, 2,
	3, 2, 2, 2, 1, 3, 1, 1, 3, 0,
	5, 5, 5, 1, 3, 0, 2, 1, 3, 3,
	2, 3, 3, 3, 4, 3, 4, 5, 6, 3,
	4, 2, 1, 1, 1, 1, 1, 1, 1, 2,
	1, 1, 3, 3, 1, 3, 1, 3, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	3, 4, 5, 4, 1, 1, 1, 1, 1, 1,
	5, 0, 1, 1, 2, 4, 0, 2, 1, 3,
	1, 1, 1, 1, 0, 3, 0, 2, 0, 3,
	1, 3, 2, 0, 1, 1, 0, 2, 4, 4,
	0, 2, 4, 0, 3, 1, 3, 0, 5, 1,
	3, 3, 0, 2, 0, 3, 0, 1, 1, 1,
	1, 1, 1, 0, 1, 0, 1, 0, 2, 1,
	0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 33, 93, 94, 96, 95, 83, 84, 86,
	87, 89, 90, 62, -14, 49, 50, 51, 52, -12,
	-75, -12, -12, -12, -12, 97, -65, 99, 103, -62,
	99, 101, 97, 97, 98, 99, 85, -12, -25, 35,
	-67, 35, -3, 17, -15, 18, -13, 29, -25, 9,
	-58, 88, -59, -41, -67, 35, -61, 102, 98, -67,
	97, -67, 35, -60, 102, -67, -60, 29, -55, 44,
	76, -16, -17, 73, -20, 35, -29, -34, -30, 67,
	44, -33, -41, -35, -40, -67, -38, -42, 20, 36,
	37, 38, 25, -39, 71, 72, 48, 102, 28, 78,
	39, -25, 33, -25, 53, -34, 44, 45, 76, 35,
	67, -67, -68, 35, -68, 100, 35, 20, 64, -67,
	-25, -32, 28, -3, -56, -41, 35, 9, 53, -18,
	-67, 19, 76, 65, 66, -31, 21, 67, 23, 24,
	22, 68, 69, 70, 71, 72, 73, 74, 75, 45,
	46, 47, 40, 41, 42, 43, -29, -34, -29, -36,
	-3, -34, -34, 44, 44, -39, 44, -45, -34, -55,
	33, -58, -28, 10, -59, 92, -34, -34, -67, -68,
	20, -66, 104, -63, 96, 94, 32, 95, 13, 35,
	35, 35, -68, -55, 33, -37, -35, 105, 53, -21,
	-22, -24, 44, 35, -39, -17, -67, 73, -29, -29,
	-34, -35, 21, 23, 24, -34, -34, 25, 67, -34,
	-34, -34, -34, -34, -34, -34, -34, 105, 105, 53,
	105, 105, -16, 18, -16, -43, -44, 79, -32, -58,
	-28, -49, 13, -29, -34, 64, -67, -68, -64, 100,
	-32, -58, 53, -41, -28, 53, -23, 54, 55, 56,
	57, 58, 60, 61, -19, 35, 19, -22, 76, -35,
	-34, -34, 65, 25, -34, 105, -16, 105, -46, -44,
	81, -29, -57, 64, -57, -49, -53, 15, 14, 35,
	35, -35, -47, 11, -22, -22, 54, 59, 54, 59,
	54, 54, 54, -26, 62, 101, 63, 35, 105, 35,
	65, -34, 105, 82, -34, 80, 30, -53, -34, -50,
	-51, -34, -68, -48, 12, 14, 64, 54, 54, 98,
	98, 98, -34, -34, 31, 53, 91, 53, -52, 26,
	27, -49, -29, -36, -29, 44, 44, 44, 7, -34,
	-34, -51, -53, -27, -67, -27, -27, -58, -54, 16,
	34, 105, 53, 105, 105, 7, 21, -67, -67, -67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 45, 45,
	45, 45, 45, 205, 196, 0, 0, 29, 0, 31,
	32, 45, 0, 0, 0, 49, 51, 52, 53, 54,
	47, 0, 0, 0, 0, 194, 0, 0, 206, 0,
	0, 197, 0, 192, 0, 192, 30, 0, 183, 87,
	34, 209, 19, 50, 0, 55, 46, 0, 0, 0,
	26, 0, 189, 0, 158, 209, 0, 0, 0, 210,
	0, 210, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 17, 56, 58, 63, 209, 61, 62, 97, 0,
	0, 128, 129, 130, 0, 158, 0, 144, 0, 160,
	161, 162, 163, 124, 147, 148, 149, 145, 146, 151,
	48, 183, 0, 95, 0, 27, 0, 0, 0, 210,
	0, 207, 37, 0, 40, 0, 42, 193, 0, 210,
	183, 33, 0, 120, 0, 185, 88, 0, 0, 59,
	64, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 112,
	113, 114, 115, 116, 117, 118, 100, 0, 0, 0,
	0, 126, 139, 0, 0, 111, 0, 0, 152, 0,
	0, 95, 168, 0, 190, 0, 126, 191, 159, 35,
	195, 0, 0, 210, 203, 198, 199, 200, 201, 202,
	41, 43, 44, 0, 0, 119, 121, 184, 0, 95,
	66, 72, 0, 84, 86, 57, 65, 60, 98, 99,
	102, 103, 0, 0, 0, 105, 0, 109, 0, 131,
	132, 133, 134, 135, 136, 137, 138, 101, 123, 0,
	125, 140, 0, 0, 0, 156, 153, 0, 187, 187,
	168, 176, 0, 96, 28, 0, 208, 38, 0, 204,
	22, 23, 0, 186, 164, 0, 0, 75, 76, 0,
	0, 0, 0, 0, 89, 73, 0, 0, 0, 104,
	106, 0, 0, 110, 127, 141, 0, 143, 0, 154,
	0, 0, 20, 0, 21, 176, 25, 0, 0, 210,
	39, 122, 166, 0, 67, 70, 77, 0, 79, 0,
	81, 82, 83, 68, 0, 0, 0, 74, 69, 85,
	0, 107, 142, 150, 157, 0, 0, 24, 177, 169,
	170, 173, 36, 168, 0, 0, 0, 78, 80, 0,
	0, 0, 108, 155, 0, 0, 0, 0, 172, 174,
	175, 176, 167, 165, 71, 0, 0, 0, 0, 178,
	179, 171, 180, 0, 93, 0, 0, 188, 18, 0,
	0, 90, 0, 91, 92, 181, 0, 94, 0, 182,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 75, 68, 3,
	44, 105, 73, 71, 53, 72, 76, 74, 3, 3,
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
	97, 98, 99, 100, 101, 102, 103, 104,
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
		//line sql.y:181
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:187
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:207
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs}
		}
	case 18:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:211
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:215
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 20:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:222
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 21:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:226
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
		//line sql.y:238
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 23:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:242
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
		//line sql.y:255
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:261
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:267
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:271
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: yyDollar[4].valExpr}}}
		}
	case 28:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:275
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
		//line sql.y:291
		{
			yyVAL.statement = &Begin{}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:295
		{
			yyVAL.statement = &Begin{}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:302
		{
			yyVAL.statement = &Commit{}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:308
		{
			yyVAL.statement = &Rollback{}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:314
		{
			yyVAL.statement = &Admin{Region: yyDollar[2].tableName, Columns: yyDollar[3].columns, Rows: yyDollar[4].insRows}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:320
		{
			yyVAL.statement = &UseDB{DB: string(yyDollar[2].bytes)}
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:326
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 36:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:330
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:335
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 38:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:341
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 39:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:345
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:350
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 41:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:356
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:362
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:366
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 44:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:371
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:376
		{
			SetAllowComments(yylex, true)
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:380
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:386
		{
			yyVAL.bytes2 = nil
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:390
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:396
		{
			yyVAL.str = AST_UNION
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:400
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:404
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:408
		{
			yyVAL.str = AST_EXCEPT
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:412
		{
			yyVAL.str = AST_INTERSECT
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:417
		{
			yyVAL.str = ""
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:421
		{
			yyVAL.str = AST_DISTINCT
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:427
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:431
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:437
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:441
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:445
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:451
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:455
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 63:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:460
		{
			yyVAL.bytes = nil
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:464
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:468
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:474
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:478
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:484
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:488
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:492
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 71:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:496
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 72:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:501
		{
			yyVAL.bytes = nil
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:505
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:509
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:515
		{
			yyVAL.str = AST_JOIN
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:519
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:523
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:527
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:531
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:535
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:539
		{
			yyVAL.str = AST_JOIN
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:543
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:547
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:553
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:557
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:561
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:567
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:571
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 89:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:576
		{
			yyVAL.indexHints = nil
		}
	case 90:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:580
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 91:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:584
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:588
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:594
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:598
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 95:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:603
		{
			yyVAL.boolExpr = nil
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:607
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:614
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:618
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:622
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:626
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:632
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:636
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 104:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:640
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:644
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 106:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:648
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 107:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:652
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 108:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:656
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:660
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:664
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:668
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:674
		{
			yyVAL.str = AST_EQ
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:678
		{
			yyVAL.str = AST_LT
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:682
		{
			yyVAL.str = AST_GT
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:686
		{
			yyVAL.str = AST_LE
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:690
		{
			yyVAL.str = AST_GE
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:694
		{
			yyVAL.str = AST_NE
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:698
		{
			yyVAL.str = AST_NSE
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:704
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:708
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:714
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:718
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:724
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:728
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:734
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:740
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:744
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:750
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:754
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:758
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:762
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:766
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:770
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:774
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:778
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:782
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:786
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:790
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:794
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
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:809
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 141:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:813
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 142:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:817
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 143:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:821
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:825
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:831
		{
			yyVAL.bytes = IF_BYTES
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:835
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:841
		{
			yyVAL.byt = AST_UPLUS
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:845
		{
			yyVAL.byt = AST_UMINUS
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:849
		{
			yyVAL.byt = AST_TILDA
		}
	case 150:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:855
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:860
		{
			yyVAL.valExpr = nil
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:864
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:870
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 154:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:874
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 155:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:880
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:885
		{
			yyVAL.valExpr = nil
		}
	case 157:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:889
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:895
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:899
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:905
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:909
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:913
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:917
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 164:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:922
		{
			yyVAL.valExprs = nil
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:926
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 166:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:931
		{
			yyVAL.boolExpr = nil
		}
	case 167:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:935
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 168:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:940
		{
			yyVAL.orderBy = nil
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:944
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:950
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:954
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 172:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:960
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 173:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:965
		{
			yyVAL.str = AST_ASC
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:969
		{
			yyVAL.str = AST_ASC
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:973
		{
			yyVAL.str = AST_DESC
		}
	case 176:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:978
		{
			yyVAL.limit = nil
		}
	case 177:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:982
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 178:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:986
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 179:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:990
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 180:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:995
		{
			yyVAL.str = ""
		}
	case 181:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:999
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 182:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1003
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
	case 183:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1016
		{
			yyVAL.columns = nil
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1020
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1026
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1030
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 187:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1035
		{
			yyVAL.updateExprs = nil
		}
	case 188:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1039
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1045
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 190:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1049
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 191:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1055
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 192:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1060
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1062
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1065
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1067
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1070
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1072
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1076
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1078
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1080
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1082
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1084
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1087
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1089
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1092
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1094
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1097
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1099
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1103
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 210:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1108
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
