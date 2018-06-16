
package lex

const (
    TK_AUTO = iota
    TK_EXTERN
    TK_REGISTER
    TK_STATIC
    TK_TYPEDEF
    TK_CONST
    TK_VOLATILE
    TK_SIGNED
    TK_UNSIGNED
    TK_SHORT
    TK_LONG
    TK_CHAR
    TK_INT
    TK_INT64
    TK_FLOAT
    TK_DOUBLE
    TK_ENUM
    TK_STRUCT
    TK_UNION
    TK_VOID
    TK_BREAK
    TK_CASE
    TK_CONTINUE
    TK_DEFAULT
    TK_DO
    TK_ELSE
    TK_FOR
    TK_GOTO
    TK_IF
    TK_RETURN
    TK_SWITCH
    TK_WHILE
    TK_SIZEOF

    //identifier
    TK_ID

    //constant
    TK_INTCONST
    TK_UINTCONST
    TK_LONGCONST
    TK_ULONGCONST
    TK_LLONGCONST
    TK_ULLONGCONST
    TK_FLOATCONST
    TK_DOUBLECONST
    TK_LDOUBLECONST
    TK_STRING
    TK_WIDESTRING

    //operators
    TK_COMMA
    TK_QUESTION
    TK_COLON
    TK_ASSIGN
    TK_BITOR_ASSIGN
    TK_BITXOR_ASSIGN
    TK_BITAND_ASSIGN
    TK_LSHIFT_ASSIGN
    TK_RSHIFT_ASSIGN
    TK_ADD_ASSIGN
    TK_SUB_ASSIGN
    TK_MUL_ASSIGN
    TK_DIV_ASSIGN
    TK_MOD_ASSIGN
    TK_OR
    TK_AND
    TK_BITOR
    TK_BITXOR
    TK_BITAND
    TK_EQUAL
    TK_UNEQUAL
    TK_GREAT
    TK_LESS
    TK_GREAT_EQ
    TK_LESS_EQ
    TK_LSHIFT
    TK_RSHIFT
    TK_ADD
    TK_SUB
    TK_MUL
    TK_DIV
    TK_MOD
    TK_INC
    TK_DEC
    TK_NOT
    TK_COMP
    TK_DOT
    TK_POINTER
    TK_LPAREN
    TK_RPAREN
    TK_LBRACKET
    TK_RBRACKET

    //punctuators
    TK_LBRACE
    TK_RBRACE
    TK_SEMICOLON
    TK_ELLIPSIS
    TK_POUND
    TK_NEWLINE

    TK_END
)

