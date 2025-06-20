<script setup>
import { ref, onMounted } from "vue";
import axiosInstance from "~/api/axios";
import { Plus, Trash } from "lucide-vue-next";

const props = defineProps({
  company: {
    type: Object,
    required: true,
  },
});

const stores = ref([]);
const isOpen = ref(false);
const isOpenDelete = ref(false);
const isOpenEdit = ref(false);
const selectedStore = ref(null);

function sendData(data) {
  getStores();
  isOpen.value = false;
  console.log("Loja cadastrada", data);
}

function sync(data) {
  getStores();
  isOpenEdit.value = false;
  console.log("Loja Editada", data);
}

async function getStores() {
  const response = await axiosInstance.get("/lojas");
  stores.value = response.data.filter(
    (loja) => loja.estabelecimento_id == props.company.id
  );
}

function handleConfirm(store) {
  console.log(store.nome);
  selectedStore.value = store;
  isOpenDelete.value = true;
}

async function handleDelete(store) {
  const response = await axiosInstance.delete(`/lojas/${store.id}`);
  getStores();
  isOpenDelete.value = false;
}

function handleEdit(store) {
  isOpenEdit.value = true;
  selectedStore.value = store;
}

onMounted(getStores);
</script>

<template>
  <div class="flex flex-col w-full gap-6">
    <div class="flex justify-end">
      <DefaultButton @click="isOpen = true">
        <Plus class="w-4 h-4 mr-2" /> Nova Loja
      </DefaultButton>
    </div>

    <!-- Modal de cadastro -->
    <Modal
      :isOpen="isOpen"
      @close="isOpen = false"
      :title="'Cadastrar Loja'"
      :subtitle="'Preencha o formulário abaixo para cadastrar uma nova loja'"
    >
      <StoreForm
        :isOpen="isOpen"
        :company="props.company"
        @submitted="
          (data) => {
            sendData(data);
          }
        "
      />
    </Modal>
    <!-- Modal de edição -->
    <Modal
      :isOpen="isOpenEdit"
      @close="isOpenEdit = false"
      :title="'Editar Loja'"
      :subtitle="'Preencha o formulário abaixo para editar aloja'"
    >
      <StoreForm
        :isOpen="isOpen"
        :company="props.company"
        :selectedStore="selectedStore"
        @submitted="
          (data) => {
            sync(data);
          }
        "
      />
    </Modal>
    <!-- Modal de remoção -->
    <Modal
      :isOpen="isOpenDelete"
      @close="isOpenDelete = false"
      :title="'Apagar Loja'"
      :subtitle="'Essa ação não poderá ser desfeita'"
    >
      <p class="mb-4">
        A loja
        <strong class="text-emerald-500">{{ selectedStore.nome }}</strong> será
        apagada permanentemente!
      </p>
      <DefaultButton @click="handleDelete(selectedStore)">Apagar</DefaultButton>
    </Modal>
    <div class="rounded-md">
      <table
        class="min-w-full divide-y divide-gray-600 text-sm text-left text-white bg-gray-800 rounded-md"
      >
        <thead class="bg-gray-700 text-xs uppercase text-gray-300">
          <tr>
            <th class="px-4 py-3">Nome</th>
            <th class="px-4 py-3">Número</th>
            <th class="px-4 py-3">Razão Social</th>
            <th class="px-4 py-3">Endereço</th>
            <th class="px-4 py-3">Cidade</th>
            <th class="px-4 py-3">Estado</th>
            <th class="px-4 py-3">CEP</th>
            <th></th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-600">
          <tr
            v-for="store in stores"
            :key="store.id"
            class="hover:bg-gray-900 cursor-pointer"
            @click="handleEdit(store)"
          >
            <td class="px-4 py-2">{{ store.nome }}</td>
            <td class="px-4 py-2">{{ store.numero_loja }}</td>
            <td class="px-4 py-2">{{ store.razao_social || "-" }}</td>
            <td class="px-4 py-2">
              {{ store.endereco }}{{ store.numero ? ", " + store.numero : "" }}
            </td>
            <td class="px-4 py-2">{{ store.cidade }}</td>
            <td class="px-4 py-2">{{ store.estado }}</td>
            <td class="px-4 py-2">{{ store.cep }}</td>
            <td class="px-4 py-2">
              <Trash
                class="text-xs hover:text-red-500"
                @click="handleConfirm(store)"
              />
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
