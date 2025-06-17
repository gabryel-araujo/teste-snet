<template>
  <div class="h-full p-4 flex flex-col gap-4">
    <div class="flex flex-col gap-4">
      <p class="text-2xl font-bold">Suas Organizações</p>
      <div class="flex gap-2">
        <DefaultButton @click="isOpen = true"
          ><Plus /> Nova organização
        </DefaultButton>
        <Modal
          :isOpen="isOpen"
          @close="isOpen = false"
          :title="'Cadastrar Empresa'"
          :subtitle="'Preencha o formulário abaixo com os dados da nova empresa'"
        >
          <CompanyForm />
        </Modal>
        <DefaultInput :type="'search'" :placeholder="'Digite aqui'" />
      </div>
    </div>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <slot />
    </div>
  </div>
</template>

<script setup>
import { Plus } from "lucide-vue-next";
import { ref } from "vue";
const isOpen = ref(false);
</script>
