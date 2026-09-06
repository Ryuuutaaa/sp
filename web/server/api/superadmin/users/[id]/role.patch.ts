export default defineEventHandler((event) => proxyPatch(event, `/api/users/${getRouterParam(event, "id")}/role`))
