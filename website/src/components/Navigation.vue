<template>
  <nav class="bg-zettl-gray border-b border-gray-800 sticky top-0 z-50 backdrop-blur-sm bg-opacity-95">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between h-16">
        <div class="flex items-center">
          <router-link to="/" class="flex items-center space-x-3">
            <div class="text-3xl font-bold bg-gradient-to-r from-zettl-purple to-purple-400 bg-clip-text text-transparent">
              ╔═══╗
            </div>
            <span class="text-2xl font-bold text-white">Zettl</span>
          </router-link>
        </div>

        <div class="hidden md:flex items-center space-x-8">
          <router-link
            v-for="item in navItems"
            :key="item.path"
            :to="item.path"
            class="text-gray-300 hover:text-zettl-purple transition-colors duration-200 font-medium"
            active-class="text-zettl-purple"
          >
            {{ item.name }}
          </router-link>
          <a
            href="https://github.com/thisguymartin/zettl"
            target="_blank"
            class="btn-primary"
          >
            ⭐ Star on GitHub
          </a>
        </div>

        <div class="md:hidden flex items-center">
          <button
            @click="mobileMenuOpen = !mobileMenuOpen"
            class="text-gray-300 hover:text-white p-2"
          >
            <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path v-if="!mobileMenuOpen" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
              <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Mobile menu -->
    <div v-if="mobileMenuOpen" class="md:hidden bg-zettl-gray border-t border-gray-800">
      <div class="px-2 pt-2 pb-3 space-y-1">
        <router-link
          v-for="item in navItems"
          :key="item.path"
          :to="item.path"
          class="block px-3 py-2 text-gray-300 hover:text-zettl-purple transition-colors duration-200"
          active-class="text-zettl-purple bg-gray-800"
          @click="mobileMenuOpen = false"
        >
          {{ item.name }}
        </router-link>
        <a
          href="https://github.com/thisguymartin/zettl"
          target="_blank"
          class="block px-3 py-2 text-center btn-primary"
        >
          ⭐ Star on GitHub
        </a>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref } from 'vue'

const mobileMenuOpen = ref(false)

const navItems = [
  { name: 'Home', path: '/' },
  { name: 'Features', path: '/features' },
  { name: 'Pricing', path: '/pricing' },
  { name: 'Docs', path: '/docs' },
]
</script>
