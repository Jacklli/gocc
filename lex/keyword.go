
package lex

type KeyWord struct {
    name string
    len int
    tok int
}


var (
    Keywords_ = []KeyWord {
    }

    KeywordsA = []KeyWord {
    }

    KeywordsB = []KeyWord {
    }

    KeywordsC = []KeyWord {
    }

    KeywordsD = []KeyWord {
    }
   
    KeywordsE = []KeyWord {
        KeyWord {
            "else",
            4,
            TK_ELSE,
        },
    }

    KeywordsF = []KeyWord {
        KeyWord {
            "for",
            3,
            TK_FOR,
        },
    }
    
    KeywordsG = []KeyWord {
    }

    KeywordsH = []KeyWord {
    }

    KeywordsI = []KeyWord {
        KeyWord {
            "if",
            2,
            TK_IF,
        },
        KeyWord {
            "int",
            3,
            TK_INTCONST,
        },
    }

    KeywordsJ = []KeyWord {
    }

    KeywordsK = []KeyWord {
    }

    KeywordsL = []KeyWord {
    }

    KeywordsM = []KeyWord {
    }

    KeywordsN = []KeyWord {
    }

    KeywordsO = []KeyWord {
    }

    KeywordsP = []KeyWord {
    }

    KeywordsQ = []KeyWord {
    }

    KeywordsR = []KeyWord {
        KeyWord {
            "return",
            6,
            TK_RETURN,
        },
    }

    KeywordsS = []KeyWord {
    }
    
    KeywordsT = []KeyWord {
    }

    KeywordsU = []KeyWord {
    }

    KeywordsV = []KeyWord {
    }

    KeywordsW = []KeyWord {
    }

    KeywordsX = []KeyWord {
    }

    KeywordsY = []KeyWord {
    }

    KeywordsZ = []KeyWord {
    }

    Keywords =  [][]KeyWord {
        Keywords_, KeywordsA, KeywordsB, KeywordsC,
        KeywordsD, KeywordsE, KeywordsF, KeywordsG,
        KeywordsH, KeywordsI, KeywordsJ, KeywordsK,
        KeywordsL, KeywordsM, KeywordsN, KeywordsO,
        KeywordsP, KeywordsQ, KeywordsR, KeywordsS,
        KeywordsT, KeywordsU, KeywordsV, KeywordsW,
        KeywordsX, KeywordsY, KeywordsZ,
    }

/*
    Keywords_ = []KeyWord {
        KeyWord {
            "__int64",
            0,
            TK_INT64,
        },
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsA = []KeyWord {
        KeyWord {
            "auto",
            4,
            TK_AUTO,
        },
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsB = []KeyWord {
        KeyWord {
            "break",
            5,
            TK_BREAK,
        },
        KeyWord {
            "NULL",
            0,
            TK_ID,
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
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsD = []KeyWord {
        KeyWord {
            "default",
            7,
            TK_DEFAULT,
        },
        KeyWord {
            "do",
            2,
            TK_DO,
        },
        KeyWord {
            "double",
            6,
            TK_DOUBLE,
        },
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }
   
    KeywordsE = []KeyWord {
        KeyWord {
            "else",
            4,
            TK_ELSE,
        },
        KeyWord {
            "enum",
            4,
            TK_ENUM,
        },
        KeyWord {
            "extern",
            6,
            TK_EXTERN,
        },
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsF = []KeyWord {
        KeyWord {
            "float",
            5,
            TK_FLOAT,
        },
        KeyWord {
            "for",
            3,
            TK_FOR,
        },
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }
    
    KeywordsG = []KeyWord {
        KeyWord {
            "goto",
            4,
            TK_GOTO,
        },
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsH = []KeyWord {
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsI = []KeyWord {
        KeyWord {
            "if",
            2,
            TK_IF,
        },
    }

    KeywordsJ = []KeyWord {
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsK = []KeyWord {
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsL = []KeyWord {
        KeyWord {
            "long",
            4,
            TK_LONG,
        },
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsM = []KeyWord {
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsN = []KeyWord {
        KeyWord {
            "NULL",
            0, 
            TK_ID,
        },
    }

    KeywordsO = []KeyWord {
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsP = []KeyWord {
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsQ = []KeyWord {
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsR = []KeyWord {
        KeyWord {
            "register",
            8,
            TK_REGISTER,
        },
        KeyWord {
            "return",
            6,
            TK_RETURN,
        },
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsS = []KeyWord {
        KeyWord {
            "short",
            5,
            TK_SHORT,
        },
        KeyWord {
            "signed",
            6,
            TK_SIGNED,
        },
        KeyWord {
            "sizeof",
            6,
            TK_SIZEOF,
        },
        KeyWord {
            "static",
            6,
            TK_STATIC,
        },
        KeyWord {
            "struct",
            6,
            TK_STRUCT,
        },
        KeyWord {
            "switch",
            6,
            TK_SWITCH,
        },
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }
    
    KeywordsT = []KeyWord {
        KeyWord {
            "typedef",
            7,
            TK_TYPEDEF,
        },
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsU = []KeyWord {
        KeyWord {
            "union",
            5,
            TK_UNION,
        },
        KeyWord {
            "unsigned",
            8,
            TK_UNSIGNED,
        },
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsV = []KeyWord {
        KeyWord {
            "void",
            4,
            TK_VOID,
        },
        KeyWord {
            "volatile",
            8,
            TK_VOLATILE,
        },
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsW = []KeyWord {
        KeyWord {
            "while",
            5,
            TK_WHILE,
        },
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsX = []KeyWord {
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsY = []KeyWord {
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    KeywordsZ = []KeyWord {
        KeyWord {
            "NULL",
            0,
            TK_ID,
        },
    }

    Keywords =  [][]KeyWord {
        Keywords_, KeywordsA, KeywordsB, KeywordsC,
        KeywordsD, KeywordsE, KeywordsF, KeywordsG,
        KeywordsH, KeywordsI, KeywordsJ, KeywordsK,
        KeywordsL, KeywordsM, KeywordsN, KeywordsO,
        KeywordsP, KeywordsQ, KeywordsR, KeywordsS,
        KeywordsT, KeywordsU, KeywordsV, KeywordsW,
        KeywordsX, KeywordsY, KeywordsZ,
    }
*/
)

