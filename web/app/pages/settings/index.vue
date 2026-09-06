<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">Pengaturan Sistem</h1>
    <table class="w-full bg-white rounded shadow text-sm max-w-3xl">
      <thead><tr class="border-b text-left"><th class="p-2">Key</th><th class="p-2">Value</th><th class="p-2">Aksi</th></tr></thead>
      <tbody>
        <tr v-for="s in settings" :key="s.id" class="border-b">
          <td class="p-2 font-mono">{{ s.key }}</td>
          <td class="p-2">
            <input v-model="draft[s.key]" class="border px-2 py-1 rounded w-full" />
          </td>
          <td class="p-2"><button class="text-blue-600 underline" @click="save(s.key)">Simpan</button></td>
        </tr>
      </tbody>
    </table>
    <div class="bg-white p-4 rounded shadow mt-4 max-w-3xl">
      <h2 class="font-semibold mb-2">Tambah Setting Baru</h2>
      <form class="flex gap-2" @submit.prevent="save(newKey)">
        <input v-model="newKey" placeholder="key_baru" class="border px-3 py-2 rounded" required />
        <input v-model="newValue" placeholder="value" class="border px-3 py-2 rounded" required />
        <button class="px-4 py-2 bg-blue-600 text-white rounded">Tambah</button>
      </form>
    </div>
    <p v-if="msg" class="text-sm mt-2 text-green-600">{{ msg }}</p>
  </div>
</template>

<script setup lang="ts">
definePageMeta({ middleware: ["auth", "role"], allowedRoles: ["super_admin"] })

const msg = ref("")
const newKey = ref("")
const newValue = ref("")
const draft = reactive<Record<string, string>>({})
const { data: settings, refresh } = await useFetch<any[]>("/api/superadmin/settings")

watchEffect(() => {
  for (const s of (settings.value as any[]) || []) draft[s.key] = s.value
})

async function save(key: string) {
  const value = key === newKey.value ? newValue.value : draft[key]
  await $fetch(`/api/superadmin/settings/${key}`, { method: "PUT", body: { value } })
  msg.value = `Setting ${key} tersimpan`
  newKey.value = ""
  newValue.value = ""
  await refresh()
}
</script>
