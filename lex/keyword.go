
package lex

type KeyWord struct {
    name string
    len int
    tok int
}


var (
    keywordsA = []KeyWord {
        KeyWord {
            "auto",
            4,
            TK_AUTO,
        },
        KeyWord {
            NULL,
            0,
            TK_ID,
        },
    }
)
