import math
numbers = list(range(math.floor(math.sqrt(600851475143)+1)))
for divisor in range(2,len(numbers)):
	adder = divisor
	if numbers[divisor] != 0:
		while divisor < len(numbers) - adder:
			divisor += adder
			numbers[divisor] = 0
largest = 0
for value in numbers:
	if value > largest and 600851475143 % value == 0:
		largest = value

print(largest)
