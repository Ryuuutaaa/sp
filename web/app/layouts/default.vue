<template>
  <div class="min-h-screen flex font-sans">
    <aside class="w-60 bg-gray-900 text-white p-4 flex flex-col gap-2">
      <h2 class="text-lg font-bold mb-4">Koperasi SP</h2>
      <NuxtLink to="/dashboard" class="px-3 py-2 rounded hover:bg-gray-700">Dashboard</NuxtLink>
      <NuxtLink to="/members" class="px-3 py-2 rounded hover:bg-gray-700">Anggota</NuxtLink>
      <NuxtLink to="/savings" class="px-3 py-2 rounded hover:bg-gray-700">Simpanan</NuxtLink>
      <NuxtLink to="/loans" class="px-3 py-2 rounded hover:bg-gray-700">Pinjaman</NuxtLink>
      <button class="mt-auto px-3 py-2 rounded bg-red-600 hover:bg-red-700" @click="logout">Logout</button>
    </aside>
    <main class="flex-1 p-6 bg-gray-50">
      <slot />
    </main>
  </div>
</template>

<script setup lang="ts">
import { authClient } from "~/utils/auth-client"

async function logout() {
  await authClient.signOut()
  await navigateTo("/auth/login")
}
</script>
