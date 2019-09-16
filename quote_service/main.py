from concurrent import futures
import grpc
import time
import logging

import quote_server
import quote_pb2
import quote_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

logger = logging.getLogger('quote_service')
logger.setLevel(logging.DEBUG)

# create console handler and set level to debug
ch = logging.StreamHandler()
ch.setLevel(logging.INFO)

# create formatter
formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')

# add formatter to ch
ch.setFormatter(formatter)

# add ch to logger
logger.addHandler(ch)
def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    quote_pb2_grpc.add_QuoteServiceServicer_to_server(quote_server.QuoteServer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    logger.info('Listening on port :50051')
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == "__main__":
    serve()
