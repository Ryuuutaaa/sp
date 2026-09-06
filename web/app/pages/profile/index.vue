<template>
  <div class="max-w-lg">
    <h1 class="text-2xl font-bold mb-4">Profil Saya</h1>
    <div class="bg-white p-6 rounded shadow text-sm flex flex-col gap-2">
      <div><p class="text-gray-500">Nama</p><p class="font-medium">{{ user?.name }}</p></div>
      <div><p class="text-gray-500">Email</p><p class="font-medium">{{ user?.email }}</p></div>
      <div><p class="text-gray-500">Peran</p><p class="font-medium">{{ user?.role }}</p></div>
      <template v-if="member">
        <hr />
        <div><p class="text-gray-500">Nomor Anggota</p><p class="font-medium">{{ member.memberNumber }}</p></div>
        <div><p class="text-gray-500">NIK</p><p class="font-medium">{{ member.nik }}</p></div>
        <div><p class="text-gray-500">Status Verifikasi</p><p class="font-medium">{{ member.verificationStatus }}</p></div>
      </template>
      <p v-else class="text-orange-600">Profil anggota belum tertaut. <NuxtLink to="/auth/register" class="underline">Lengkapi data anggota</NuxtLink></p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { authClient } from "~/utils/auth-client"

definePageMeta({ middleware: ["auth", "role"], allowedRoles: ["anggota"] })

const { data: session } = await authClient.getSession()
const user = session?.user as any
const { data: member } = await useFetch("/api/anggota/profile")
</script>