/* comments for tokens
    TOKEN(TK_AUTO,      "auto")
    TOKEN(TK_EXTERN,    "extern")
    TOKEN(TK_REGISTER,  "register")
    TOKEN(TK_STATIC,    "static")
    TOKEN(TK_TYPEDEF,   "typedef")
    TOKEN(TK_CONST,     "const")
    TOKEN(TK_VOLATILE,  "volatile")
    TOKEN(TK_SIGNED,    "signed")
    TOKEN(TK_UNSIGNED,  "unsigned")
    TOKEN(TK_SHORT,     "short")
    TOKEN(TK_LONG,      "long")
    TOKEN(TK_CHAR,      "char")
    TOKEN(TK_INT,       "int")
    TOKEN(TK_INT64,     "__int64")
    TOKEN(TK_FLOAT,     "float")
    TOKEN(TK_DOUBLE,    "double")
    TOKEN(TK_ENUM,      "enum")
    TOKEN(TK_STRUCT,    "struct")
    TOKEN(TK_UNION,     "union")
    TOKEN(TK_VOID,      "void")
    TOKEN(TK_BREAK,     "break")
    TOKEN(TK_CASE,      "case")
    TOKEN(TK_CONTINUE,  "continue")
    TOKEN(TK_DEFAULT,   "default")
    TOKEN(TK_DO,        "do")
    TOKEN(TK_ELSE,      "else")
    TOKEN(TK_FOR,       "for")
    TOKEN(TK_GOTO,      "goto")
    TOKEN(TK_IF,        "if")
    TOKEN(TK_RETURN,    "return")
    TOKEN(TK_SWITCH,    "switch")
    TOKEN(TK_WHILE,     "while")
    TOKEN(TK_SIZEOF,    "sizeof")

    //identifier
    TOKEN(TK_ID,        "ID")

    //constant
    TOKEN(TK_INTCONST,     "int")
    TOKEN(TK_UINTCONST,    "unsigned int")
    TOKEN(TK_LONGCONST,    "long")
    TOKEN(TK_ULONGCONST,   "unsigned long")
    TOKEN(TK_LLONGCONST,   "long long")
    TOKEN(TK_ULLONGCONST,  "unsigned long long")
    TOKEN(TK_FLOATCONST,   "float")
    TOKEN(TK_DOUBLECONST,  "double")
    TOKEN(TK_LDOUBLECONST, "long double")
    TOKEN(TK_STRING,       "STR")
    TOKEN(TK_WIDESTRING,   "WSTR")

    //operators
    TOKEN(TK_COMMA,         ",")
    TOKEN(TK_QUESTION,      "?")
    TOKEN(TK_COLON,         ":")
    TOKEN(TK_ASSIGN,        "=")
    TOKEN(TK_BITOR_ASSIGN,  "|=")
    TOKEN(TK_BITXOR_ASSIGN, "^=")
    TOKEN(TK_BITAND_ASSIGN, "&=")
    TOKEN(TK_LSHIFT_ASSIGN, "<<=")
    TOKEN(TK_RSHIFT_ASSIGN, ">>=")
    TOKEN(TK_ADD_ASSIGN,    "+=")
    TOKEN(TK_SUB_ASSIGN,    "-=")
    TOKEN(TK_MUL_ASSIGN,    "*=")
    TOKEN(TK_DIV_ASSIGN,    "/=")
    TOKEN(TK_MOD_ASSIGN,    "%=")
    TOKEN(TK_OR,            "||")
    TOKEN(TK_AND,           "&&")
    TOKEN(TK_BITOR,         "|")
    TOKEN(TK_BITXOR,        "^")
    TOKEN(TK_BITAND,        "&")
    TOKEN(TK_EQUAL,         "==")
    TOKEN(TK_UNEQUAL,       "!=")
    TOKEN(TK_GREAT,         ">")
    TOKEN(TK_LESS,          "<")
    TOKEN(TK_GREAT_EQ,      ">=")
    TOKEN(TK_LESS_EQ,       "<=")
    TOKEN(TK_LSHIFT,        "<<")
    TOKEN(TK_RSHIFT,        ">>")
    TOKEN(TK_ADD,           "+")
    TOKEN(TK_SUB,           "-")
    TOKEN(TK_MUL,           "*")
    TOKEN(TK_DIV,           "/")
    TOKEN(TK_MOD,           "%")
    TOKEN(TK_INC,           "++")
    TOKEN(TK_DEC,           "--")
    TOKEN(TK_NOT,           "!")
    TOKEN(TK_COMP,          "~")
    TOKEN(TK_DOT,           ".")
    TOKEN(TK_POINTER,       "->")
    TOKEN(TK_LPAREN,        "(")
    TOKEN(TK_RPAREN,        ")")
    TOKEN(TK_LBRACKET,      "[")
    TOKEN(TK_RBRACKET,      "]")

    //punctuators
    TOKEN(TK_LBRACE,        "{")
    TOKEN(TK_RBRACE,        "}")
    TOKEN(TK_SEMICOLON,     ";")
    TOKEN(TK_ELLIPSIS,       "...")
    TOKEN(TK_POUND,         "#")
    TOKEN(TK_NEWLINE,       "\n")

    TOKEN(TK_END,           "EOF")
*/

type TK_value struct {
    i int
    f float32
    d float64
    p *string
    kd int // 1 for int, 2 for float, 3 for double, 4 for *string
}

