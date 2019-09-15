from concurrent import futures
import time
import grpc

import quote_server
import quote_service_pb2
import quote_service_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24
def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    quote_service_pb2_grpc.add_QuoteServiceServicer_to_server(quote_server.QuoteServer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print('Running on port 50051')
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == "__main__":
    serve()
