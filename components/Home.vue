<template>
  <div class="h-full p-4 flex flex-col gap-4">
    <div class="flex flex-col gap-4">
      <p class="text-2xl font-bold">Suas Organizações</p>
      <div class="flex gap-2">
        <ActionButton><Plus /> Nova organização </ActionButton>
        <!-- todo: criar um componente com esse input -->
        <input
          type="search"
          placeholder="Busque aqui..."
          class="border border-gray-600 rounded-md p-1 focus:ring-emerald-400 focus:ring-1 outline-none text-sm w-1/4"
        />
      </div>
    </div>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <slot />
    </div>
  </div>
</template>

<script setup>
import { Plus } from "lucide-vue-next";
</script>
