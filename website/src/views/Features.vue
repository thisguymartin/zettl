<template>
  <div class="features py-20">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- Header -->
      <div class="text-center mb-20">
        <h1 class="section-title mb-6">Powerful Features</h1>
        <p class="text-xl text-gray-400 max-w-3xl mx-auto">
          Everything you need to manage your notes with AI-powered intelligence
        </p>
      </div>

      <!-- Core Features -->
      <div class="mb-20">
        <h2 class="text-3xl font-bold text-white mb-10 text-center">Core Functionality</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div
            v-for="feature in coreFeatures"
            :key="feature.title"
            class="card"
          >
            <div class="text-4xl mb-4">{{ feature.icon }}</div>
            <h3 class="text-2xl font-bold text-white mb-3">{{ feature.title }}</h3>
            <p class="text-gray-400 mb-4">{{ feature.description }}</p>
            <ul class="space-y-2">
              <li
                v-for="detail in feature.details"
                :key="detail"
                class="flex items-start text-gray-300"
              >
                <span class="text-zettl-purple mr-2">✓</span>
                <span>{{ detail }}</span>
              </li>
            </ul>
          </div>
        </div>
      </div>

      <!-- AI Features -->
      <div class="mb-20 bg-gradient-to-r from-purple-900 to-indigo-900 rounded-2xl p-8 md:p-12">
        <h2 class="text-3xl font-bold text-white mb-10 text-center">AI-Powered Features</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div
            v-for="feature in aiFeatures"
            :key="feature.title"
            class="bg-black bg-opacity-30 rounded-lg p-6 backdrop-blur-sm"
          >
            <div class="text-4xl mb-4">{{ feature.icon }}</div>
            <h3 class="text-2xl font-bold text-white mb-3">{{ feature.title }}</h3>
            <p class="text-purple-100 mb-4">{{ feature.description }}</p>
            <div class="bg-black bg-opacity-40 rounded p-4 font-mono text-sm text-green-400">
              {{ feature.code }}
            </div>
          </div>
        </div>
      </div>

      <!-- Technical Specs -->
      <div class="mb-20">
        <h2 class="text-3xl font-bold text-white mb-10 text-center">Technical Specifications</h2>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div
            v-for="spec in techSpecs"
            :key="spec.label"
            class="card text-center"
          >
            <div class="text-2xl font-bold text-zettl-purple mb-2">{{ spec.value }}</div>
            <div class="text-gray-400">{{ spec.label }}</div>
          </div>
        </div>
      </div>

      <!-- Keyboard Shortcuts -->
      <div class="mb-20">
        <h2 class="text-3xl font-bold text-white mb-10 text-center">Keyboard-First Navigation</h2>
        <div class="card max-w-3xl mx-auto">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div
              v-for="shortcut in shortcuts"
              :key="shortcut.key"
              class="flex justify-between items-center py-2 border-b border-gray-700 last:border-0"
            >
              <span class="text-gray-300">{{ shortcut.action }}</span>
              <kbd class="px-3 py-1 bg-gray-800 text-zettl-purple rounded font-mono text-sm border border-gray-600">
                {{ shortcut.key }}
              </kbd>
            </div>
          </div>
        </div>
      </div>

      <!-- Integration -->
      <div class="text-center">
        <h2 class="text-3xl font-bold text-white mb-6">Built With Industry Leaders</h2>
        <div class="flex flex-wrap justify-center items-center gap-8 mb-10">
          <div class="text-gray-400 font-semibold text-lg">Go 1.24+</div>
          <div class="text-gray-600">•</div>
          <div class="text-gray-400 font-semibold text-lg">Bubble Tea</div>
          <div class="text-gray-600">•</div>
          <div class="text-gray-400 font-semibold text-lg">SQLite</div>
          <div class="text-gray-600">•</div>
          <div class="text-gray-400 font-semibold text-lg">OpenAI</div>
        </div>
        <router-link to="/docs" class="btn-primary text-lg">
          View Full Documentation →
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
const coreFeatures = [
  {
    icon: '📝',
    title: 'CRUD Operations',
    description: 'Complete note management with an intuitive interface',
    details: [
      'Create notes with automatic title generation',
      'Edit notes with multi-line support',
      'Delete notes with confirmation',
      'Browse all notes with timestamps',
    ],
  },
  {
    icon: '🔍',
    title: 'Smart Search',
    description: 'Multiple search methods to find exactly what you need',
    details: [
      'Fuzzy search across titles and content',
      'Tag-based filtering',
      'Real-time search results',
      'Semantic search with embeddings',
    ],
  },
  {
    icon: '💾',
    title: 'Local Storage',
    description: 'All your data stored securely on your machine',
    details: [
      'SQLite database for reliability',
      'No cloud dependencies',
      'Fast queries with indexing',
      'Automatic backups',
    ],
  },
  {
    icon: '📤',
    title: 'Export & Import',
    description: 'Take your data anywhere with flexible formats',
    details: [
      'JSON export with metadata',
      'Markdown files with frontmatter',
      'Batch import capabilities',
      'Preserve all note attributes',
    ],
  },
]

const aiFeatures = [
  {
    icon: '🤖',
    title: 'AI Chat',
    description: 'Conversational AI that understands your notes',
    code: 'AI: Based on your notes, here are 3 ideas...',
  },
  {
    icon: '🎯',
    title: 'Embeddings',
    description: 'Vector representations for semantic understanding',
    code: '1536-dimensional vectors via OpenAI',
  },
  {
    icon: '🔗',
    title: 'Similarity Search',
    description: 'Find related notes using cosine similarity',
    code: 'similarity_score: 0.87 (highly related)',
  },
  {
    icon: '💬',
    title: 'Chat History',
    description: 'Persistent conversations across sessions',
    code: 'Conversation saved to ~/.zettl/zettl.db',
  },
]

const techSpecs = [
  { value: 'text-embedding-3-small', label: 'Embedding Model' },
  { value: '1536 dims', label: 'Vector Size' },
  { value: 'gpt-4o-mini', label: 'Chat Model' },
  { value: 'SQLite', label: 'Database' },
  { value: 'Bubble Tea', label: 'TUI Framework' },
  { value: '<10MB', label: 'Binary Size' },
]

const shortcuts = [
  { action: 'Create Note', key: 'n' },
  { action: 'List Notes', key: 'l' },
  { action: 'Search', key: 's' },
  { action: 'AI Features', key: 'a' },
  { action: 'Navigate Up', key: '↑ / k' },
  { action: 'Navigate Down', key: '↓ / j' },
  { action: 'Open Note', key: 'Enter' },
  { action: 'Save', key: 'Ctrl+S' },
  { action: 'Back', key: 'Esc' },
  { action: 'Quit', key: 'q' },
]
</script>
