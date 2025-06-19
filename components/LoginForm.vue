<script setup>
import { Lock, Mail } from "lucide-vue-next";

function goHome() {
  navigateTo("/app");
}
</script>

<template>
  <div
    class="border border-gray-700 h-[500px] w-[800px] flex rounded-2xl overflow-hidden shadow-lg"
  >
    <div class="w-1/2 p-8 flex flex-col justify-center">
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-emerald-500 mb-2">
          Bem-vindo de volta
        </h1>
        <p class="text-gray-400">Acesse sua conta para continuar</p>
      </div>

      <div class="space-y-6">
        <div class="space-y-1">
          <label class="block text-sm font-medium text-gray-300" for="email"
            >Email</label
          >
          <div class="relative">
            <div
              class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
            >
              <Mail class="h-5 w-5 text-gray-400" />
            </div>
            <DefaultInput
              id="email"
              type="email"
              placeholder="seu@email.com"
              class="w-full pl-10"
            />
          </div>
        </div>

        <div class="space-y-1">
          <label class="block text-sm font-medium text-gray-300" for="senha"
            >Senha</label
          >
          <div class="relative">
            <div
              class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
            >
              <Lock class="h-5 w-5 text-gray-400" />
            </div>
            <DefaultInput
              id="senha"
              type="password"
              placeholder="***********"
              class="w-full pl-10"
            />
          </div>
          <div class="w-full justify-between">
            <section class="inline-flex gap-2">
              <input type="checkbox" id="lembrar" />
              <p class="block text-sm font-medium text-gray-400" for="lembrar">
                Lembrar-me
              </p>
            </section>

            <a
              href="#"
              class="text-xs text-emerald-400 hover:text-emerald-300 float-right mt-1"
              >Esqueceu a senha?</a
            >
          </div>
        </div>

        <DefaultButton class="w-full mt-6 py-3" @click="goHome">
          <span class="font-semibold">Entrar</span>
        </DefaultButton>
      </div>

      <div class="text-center mt-8 text-sm text-gray-400">
        Não tem uma conta?
        <a href="#" class="text-emerald-400 hover:text-emerald-300 font-medium"
          >Cadastre-se</a
        >
      </div>
    </div>

    <div
      class="w-1/2 bg-gradient-to-br from-emerald-500 to-emerald-600 flex justify-center items-center p-8"
    >
      <img
        src="/login.png"
        alt="Ilustração de um usuário fazendo login"
        class="max-h-[80%] object-contain"
      />
    </div>
  </div>
</template>
