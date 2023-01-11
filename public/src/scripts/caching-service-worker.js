(function () {
  if (!self.caches) {
    return;
  }

  const cacheName = 'v1';

  async function putInCache(req, res) {
    const cache = await self.caches.open(cacheName);
    await cache.put(req, res);
  }

  async function serverFirst(e) {
    try {
      const res = await fetch(e.request.url);
      if (res) {
        console.log("From server");
        e.waitUntil(putInCache(e.request, res.clone()));
        return res;
      }
    } catch (err) {
      const cacheRes = await self.caches.match(e.request);
      if (cacheRes) {
        console.log("From cache");
        return cacheRes;
      }
      return new Request('Resource not found', {
        status: 408,
        headers: {
          'Content-Type': 'text/html',
        },
      })
    }
  }

  self.addEventListener('fetch', async (e) => {
    if (e.request.method === "GET") {
      console.log(e.request.url);
      e.respondWith(serverFirst(e));
    }
  });
})()