numbers = new Array(1000000)
for (i = 0; i < numbers.length; i++){
	numbers[i] = i;
}

numbers[1] = 0;
for (i = 2; i < numbers.length; i++ ){
	increment = i;
	for (increment += i; increment < numbers.length; increment += i){
		numbers[increment] = 0;
	}
}

count = 0;

for (i = 0; i < numbers.length; i++){
	if (numbers[i] != 0){
		count++;
	}
	if (count == 10001){
		console.log(numbers[i]);
		break;
	}
}
