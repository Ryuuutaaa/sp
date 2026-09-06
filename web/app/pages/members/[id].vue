<template>
  <div>
    <NuxtLink to="/members" class="text-sm text-blue-600 underline">← Kembali</NuxtLink>
    <h1 class="text-2xl font-bold my-4">Detail Anggota</h1>
    <div v-if="member" class="bg-white p-6 rounded shadow grid md:grid-cols-2 gap-4 text-sm">
      <div><p class="text-gray-500">Nomor</p><p class="font-medium">{{ member.memberNumber }}</p></div>
      <div><p class="text-gray-500">Nama</p><p class="font-medium">{{ member.name }}</p></div>
      <div><p class="text-gray-500">NIK</p><p class="font-medium">{{ member.nik }}</p></div>
      <div><p class="text-gray-500">HP</p><p class="font-medium">{{ member.phone }}</p></div>
      <div><p class="text-gray-500">Alamat</p><p class="font-medium">{{ member.address }}</p></div>
      <div><p class="text-gray-500">Pekerjaan</p><p class="font-medium">{{ member.occupation }}</p></div>
      <div><p class="text-gray-500">Verifikasi</p><p class="font-medium">{{ member.verificationStatus }}</p></div>
      <div><p class="text-gray-500">Status</p><p class="font-medium">{{ member.status }}</p></div>
      <div>
        <p class="text-gray-500 mb-1">Foto KTP</p>
        <img :src="member.ktpPhotoUrl" alt="KTP" class="h-40 object-contain border rounded" />
      </div>
      <div>
        <p class="text-gray-500 mb-1">Selfie + KTP</p>
        <img :src="member.selfieKtpPhotoUrl" alt="Selfie" class="h-40 object-contain border rounded" />
      </div>
    </div>
    <div v-if="canManage && member" class="flex gap-2 mt-4">
      <button v-if="member.verificationStatus === 'pending'" class="px-4 py-2 bg-green-600 text-white rounded" @click="verify('verified')">Verifikasi</button>
      <button v-if="member.verificationStatus === 'pending'" class="px-4 py-2 border rounded" @click="verify('rejected')">Tolak</button>
      <button v-if="member.status === 'active'" class="px-4 py-2 bg-red-600 text-white rounded" @click="setStatus('inactive')">Nonaktifkan</button>
      <button v-else class="px-4 py-2 bg-blue-600 text-white rounded" @click="setStatus('active')">Aktifkan</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { authClient } from "~/utils/auth-client"

definePageMeta({ middleware: ["auth", "role"], allowedRoles: ["super_admin", "admin", "teller"] })

const route = useRoute()
const id = route.params.id as string
const { data: session } = await authClient.getSession()
const role = (session?.user as any)?.role
const canManage = role === "super_admin" || role === "admin"

const { data: member, refresh } = await useFetch(`/api/admin/members/${id}`)

async function verify(status: string) {
  await $fetch(`/api/admin/members/${id}/verify`, { method: "PATCH", body: { status, verifiedBy: (session?.user as any)?.id } })
  await refresh()
}

async function setStatus(status: string) {
  await $fetch(`/api/admin/members/${id}/status`, { method: "PATCH", body: { status } })
  await refresh()
}
</script>
