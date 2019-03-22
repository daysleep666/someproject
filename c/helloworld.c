#include <stdio.h>

void test() {
    printf("我会在exit时被调用 \n");
}

int main() {
    atexit(&test);  // int atexit(void (*func)(void));   exit函数执行时会调用它
    printf("hello world\n");
    _exit(0);
    return 0;
}