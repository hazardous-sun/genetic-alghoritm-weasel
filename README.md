# Genetic Alghoritm for the Weasel program

This project implements a genetic algorithm using the roulette selection method to solve
the "[Weasel problem](https://en.wikipedia.org/wiki/Weasel_program)". The goal is to evolve a population of strings
towards a target string, "METHINKS IT IS LIKE A WEASEL".

## Key Concepts

- Genetic Algorithm (GA): A search heuristic inspired by natural selection. It iteratively evolves a population of
  candidate solutions towards an optimal solution based on fitness.
- Chromosome: A string representing a candidate solution.
- Mutation: Randomly changing a character in a chromosome.
- Crossover: Combining parts of two chromosomes to create new offspring.
- Adaptability Index: A measure of how close a chromosome is to the target string.

## Parameters:

- ChromossomesCount: Number of chromosomes in the population (default: 100).
- MutationFactor: Number of mutations to introduce per cycle (default: 4).
- Destiny: The target string ("METHINKS IT IS LIKE A WEASEL").
- Prize: The value received after each successfully found part of the sentence.

## Customization

- You can adjust parameters like ChromossomesCount and MutationFactor to experiment with the algorithm's behavior.
- You can modify the Destiny constant to solve a different string optimization problem.

# References

- https://en.wikipedia.org/wiki/Weasel_program