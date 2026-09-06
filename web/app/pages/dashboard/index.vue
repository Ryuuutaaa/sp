<template>
  <div>
    <h1 class="text-2xl font-bold mb-1">Dashboard</h1>
    <p class="text-sm text-gray-500 mb-4">Peran: {{ role || "-" }}</p>

    <!-- ADMIN / SUPER ADMIN -->
    <div v-if="isStaff" class="grid grid-cols-3 gap-4">
      <div class="bg-white p-4 rounded shadow">Anggota: {{ members?.length ?? "-" }}</div>
      <div class="bg-white p-4 rounded shadow">Pinjaman: {{ loans?.length ?? "-" }}</div>
      <div class="bg-white p-4 rounded shadow">Tipe Simpanan: {{ types?.length ?? "-" }}</div>
      <NuxtLink to="/members" class="bg-white p-4 rounded shadow hover:bg-gray-100">Verifikasi anggota →</NuxtLink>
      <NuxtLink to="/loans" class="bg-white p-4 rounded shadow hover:bg-gray-100">Review pinjaman →</NuxtLink>
      <NuxtLink to="/savings" class="bg-white p-4 rounded shadow hover:bg-gray-100">Transaksi simpanan →</NuxtLink>
    </div>

    <!-- TELLER -->
    <div v-else-if="role === 'teller'" class="grid grid-cols-2 gap-4">
      <NuxtLink to="/savings" class="bg-white p-4 rounded shadow hover:bg-gray-100">Setor / tarik simpanan →</NuxtLink>
      <NuxtLink to="/loans" class="bg-white p-4 rounded shadow hover:bg-gray-100">Input bayar angsuran →</NuxtLink>
      <div class="bg-white p-4 rounded shadow">Total anggota: {{ members?.length ?? "-" }}</div>
      <div class="bg-white p-4 rounded shadow">Pinjaman aktif: {{ loans?.length ?? "-" }}</div>
    </div>

    <!-- ANGGOTA -->
    <div v-else class="grid grid-cols-2 gap-4">
      <div class="bg-white p-4 rounded shadow">Transaksi saya: {{ mySavings?.length ?? "-" }}</div>
      <div class="bg-white p-4 rounded shadow">Tipe simpanan: {{ types?.length ?? "-" }}</div>
      <NuxtLink to="/loans" class="bg-white p-4 rounded shadow hover:bg-gray-100">Ajukan / lihat pinjaman →</NuxtLink>
      <NuxtLink to="/savings" class="bg-white p-4 rounded shadow hover:bg-gray-100">Riwayat simpanan →</NuxtLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import { authClient } from "~/utils/auth-client"

definePageMeta({ middleware: ["auth"] })

const headers = import.meta.server ? useRequestHeaders(["cookie"]) : undefined
const { data: session } = await authClient.getSession({
  fetchOptions: headers ? { headers } : {},
})
const role = (session?.user as any)?.role as string | undefined
const memberId = (session?.user as any)?.memberId as string | undefined
const isStaff = computed(() => role === "super_admin" || role === "admin")

const { data: members } = await useMembers()
const { data: loans } = await useLoans()
const { data: types } = await useSavingsTypes()
const { data: mySavings } = await useFetch(
  () => (role === "anggota" && memberId ? `/api/anggota/savings?memberId=${memberId}` : null),
)
</script>
