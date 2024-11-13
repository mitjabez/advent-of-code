#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int score(char c) {
    if (c >= 'a' && c <= 'z') return (int) (c - 96);
    else return (int) (c - 38);
}

void part1(char *argv[]) {
    char *line = NULL;
    size_t bufsize = 0;
    ssize_t len;
    int total = 0;
    char badges[256];
    FILE *f = fopen(argv[1], "r");

    while((len = getline(&line, &bufsize, f)) != -1) {
        memset(badges, 0, 256);
        int i;
        for (i = 0; i < (len-1) / 2; i++)
            badges[line[i]] = 1;
        for (int j = i; j < len; j++) {
            char c = line[j];
            if (badges[c]) {
                total += score(c);
                break;
            }
        }
    }
    free(line);
    fclose(f);
    printf("%d\n",total);
}

void part2(char *argv[]) {
    char *line = NULL;
    size_t bufsize = 0;
    ssize_t len;
    int total = 0;
    unsigned char badges[256];
    int lineno = 0;
    FILE *f = fopen(argv[1], "r");

    while((len = getline(&line, &bufsize, f)) != -1) {
        if (lineno % 3 == 0)
            memset(badges, 0, 256);

        for (int i = 0; i < len - 1; i++) {
            char c = line[i];
            badges[c] |= 0x01 << (lineno % 3);
            if (badges[c] == 0x07) {
                total += score(c);
                break;
            }
        }
        lineno++;
    }
    free(line);
    fclose(f);
    printf("%d\n",total);
}

int main(int argc, char *argv[]) {
    part1(argv);
    part2(argv);

    return 0;
}
