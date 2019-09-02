import quote_service_pb2
import quote_service_pb2_grpc
from YahooFinanceCrawler import YahooFinanceCrawler

class QuoteServer(quote_service_pb2_grpc.QuoteServiceServicer):
    def ListQuoteHistory(self, request, cxt):
        quoteResponse = quote_service_pb2.QuoteResponse()
        quotes = YahooFinanceCrawler.getQuoteHistory(request.symbol)
        quoteResponse.quotes.extend(quotes)
        return quoteResponse
