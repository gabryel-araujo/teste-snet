<script setup>
definePageMeta({
  middleware: "fake-auth",
});
</script>

<template>
  <!-- Comentário de Gabryel: Essa parte foi gerada com IA apenas para gerir uma documentação estética do projeto! -->
  <div class="min-h-screen bg-gray-900 text-white">
    <!-- Hero Section -->
    <section
      class="relative overflow-hidden bg-gradient-to-br from-emerald-500 to-emerald-700 py-16 sm:py-24"
    >
      <div class="absolute inset-0 bg-black opacity-10"></div>
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center">
          <h1 class="text-4xl sm:text-5xl lg:text-6xl font-bold mb-6">
            Sistema de Gerenciamento
            <span class="block text-emerald-200">de Estabelecimentos</span>
          </h1>
          <p class="text-lg sm:text-xl text-emerald-100 max-w-3xl mx-auto mb-8">
            Uma aplicação fullstack moderna construída com Nuxt 3 e Golang,
            projetada para gerenciar estabelecimentos e suas lojas de forma
            eficiente.
          </p>
          <p class="text-lg sm:text-xl text-emerald-100 max-w-3xl mx-auto mb-8">
            Essa documentação foi gerada por Inteligência artificial!
          </p>
          <div class="flex flex-col sm:flex-row gap-4 justify-center">
            <div
              class="bg-white bg-opacity-20 backdrop-blur-sm rounded-lg px-4 py-2"
            >
              <span class="text-emerald-700 font-semibold"
                >Frontend: Nuxt 3 + Vue 3</span
              >
            </div>
            <div
              class="bg-white bg-opacity-20 backdrop-blur-sm rounded-lg px-4 py-2"
            >
              <span class="text-emerald-700 font-semibold"
                >Backend: Golang + Echo</span
              >
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Features Section -->
    <section class="py-16 sm:py-20">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12">
          <h2 class="text-3xl sm:text-4xl font-bold text-white mb-4">
            Funcionalidades Principais
          </h2>
          <p class="text-gray-400 text-lg max-w-2xl mx-auto">
            Um sistema completo para gerenciar estabelecimentos e lojas com
            interface moderna e API robusta
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          <!-- Autenticação -->
          <div
            class="bg-gray-800 rounded-xl p-6 border border-gray-700 hover:border-emerald-500 transition-colors"
          >
            <div
              class="w-12 h-12 bg-emerald-500 rounded-lg flex items-center justify-center mb-4"
            >
              <svg
                class="w-6 h-6 text-white"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"
                ></path>
              </svg>
            </div>
            <h3 class="text-xl font-semibold text-white mb-2">
              Sistema de Login
            </h3>
            <p class="text-gray-400">
              Login simulado com interface moderna e responsiva para acesso ao
              sistema.
            </p>
          </div>

          <!-- CRUD Estabelecimentos -->
          <div
            class="bg-gray-800 rounded-xl p-6 border border-gray-700 hover:border-emerald-500 transition-colors"
          >
            <div
              class="w-12 h-12 bg-emerald-500 rounded-lg flex items-center justify-center mb-4"
            >
              <svg
                class="w-6 h-6 text-white"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"
                ></path>
              </svg>
            </div>
            <h3 class="text-xl font-semibold text-white mb-2">
              Gestão de Estabelecimentos
            </h3>
            <p class="text-gray-400">
              CRUD completo para estabelecimentos com validação de campos e
              integração com CEP.
            </p>
          </div>

          <!-- CRUD Lojas -->
          <div
            class="bg-gray-800 rounded-xl p-6 border border-gray-700 hover:border-emerald-500 transition-colors"
          >
            <div
              class="w-12 h-12 bg-emerald-500 rounded-lg flex items-center justify-center mb-4"
            >
              <svg
                class="w-6 h-6 text-white"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"
                ></path>
              </svg>
            </div>
            <h3 class="text-xl font-semibold text-white mb-2">
              Controle de Lojas
            </h3>
            <p class="text-gray-400">
              Gerenciamento completo de lojas vinculadas aos estabelecimentos
              com relacionamentos.
            </p>
          </div>

          <!-- Dashboard -->
          <div
            class="bg-gray-800 rounded-xl p-6 border border-gray-700 hover:border-emerald-500 transition-colors"
          >
            <div
              class="w-12 h-12 bg-emerald-500 rounded-lg flex items-center justify-center mb-4"
            >
              <svg
                class="w-6 h-6 text-white"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
                ></path>
              </svg>
            </div>
            <h3 class="text-xl font-semibold text-white mb-2">
              Dashboard Intuitivo
            </h3>
            <p class="text-gray-400">
              Interface centralizada para visualizar e gerenciar todos os
              estabelecimentos e lojas.
            </p>
          </div>

          <!-- Responsivo -->
          <div
            class="bg-gray-800 rounded-xl p-6 border border-gray-700 hover:border-emerald-500 transition-colors"
          >
            <div
              class="w-12 h-12 bg-emerald-500 rounded-lg flex items-center justify-center mb-4"
            >
              <svg
                class="w-6 h-6 text-white"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 18h.01M8 21h8a1 1 0 001-1v-1H7v1a1 1 0 001 1zm8-10V7a4 4 0 10-8 0v4h8z"
                ></path>
              </svg>
            </div>
            <h3 class="text-xl font-semibold text-white mb-2">
              Design Responsivo
            </h3>
            <p class="text-gray-400">
              Interface adaptável que funciona perfeitamente em desktop, tablet
              e mobile.
            </p>
          </div>

          <!-- API Robusta -->
          <div
            class="bg-gray-800 rounded-xl p-6 border border-gray-700 hover:border-emerald-500 transition-colors"
          >
            <div
              class="w-12 h-12 bg-emerald-500 rounded-lg flex items-center justify-center mb-4"
            >
              <svg
                class="w-6 h-6 text-white"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M8.111 16.404a5.5 5.5 0 017.778 0M12 20h.01m-7.08-7.071c3.904-3.905 10.236-3.905 14.141 0M1.394 9.393c5.857-5.857 15.355-5.857 21.213 0"
                ></path>
              </svg>
            </div>
            <h3 class="text-xl font-semibold text-white mb-2">API RESTful</h3>
            <p class="text-gray-400">
              Backend estruturado em camadas com endpoints padronizados e regras
              de negócio.
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- Tech Stack Section -->
    <section class="py-16 sm:py-20 bg-gray-800">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12">
          <h2 class="text-3xl sm:text-4xl font-bold text-white mb-4">
            Tecnologias Utilizadas
          </h2>
          <p class="text-gray-400 text-lg">
            Stack moderna com as melhores práticas de desenvolvimento
          </p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-12">
          <!-- Frontend -->
          <div class="bg-gray-900 rounded-xl p-8 border border-gray-700">
            <h3
              class="text-2xl font-bold text-emerald-400 mb-6 flex items-center"
            >
              <svg class="w-8 h-8 mr-3" fill="currentColor" viewBox="0 0 24 24">
                <path
                  d="M24,1.61H14.06L12,5.16,9.94,1.61H0L12,22.39ZM12,14.08,5.16,2.23H9.59L12,6.41l2.41-4.18h4.43Z"
                />
              </svg>
              Frontend
            </h3>
            <div class="space-y-4">
              <div class="flex items-center">
                <div class="w-2 h-2 bg-emerald-400 rounded-full mr-3"></div>
                <span class="text-white"
                  ><strong>Nuxt 3:</strong> Framework Vue.js com SSR e SSG</span
                >
              </div>
              <div class="flex items-center">
                <div class="w-2 h-2 bg-emerald-400 rounded-full mr-3"></div>
                <span class="text-white"
                  ><strong>Vue 3:</strong> Composition API e reatividade</span
                >
              </div>
              <div class="flex items-center">
                <div class="w-2 h-2 bg-emerald-400 rounded-full mr-3"></div>
                <span class="text-white"
                  ><strong>Tailwind CSS:</strong> Estilização utilitária</span
                >
              </div>
              <div class="flex items-center">
                <div class="w-2 h-2 bg-emerald-400 rounded-full mr-3"></div>
                <span class="text-white"
                  ><strong>Vee-Validate:</strong> Validação de formulários</span
                >
              </div>
              <div class="flex items-center">
                <div class="w-2 h-2 bg-emerald-400 rounded-full mr-3"></div>
                <span class="text-white"
                  ><strong>Axios:</strong> Cliente HTTP para APIs</span
                >
              </div>
            </div>
          </div>

          <!-- Backend -->
          <div class="bg-gray-900 rounded-xl p-8 border border-gray-700">
            <h3 class="text-2xl font-bold text-blue-400 mb-6 flex items-center">
              <svg class="w-8 h-8 mr-3" fill="currentColor" viewBox="0 0 24 24">
                <path
                  d="M1.811 10.231c-.047 0-.058-.023-.035-.059l.246-.315c.023-.035.081-.058.128-.058h4.172c.046 0 .058.035.035.082l-.199.303c-.023.047-.082.07-.117.07H1.811zm-.105 2.357c-.047 0-.058-.023-.035-.058l.211-.316c.023-.035.081-.058.129-.058h4.691c.047 0 .07.035.058.081l-.093.28c-.012.047-.058.07-.105.07H1.706zm-.211 2.357c-.047 0-.058-.024-.035-.059l.246-.315c.023-.035.081-.059.128-.059h4.172c.047 0 .059.035.035.082l-.199.303c-.023.047-.081.07-.117.07H1.495zm8.698-2.357c-.047 0-.058-.023-.035-.058l.211-.316c.023-.035.081-.058.128-.058h4.691c.047 0 .07.035.059.081l-.094.28c-.011.047-.058.07-.105.07H10.193z"
                />
              </svg>
              Backend
            </h3>
            <div class="space-y-4">
              <div class="flex items-center">
                <div class="w-2 h-2 bg-blue-400 rounded-full mr-3"></div>
                <span class="text-white"
                  ><strong>Golang:</strong> Linguagem de alta performance</span
                >
              </div>
              <div class="flex items-center">
                <div class="w-2 h-2 bg-blue-400 rounded-full mr-3"></div>
                <span class="text-white"
                  ><strong>Echo Framework:</strong> Web framework
                  minimalista</span
                >
              </div>
              <div class="flex items-center">
                <div class="w-2 h-2 bg-blue-400 rounded-full mr-3"></div>
                <span class="text-white"
                  ><strong>PostgreSQL:</strong> Banco de dados relacional</span
                >
              </div>
              <div class="flex items-center">
                <div class="w-2 h-2 bg-blue-400 rounded-full mr-3"></div>
                <span class="text-white"
                  ><strong>Arquitetura em Camadas:</strong> Handler, Service,
                  Repository</span
                >
              </div>
              <div class="flex items-center">
                <div class="w-2 h-2 bg-blue-400 rounded-full mr-3"></div>
                <span class="text-white"
                  ><strong>Docker:</strong> Containerização e orquestração</span
                >
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Architecture Section -->
    <section class="py-16 sm:py-20">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12">
          <h2 class="text-3xl sm:text-4xl font-bold text-white mb-4">
            Arquitetura do Sistema
          </h2>
          <p class="text-gray-400 text-lg max-w-3xl mx-auto">
            Sistema estruturado seguindo boas práticas de desenvolvimento e
            princípios SOLID
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <!-- Frontend Layer -->
          <div
            class="bg-gradient-to-br from-emerald-500 to-emerald-600 rounded-xl p-6 text-white"
          >
            <h3 class="text-xl font-bold mb-4">Camada de Apresentação</h3>
            <ul class="space-y-2 text-emerald-100">
              <li>• Interface de usuário responsiva</li>
              <li>• Validação de formulários</li>
              <li>• Gerenciamento de estado</li>
              <li>• Integração com APIs</li>
            </ul>
          </div>

          <!-- API Layer -->
          <div
            class="bg-gradient-to-br from-blue-500 to-blue-600 rounded-xl p-6 text-white"
          >
            <h3 class="text-xl font-bold mb-4">Camada de API</h3>
            <ul class="space-y-2 text-blue-100">
              <li>• Endpoints RESTful</li>
              <li>• Middleware de autenticação</li>
              <li>• Validação de dados</li>
              <li>• Tratamento de erros</li>
            </ul>
          </div>

          <!-- Database Layer -->
          <div
            class="bg-gradient-to-br from-purple-500 to-purple-600 rounded-xl p-6 text-white"
          >
            <h3 class="text-xl font-bold mb-4">Camada de Dados</h3>
            <ul class="space-y-2 text-purple-100">
              <li>• Modelo relacional</li>
              <li>• Integridade referencial</li>
              <li>• Queries otimizadas</li>
              <li>• Transações ACID</li>
            </ul>
          </div>
        </div>
      </div>
    </section>

    <!-- Business Rules Section -->
    <section class="py-16 sm:py-20 bg-gray-800">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12">
          <h2 class="text-3xl sm:text-4xl font-bold text-white mb-4">
            Regras de Negócio
          </h2>
          <p class="text-gray-400 text-lg">
            Lógica implementada para garantir consistência dos dados
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div class="bg-gray-900 rounded-xl p-6 border border-gray-700">
            <h3 class="text-xl font-semibold text-emerald-400 mb-4">
              Relacionamento Estabelecimento-Loja
            </h3>
            <p class="text-gray-300 mb-4">
              Um estabelecimento pode ter múltiplas lojas vinculadas, criando um
              relacionamento hierárquico no sistema.
            </p>
            <div class="bg-gray-800 rounded-lg p-4">
              <code class="text-emerald-400 text-sm">
                Estabelecimento 1:N Lojas
              </code>
            </div>
          </div>

          <div class="bg-gray-900 rounded-xl p-6 border border-gray-700">
            <h3 class="text-xl font-semibold text-red-400 mb-4">
              Proteção de Integridade
            </h3>
            <p class="text-gray-300 mb-4">
              Não é possível excluir um estabelecimento que possui lojas
              vinculadas, garantindo a integridade dos dados.
            </p>
            <div class="bg-gray-800 rounded-lg p-4">
              <code class="text-red-400 text-sm">
                DELETE CASCADE: Protegido
              </code>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Footer -->
    <footer class="py-12 border-t border-gray-800">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center">
          <p class="text-gray-400 mb-4">
            Sistema desenvolvido como desafio técnico demonstrando conhecimentos
            em desenvolvimento fullstack
          </p>
          <div class="flex justify-center space-x-6">
            <span class="text-gray-500 text-sm">Docker + Docker Compose</span>

            <span class="text-gray-500 text-sm">•</span>
            <span class="text-gray-500 text-sm">Documentação Completa</span>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>
