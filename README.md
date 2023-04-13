## Prerequisites

* Install [go](https://golang.org/doc/install) compiler.

## Instructions

1. Create a project dir under `$GOPATH/src`

```sh
mkdir -p $GOPATH/src/github.com/bwwu/stockbuddy
cd $GOPATH/src/github.com/bwwu/stockbuddy
git clone git@github.com:bwwu/stockbuddy.git
```

2. Compile stockbuddy.

```sh
go build .
```

3. Set TradingBot password as an env var.

```sh
export STOCKBUDDY_PASSWORD=<password>
```

4. Run stockbuddy with appropriate flags, for example:

```sh
./stockbuddy --mail_to="foo@example.com,bar@example.com"
```

## Runtime flags

|	Name	|	Description		|Example Usage		|
|---------------|-------------------------------|-----------------------|
|mail_to	|comma-separated list of emails	|"a@foo.com,b@bar.com"	|
|nomail		|run stockbuddy without sending	| --nomail=True		|
|use_watchlist	|path to csvs of ticker symbols |"path/to/stocks.csv"	|

Default watchlist is located at [watchlists/default.txt](
https://github.com/bwwu/stockbuddy/blob/master/watchlists/default.txt).
