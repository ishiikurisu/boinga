#include <stdio.h>
#include <stdlib.h>
#include "to.h"
#include "from.h"

int main(int argc, char* argv[]) {
    char* expected = "this is another string that, hopefully, will show everything works.";
    char* boinged = to_boinga(expected);
    char* deboinged = NULL;

    printf("%s ->\n%s\n", expected, boinged);

    deboinged = from_boinga(boinged);
    printf("%s ->\n%s\n", boinged, deboinged);

    return 0;
}
