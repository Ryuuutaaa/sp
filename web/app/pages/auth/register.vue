<template>
  <div>
    <h1 class="text-xl font-bold mb-1">Daftar Anggota</h1>
    <p class="text-sm text-gray-500 mb-4">Langkah {{ step }} dari 2</p>

    <form v-if="step === 1" class="flex flex-col gap-3" @submit.prevent="nextStep">
      <input v-model="form.name" placeholder="Nama lengkap" class="border px-3 py-2 rounded" required />
      <input v-model="form.email" type="email" placeholder="Email" class="border px-3 py-2 rounded" required />
      <input v-model="form.password" type="password" placeholder="Password (min 8)" class="border px-3 py-2 rounded" required />
      <button class="bg-blue-600 text-white px-4 py-2 rounded">Lanjut</button>
      <p v-if="error" class="text-red-600 text-sm">{{ error }}</p>
    </form>

    <form v-else class="flex flex-col gap-3" @submit.prevent="submit">
      <input v-model="form.memberNumber" placeholder="Nomor anggota" class="border px-3 py-2 rounded" required />
      <input v-model="form.nik" placeholder="NIK (16 digit)" class="border px-3 py-2 rounded" required />
      <input v-model="form.phone" placeholder="No. HP" class="border px-3 py-2 rounded" required />
      <input v-model="form.address" placeholder="Alamat" class="border px-3 py-2 rounded" required />
      <input v-model="form.occupation" placeholder="Pekerjaan" class="border px-3 py-2 rounded" required />
      <input v-model="form.ktpPhotoUrl" placeholder="URL foto KTP" class="border px-3 py-2 rounded" required />
      <input v-model="form.selfieKtpPhotoUrl" placeholder="URL selfie + KTP" class="border px-3 py-2 rounded" required />
      <div class="flex gap-2">
        <button type="button" class="px-4 py-2 rounded border" @click="step = 1">Kembali</button>
        <button class="flex-1 bg-blue-600 text-white px-4 py-2 rounded" :disabled="loading">
          {{ loading ? "Mendaftar..." : "Daftar" }}
        </button>
      </div>
      <p v-if="error" class="text-red-600 text-sm">{{ error }}</p>
      <p v-if="success" class="text-green-600 text-sm">Pendaftaran berhasil! Menunggu verifikasi admin. <NuxtLink to="/auth/login" class="underline">Login</NuxtLink></p>
    </form>

    <p class="text-sm mt-4">Sudah punya akun? <NuxtLink to="/auth/login" class="text-blue-600 underline">Login</NuxtLink></p>
  </div>
</template>

<script setup lang="ts">
import { z } from "zod"
import { authClient } from "~/utils/auth-client"

definePageMeta({ layout: "auth" })

const step = ref(1)
const loading = ref(false)
const error = ref("")
const success = ref(false)

const form = reactive({
  name: "", email: "", password: "",
  memberNumber: "", nik: "", phone: "", address: "", occupation: "",
  ktpPhotoUrl: "", selfieKtpPhotoUrl: "",
})

const accountSchema = z.object({
  name: z.string().min(3),
  email: z.string().email(),
  password: z.string().min(8),
})

const profileSchema = z.object({
  memberNumber: z.string().min(1),
  nik: z.string().length(16),
  phone: z.string().min(9),
  address: z.string().min(5),
  occupation: z.string().min(2),
  ktpPhotoUrl: z.string().url(),
  selfieKtpPhotoUrl: z.string().url(),
})

function nextStep() {
  error.value = ""
  if (!accountSchema.safeParse({ name: form.name, email: form.email, password: form.password }).success) {
    error.value = "Nama min 3 huruf, email valid, password min 8 karakter"
    return
  }
  step.value = 2
}

async function submit() {
  error.value = ""
  success.value = false
  if (!profileSchema.safeParse({
    memberNumber: form.memberNumber, nik: form.nik, phone: form.phone,
    address: form.address, occupation: form.occupation,
    ktpPhotoUrl: form.ktpPhotoUrl, selfieKtpPhotoUrl: form.selfieKtpPhotoUrl,
  }).success) {
    error.value = "Periksa lagi: NIK 16 digit, URL foto valid, semua wajib diisi"
    return
  }
  loading.value = true
  try {
    const signup = await authClient.signUp.email({ email: form.email, password: form.password, name: form.name })
    if (signup.error) throw new Error(signup.error.message || "Akun gagal dibuat")
    await $fetch("/api/members/register", {
      method: "POST",
      body: {
        memberNumber: form.memberNumber, name: form.name, nik: form.nik,
        address: form.address, phone: form.phone, occupation: form.occupation,
        ktpPhotoUrl: form.ktpPhotoUrl, selfieKtpPhotoUrl: form.selfieKtpPhotoUrl,
      },
    })
    success.value = true
  } catch (e: any) {
    error.value = e?.data?.message || e?.message || "Pendaftaran gagal"
  } finally {
    loading.value = false
  }
}
</script>
