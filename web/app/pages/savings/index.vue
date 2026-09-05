<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">Simpanan</h1>
    <div class="bg-white p-4 rounded shadow">
      <h2 class="font-semibold mb-2">Jenis Simpanan</h2>
      <ul class="list-disc ml-5">
        <li v-for="t in types" :key="t.id">{{ t.name }} — {{ t.defaultAmount }}</li>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({ middleware: ["auth"] })
const { data: types } = await useSavingsTypes()
</script>
