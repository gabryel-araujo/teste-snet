<script setup>
import axiosInstance from "~/api/axios";

const props = defineProps({
  company: {
    type: Object,
    required: true,
  },
});

const stores = ref([]);

onMounted(async () => {
  const response = await axiosInstance.get("/lojas");
  stores.value = response.data;
  console.log(response.data);
});

async function getEstabs() {
  const response = await axiosInstance.get("/lojas");
  stores.value = response.data;
}
</script>

<template>
  <div class="flex flex-col w-full gap-4">
    <div class="rounded-md">
      <table
        class="min-w-full divide-y divide-gray-600 text-sm text-left text-white bg-gray-800 rounded-md"
      >
        <thead class="bg-gray-700 text-xs uppercase text-gray-300">
          <tr>
            <th class="px-4 py-3" @click="isOpen = true">Nome</th>
            <th class="px-4 py-3">Número</th>
            <th class="px-4 py-3">Razão Social</th>
            <th class="px-4 py-3">Endereço</th>
            <th class="px-4 py-3">Cidade</th>
            <th class="px-4 py-3">Estado</th>
            <th class="px-4 py-3">CEP</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-600">
          <tr v-for="store in stores" :key="store.id" class="hover:bg-gray-700">
            <td class="px-4 py-2">{{ store.nome }}</td>
            <td class="px-4 py-2">{{ store.numero_loja }}</td>
            <td class="px-4 py-2">{{ store.razao_social || "-" }}</td>
            <td class="px-4 py-2">
              {{ store.endereco }}{{ store.numero ? ", " + store.numero : "" }}
            </td>
            <td class="px-4 py-2">{{ store.cidade }}</td>
            <td class="px-4 py-2">{{ store.estado }}</td>
            <td class="px-4 py-2">{{ store.cep }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
