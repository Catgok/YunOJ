const {CompressionTypes, Kafka, CompressionCodecs} = require('kafkajs')
const SnappyCodec = require('kafkajs-snappy')
const {WebSocketServer} = require('ws')
const url = require('url')
const axios = require('axios')

CompressionCodecs[CompressionTypes.Snappy] = SnappyCodec

const wss = new WebSocketServer({port: 8888});

const kafka = new Kafka({
    clientId: 'YUN-OJ-Kafka-Client-Agent',
    brokers: ['serverhost:9092']
});

const consumer = kafka.consumer({groupId: 'submit-change-notice-consumer-font'})
// 处理从 Kafka 接收到的消息，发送到所有 WebSocket 客户端
const run = async () => {
    await consumer.connect();
    await consumer.subscribe({topic: 'submit-change-notice', fromBeginning: true});

    await consumer.run({
        eachMessage: async ({topic, partition, message}) => {
            const payload = message.value.toString();
            for (let client of clients) {
                // console.log(client.readyState, client.userId, JSON.parse(payload).UserId)
                if (client.readyState === 1 && JSON.parse(payload).UserId === client.userId) client.send(payload);
            }
        },
    });
};

// 保存所有已连接的 WebSocket 客户端
const clients = new Set();

wss.on('connection', async function connection(ws, req) {
    // console.log(req)
    const queryParams = url.parse(req.url, true).query;
    const uToken = queryParams.uToken;
    await axios.post('http://serverhost:12030/v1/user/getUserInfoByUToken', {}, {
        headers: {'U-Token': uToken}
    }).then(response => {
        const res = JSON.parse(JSON.stringify(response.data))
        ws.userId = res.data.userId
    }).catch(error => {
        console.error(error);
    })
    clients.add(ws);
    ws.on('close', () => {
        clients.delete(ws);
    });
});


run().catch(console.error);

