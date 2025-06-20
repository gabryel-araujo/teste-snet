<script setup lang="ts">
import { Building2, EllipsisVertical, ChevronRight } from "lucide-vue-next";
import type { Estabelecimentos } from "~/types/Estabelecimentos";

const props = defineProps<{
  company: Estabelecimentos;
}>();
</script>

<template>
  <div class="relative group" @click="$emit('click')">
    <div
      class="flex border border-gray-700 p-6 rounded-xl bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900 justify-between items-center hover:border-emerald-400/50 transition-all duration-300 hover:cursor-pointer"
    >
      <div
        class="absolute inset-0 rounded-xl bg-gradient-to-r from-emerald-400/5 to-blue-400/5 opacity-0 transition-opacity duration-300"
      ></div>

      <section class="inline-flex gap-4 items-center relative z-10">
        <div
          class="bg-gradient-to-br from-emerald-500 to-emerald-600 p-3 rounded-xl shadow-lg group-hover:shadow-emerald-500/20 transition-all duration-300"
        >
          <Building2 class="text-white w-5 h-5" />
        </div>

        <section>
          <p
            class="font-semibold text-white text-lg group-hover:text-emerald-50 transition-colors duration-300"
          >
            {{ props.company.nome }}
          </p>
          <p class="text-gray-400 text-sm flex items-center gap-2 truncate">
            <span class="w-1 h-1 bg-emerald-400 rounded-full opacity-60"></span>
            <strong>Razão Social:</strong>{{ props.company.razao_social }}
          </p>
          <p class="text-gray-400 text-sm flex items-center gap-2 truncate">
            <span class="w-1 h-1 bg-emerald-400 rounded-full opacity-60"></span>
            <strong>Nº Estabelecimento:</strong
            >{{ props.company.numero_estabelecimento }}
          </p>
          <p class="text-gray-400 text-sm flex items-center gap-2 truncate">
            <span class="w-1 h-1 bg-emerald-400 rounded-full opacity-60"></span>
            <strong>Endereço:</strong>{{ props.company.endereco }},{{
              props.company.numero
            }}
          </p>
          <p class="text-gray-400 text-sm flex items-center gap-2 truncate">
            <span class="w-1 h-1 bg-emerald-400 rounded-full opacity-60"></span>
            <strong>Cidade:</strong>{{ props.company.cidade }}
          </p>
          <p class="text-gray-400 text-sm flex items-center gap-2 truncate">
            <span class="w-1 h-1 bg-emerald-400 rounded-full opacity-60"></span>
            <strong>Estado:</strong>{{ props.company.estado }}
          </p>
        </section>
      </section>

      <div
        class="text-gray-500 group-hover:text-emerald-400 transition-all duration-300 group-hover:translate-x-1"
      >
        <ChevronRight class="w-5 h-5" />
      </div>
    </div>
  </div>
</template>
