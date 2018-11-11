class Main{
	public static void main(String[] args) {
		int i = 1, j = 2, k = 0, sum = 0;
		while ( j < 4000000 ){
			if ( j % 2 == 0 ){
				sum += j;
			}
			k = i;
			i = j;
			j += k;
		}
	System.out.println(sum);
	}
}