import quote_pb2
import quote_pb2_grpc
from yahoo_finance_crawler import YahooFinanceCrawler

class QuoteServer(quote_pb2_grpc.QuoteServiceServicer):
    def ListQuoteHistory(self, request, cxt):
        quoteResponse = quote_pb2.QuoteResponse()
        quotes = YahooFinanceCrawler.getQuoteHistory(request.symbol)
        quoteResponse.quotes.extend(quotes)
        return quoteResponse
