<template>
  <div class="max-w-lg">
    <h1 class="text-2xl font-bold mb-4">Profil Koperasi</h1>
    <form class="bg-white p-6 rounded shadow flex flex-col gap-3" @submit.prevent="save">
      <label class="text-sm">Nama Koperasi<input v-model="form.kop_name" class="border px-3 py-2 rounded w-full mt-1" /></label>
      <label class="text-sm">Alamat<input v-model="form.kop_address" class="border px-3 py-2 rounded w-full mt-1" /></label>
      <label class="text-sm">Logo (URL)<input v-model="form.kop_logo" class="border px-3 py-2 rounded w-full mt-1" /></label>
      <button class="bg-blue-600 text-white px-4 py-2 rounded">Simpan</button>
      <p v-if="msg" class="text-green-600 text-sm">{{ msg }}</p>
    </form>
  </div>
</template>

<script setup lang="ts">
definePageMeta({ middleware: ["auth", "role"], allowedRoles: ["super_admin"] })

const msg = ref("")
const form = reactive({ kop_name: "", kop_address: "", kop_logo: "" })
const settings = await $fetch<any[]>("/api/superadmin/settings")
for (const s of settings) {
  if (s.key in form) (form as any)[s.key] = s.value
}

async function save() {
  for (const key of Object.keys(form)) {
    await $fetch(`/api/superadmin/settings/${key}`, { method: "PUT", body: { value: (form as any)[key] } })
  }
  msg.value = "Profil koperasi tersimpan"
}
</script>
