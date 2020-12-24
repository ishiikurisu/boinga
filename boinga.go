package boinga

/*
 * FROM BOINGA
 */

// checks if a byte is a lowercase ASCII char
func isLower(c byte) bool {
    return c >= 97 && c <= 122
}

// gets the char a boinga string represents
func toChar(boinga []byte) byte {
    var c byte = 0x0

    for i := 5; i >= 0; i-- {
        c <<= 0x1
        if !isLower(boinga[i]) {
            c |= 0x1
        }
    }

    return c + ' '
}

// translate boinga language
func FromBoinga(inlet string) string {
    length := len(inlet)
    boinga := []byte("boinga")
    outlet := []byte{}

    j := 0
    for i := 0; i < length; i++ {
        c := inlet[i]
        if c == ' ' {
            outlet = append(outlet, toChar(boinga))
            j = 0
        } else {
            boinga[j] = inlet[i]
            j++
        }
    }

    if j > 0 {
        outlet = append(outlet, toChar(boinga))
    }

    return string(outlet[:])
}

/*
 * TO BOINGA
 */

// the idea is to take only the ASCII bytes from 32 (space) to  ('Z').
// first we make sure that the byte 0 is actually space by subtracting 32 from
// all bytes and, if the byte is between 97-32 and 122-32 (that is, if the byte
// is lower case), we convert it to uppercase by subtracting 32 again.
// we then convert the resulting byte into its binary representation and use
// it to write boinga.
func fromChar(b byte) []byte {
    const lowerBoinga string = "boinga"
    const upperBoinga string = "BOINGA"
    outlet := make([]byte, 6)

    // ensuring byte is on range
    b -= 32
    if b >= 97 - 32 && b <= 122 - 32 {
        b -= 32
    }

    // converting byte to binary representation to then convert it to boinga
    for i := 0; i < 6; i++ {
        if b & 0x1 != 0 {
            outlet[i] = upperBoinga[i]
        } else {
            outlet[i] = lowerBoinga[i]
        }
        b >>= 1
    }

    return outlet
}

// translate to boinga language
func ToBoinga(inlet string) string {
    length := len(inlet)
    outlet := make([]byte, length * 7)

    for i := 0; i < length; i++ {
        boinga := fromChar(inlet[i])
        for j := 0; j < 6; j++ {
            outlet[i * 7 + j] = boinga[j]
        }
        outlet[i * 7 + 6] = ' ';
    }

    return string(outlet[:len(outlet) - 1])
}
