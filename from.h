#ifndef FROM_H
#define FROM_H
#include <stdlib.h>
#include <string.h>

int is_lower(char c) {
    return c >= 97 && c <= 122;
}

char boinga_to_char(char* boinga) {
    int i = 0;
    char c = 0;

    for (i = 5; i >= 0; i--) {
        c <<= 1;
        if (!is_lower(boinga[i])) {
            c |= 0x1;
        }
    }

    return c + ' ';
}

char* from_boinga(char* inlet) {
    int length = strlen(inlet);
    char* boinga = (char*) malloc(sizeof(char) * 7);
    char* outlet = (char*) malloc(sizeof(char) * 2);
    int i, j, k;
    char c;

    j = 0;
    k = 0;
    for (i = 0; i < length; i++) {
        c = inlet[i];
        if (c == ' ' || c == '\0') {
            outlet[k] = boinga_to_char(boinga);
            j = 0;
            k++;
        } else {
            boinga[j] = inlet[i];
            j++;
        }
    }

    outlet[k] = '\0';

    return outlet;
}

#endif /* end of include guard: FROM_H */
