<template>
  <div class="pricing py-20">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- Header -->
      <div class="text-center mb-20">
        <h1 class="section-title mb-6">Simple, Transparent Pricing</h1>
        <p class="text-xl text-gray-400 max-w-3xl mx-auto">
          Zettl is free and open source. You only pay for the AI features you use.
        </p>
      </div>

      <!-- Pricing Cards -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-20">
        <!-- Free Tier -->
        <div class="card">
          <div class="text-center mb-6">
            <h3 class="text-2xl font-bold text-white mb-2">Free</h3>
            <div class="text-4xl font-bold text-zettl-purple mb-2">$0</div>
            <div class="text-gray-400">Forever</div>
          </div>
          <ul class="space-y-3 mb-8">
            <li class="flex items-start text-gray-300">
              <span class="text-green-500 mr-2">✓</span>
              <span>Unlimited notes</span>
            </li>
            <li class="flex items-start text-gray-300">
              <span class="text-green-500 mr-2">✓</span>
              <span>Local SQLite storage</span>
            </li>
            <li class="flex items-start text-gray-300">
              <span class="text-green-500 mr-2">✓</span>
              <span>Full-text search</span>
            </li>
            <li class="flex items-start text-gray-300">
              <span class="text-green-500 mr-2">✓</span>
              <span>Tag system</span>
            </li>
            <li class="flex items-start text-gray-300">
              <span class="text-green-500 mr-2">✓</span>
              <span>Export/Import</span>
            </li>
            <li class="flex items-start text-gray-400">
              <span class="text-gray-600 mr-2">✗</span>
              <span>AI features</span>
            </li>
          </ul>
          <a href="https://github.com/thisguymartin/zettl" target="_blank" class="btn-secondary w-full block text-center">
            Download Free
          </a>
        </div>

        <!-- Pay as You Go -->
        <div class="card border-2 border-zettl-purple relative">
          <div class="absolute -top-4 left-1/2 transform -translate-x-1/2 bg-zettl-purple text-white px-4 py-1 rounded-full text-sm font-semibold">
            Most Popular
          </div>
          <div class="text-center mb-6">
            <h3 class="text-2xl font-bold text-white mb-2">Pay as You Go</h3>
            <div class="text-4xl font-bold text-zettl-purple mb-2">~$5-15</div>
            <div class="text-gray-400">per month</div>
          </div>
          <ul class="space-y-3 mb-8">
            <li class="flex items-start text-gray-300">
              <span class="text-green-500 mr-2">✓</span>
              <span>Everything in Free</span>
            </li>
            <li class="flex items-start text-gray-300">
              <span class="text-green-500 mr-2">✓</span>
              <span>OpenAI embeddings</span>
            </li>
            <li class="flex items-start text-gray-300">
              <span class="text-green-500 mr-2">✓</span>
              <span>Semantic search</span>
            </li>
            <li class="flex items-start text-gray-300">
              <span class="text-green-500 mr-2">✓</span>
              <span>AI chat</span>
            </li>
            <li class="flex items-start text-gray-300">
              <span class="text-green-500 mr-2">✓</span>
              <span>Context-aware responses</span>
            </li>
          </ul>
          <a href="https://platform.openai.com/signup" target="_blank" class="btn-primary w-full block text-center">
            Get API Key
          </a>
        </div>

        <!-- Self-Hosted -->
        <div class="card">
          <div class="text-center mb-6">
            <h3 class="text-2xl font-bold text-white mb-2">Self-Hosted</h3>
            <div class="text-4xl font-bold text-zettl-purple mb-2">$0</div>
            <div class="text-gray-400">One-time setup</div>
          </div>
          <ul class="space-y-3 mb-8">
            <li class="flex items-start text-gray-300">
              <span class="text-green-500 mr-2">✓</span>
              <span>Everything in Free</span>
            </li>
            <li class="flex items-start text-gray-300">
              <span class="text-green-500 mr-2">✓</span>
              <span>Local LLM support</span>
            </li>
            <li class="flex items-start text-gray-300">
              <span class="text-green-500 mr-2">✓</span>
              <span>Ollama integration</span>
            </li>
            <li class="flex items-start text-gray-300">
              <span class="text-green-500 mr-2">✓</span>
              <span>100% private</span>
            </li>
            <li class="flex items-start text-gray-400">
              <span class="text-yellow-500 mr-2">⚠</span>
              <span>Requires GPU</span>
            </li>
          </ul>
          <router-link to="/docs" class="btn-secondary w-full block text-center">
            View Docs
          </router-link>
        </div>
      </div>

      <!-- Cost Breakdown -->
      <div class="mb-20">
        <h2 class="text-3xl font-bold text-white mb-10 text-center">Cost Breakdown</h2>
        <div class="max-w-4xl mx-auto">
          <div class="card mb-8">
            <h3 class="text-2xl font-bold text-white mb-4">🎯 Embeddings</h3>
            <div class="space-y-3 text-gray-300">
              <div class="flex justify-between">
                <span>Cost per 1,000 tokens</span>
                <span class="text-zettl-purple font-semibold">$0.0001</span>
              </div>
              <div class="flex justify-between">
                <span>Average note size</span>
                <span class="text-gray-400">100-500 tokens</span>
              </div>
              <div class="flex justify-between">
                <span>1,000 notes (one-time)</span>
                <span class="text-green-500 font-semibold">~$0.01-0.05</span>
              </div>
            </div>
            <div class="mt-4 p-4 bg-green-900 bg-opacity-20 border border-green-700 rounded">
              <p class="text-green-400 text-sm">
                💡 Embeddings are generated once per note. Updates only re-generate changed notes.
              </p>
            </div>
          </div>

          <div class="card">
            <h3 class="text-2xl font-bold text-white mb-4">💬 Chat</h3>
            <div class="space-y-3 text-gray-300 mb-4">
              <div class="flex justify-between">
                <span>Cost per conversation</span>
                <span class="text-zettl-purple font-semibold">$0.01-0.03</span>
              </div>
              <div class="flex justify-between">
                <span>With context (5 notes)</span>
                <span class="text-zettl-purple font-semibold">$0.05-0.15</span>
              </div>
              <div class="flex justify-between">
                <span>100 conversations/month</span>
                <span class="text-yellow-500 font-semibold">~$5-15</span>
              </div>
            </div>
            <div class="p-4 bg-blue-900 bg-opacity-20 border border-blue-700 rounded">
              <p class="text-blue-400 text-sm">
                💡 More context = higher cost. Limit context to relevant notes for best value.
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- Example Scenarios -->
      <div class="mb-20">
        <h2 class="text-3xl font-bold text-white mb-10 text-center">Example Scenarios</h2>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="card">
            <h4 class="text-xl font-bold text-white mb-3">💼 Professional</h4>
            <div class="text-sm text-gray-400 space-y-2 mb-4">
              <div>• 2,000 notes</div>
              <div>• 200 chats/month</div>
              <div>• Heavy AI usage</div>
            </div>
            <div class="text-2xl font-bold text-zettl-purple">~$20-30/mo</div>
          </div>

          <div class="card">
            <h4 class="text-xl font-bold text-white mb-3">👨‍💻 Developer</h4>
            <div class="text-sm text-gray-400 space-y-2 mb-4">
              <div>• 500 notes</div>
              <div>• 50 chats/month</div>
              <div>• Moderate usage</div>
            </div>
            <div class="text-2xl font-bold text-zettl-purple">~$3-7/mo</div>
          </div>

          <div class="card">
            <h4 class="text-xl font-bold text-white mb-3">📚 Student</h4>
            <div class="text-sm text-gray-400 space-y-2 mb-4">
              <div>• 200 notes</div>
              <div>• 20 chats/month</div>
              <div>• Light usage</div>
            </div>
            <div class="text-2xl font-bold text-zettl-purple">~$1-3/mo</div>
          </div>
        </div>
      </div>

      <!-- FAQ -->
      <div class="max-w-3xl mx-auto">
        <h2 class="text-3xl font-bold text-white mb-10 text-center">Frequently Asked Questions</h2>
        <div class="space-y-6">
          <div
            v-for="faq in faqs"
            :key="faq.question"
            class="card"
          >
            <h3 class="text-lg font-bold text-white mb-2">{{ faq.question }}</h3>
            <p class="text-gray-400">{{ faq.answer }}</p>
          </div>
        </div>
      </div>

      <!-- CTA -->
      <div class="text-center mt-20">
        <h2 class="text-3xl font-bold text-white mb-6">Ready to Get Started?</h2>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <a href="https://github.com/thisguymartin/zettl" target="_blank" class="btn-primary text-lg px-8 py-4">
            Download Zettl
          </a>
          <router-link to="/docs" class="btn-secondary text-lg px-8 py-4">
            Read Documentation
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const faqs = [
  {
    question: 'Is Zettl really free?',
    answer: 'Yes! Zettl is completely free and open source. You only pay for OpenAI API usage if you choose to use AI features. All core note-taking functionality is free forever.',
  },
  {
    question: 'Can I use Zettl without AI features?',
    answer: 'Absolutely! Zettl works perfectly without any AI features. You get full note management, search, tags, and export/import at no cost.',
  },
  {
    question: 'How much does OpenAI API cost?',
    answer: 'OpenAI charges based on usage. For typical use (1000 notes, 100 chats/month), expect $5-15/month. You only pay for what you use.',
  },
  {
    question: 'Can I use local AI models?',
    answer: 'Local LLM support is coming soon! You will be able to use Ollama or other local models for completely free AI features.',
  },
  {
    question: 'Where is my data stored?',
    answer: 'All your data is stored locally in a SQLite database at ~/.zettl/zettl.db. Your notes never leave your machine unless you explicitly use AI features.',
  },
  {
    question: 'Can I export my notes?',
    answer: 'Yes! Export to JSON (with full metadata) or individual Markdown files. Your data is always portable.',
  },
]
</script>
