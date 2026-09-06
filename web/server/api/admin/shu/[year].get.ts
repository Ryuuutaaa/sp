export default defineEventHandler((event) => proxyGet(event, `/api/shu/${getRouterParam(event, "year")}`))
