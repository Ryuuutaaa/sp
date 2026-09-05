<template>
  <div>
    <h1 class="text-xl font-bold mb-4">Login Koperasi</h1>
    <form class="flex flex-col gap-3" @submit.prevent="login">
      <input v-model="email" type="email" placeholder="Email" class="border px-3 py-2 rounded" required />
      <input v-model="password" type="password" placeholder="Password" class="border px-3 py-2 rounded" required />
      <button class="bg-blue-600 text-white px-4 py-2 rounded">Login</button>
      <p v-if="error" class="text-red-600 text-sm">{{ error }}</p>
      <p class="text-sm">Belum punya akun? <NuxtLink to="/auth/register" class="text-blue-600 underline">Daftar anggota</NuxtLink></p>
    </form>
  </div>
</template>

<script setup lang="ts">
import { z } from "zod"
import { authClient } from "~/utils/auth-client"

definePageMeta({ layout: "auth" })

const email = ref("")
const password = ref("")
const error = ref("")

const schema = z.object({ email: z.string().email(), password: z.string().min(8) })

async function login() {
  error.value = ""
  const parsed = schema.safeParse({ email: email.value, password: password.value })
  if (!parsed.success) {
    error.value = "Email / password tidak valid"
    return
  }
  const res = await authClient.signIn.email({ email: email.value, password: password.value })
  if (res.error) {
    error.value = res.error.message || "Login gagal"
    return
  }
  await navigateTo("/dashboard")
}
</script>
