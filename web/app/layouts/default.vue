<template>
  <div class="min-h-screen flex font-sans">
    <aside class="w-60 bg-gray-900 text-white p-4 flex flex-col gap-2">
      <h2 class="text-lg font-bold mb-4">Koperasi SP</h2>
      <NuxtLink to="/dashboard" class="px-3 py-2 rounded hover:bg-gray-700">Dashboard</NuxtLink>
      <template v-if="isStaff || role === 'teller'">
        <NuxtLink to="/members" class="px-3 py-2 rounded hover:bg-gray-700">Anggota</NuxtLink>
      </template>
      <NuxtLink to="/savings" class="px-3 py-2 rounded hover:bg-gray-700">Simpanan</NuxtLink>
      <NuxtLink to="/loans" class="px-3 py-2 rounded hover:bg-gray-700">Pinjaman</NuxtLink>
      <NuxtLink to="/installments" class="px-3 py-2 rounded hover:bg-gray-700">Angsuran</NuxtLink>
      <template v-if="isAdmin">
        <NuxtLink to="/cash" class="px-3 py-2 rounded hover:bg-gray-700">Kas</NuxtLink>
        <NuxtLink to="/shu" class="px-3 py-2 rounded hover:bg-gray-700">SHU</NuxtLink>
        <NuxtLink to="/reports" class="px-3 py-2 rounded hover:bg-gray-700">Laporan</NuxtLink>
      </template>
      <NuxtLink v-if="role === 'teller'" to="/reports/harian" class="px-3 py-2 rounded hover:bg-gray-700">Harian</NuxtLink>
      <NuxtLink v-if="role === 'anggota'" to="/shu" class="px-3 py-2 rounded hover:bg-gray-700">SHU Saya</NuxtLink>
      <NuxtLink v-if="role === 'anggota'" to="/profile" class="px-3 py-2 rounded hover:bg-gray-700">Profil</NuxtLink>
      <template v-if="role === 'super_admin'">
        <NuxtLink to="/settings" class="px-3 py-2 rounded hover:bg-gray-700">Pengaturan</NuxtLink>
        <NuxtLink to="/users" class="px-3 py-2 rounded hover:bg-gray-700">Users</NuxtLink>
        <NuxtLink to="/koperasi" class="px-3 py-2 rounded hover:bg-gray-700">Koperasi</NuxtLink>
      </template>
      <button class="mt-auto px-3 py-2 rounded bg-red-600 hover:bg-red-700" @click="logout">Logout</button>
    </aside>
    <main class="flex-1 p-6 bg-gray-50">
      <slot />
    </main>
  </div>
</template>

<script setup lang="ts">
import { authClient } from "~/utils/auth-client"

const { data: session } = await authClient.getSession()
const role = (session?.user as any)?.role as string | undefined
const isStaff = computed(() => role === "super_admin" || role === "admin")
const isAdmin = computed(() => role === "super_admin" || role === "admin")

async function logout() {
  await authClient.signOut()
  await navigateTo("/auth/login")
}
</script>
