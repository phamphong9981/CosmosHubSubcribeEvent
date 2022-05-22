import json
import websockets
import asyncio
import requests
import redis
import time

HOST = 'localhost'
PORT = '6379'

if __name__ == '__main__':
    r = redis.Redis(host=HOST, port=PORT)
    pub = r.pubsub()
    pub.psubscribe("*")

    while True:
        data = pub.get_message()
        if data:
            message = data['data']
            if message and message != 1:
                requests.post("https://discordapp.com/api/webhooks/965300320194408488/YQpHtPDU9j31k2eCa_ScJtSaJ6cxtRRk5vNGv75AFCMR4YTOgtFnxvFvfJHVq5RVblnv", { "content": message })

        time.sleep(1)

# pubsub.run_in_thread(sleep_time=.01)
# command={"jsonrpc": "2.0","method":"subscribe","id": 0,"params": {"query": "tm.event='Tx' AND unbond.validator='cosmosvaloper178h4s6at5v9cd8m9n7ew3hg7k9eh0s6wptxpcn'"}}
# async def hello():
#     r = redis.Redis(host=HOST, port=PORT)
#     pub = r.pubsub()
#     pub.subscribe(CHANNEL)
#     async with websockets.connect("wss://rpc-atom-testnet.aura.network/websocket") as websocket:
#         await websocket.send(json.dumps(command))
#         async for message in websocket:
#             requests.post("https://discordapp.com/api/webhooks/965300320194408488/YQpHtPDU9j31k2eCa_ScJtSaJ6cxtRRk5vNGv75AFCMR4YTOgtFnxvFvfJHVq5RVblnv", { "content": message })

# asyncio.run(hello())



# requests.post("https://discordapp.com/api/webhooks/965300320194408488/YQpHtPDU9j31k2eCa_ScJtSaJ6cxtRRk5vNGv75AFCMR4YTOgtFnxvFvfJHVq5RVblnv", { "content": message })