#include </usr/local/include/gmp.h>
#include <errno.h>
#include <stdio.h>
#include <stdlib.h>

void print_prime_factors(mpz_t number);

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
	size_t n = 0;
	mpz_t number;

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

	while (getline(&buffer, &n, file) != -1)
	{
		/* initialize the number and check for errors */
		if (mpz_init_set_str(number, buffer, 10) == -1)
		{
			mpz_clear(number);
			continue; /* yes, it was invalid but just skip it */
		}
		print_prime_factors(number);
	}

	free(buffer);
	fclose(file);

	mpz_clear(number);

	if (errno == ENOMEM)
	{
		fprintf(stderr, "Error: malloc failed\n");
		exit(EXIT_FAILURE);
	}

	return (EXIT_SUCCESS);
}

/**
 * print_prime_factors - factorizes as many numbers as possible into a product
 * of two smaller numbers and prints the result
 * @number: the number to factorize
 */
void print_prime_factors(mpz_t number)
{
	mpz_t odd_prime, quotient;

	mpz_init(odd_prime);
	mpz_init(quotient);

	if (mpz_cmp_si(number, 1) <= 0)
		return; /* skip 1 and anything below it */

	/* let's check whether it's divisible by 2 and perform an early return */
	if (mpz_even_p(number) != 0)
	{
		mpz_divexact_ui(quotient, number, 2);
		gmp_printf("%Zd=%Zd*%d\n", number, quotient, 2);
		return;
	}

	/* let's use odd numbers and step through the number, it wasn't even */
	mpz_set_ui(odd_prime, 3);
	mpz_sqrt(quotient, number);

	while (mpz_cmp(odd_prime, quotient) <= 0)
	{
		if (mpz_divisible_p(number, odd_prime) != 0)
		{
			mpz_divexact(quotient, number, odd_prime);
			/* print the result*/
			gmp_printf("%Zd=%Zd*%Zd\n", number, quotient, odd_prime);

			/* clean up and return */
			mpz_clear(odd_prime);
			mpz_clear(quotient);
			return;
		}

		/* increment odd_prime by 2 */
		mpz_add_ui(odd_prime, odd_prime, 2); /* odd_prime += 2*/
	}

	/* the number is a prime number */
	gmp_printf("%Zd=%Zd*%d\n", number, number, 1);

	/* clean up what was used */
	mpz_clear(odd_prime);
	mpz_clear(quotient);
}
