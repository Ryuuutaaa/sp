<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">Daftar Anggota</h1>
    <input v-model="q" placeholder="Cari nama / nomor / NIK" class="border px-3 py-2 rounded mb-4 w-full max-w-md" />
    <table class="w-full bg-white rounded shadow text-sm">
      <thead><tr class="border-b text-left"><th class="p-2">No</th><th class="p-2">Nama</th><th class="p-2">NIK</th><th class="p-2">Verifikasi</th><th class="p-2">Status</th><th class="p-2">Aksi</th></tr></thead>
      <tbody>
        <tr v-for="m in filtered" :key="m.id" class="border-b">
          <td class="p-2">{{ m.memberNumber }}</td>
          <td class="p-2">{{ m.name }}</td>
          <td class="p-2">{{ m.nik }}</td>
          <td class="p-2">{{ m.verificationStatus }}</td>
          <td class="p-2">{{ m.status }}</td>
          <td class="p-2 flex gap-2">
            <NuxtLink :to="`/members/${m.id}`" class="text-blue-600 underline">Detail</NuxtLink>
            <button v-if="canManage && m.verificationStatus === 'pending'" class="text-green-600 underline" @click="verify(m.id, 'verified')">Verifikasi</button>
            <button v-if="canManage && m.status === 'active'" class="text-red-600 underline" @click="setStatus(m.id, 'inactive')">Nonaktifkan</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { authClient } from "~/utils/auth-client"

definePageMeta({ middleware: ["auth", "role"], allowedRoles: ["super_admin", "admin", "teller"] })

const { data: session } = await authClient.getSession()
const role = (session?.user as any)?.role
const canManage = role === "super_admin" || role === "admin"

const q = ref("")
const { data: members, refresh } = await useMembers()
const filtered = computed(() => {
  const list = (members.value as any[]) || []
  if (!q.value) return list
  const s = q.value.toLowerCase()
  return list.filter((m: any) => [m.name, m.memberNumber, m.nik].join(" ").toLowerCase().includes(s))
})

async function verify(id: string, status: string) {
  const { data: s } = await authClient.getSession()
  await $fetch(`/api/admin/members/${id}/verify`, { method: "PATCH", body: { status, verifiedBy: (s?.user as any)?.id } })
  await refresh()
}

async function setStatus(id: string, status: string) {
  await $fetch(`/api/admin/members/${id}/status`, { method: "PATCH", body: { status } })
  await refresh()
}
</script>
