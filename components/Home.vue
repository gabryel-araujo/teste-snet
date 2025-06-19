<script setup lang="ts">
import { Plus } from "lucide-vue-next";
import { ref } from "vue";
import { Pencil } from "lucide-vue-next";
import type { Estabelecimentos } from "~/types/Estabelecimentos";
const isOpen = ref(false);

const emit = defineEmits(["submitted"]);
const isOpenEdit = ref(false);

const props = defineProps<{
  company: Estabelecimentos;
}>();

function sendData() {
  emit("submitted");
}
</script>

<template>
  <div class="h-full p-4 flex flex-col gap-4">
    <div class="flex flex-col gap-4">
      <section class="flex gap-4">
        <div>
          <div class="flex gap-2 items-center">
            <p class="text-2xl font-bold" v-if="company">
              {{ props.company.nome }} -
              {{ props.company.numero_estabelecimento }}
            </p>
            <Pencil
              @click="isOpenEdit = true"
              class="hover:text-emerald-500 cursor-pointer transition"
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
            :isOpen="isOpen"
            @submitted="
              () => {
                sendData();
                isOpen = false;
              }
            "
          />
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
              () => {
                sendData();
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
</template>
