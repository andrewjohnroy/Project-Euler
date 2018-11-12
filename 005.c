#include <stdio.h>

int main(){
	int dividend = 20;
	while(1){
		int divisible = 1;
		for (int divisor = 1; divisor <= 20; divisor++){
			if (dividend % divisor != 0){
				divisible = 0;
				break;
			}
		}
		if (divisible){
			printf("%d", dividend);
			break;
		}else{
			dividend++;
		}
	}
}
