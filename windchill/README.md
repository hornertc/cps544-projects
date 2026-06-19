[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/Jd7GRuX5)
# Assignment: Wind Chill

![Wind Chill Chart](https://www.weather.gov/images/ama/Wind_Chill/windchill.gif)

## Overview

How cold is it outside? The temperature alone is not enough to provide the answer. Other factors including wind speed, relative humidity, and sunshine play important roles in determining coldness outside. In 2001, the National Weather Service (NWS) implemented the new wind-chill temperature to measure the coldness using temperature and wind speed. The formula is

$$
t_{wc} = 35.74 + 0.6215t_a - 35.75v^{0.16} + 0.4275t_a v^{0.16}
$$

where $t_a$ is the outside temperature measured in degrees Fahrenheit, $v$ is the speed measured in miles per hour, and $t_{wc}$ is the wind-chill temperature. The formula cannot be used for wind speeds below 2 mph or temperatures below -58°F or above 41°F.

## Learning Objectives

- Parse standard input
- Output formatting
- Mathematical computations
- Flag parsing

## Requirements

1. Write a Go program and put your file(s) in `cmd/windchill/`.

1. The program must prompt the user to enter a temperature between -58°F  and 41°F then a wind speed greater than or equal to 2 then displays the compute wind-chill temperature.  If the user enters an invalid number due it being out of range or not a number, prompt the user again until they enter a valid value.  Output two digits past the decimal place (hundredths of °F).

1. Add a flag to the program called `-table` that displays a table instead of prompting.  The table should have the same bounds and format as the above table except that it should show the wind chill to the tenths place.  The decimal place must be vertically aligned (see the example) so it is easy to read.

1. The table must be computed from the formula above.  Do not store the table in your code.

1. Do not use any library other than the Go standard library.

1. The source code must compile with the most recent version of the Go compiler.

1. The program must not panic under any circumstances.

1. Make sure your code is "gofmt'd".  See `gofmt` or better use `goimports` or better yet configure IDE to do this formatting on file save.

## Hints

- The `math`, `fmt`, and `strconv` packages might be useful for this assignment.
- Use the `Makefile`.  For example, you can run tests with `make test -B`.  To make sure your code is properly formatted run `make fix`.

## Example Output

```shell
$ go run windchill.go
Enter the temperature in °F: 6.3
Enter the wind speed in miles per hour: 4  
The wind chill index is -1.61°F

$ go run windchill.go -table
 36.5  30.6  24.7  18.9  13.0   7.1   1.2  -4.6 -10.5 -16.4 -22.3 -28.1 -34.0 -39.9 -45.7 -51.6 -57.5 -63.4 
 33.6  27.4  21.2  15.1   8.9   2.7  -3.5  -9.7 -15.9 -22.1 -28.3 -34.5 -40.7 -46.9 -53.1 -59.3 -65.5 -71.7 
 31.8  25.4  19.0  12.6   6.2  -0.2  -6.6 -13.0 -19.4 -25.8 -32.2 -38.6 -45.0 -51.4 -57.8 -64.2 -70.6 -77.0 
 30.5  23.9  17.4  10.8   4.2  -2.3  -8.9 -15.4 -22.0 -28.6 -35.1 -41.7 -48.2 -54.8 -61.4 -67.9 -74.5 -81.0 
 29.4  22.7  16.0   9.3   2.6  -4.0 -10.7 -17.4 -24.1 -30.8 -37.5 -44.1 -50.8 -57.5 -64.2 -70.9 -77.6 -84.3 
 28.5  21.7  14.9   8.1   1.3  -5.5 -12.3 -19.1 -25.9 -32.7 -39.4 -46.2 -53.0 -59.8 -66.6 -73.4 -80.2 -87.0 
 27.7  20.8  13.9   7.0   0.1  -6.8 -13.6 -20.5 -27.4 -34.3 -41.2 -48.1 -54.9 -61.8 -68.7 -75.6 -82.5 -89.3 
 26.9  20.0  13.0   6.1  -0.9  -7.9 -14.8 -21.8 -28.8 -35.7 -42.7 -49.7 -56.6 -63.6 -70.6 -77.5 -84.5 -91.4 
 26.3  19.3  12.2   5.2  -1.8  -8.9 -15.9 -23.0 -30.0 -37.0 -44.1 -51.1 -58.1 -65.2 -72.2 -79.3 -86.3 -93.3 
 25.7  18.6  11.5   4.4  -2.7  -9.8 -16.9 -24.0 -31.1 -38.2 -45.3 -52.4 -59.5 -66.6 -73.7 -80.8 -87.9 -95.1 
 25.2  18.0  10.9   3.7  -3.5 -10.6 -17.8 -25.0 -32.1 -39.3 -46.5 -53.6 -60.8 -68.0 -75.1 -82.3 -89.5 -96.6 
 24.7  17.5  10.2   3.0  -4.2 -11.4 -18.6 -25.9 -33.1 -40.3 -47.5 -54.8 -62.0 -69.2 -76.4 -83.7 -90.9 -98.1

$ go run windchill.go
Enter the temperature in °F: -92
-92.000000 is out of range.  Valid temperatures are between -58°F and 41°F.
Enter the temperature in °F: 13
Enter the wind speed in miles per hour: 1
1.000000 is out of range.  Valid wind speeds are above 2 mph.
Enter the wind speed in miles per hour: x
strconv.ParseFloat: parsing "": invalid syntax
Enter the wind speed in miles per hour: unexpected newline
Enter the wind speed in miles per hour: 4
The wind chill index is 6.13°F 
```

## Submission

- Commit and push your working code to your GIT repository.
- Ensure all tests pass otherwise you may receive no credit.
