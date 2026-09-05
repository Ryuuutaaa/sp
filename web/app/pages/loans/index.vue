<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">Pinjaman</h1>
    <table class="w-full bg-white rounded shadow text-sm">
      <thead><tr class="border-b"><th class="p-2 text-left">No Pinjaman</th><th class="p-2 text-left">Jumlah</th><th class="p-2 text-left">Tenor</th><th class="p-2 text-left">Status</th></tr></thead>
      <tbody>
        <tr v-for="l in loans" :key="l.id" class="border-b">
          <td class="p-2">{{ l.loanNumber }}</td>
          <td class="p-2">{{ l.amount }}</td>
          <td class="p-2">{{ l.tenorMonths }} bln</td>
          <td class="p-2">{{ l.status }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
definePageMeta({ middleware: ["auth"] })
const { data: loans } = await useLoans()
</script>
