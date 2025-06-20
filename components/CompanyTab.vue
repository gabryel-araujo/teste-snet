<script setup>
import { ref, onMounted } from "vue";
import axiosInstance from "~/api/axios";
import { Plus, Trash } from "lucide-vue-next";
import Swal from "sweetalert2";

const props = defineProps({
  company: {
    type: Object,
    required: true,
  },
});

const toast = Swal.mixin({
  toast: true,
  position: "top-end",
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: (toast) => {
    toast.onmouseenter = Swal.stopTimer;
    toast.onmouseleave = Swal.resumeTimer;
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
  toast.fire({
    icon: "success",
    title: "Loja cadastrada!",
  });
}

function sync(data) {
  getStores();
  isOpenEdit.value = false;
  console.log("Loja Editada", data);
  toast.fire({
    icon: "success",
    title: "Loja editada!",
  });
}

async function getStores() {
  try {
    const response = await axiosInstance.get("/lojas");
    stores.value = response.data.filter(
      (loja) => loja.estabelecimento_id == props.company.id
    );
  } catch (error) {
    console.error("Erro ao carregar lojas:", error);
    stores.value = [];
  }
}

function handleConfirm(store) {
  console.log(store.nome);
  selectedStore.value = store;
  isOpenDelete.value = true;
}

async function handleDelete(store) {
  try {
    await axiosInstance.delete(`/lojas/${store.id}`);
    stores.value = stores.value.filter((s) => s.id !== store.id);
    isOpenDelete.value = false;
    toast.fire({
      icon: "success",
      title: "Loja removida com sucesso!",
    });
  } catch (error) {
    console.error("Erro ao deletar loja:", error);
  }
}

function handleEdit(store) {
  isOpenEdit.value = true;
  selectedStore.value = store;
}

onMounted(getStores);
</script>

<template>
  <div class="flex flex-col w-full gap-6 p-4">
    <!-- Botão Nova Loja -->
    <div class="flex justify-between items-center">
      <p class="text-lg font-semibold">Listagem de Lojas</p>
      <DefaultButton @click="isOpen = true" class="w-full sm:w-auto">
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
      :subtitle="'Preencha o formulário abaixo para editar a loja'"
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

    <!-- Tabela para o pc -->
    <div class="hidden lg:block rounded-md overflow-hidden">
      <div class="max-h-96 overflow-y-auto">
        <table
          class="min-w-full divide-y divide-gray-600 text-sm text-left text-white bg-gray-800"
        >
          <thead
            class="bg-gray-700 text-xs uppercase text-gray-300 sticky top-0"
          >
            <tr>
              <th class="px-4 py-3">Nome</th>
              <th class="px-4 py-3">Número</th>
              <th class="px-4 py-3">Razão Social</th>
              <th class="px-4 py-3">Endereço</th>
              <th class="px-4 py-3">Cidade</th>
              <th class="px-4 py-3">Estado</th>
              <th class="px-4 py-3">CEP</th>
              <th class="px-4 py-3">Ações</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-600">
            <tr
              v-for="store in stores"
              :key="store.id"
              class="hover:bg-gray-900 cursor-pointer"
              @click="handleEdit(store)"
            >
              <td class="px-4 py-3 font-medium">{{ store.nome }}</td>
              <td class="px-4 py-3">{{ store.numero_loja }}</td>
              <td class="px-4 py-3">{{ store.razao_social || "-" }}</td>
              <td class="px-4 py-3">
                {{ store.endereco
                }}{{ store.numero ? ", " + store.numero : "" }}
              </td>
              <td class="px-4 py-3">{{ store.cidade }}</td>
              <td class="px-4 py-3">{{ store.estado }}</td>
              <td class="px-4 py-3">{{ store.cep }}</td>
              <td class="px-4 py-3">
                <button
                  @click.stop="handleConfirm(store)"
                  class="p-1 hover:text-red-500 transition-colors cursor-pointer"
                >
                  <Trash class="w-4 h-4" />
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="stores.length === 0" class="text-center py-12 text-gray-400">
        <div class="mb-4">
          <Plus class="w-12 h-12 mx-auto opacity-50" />
        </div>
        <p class="text-lg font-medium mb-2">Nenhuma loja cadastrada</p>
        <p class="text-sm">Clique no botão "Nova Loja" para começar</p>
      </div>
    </div>

    <!-- Card para o mobile -->
    <div class="lg:hidden">
      <div class="max-h-96 overflow-y-auto space-y-4 pr-2">
        <div
          v-for="store in stores"
          :key="store.id"
          class="bg-gray-800 rounded-lg p-4 shadow-lg border border-gray-700 hover:bg-gray-750 transition-colors cursor-pointer"
          @click="handleEdit(store)"
        >
          <div class="flex justify-between items-start mb-3">
            <div>
              <h3 class="font-semibold text-white text-lg">{{ store.nome }}</h3>
              <p class="text-sm text-gray-400">Nº {{ store.numero_loja }}</p>
            </div>
            <button
              @click="handleConfirm(store)"
              class="p-2 hover:text-red-500 transition-colors"
            >
              <Trash class="w-4 h-4" />
            </button>
          </div>
          <div class="space-y-2">
            <div
              v-if="store.razao_social"
              class="flex flex-col sm:flex-row sm:justify-between"
            >
              <span class="text-xs text-gray-400 uppercase font-medium"
                >Razão Social</span
              >
              <span class="text-sm text-white">{{ store.razao_social }}</span>
            </div>
            <div class="flex flex-col sm:flex-row sm:justify-between">
              <span class="text-xs text-gray-400 uppercase font-medium"
                >Endereço</span
              >
              <span class="text-sm text-white">
                {{ store.endereco
                }}{{ store.numero ? ", " + store.numero : "" }}
              </span>
            </div>
            <div class="flex flex-col sm:flex-row sm:justify-between">
              <span class="text-xs text-gray-400 uppercase font-medium"
                >Cidade/Estado</span
              >
              <span class="text-sm text-white"
                >{{ store.cidade }}, {{ store.estado }}</span
              >
            </div>
            <div class="flex flex-col sm:flex-row sm:justify-between">
              <span class="text-xs text-gray-400 uppercase font-medium"
                >CEP</span
              >
              <span class="text-sm text-white font-mono">{{ store.cep }}</span>
            </div>
          </div>
          <div class="mt-3 pt-3 border-t border-gray-700">
            <p class="text-xs text-gray-500 text-center">Toque para editar</p>
          </div>
        </div>
      </div>

      <div v-if="stores.length === 0" class="text-center py-12 text-gray-400">
        <div class="mb-4">
          <Plus class="w-12 h-12 mx-auto opacity-50" />
        </div>
        <p class="text-lg font-medium mb-2">Nenhuma loja cadastrada</p>
        <p class="text-sm">Clique no botão "Nova Loja" para começar</p>
      </div>
    </div>
  </div>
</template>
