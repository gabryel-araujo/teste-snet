<script setup lang="ts">
import { Building2, EllipsisVertical } from "lucide-vue-next";
import { Menu, MenuButton, MenuItems, MenuItem } from "@headlessui/vue";
import type { Estabelecimentos } from "~/types/Estabelecimentos";

const props = defineProps<{
  company: Estabelecimentos;
}>();

function editar() {}

function excluir() {}
</script>

<template>
  <div class="relative" @click="$emit('click')">
    <Menu>
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
        <MenuButton
          class="w-8 h-8 hover:bg-gray-600 rounded-lg flex items-center justify-center group relative z-10"
        >
          <EllipsisVertical
            class="text-gray-500 group-hover:text-slate-100 hover:cursor-pointer transition"
          />
        </MenuButton>
      </div>

      <transition
        enter-active-class="transition duration-100 ease-out"
        enter-from-class="transform scale-95 opacity-0"
        enter-to-class="transform scale-100 opacity-100"
        leave-active-class="transition duration-75 ease-in"
        leave-from-class="transform scale-100 opacity-100"
        leave-to-class="transform scale-95 opacity-0"
      >
        <MenuItems
          class="absolute right-4 top-full mt-1 w-56 origin-top-right divide-y divide-gray-700 rounded-md bg-gray-800 shadow-lg ring-1 ring-gray-600 focus:outline-none z-20"
        >
          <div class="px-1 py-1">
            <MenuItem v-slot="{ active }" as="template">
              <button
                :class="[
                  active ? 'bg-emerald-500 text-white' : 'text-gray-300',
                  'group flex w-full items-center rounded-md px-2 py-2 text-sm cursor-pointer',
                ]"
                @click="editar"
              >
                Editar
              </button>
            </MenuItem>
            <MenuItem v-slot="{ active }">
              <button
                :class="[
                  active ? 'bg-emerald-500 text-white' : 'text-gray-300',
                  'group flex w-full items-center rounded-md px-2 py-2 text-sm cursor-pointer',
                ]"
                @click="excluir"
              >
                Excluir
              </button>
            </MenuItem>
          </div>
        </MenuItems>
      </transition>
    </Menu>
  </div>
</template>
