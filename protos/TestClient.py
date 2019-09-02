# A simple stubby client for testing out QuoteService
import grpc
import quote_service_pb2
import quote_service_pb2_grpc

def main():
    channel = grpc.insecure_channel('localhost:50051')
    stub = quote_service_pb2_grpc.QuoteServiceStub(channel)
    req = quote_service_pb2.QuoteRequest()
    req.symbol = 'GOOG'
    googQuotes = stub.ListQuoteHistory(req)

    for googQuote in googQuotes.quotes:
        print(googQuote.close)

if __name__ == '__main__':
    main()
