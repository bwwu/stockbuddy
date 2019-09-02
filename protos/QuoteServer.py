import quote_service_pb2
import quote_service_pb2_grpc
import YahooFinanceCrawler

class QuoteServer(quote_service_pb2_grpc.QuoteServiceServicer):
    def ListQuoteHistory(self, request, cxt):
        quoteResponse = quote_service_pb2.QuoteResponse()
        quoteResponse.quotes = YahooFinanceCrawler(request.symbol)
        return quoteResponse
