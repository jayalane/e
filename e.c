#include <stdio.h>

static unsigned char digits[8192];
static int num_terms = 3400;

static const char * onetwothrees[]  = {
  "0",
  "1",
  "2",
  "3",
  "4",
  "5",
  "6",
  "7",
  "8",
  "9"
};

void divide(int n) {
    int i;
    int val = 0;

    for (i = 0; i<sizeof(digits); i++) {
        val = val * 10 + digits[i];
        if (val < n) {
            digits[i] = 0;
            continue;
        } 
        digits[i] = val / n;
        val = val % n;
    } 
}

int
main (int argc, char ** argv)
{
    int loop = 0;
    int loop2 = 0;
    digits[0] = 1;
    for (loop = num_terms; loop > 0; loop--) {
        divide(loop);
        digits[0] = digits[0] + 1;
    }
    for (loop2 = 0; loop2 < sizeof(digits); loop2++) {
      printf(onetwothrees[digits[loop2]]);
    }
    printf("\n");

    return 0;
}
