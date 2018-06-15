
package lex

type KeyWord struct {
    name string
    len int
    tok int
}

var (
    Keywords_ = []KeyWord {
        KeyWord {
            "__int64",
            0,
            TK_INT64,
        },
    }

    KeywordsA = []KeyWord {
        KeyWord {
            "auto",
            4,
            TK_AUTO,
        },
    }

    KeywordsB = []KeyWord {
        KeyWord {
            "break",
            5,
            TK_BREAK,
        },
    }

    KeywordsC = []KeyWord {
        KeyWord {
            "case",
            4,
            TK_CASE,
        },
        KeyWord {
            "char",
            4,
            TK_CHAR,
        },
        KeyWord {
            "const",
            5,
            TK_CONST,
        },
        KeyWord {
            "continue",
            8,
            TK_CONTINUE,
        },
    }
)

