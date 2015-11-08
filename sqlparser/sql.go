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
const COLLATE = 57418
const CREATE = 57419
const ALTER = 57420
const DROP = 57421
const RENAME = 57422
const TABLE = 57423
const INDEX = 57424
const VIEW = 57425
const TO = 57426
const IGNORE = 57427
const IF = 57428
const UNIQUE = 57429
const USING = 57430

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

const yyNprod = 210
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 619

var yyAct = [...]int{

	105, 306, 179, 70, 340, 302, 371, 132, 261, 102,
	96, 113, 141, 177, 220, 256, 192, 180, 3, 92,
	88, 91, 380, 72, 153, 154, 380, 103, 35, 36,
	37, 38, 202, 380, 60, 84, 277, 278, 279, 280,
	281, 148, 282, 283, 77, 74, 148, 148, 79, 249,
	218, 81, 62, 269, 73, 85, 45, 135, 47, 50,
	351, 51, 48, 247, 324, 326, 53, 54, 55, 208,
	130, 350, 349, 382, 78, 75, 56, 381, 97, 131,
	80, 52, 333, 250, 379, 125, 328, 139, 206, 134,
	74, 209, 332, 153, 154, 150, 257, 297, 295, 145,
	248, 217, 325, 257, 237, 300, 143, 288, 335, 152,
	176, 178, 128, 90, 181, 166, 167, 168, 182, 346,
	185, 303, 61, 74, 124, 74, 191, 265, 71, 198,
	153, 154, 73, 188, 73, 303, 138, 199, 318, 316,
	196, 197, 189, 319, 317, 83, 238, 212, 194, 205,
	207, 204, 226, 198, 164, 165, 166, 167, 168, 224,
	227, 213, 97, 348, 228, 229, 347, 322, 225, 230,
	216, 58, 235, 236, 321, 239, 240, 241, 242, 243,
	244, 245, 246, 320, 231, 147, 277, 278, 279, 280,
	281, 74, 282, 283, 259, 193, 124, 97, 97, 249,
	73, 86, 258, 266, 263, 252, 254, 143, 260, 264,
	193, 267, 356, 272, 68, 74, 120, 127, 271, 74,
	35, 36, 37, 38, 73, 253, 270, 108, 273, 148,
	366, 143, 112, 365, 224, 118, 274, 287, 275, 121,
	223, 123, 95, 109, 110, 111, 214, 290, 291, 222,
	190, 100, 364, 124, 126, 116, 186, 89, 184, 140,
	289, 89, 183, 294, 89, 304, 61, 97, 301, 305,
	286, 299, 75, 18, 99, 296, 329, 327, 114, 115,
	93, 358, 359, 377, 310, 119, 285, 224, 224, 151,
	314, 315, 161, 162, 163, 164, 165, 166, 167, 168,
	311, 378, 354, 223, 309, 61, 331, 337, 117, 211,
	210, 251, 222, 146, 334, 195, 355, 342, 59, 39,
	136, 338, 341, 161, 162, 163, 164, 165, 166, 167,
	168, 161, 162, 163, 164, 165, 166, 167, 168, 41,
	42, 43, 44, 133, 352, 129, 82, 122, 362, 353,
	336, 57, 360, 18, 87, 361, 67, 363, 293, 196,
	384, 369, 370, 200, 137, 372, 372, 372, 74, 368,
	341, 375, 373, 374, 108, 65, 142, 73, 63, 112,
	18, 385, 118, 307, 345, 386, 308, 387, 344, 95,
	109, 110, 111, 262, 232, 108, 233, 234, 100, 313,
	112, 193, 116, 118, 69, 383, 367, 18, 40, 17,
	75, 109, 110, 111, 16, 15, 14, 13, 12, 100,
	201, 99, 46, 116, 268, 114, 115, 93, 203, 49,
	76, 144, 119, 108, 376, 357, 18, 339, 112, 343,
	312, 118, 99, 298, 187, 255, 114, 115, 75, 109,
	110, 111, 107, 119, 104, 117, 112, 100, 106, 118,
	215, 116, 101, 155, 98, 323, 75, 109, 110, 111,
	221, 276, 219, 94, 284, 126, 117, 149, 64, 116,
	99, 34, 66, 11, 114, 115, 10, 112, 9, 8,
	118, 119, 18, 19, 20, 21, 7, 75, 109, 110,
	111, 6, 114, 115, 5, 4, 126, 2, 1, 119,
	116, 0, 0, 0, 117, 0, 0, 0, 330, 0,
	22, 161, 162, 163, 164, 165, 166, 167, 168, 0,
	0, 0, 117, 114, 115, 0, 0, 0, 292, 0,
	119, 161, 162, 163, 164, 165, 166, 167, 168, 33,
	161, 162, 163, 164, 165, 166, 167, 168, 0, 0,
	0, 0, 0, 117, 156, 160, 158, 159, 0, 0,
	27, 28, 0, 29, 30, 0, 31, 32, 0, 23,
	24, 26, 25, 172, 173, 174, 175, 0, 169, 170,
	171, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	157, 161, 162, 163, 164, 165, 166, 167, 168,
}
var yyPact = [...]int{

	487, -1000, -1000, 171, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -40, -39, -15, -30, -1000, -9, -1000,
	-1000, -1000, 283, 231, 402, 361, -1000, -1000, -1000, 357,
	-1000, 327, 283, 395, 40, -57, -23, 231, -1000, -16,
	231, -1000, 311, -66, 231, -66, -1000, 325, 220, 37,
	-1000, -1000, -1000, -1000, 354, -1000, 177, 283, 314, 283,
	143, 462, -1000, 172, -1000, 36, 310, 3, 231, -1000,
	308, -1000, -42, 285, 344, 72, 231, 283, 348, 237,
	278, 176, -1000, -1000, 270, 33, 65, 543, -1000, 413,
	375, -1000, -1000, -1000, 462, 218, 214, -1000, 212, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 462,
	-1000, 217, 237, 391, 237, 224, 431, 462, 231, -1000,
	343, -71, -1000, 56, -1000, 275, -1000, -1000, 274, -1000,
	213, -1000, 210, 171, -3, -1000, -1000, 205, 354, -1000,
	-1000, 231, 87, 413, 413, 462, 210, 373, 462, 462,
	79, 462, 462, 462, 462, 462, 462, 462, 462, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 543, -41, -4,
	-21, 543, -1000, 207, 354, -1000, 402, 17, 482, 348,
	237, 200, 380, 413, -1000, 462, 482, 482, -1000, -1000,
	-1000, 63, 231, -1000, -46, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 348, 237, 160, -1000, -1000, 237, 185,
	132, 251, 268, 31, -1000, -1000, -1000, -1000, -1000, -1000,
	482, -1000, 210, 462, 462, 482, 473, -1000, 333, 83,
	83, 83, 42, 42, -1000, -1000, -1000, -1000, -1000, 462,
	-1000, -1000, -6, 354, -7, 24, -1000, 413, 57, 71,
	380, 368, 372, 65, 482, 269, -1000, -1000, 249, -1000,
	-1000, 143, 210, -1000, 388, 205, 205, -1000, -1000, 85,
	84, 129, 120, 113, 2, -1000, 242, -18, 241, -1000,
	482, 453, 462, -1000, 482, -1000, -12, -1000, 0, -1000,
	462, 28, -1000, 320, -1000, 368, -1000, 462, 462, -1000,
	-1000, -1000, 376, 370, 132, 55, -1000, 112, -1000, 109,
	-1000, -1000, -1000, -1000, -25, -26, -37, -1000, -1000, -1000,
	462, 482, -1000, -1000, 482, 462, 271, -1000, 263, 159,
	-1000, 255, -1000, 380, 413, 462, 413, -1000, -1000, 208,
	189, 186, 482, 482, 399, 462, 462, -1000, -1000, -1000,
	368, 65, 146, 65, 231, 231, 231, 237, 482, -1000,
	267, -20, -1000, -27, -31, 143, -1000, 398, 339, -1000,
	231, -1000, -1000, -1000, 231, -1000, 231, -1000,
}
var yyPgo = [...]int{

	0, 508, 507, 17, 505, 504, 501, 496, 489, 488,
	486, 483, 319, 482, 481, 478, 21, 19, 477, 474,
	473, 472, 14, 471, 470, 171, 465, 6, 16, 10,
	464, 463, 12, 462, 13, 27, 2, 460, 458, 11,
	454, 9, 452, 445, 15, 444, 443, 440, 439, 8,
	437, 4, 435, 1, 434, 20, 431, 5, 3, 23,
	145, 430, 429, 428, 424, 422, 420, 0, 7, 418,
	417, 416, 415, 414, 409, 408,
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
	50, 50, 51, 52, 52, 52, 53, 53, 53, 54,
	54, 54, 55, 55, 56, 56, 57, 57, 58, 58,
	59, 60, 60, 61, 61, 62, 62, 63, 63, 63,
	63, 63, 64, 64, 65, 65, 66, 66, 67, 68,
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
	1, 3, 2, 0, 1, 1, 0, 2, 4, 0,
	2, 4, 0, 3, 1, 3, 0, 5, 1, 3,
	3, 0, 2, 0, 3, 0, 1, 1, 1, 1,
	1, 1, 0, 1, 0, 1, 0, 2, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 33, 92, 93, 95, 94, 83, 84, 86,
	87, 89, 90, 62, -14, 49, 50, 51, 52, -12,
	-75, -12, -12, -12, -12, 96, -65, 98, 102, -62,
	98, 100, 96, 96, 97, 98, 85, -12, -25, 35,
	-67, 35, -3, 17, -15, 18, -13, 29, -25, 9,
	-58, 88, -59, -41, -67, 35, -61, 101, 97, -67,
	96, -67, 35, -60, 101, -67, -60, 29, -55, 44,
	76, -16, -17, 73, -20, 35, -29, -34, -30, 67,
	44, -33, -41, -35, -40, -67, -38, -42, 20, 36,
	37, 38, 25, -39, 71, 72, 48, 101, 28, 78,
	39, -25, 33, -25, 53, -34, 44, 45, 76, 35,
	67, -67, -68, 35, -68, 99, 35, 20, 64, -67,
	-25, -32, 28, -3, -56, -41, 35, 9, 53, -18,
	-67, 19, 76, 65, 66, -31, 21, 67, 23, 24,
	22, 68, 69, 70, 71, 72, 73, 74, 75, 45,
	46, 47, 40, 41, 42, 43, -29, -34, -29, -36,
	-3, -34, -34, 44, 44, -39, 44, -45, -34, -55,
	33, -58, -28, 10, -59, 91, -34, -34, -67, -68,
	20, -66, 103, -63, 95, 93, 32, 94, 13, 35,
	35, 35, -68, -55, 33, -37, -35, 104, 53, -21,
	-22, -24, 44, 35, -39, -17, -67, 73, -29, -29,
	-34, -35, 21, 23, 24, -34, -34, 25, 67, -34,
	-34, -34, -34, -34, -34, -34, -34, 104, 104, 53,
	104, 104, -16, 18, -16, -43, -44, 79, -32, -58,
	-28, -49, 13, -29, -34, 64, -67, -68, -64, 99,
	-32, -58, 53, -41, -28, 53, -23, 54, 55, 56,
	57, 58, 60, 61, -19, 35, 19, -22, 76, -35,
	-34, -34, 65, 25, -34, 104, -16, 104, -46, -44,
	81, -29, -57, 64, -57, -49, -53, 15, 14, 35,
	35, -35, -47, 11, -22, -22, 54, 59, 54, 59,
	54, 54, 54, -26, 62, 100, 63, 35, 104, 35,
	65, -34, 104, 82, -34, 80, 30, -53, -34, -50,
	-51, -34, -68, -48, 12, 14, 64, 54, 54, 97,
	97, 97, -34, -34, 31, 53, 53, -52, 26, 27,
	-49, -29, -36, -29, 44, 44, 44, 7, -34, -51,
	-53, -27, -67, -27, -27, -58, -54, 16, 34, 104,
	53, 104, 104, 7, 21, -67, -67, -67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 45, 45,
	45, 45, 45, 204, 195, 0, 0, 29, 0, 31,
	32, 45, 0, 0, 0, 49, 51, 52, 53, 54,
	47, 0, 0, 0, 0, 193, 0, 0, 205, 0,
	0, 196, 0, 191, 0, 191, 30, 0, 182, 87,
	34, 208, 19, 50, 0, 55, 46, 0, 0, 0,
	26, 0, 188, 0, 158, 208, 0, 0, 0, 209,
	0, 209, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 17, 56, 58, 63, 208, 61, 62, 97, 0,
	0, 128, 129, 130, 0, 158, 0, 144, 0, 160,
	161, 162, 163, 124, 147, 148, 149, 145, 146, 151,
	48, 182, 0, 95, 0, 27, 0, 0, 0, 209,
	0, 206, 37, 0, 40, 0, 42, 192, 0, 209,
	182, 33, 0, 120, 0, 184, 88, 0, 0, 59,
	64, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 112,
	113, 114, 115, 116, 117, 118, 100, 0, 0, 0,
	0, 126, 139, 0, 0, 111, 0, 0, 152, 0,
	0, 95, 168, 0, 189, 0, 126, 190, 159, 35,
	194, 0, 0, 209, 202, 197, 198, 199, 200, 201,
	41, 43, 44, 0, 0, 119, 121, 183, 0, 95,
	66, 72, 0, 84, 86, 57, 65, 60, 98, 99,
	102, 103, 0, 0, 0, 105, 0, 109, 0, 131,
	132, 133, 134, 135, 136, 137, 138, 101, 123, 0,
	125, 140, 0, 0, 0, 156, 153, 0, 186, 186,
	168, 176, 0, 96, 28, 0, 207, 38, 0, 203,
	22, 23, 0, 185, 164, 0, 0, 75, 76, 0,
	0, 0, 0, 0, 89, 73, 0, 0, 0, 104,
	106, 0, 0, 110, 127, 141, 0, 143, 0, 154,
	0, 0, 20, 0, 21, 176, 25, 0, 0, 209,
	39, 122, 166, 0, 67, 70, 77, 0, 79, 0,
	81, 82, 83, 68, 0, 0, 0, 74, 69, 85,
	0, 107, 142, 150, 157, 0, 0, 24, 177, 169,
	170, 173, 36, 168, 0, 0, 0, 78, 80, 0,
	0, 0, 108, 155, 0, 0, 0, 172, 174, 175,
	176, 167, 165, 71, 0, 0, 0, 0, 178, 171,
	179, 0, 93, 0, 0, 187, 18, 0, 0, 90,
	0, 91, 92, 180, 0, 94, 0, 181,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 75, 68, 3,
	44, 104, 73, 71, 53, 72, 76, 74, 3, 3,
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
	97, 98, 99, 100, 101, 102, 103,
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
		//line sql.y:165
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:171
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:191
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs}
		}
	case 18:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:195
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:199
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 20:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:206
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 21:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:210
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
		//line sql.y:222
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 23:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:226
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
		//line sql.y:239
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:245
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:251
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:255
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: yyDollar[4].valExpr}}}
		}
	case 28:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:259
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
		//line sql.y:275
		{
			yyVAL.statement = &Begin{}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:279
		{
			yyVAL.statement = &Begin{}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:286
		{
			yyVAL.statement = &Commit{}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:292
		{
			yyVAL.statement = &Rollback{}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:298
		{
			yyVAL.statement = &Admin{Region: yyDollar[2].tableName, Columns: yyDollar[3].columns, Rows: yyDollar[4].insRows}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:304
		{
			yyVAL.statement = &UseDB{DB: string(yyDollar[2].bytes)}
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:310
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 36:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:314
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:319
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 38:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:325
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 39:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:329
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:334
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 41:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:340
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:346
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:350
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 44:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:355
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:360
		{
			SetAllowComments(yylex, true)
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:364
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:370
		{
			yyVAL.bytes2 = nil
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:374
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:380
		{
			yyVAL.str = AST_UNION
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:384
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:388
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:392
		{
			yyVAL.str = AST_EXCEPT
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:396
		{
			yyVAL.str = AST_INTERSECT
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:401
		{
			yyVAL.str = ""
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:405
		{
			yyVAL.str = AST_DISTINCT
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:411
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:415
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:421
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:425
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:429
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:435
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:439
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 63:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:444
		{
			yyVAL.bytes = nil
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:448
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:452
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:458
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:462
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:468
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:472
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:476
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 71:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:480
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 72:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:485
		{
			yyVAL.bytes = nil
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:489
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:493
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:499
		{
			yyVAL.str = AST_JOIN
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:503
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:507
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:511
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:515
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:519
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:523
		{
			yyVAL.str = AST_JOIN
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:527
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:531
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:537
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:541
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:545
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:551
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:555
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 89:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:560
		{
			yyVAL.indexHints = nil
		}
	case 90:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:564
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 91:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:568
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:572
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:578
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:582
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 95:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:587
		{
			yyVAL.boolExpr = nil
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:591
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:598
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:602
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:606
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:610
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:616
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:620
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 104:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:624
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:628
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 106:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:632
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 107:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:636
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 108:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:640
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:644
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:648
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:652
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:658
		{
			yyVAL.str = AST_EQ
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:662
		{
			yyVAL.str = AST_LT
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:666
		{
			yyVAL.str = AST_GT
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:670
		{
			yyVAL.str = AST_LE
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:674
		{
			yyVAL.str = AST_GE
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:678
		{
			yyVAL.str = AST_NE
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:682
		{
			yyVAL.str = AST_NSE
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:688
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:692
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:698
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:702
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:708
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:712
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:718
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:724
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:728
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:734
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:738
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:742
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:746
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:750
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:754
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:758
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:762
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:766
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:770
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:774
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:778
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
		//line sql.y:793
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 141:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:797
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 142:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:801
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 143:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:805
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:809
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:815
		{
			yyVAL.bytes = IF_BYTES
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:819
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:825
		{
			yyVAL.byt = AST_UPLUS
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:829
		{
			yyVAL.byt = AST_UMINUS
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:833
		{
			yyVAL.byt = AST_TILDA
		}
	case 150:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:839
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:844
		{
			yyVAL.valExpr = nil
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:848
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:854
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 154:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:858
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 155:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:864
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:869
		{
			yyVAL.valExpr = nil
		}
	case 157:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:873
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:879
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:883
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:889
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:893
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:897
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:901
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 164:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:906
		{
			yyVAL.valExprs = nil
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:910
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 166:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:915
		{
			yyVAL.boolExpr = nil
		}
	case 167:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:919
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 168:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:924
		{
			yyVAL.orderBy = nil
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:928
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:934
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:938
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 172:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:944
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 173:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:949
		{
			yyVAL.str = AST_ASC
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:953
		{
			yyVAL.str = AST_ASC
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:957
		{
			yyVAL.str = AST_DESC
		}
	case 176:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:962
		{
			yyVAL.limit = nil
		}
	case 177:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:966
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 178:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:970
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 179:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:975
		{
			yyVAL.str = ""
		}
	case 180:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:979
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 181:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:983
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
	case 182:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:996
		{
			yyVAL.columns = nil
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1000
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1006
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1010
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 186:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1015
		{
			yyVAL.updateExprs = nil
		}
	case 187:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1019
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1025
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 189:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1029
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 190:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1035
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 191:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1040
		{
			yyVAL.empty = struct{}{}
		}
	case 192:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1042
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1045
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1047
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1050
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1052
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1056
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1058
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1060
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1062
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1064
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1067
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1069
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1072
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1074
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1077
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1079
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1083
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 209:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1088
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
