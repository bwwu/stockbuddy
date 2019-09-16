import logging

import quote_pb2
import quote_pb2_grpc
from yahoo_finance_crawler import YahooFinanceCrawler

logger = logging.getLogger('quote_service')

class QuoteServer(quote_pb2_grpc.QuoteServiceServicer):
    def ListQuoteHistory(self, request, cxt):
        logger.info('ListQuoteHistory: Symbol=%s, Period=%d' % (request.symbol, request.period))
        quoteResponse = quote_pb2.QuoteResponse()
        quotes = YahooFinanceCrawler.getQuoteHistory(request.symbol, request.period)
        quoteResponse.quotes.extend(quotes)
        logger.info('ListQuoteHistory: Found %d quotes for %s' % (len(quotes), request.symbol))
        return quoteResponse
