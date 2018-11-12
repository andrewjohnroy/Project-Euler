using System;

namespace Project_Euler_Solutions
{
    class Program
    {
        static void Main(string[] args)
        {
            int sumofsquares = 0;
			int squareofsum = 0;
			
			for (int i = 1; i <= 100; i++){
				sumofsquares += i * i;
				squareofsum += i;
			}

			squareofsum *= squareofsum;
			Console.WriteLine(squareofsum - sumofsquares);
        }
    }
}
