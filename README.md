# linear-stats

This program profect statistical calculations focusing on the Linear Regression Line and the Pearson Correlation Coefficient. The data is read from a file, and the program prints the results of each statistical calculation.

## Overview

The program reads numerical data from a specified file and computes two important statistical measures:

1. **Linear Regression Line**: A linear equation representing the best fit line through the data points.
2. **Pearson Correlation Coefficient**: A measure of the linear correlation between two datasets.

## Data Format

The data should be in a text file where each line contains a single numeric value. The values represent the y-axis coordinates of a graph, with the x-axis coordinates being the line numbers (0, 1, 2, 3, ...).

### Example of data.txt

```
189
113 
121 
114 
145 
110
...
```

## Cloning the Repository

To clone this repository, use the following command in your terminal:

```bash
git clone https://learn.zone01kisumu.ke/git/hiombima/linear-stats.git
cd linear-stats
```

## Usage

To run the program, use the following command in your terminal:

```go
go run . data.txt
```

### Output Format
The program will print the results in the following format:
```
Linear Regression Line: y = <value>x + <value>
Pearson Correlation Coefficient: <value>
```

- The values for the Linear Regression Line will be displayed with 6 decimal places.
- The Pearson Correlation Coefficient value will be displayed with 10 decimal places.

### Example Output
```
Linear Regression Line: y = 1.200000x + 100.000000
Pearson Correlation Coefficient: 0.8944271909
```
## Functions

Reads numeric data from a specified file and returns a slice of float64 values.

### LinearRegression
Calculates the slope and intercept of the best fit line using the linear regression formula.

### Correlation
Calculates the Pearson correlation coefficient for the provided dataset.

## Contribution

Follow these steps if you would like to contribute to this project or if you've discovered a bug.

1. **Fork the repository:** Click the "Fork" button at the top-right corner of the repository page.
2. **Clone your fork:** Clone the forked repository to your local machine using:

```bas
git clone https://github.com/<your-username>/<your-forked-repo>.git
```
3. **Create a new branch:** Create a branch for your feature or bug fix.

```bash
git checkout -b feature/your-feature-name
```
4. **Make your changes:** Implement your feature or bug fix, ensuring code quality.
5. **Run tests:** Make sure your changes do not break existing functionality. Write tests where applicable.
6. **Commit your changes:** Follow good commit message practices (e.g., `git commit -m "Add feature X`").
7. **Push your branch:** Push your changes to your fork.
8. **Create a Pull Request:** Go to the original repository, and click "Compare & pull request." Describe your changes in detail.

## License
This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.