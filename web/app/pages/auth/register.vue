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

      <label class="text-sm font-medium">Foto KTP (gambar, maks 2MB)</label>
      <input type="file" accept="image/*" class="border px-3 py-2 rounded text-sm" @change="onFile($event, 'ktp')" />
      <img v-if="preview.ktp" :src="preview.ktp" alt="Pratinjau KTP" class="h-32 object-contain border rounded" />

      <label class="text-sm font-medium">Selfie + KTP (gambar, maks 2MB)</label>
      <input type="file" accept="image/*" class="border px-3 py-2 rounded text-sm" @change="onFile($event, 'selfie')" />
      <img v-if="preview.selfie" :src="preview.selfie" alt="Pratinjau selfie" class="h-32 object-contain border rounded" />

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

const MAX_SIZE = 2 * 1024 * 1024 // 2MB

const step = ref(1)
const loading = ref(false)
const error = ref("")
const success = ref(false)

const form = reactive({
  name: "", email: "", password: "",
  memberNumber: "", nik: "", phone: "", address: "", occupation: "",
})

const files = reactive<{ ktp: File | null, selfie: File | null }>({ ktp: null, selfie: null })
const preview = reactive<{ ktp: string, selfie: string }>({ ktp: "", selfie: "" })

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
})

function nextStep() {
  error.value = ""
  if (!accountSchema.safeParse({ name: form.name, email: form.email, password: form.password }).success) {
    error.value = "Nama min 3 huruf, email valid, password min 8 karakter"
    return
  }
  step.value = 2
}

function onFile(e: Event, kind: "ktp" | "selfie") {
  error.value = ""
  const input = e.target as HTMLInputElement
  const file = input.files?.[0]
  if (!file) return
  if (!file.type.startsWith("image/")) {
    error.value = "File harus berupa gambar"
    input.value = ""
    return
  }
  if (file.size > MAX_SIZE) {
    error.value = "Ukuran file maksimal 2MB"
    input.value = ""
    return
  }
  files[kind] = file
  preview[kind] = URL.createObjectURL(file)
}

async function uploadFile(file: File) {
  const fd = new FormData()
  fd.append("file", file)
  const res = await $fetch<{ url: string }>("/api/members/upload", { method: "POST", body: fd })
  return res.url
}

async function submit() {
  error.value = ""
  success.value = false
  if (!profileSchema.safeParse({
    memberNumber: form.memberNumber, nik: form.nik, phone: form.phone,
    address: form.address, occupation: form.occupation,
  }).success) {
    error.value = "Periksa lagi: NIK 16 digit dan semua wajib diisi"
    return
  }
  if (!files.ktp || !files.selfie) {
    error.value = "Foto KTP dan selfie + KTP wajib diupload"
    return
  }
  loading.value = true
  try {
    const signup = await authClient.signUp.email({ email: form.email, password: form.password, name: form.name })
    if (signup.error) throw new Error(signup.error.message || "Akun gagal dibuat")
    const [ktpPhotoUrl, selfieKtpPhotoUrl] = await Promise.all([
      uploadFile(files.ktp),
      uploadFile(files.selfie),
    ])
    const member = await $fetch<{ id: string }>("/api/members/register", {
      method: "POST",
      body: {
        memberNumber: form.memberNumber, name: form.name, nik: form.nik,
        address: form.address, phone: form.phone, occupation: form.occupation,
        ktpPhotoUrl, selfieKtpPhotoUrl,
      },
    })
    await authClient.updateUser({ memberId: member.id } as any)
    success.value = true
  } catch (e: any) {
    error.value = e?.data?.message || e?.message || "Pendaftaran gagal"
  } finally {
    loading.value = false
  }
}
</script>
