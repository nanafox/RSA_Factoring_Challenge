#include <errno.h>
#include <math.h>
#include <stdio.h>
#include <stdlib.h>

void print_prime_factors(long long int number);

/**
 * main - Prime factorization
 * @argc: argument counter
 * @argv: argument vector
 *
 * Return: 0 on success, else non-zero
 */
int main(int argc, char *argv[])
{
	FILE *file;
	char *buffer = NULL;
	long long int n_read, number;
	size_t n = 0;

	if (argc != 2)
	{
		fprintf(stderr, "USAGE: factors file\n");
		exit(EXIT_FAILURE);
	}

	file = fopen(argv[1], "r");
	if (file == NULL)
	{
		fprintf(stderr, "Error: Can't open file %s\n", argv[1]);
		exit(EXIT_FAILURE);
	}

	while ((n_read = getline(&buffer, &n, file)) != -1)
	{
		buffer[n_read - 1] = '\0';
		number = atoll(buffer);

		print_prime_factors(number);
	}

	free(buffer);
	buffer = NULL;
	fclose(file);

	if (n_read == -1)
	{
		if (errno == ENOMEM)
		{
			fprintf(stderr, "Error: malloc failed\n");
			exit(EXIT_FAILURE);
		}
		if (errno == EINVAL)
		{
			fprintf(stderr, "Error: Invalid parameter received\n");
			exit(EXIT_FAILURE);
		}
	}

	return (EXIT_SUCCESS);
}

/**
 * print_prime_factors - factorizes as many numbers as possible into a product
 * of two smaller numbers and prints the result
 * @number: the number to factorize
 */
void print_prime_factors(long long int number)
{
	long long int odd_prime;

	if (number <= 1)
		return;

	/* let's check whether it's divisible 2 and perform an early return */
	if (number % 2 == 0)
	{
		printf("%lld=%lld*%d\n", number, number / 2, 2);
		return;
	}

	for (odd_prime = 3; odd_prime <= sqrt(number); odd_prime += 2)
	{
		if ((number % odd_prime) == 0)
		{
			printf("%lld=%lld*%lld\n", number, number / odd_prime, odd_prime);
			return;
		}
	}

	/* the number is a prime number */
	printf("%lld=%lld*%d\n", number, number, 1);
}
