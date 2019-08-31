import datetime
import requests

class YahooFinanceCrawler:
    # PRF=t%3D<SYMBOL>
    cookies = { "B": "ajch0f5elj4sp", "APID": "1Adf2ce59c-c1e2-11e9-adce-025f25c4bfdc", "APIDTS": "1566870258", "PRF": ""}


    BASE_URL = "https://query1.finance.yahoo.com/v7/finance/download/{}?period1={}&period2={}&interval=1d&events=history&crumb=pRB6UiIiFnn"

    # Get history price history for the past year for a symbol
    def getSymbolHistory(symbol):
        cookiesForRequest = YahooFinanceCrawler.cookies.copy()
        cookiesForRequest["PRF"] = "t%3d{}".format(symbol)

        # Period query params should be 1 year ago and today respectively.
        periodEnd = datetime.datetime.now()
        periodStart = periodEnd - datetime.timedelta(days=365)
        url = YahooFinanceCrawler.BASE_URL.format(symbol, int(periodStart.timestamp()), int(periodEnd.timestamp()))
        
        r = requests.get(url, cookies = cookiesForRequest)
        for line in r.text.split('\n'):
            #op,high,low,close,ajd,vol,*_ = line.split(',')
            #print(op)
            print(str(len(line.split(','))) + ' ' + line)

