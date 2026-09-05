export default defineNuxtConfig({
  compatibilityDate: '2025-01-01',
  future: {
    compatibilityVersion: 4,
  },
  modules: [
    '@unocss/nuxt',
  ],
  runtimeConfig: {
    backendUrl: process.env.BACKEND_URL || 'http://localhost:8080',
    betterAuthSecret: process.env.BETTER_AUTH_SECRET || 'secret-key-12345678901234567890123456',
    public: {
      backendUrl: process.env.PUBLIC_BACKEND_URL || 'http://localhost:8080',
    }
  }
})
