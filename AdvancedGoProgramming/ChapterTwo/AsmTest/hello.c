#include <stdio.h>


/*
A clobbered register is a register which is trashed i.e. modified in unpredictable way by inline assembler. This usually happens when you need a temp. register or use particular instruction which happens to modify some register as a by-product.
*/
//gcc -o helloasm hello.c
// int main(){

// //    int x = 1;

// //    printf("Before asm, x = %d\n",x);

// //     asm ("movl %1, %%eax;"
// //          "movl %%eax, %0;"
// //          :"=r"(x) /* x is output operand and it's related to %0 */
// //          :"r"(11)  /* 11 is input operand and it's related to %1 */
// //          :"%eax"); /* %eax is clobbered register */

// //    printf("After asm,  x = %d\n", x);
//     AsmMovl();
// }



void AsmMovl(){
     int x = 1;

   printf("Before asm, x = %d\n",x);

    asm ("movl %1, %%eax;"
         "movl %%eax, %0;"
         :"=r"(x) /* x is output operand and it's related to %0 */
         :"r"(11)  /* 11 is input operand and it's related to %1 */
         :"%eax"); /* %eax is clobbered register */

   printf("After asm,  x = %d\n", x);
}