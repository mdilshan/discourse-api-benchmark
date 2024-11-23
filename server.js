const Fastify = require('fastify');
const { GetLatestPost } = require('./discourse/discourse');
const app = Fastify();

app.get('/latest-posts', async (request, reply) => {
    const res = await GetLatestPost();
    return res;
});

app.listen({ port: 5001 }, (err, address) => {
    if (err) {
        console.error(err);
        process.exit(1);
    }
    console.log(`Server listening at ${address}`);
});

