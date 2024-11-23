const { GetCache, SetCache } = require('./cache');

const config = {
		Url: process.env.URL,          	
        ApiKey: process.env.API_KEY,
		AdminUserName: process.env.ADMIN_USERNAME
}

async function GetLatestPost() {
    const cache = await GetCache('latest-posts');
    if (cache) {
        return JSON.parse(cache);
    }

    const res = await fetch(`${config.Url}/posts.json`, {
        method: 'GET',
        headers: {
            'Api-Key': config.ApiKey,
            'Api-Username': config.AdminUserName,
        },
    });
    
    const json = await res.json();
    
    await SetCache('latest-posts', JSON.stringify(json));
    return json;
}

module.exports = {
    GetLatestPost,
};
