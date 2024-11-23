const redis = require('redis');

const client = redis.createClient();
client.connect().then(() => {
    console.log('Redis connected');
}).catch((err) => {
    console.error(err);
})

const namespace = 'discourse-node:';

function SetCache(key, value) {
    return client.set(namespace + key, JSON.stringify(value));
}

function GetCache(key) {
    return client.get(namespace + key);
}

module.exports = {
    SetCache,
    GetCache,
};
