<script setup lang="ts">
import { ref } from "vue";
import { Plus, Pencil, ArrowLeft, Trash } from "lucide-vue-next";
import type { Estabelecimentos } from "~/types/Estabelecimentos";
import axiosInstance from "~/api/axios";
import Swal from "sweetalert2";
const isOpen = ref(false);
const isOpenEdit = ref(false);
const isOpenDelete = ref(false);

const emit = defineEmits(["submitted", "close"]);

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

const props = defineProps<{
  company: Estabelecimentos;
}>();

function sendData(data: any) {
  toast.fire({ icon: "success", title: "Empresa atualizada com sucesso!" });
  emit("submitted");
  emit("close");
}

async function handleDelete() {
  try {
    const response = await axiosInstance.delete(
      `/estabelecimentos/${props.company.id}`
    );
    emit("close");
    toast.fire({
      icon: "success",
      title: "Empresa excluída com sucesso!",
    });
  } catch (error) {
    toast.fire({
      icon: "error",
      title: "Não é possível excluir uma empresa que possua lojas cadastradas",
    });
  }
  emit("submitted");
}
</script>

<template>
  <div class="h-full p-4 flex flex-col gap-4">
    <div class="flex flex-col gap-4">
      <section class="flex gap-4 items-center justify-between">
        <div>
          <div class="flex gap-2 items-center">
            <ArrowLeft
              v-if="company"
              @click="$emit('close')"
              class="hover:text-emerald-500 cursor-pointer transition"
            />
            <p class="text-2xl font-bold" v-if="company">
              {{ props.company.nome }} -
              {{ props.company.numero_estabelecimento }}
            </p>
            <Pencil
              @click="isOpenEdit = true"
              class="hover:text-emerald-500 cursor-pointer transition"
              v-if="company"
            />
          </div>

          <p class="text-md text-gray-500" v-if="company">
            {{ props.company.razao_social }}
          </p>
          <p class="text-2xl font-bold" v-else>Suas Organizações</p>
        </div>
        <Modal
          :isOpen="isOpenEdit"
          @close="isOpenEdit = false"
          :title="'Editar Empresa'"
          :subtitle="'Altere abaixo os dados da nova empresa'"
        >
          <CompanyForm
            :isOpen="isOpenEdit"
            :company="company"
            @submitted="
          (data:any) => { sendData(data); isOpenEdit = false; } "
          />
        </Modal>
        <Trash
          class="hover:text-red-500 cursor-pointer transition"
          @click="isOpenDelete = true"
          v-if="company"
        />
        <Modal
          :isOpen="isOpenDelete"
          @close="isOpenDelete = false"
          :title="'Excluir Empresa?'"
          :subtitle="'Essa ação não poderá ser desfeita'"
        >
          <p class="mb-4">
            A empresa
            <strong class="text-emerald-500">{{ props.company.nome }}</strong>
            será excluída permanentemente!
          </p>
          <DefaultButton @click="handleDelete">Excluir</DefaultButton>
        </Modal>
      </section>
      <div class="flex gap-2" v-if="company == null">
        <DefaultButton @click="isOpen = true"
          ><Plus /> Nova organização
        </DefaultButton>
        <Modal
          :isOpen="isOpen"
          @close="isOpen = false"
          :title="'Cadastrar Empresa'"
          :subtitle="'Preencha o formulário abaixo com os dados da nova empresa'"
        >
          <CompanyForm
            :isOpen="isOpen"
            @submitted="
              (data) => {
                sendData(data);
                isOpen = false;
              }
            "
          />
        </Modal>
        <!-- <DefaultInput
          :type="'search'"
          :placeholder="'Digite aqui'"
          v-model:data="text"
        /> -->
      </div>
    </div>
    <div class="max-h-96 overflow-y-auto pr-2">
      <div
        :class="[
          company == null
            ? 'grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4'
            : 'flex w-full',
        ]"
      >
        <slot />
      </div>
    </div>
  </div>
</template>
