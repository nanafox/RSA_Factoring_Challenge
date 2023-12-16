#include <errno.h>
#include <stdio.h>
#include <stdlib.h>

void print_prime_factors(size_t number);

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
	ssize_t n_read;
	size_t number, n = 0;

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
void print_prime_factors(size_t number)
{
	size_t original_number = number, odd_prime = 3;

	if (number == 1)
		return;

	if (number % 2 == 0)
	{
		printf("%lu=%lu*%d\n", number, number / 2, 2);
		return;
	}

	while ((number % odd_prime) != 0)
	{
		odd_prime += 2;
	}

	printf("%lu=%lu*%lu\n", original_number, number / odd_prime, odd_prime);
}
