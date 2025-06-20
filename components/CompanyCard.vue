<script setup lang="ts">
import { Building2, EllipsisVertical } from "lucide-vue-next";
import type { Estabelecimentos } from "~/types/Estabelecimentos";

const props = defineProps<{
  company: Estabelecimentos;
}>();
</script>

<template>
  <div class="relative" @click="$emit('click')">
    <div
      class="flex border border-gray-600 p-4 rounded-lg bg-gray-900 justify-between items-center hover:ring-1 hover:ring-emerald-400 hover:ring-opacity-70 transition hover:cursor-pointer"
    >
      <section class="inline-flex gap-2 items-center">
        <div class="bg-zinc-700 p-2 rounded-full">
          <Building2 />
        </div>
        <section>
          <p class="font-bold">{{ props.company.nome }}</p>
          <p class="text-gray-500 text-xs">
            {{ props.company.numero_estabelecimento }}
          </p>
        </section>
      </section>
    </div>
  </div>
</template>
