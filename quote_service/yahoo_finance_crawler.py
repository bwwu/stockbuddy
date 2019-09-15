import datetime
import requests
import quote_pb2

class YahooFinanceCrawler:
    # PRF=t%3D<SYMBOL>
    cookies = { "B": "ajch0f5elj4sp", "APID": "1Adf2ce59c-c1e2-11e9-adce-025f25c4bfdc", "APIDTS": "1566870258", "PRF": ""}


    BASE_URL = "https://query1.finance.yahoo.com/v7/finance/download/{}?period1={}&period2={}&interval=1d&events=history&crumb=pRB6UiIiFnn"

    # Get history price history for the past year for a symbol
    def getQuoteHistory(symbol, days = 365):
        cookiesForRequest = YahooFinanceCrawler.cookies.copy()
        cookiesForRequest["PRF"] = "t%3d{}".format(symbol)

        # Period query params should be 1 year ago and today respectively.
        periodEnd = datetime.datetime.now()
        periodStart = periodEnd - datetime.timedelta(days=days)
        url = YahooFinanceCrawler.BASE_URL.format(symbol, int(periodStart.timestamp()), int(periodEnd.timestamp()))
        
        r = requests.get(url, cookies = cookiesForRequest)

        quoteProtos = []
        for line in r.text.split('\n'):
            rowItems = line.split(',')
            if len(rowItems) == 7 and rowItems[0] != 'Date':
                quote = createQuoteProtoFromRow(symbol, rowItems)
                quoteProtos.append(quote)
        return quoteProtos
                


def createQuoteProtoFromRow(symbol, rowItems):
    quote = quote_pb2.Quote()
    quote.symbol = symbol
    quote.timestamp = int(datetime.datetime.strptime(rowItems[0], '%Y-%m-%d').timestamp())
    quote.open = float(rowItems[1])
    quote.high = float(rowItems[2])
    quote.low = float(rowItems[3])
    quote.close = float(rowItems[4])
    quote.adj_close = float(rowItems[5])
    quote.volume = int(rowItems[6])
    return quote
    
