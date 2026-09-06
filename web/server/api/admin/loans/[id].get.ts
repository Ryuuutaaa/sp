export default defineEventHandler((event) => proxyGet(event, `/api/loans/${getRouterParam(event, "id")}`))
