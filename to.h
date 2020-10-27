#ifndef TO_H
#define TO_H
#include <stdlib.h>
#include <string.h>

char* uppercase(char* s) {
    int i = 0;

    printf("# uppercase(\"%s\")\n", s);
    for (i = 0; s[i] != '\0'; i++) {
        if (s[i] >= 97 && s[i] <= 122) {
            s[i] -= 32;
        }
        printf("%d %c\n", i, s[i]);  // DEBUG
    }

    return s;
}

/*
the idea is to take only the ASCII chars from 32 (space) to  ('Z').
first we make sure that the char 0 is actually space by subtracting 32 from
all chars and, if the char is between 97-32 and 122-32 (that is, if the char
is lower case), we convert it to uppercase by subtracting 32 again.
we then convert the resulting char into its binary representation and use
it to write boinga.
*/
char* char_to_boinga(char c) {
    char* lowercase_boinga = "boinga";
    char* uppercase_boinga = "BOINGA";
    char* outlet = (char*) malloc(sizeof(char) * 7);
    char b;
    int i;

    // ensuring character is on range
    c -= 32;
    if (c >= 97 - 32 && c <= 122 - 32) {
        c -= 32;
    }

    // converting char to binary representation to then convert it to boinga
    for (i = 0; i < 6; i++) {
        outlet[i] = (c & 0x1)? uppercase_boinga[i] : lowercase_boinga[i];
        c >>= 1;
    }

    // b o i n g a 0
    // 0 1 2 3 4 5 6
    outlet[6] = '\0';

    return outlet;
}

char* to_boinga(char* inlet) {
    int length = strlen(inlet);
    char* outlet = (char*) malloc(sizeof(char) * (length + 1) * 7);
    int i, j;
    char* boinga;

    for (i = 0; inlet[i] != '\0'; i++) {
        boinga = char_to_boinga(inlet[i]);
        for (j = 0; j < 6; j++) {
            outlet[i * 7 + j] = boinga[j];
        }
        outlet[i * 7 + 6] = ' ';
    }

    return outlet;
}

#endif /* end of include guard: TO_H */
