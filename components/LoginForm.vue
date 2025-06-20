<script setup>
import { Lock, Mail } from "lucide-vue-next";
import { useForm, useField } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import { loginSchema } from "~/schemas/loginSchema";

const { handleSubmit, errors } = useForm({
  validationSchema: toTypedSchema(loginSchema),
});

const { value: email } = useField("email");
const { value: password } = useField("password");

const onSubmit = handleSubmit((values) => {
  console.log("dados do login", values);
  localStorage.setItem("auth-token", values.email);
  navigateTo("/app");
});
</script>

<template>
  <form @submit.prevent="onSubmit" class="w-full max-w-4xl">
    <div
      class="border border-gray-700 w-full flex flex-col lg:flex-row rounded-2xl overflow-hidden shadow-lg bg-gray-800"
    >
      <div class="w-full lg:w-1/2 p-6 sm:p-8 flex flex-col justify-center">
        <div class="text-center mb-8">
          <h1 class="text-2xl sm:text-3xl font-bold text-emerald-500 mb-2">
            Seja bem-vindo!
          </h1>
          <p class="text-gray-400 text-sm sm:text-base">
            Acesse sua conta para continuar
          </p>
        </div>

        <div class="space-y-4 sm:space-y-6">
          <div class="space-y-1">
            <label class="block text-sm font-medium text-gray-300" for="email">
              Email
            </label>
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
                v-model:data="email"
              />
            </div>
            <p class="text-red-500 text-xs sm:text-sm transition">
              {{ errors.email }}
            </p>
          </div>

          <div class="space-y-1">
            <label
              class="block text-sm font-medium text-gray-300"
              for="password"
            >
              Senha
            </label>
            <div class="relative">
              <div
                class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
              >
                <Lock class="h-5 w-5 text-gray-400" />
              </div>
              <DefaultInput
                id="password"
                type="password"
                placeholder="***********"
                class="w-full pl-10"
                v-model:data="password"
              />
            </div>
            <p class="text-red-500 text-xs sm:text-sm transition">
              {{ errors.password }}
            </p>
          </div>

          <div
            class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-2"
          >
            <div class="flex items-center gap-2">
              <input
                type="checkbox"
                id="lembrar"
                class="w-4 h-4 text-emerald-600 bg-gray-700 border-gray-600 rounded focus:ring-emerald-500"
              />
              <label class="text-sm font-medium text-gray-400" for="lembrar">
                Lembrar-me
              </label>
            </div>
            <a
              href="#"
              class="text-xs sm:text-sm text-emerald-400 hover:text-emerald-300 transition-colors"
            >
              Esqueceu a senha?
            </a>
          </div>

          <DefaultButton class="w-full mt-6 py-3" type="submit">
            <span class="font-semibold">Entrar</span>
          </DefaultButton>
        </div>

        <div class="text-center mt-6 sm:mt-8 text-xs sm:text-sm text-gray-400">
          Não tem uma conta?
          <a
            href="#"
            class="text-emerald-400 hover:text-emerald-300 font-medium transition-colors ml-1"
          >
            Cadastre-se!
          </a>
        </div>
      </div>

      <div
        class="hidden lg:flex lg:w-1/2 bg-gradient-to-br from-emerald-500 to-emerald-600 justify-center items-center p-8"
      >
        <img
          src="/login.png"
          alt="Ilustração de um usuário fazendo login"
          class="max-h-[80%] max-w-full object-contain"
        />
      </div>
    </div>
  </form>
</template>
