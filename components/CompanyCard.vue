<script setup>
import { Building2, EllipsisVertical } from "lucide-vue-next";

const props = defineProps({
  companyName: {
    type: String,
    required: true,
  },
  companyNumber: {
    type: Number,
    required: true,
  },
});
</script>

<template>
  <div
    class="flex border border-gray-600 p-4 rounded-lg bg-gray-900 justify-between items-center hover:ring-1 hover:ring-emerald-400 hover:ring-opacity-70 transition hover:cursor-pointer"
  >
    <section class="inline-flex gap-2 items-center">
      <div class="bg-zinc-700 p-2 rounded-full">
        <Building2 />
      </div>
      <section>
        <p class="font-bold">{{ companyName }}</p>
        <p class="text-gray-500 text-xs">{{ companyNumber }}</p>
      </section>
    </section>
    <EllipsisVertical
      class="text-gray-500 hover:text-slate-400 hover:cursor-pointer transition"
    />
  </div>
</template>
