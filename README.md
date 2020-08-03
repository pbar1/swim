# Swim

Swimming times utility

## Usage

```
Enumerates all possible NCAA Event Rank Search parameters for SCY. This
includes both male and female, all named date ranges, all conferences, and all legal
competition events. Results are saved to disk in CSV files corresponding to the
search parameters that were used for that file. No results file will be written if
the query returns an error. A "results" directory in the current working directory
must already exist!

Values for DIVISION are one of: d1|d2|d3

Usage:
  swim mirror [DIVISION] [flags]

Flags:
  -h, --help                      help for mirror
  -p, --pool-size int             Number of concurrent threads (default 10)
  -t, --search-timeout duration   Seconds to allow searches before timing out (default 6s)
```
