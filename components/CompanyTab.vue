<script setup>
const props = defineProps({
  company: {
    type: Object,
    required: true,
  },
});

import { X } from "lucide-vue-next";
</script>

<template>
  <div class="w-full bg-gray-900 rounded-md border border-gray-600 p-4">
    <div class="flex items-center justify-between">
      <h2 class="text-xl font-bold">Detalhes da Empresa</h2>
      <button
        class="text-red-500 cursor-pointer hover:text-red-400 transition"
        @click="$emit('close')"
      >
        <X />
      </button>
    </div>

    <CompanyForm />
  </div>
</template>
