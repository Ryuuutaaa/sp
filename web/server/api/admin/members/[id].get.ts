export default defineEventHandler((event) => proxyGet(event, `/api/members/${getRouterParam(event, "id")}`))
