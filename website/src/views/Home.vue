<template>
  <div class="home">
    <!-- Hero Section -->
    <section class="relative overflow-hidden">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-20 md:py-32">
        <div class="text-center">
          <h1 class="text-5xl md:text-7xl font-black mb-6 leading-tight">
            <span class="bg-gradient-to-r from-zettl-purple via-purple-400 to-pink-400 bg-clip-text text-transparent">
              AI-Powered
            </span>
            <br />
            <span class="text-white">Note-Taking</span>
            <br />
            <span class="text-gray-400">in Your Terminal</span>
          </h1>
          <p class="text-xl md:text-2xl text-gray-300 mb-10 max-w-3xl mx-auto">
            Zettl combines the simplicity of terminal-based note-taking with the power of AI.
            Create, search, and chat with your notes using semantic search and embeddings.
          </p>
          <div class="flex flex-col sm:flex-row gap-4 justify-center items-center">
            <a href="#installation" class="btn-primary text-lg px-8 py-4">
              Get Started
            </a>
            <a href="https://github.com/thisguymartin/zettl" target="_blank" class="btn-secondary text-lg px-8 py-4">
              View on GitHub
            </a>
          </div>

          <!-- Terminal Demo -->
          <div class="mt-16 bg-gray-900 rounded-lg border border-gray-700 shadow-2xl overflow-hidden max-w-4xl mx-auto">
            <div class="bg-gray-800 px-4 py-2 flex items-center space-x-2 border-b border-gray-700">
              <div class="w-3 h-3 rounded-full bg-red-500"></div>
              <div class="w-3 h-3 rounded-full bg-yellow-500"></div>
              <div class="w-3 h-3 rounded-full bg-green-500"></div>
              <span class="text-gray-400 text-sm ml-4">zettl</span>
            </div>
            <div class="p-6 font-mono text-sm md:text-base text-left">
              <pre class="text-green-400">$ zettl</pre>
              <pre class="text-gray-400 mt-4">
╔══════════════════════════════════════════════════════╗
║  ███████╗███████╗████████╗████████╗██╗              ║
║  ╚══███╔╝██╔════╝╚══██╔══╝╚══██╔══╝██║              ║
║    ███╔╝ █████╗     ██║      ██║   ██║              ║
║   ███╔╝  ██╔══╝     ██║      ██║   ██║              ║
║  ███████╗███████╗   ██║      ██║   ███████╗         ║
║  ╚══════╝╚══════╝   ╚═╝      ╚═╝   ╚══════╝         ║
╚══════════════════════════════════════════════════════╝

<span class="text-white">Welcome to your digital notebook!</span>

<span class="text-purple-400">1. New Note (n)</span>     - Create a new note
<span class="text-purple-400">2. List Notes (l)</span>   - Browse existing notes
<span class="text-purple-400">3. Search (s)</span>       - Find notes with fuzzy search
<span class="text-purple-400">4. AI Features (a)</span>  - Chat, embeddings, and semantic search

<span class="text-gray-500">Press the number or letter key to navigate.</span>
              </pre>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Features Overview -->
    <section class="py-20 bg-zettl-gray">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="section-title">Why Zettl?</h2>
          <p class="text-xl text-gray-400">Powerful features designed for productivity</p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div class="card">
            <div class="text-4xl mb-4">🚀</div>
            <h3 class="text-2xl font-bold text-white mb-3">Lightning Fast</h3>
            <p class="text-gray-400">
              Built with Go and Bubble Tea for incredible performance.
              No loading screens, no lag - just pure speed.
            </p>
          </div>

          <div class="card">
            <div class="text-4xl mb-4">🤖</div>
            <h3 class="text-2xl font-bold text-white mb-3">AI-Powered</h3>
            <p class="text-gray-400">
              Semantic search using OpenAI embeddings. Chat with your notes
              and get intelligent responses based on context.
            </p>
          </div>

          <div class="card">
            <div class="text-4xl mb-4">💾</div>
            <h3 class="text-2xl font-bold text-white mb-3">Local First</h3>
            <p class="text-gray-400">
              All data stored locally in SQLite. Your notes are yours.
              Works offline, syncs when you want.
            </p>
          </div>

          <div class="card">
            <div class="text-4xl mb-4">🔍</div>
            <h3 class="text-2xl font-bold text-white mb-3">Smart Search</h3>
            <p class="text-gray-400">
              Find notes by meaning, not just keywords. Vector similarity
              search finds related content intelligently.
            </p>
          </div>

          <div class="card">
            <div class="text-4xl mb-4">📝</div>
            <h3 class="text-2xl font-bold text-white mb-3">Clean Interface</h3>
            <p class="text-gray-400">
              Beautiful TUI with keyboard-first navigation. No mouse needed.
              Distraction-free writing experience.
            </p>
          </div>

          <div class="card">
            <div class="text-4xl mb-4">🔒</div>
            <h3 class="text-2xl font-bold text-white mb-3">Privacy Focused</h3>
            <p class="text-gray-400">
              AI features are optional. Use local embeddings or skip AI entirely.
              Your data, your choice.
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- Installation Section -->
    <section id="installation" class="py-20">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="section-title">Get Started in Seconds</h2>
          <p class="text-xl text-gray-400">Simple installation, powerful results</p>
        </div>

        <div class="max-w-3xl mx-auto space-y-8">
          <div class="card">
            <div class="flex items-start space-x-4">
              <div class="flex-shrink-0 w-8 h-8 bg-zettl-purple rounded-full flex items-center justify-center text-white font-bold">
                1
              </div>
              <div class="flex-grow">
                <h3 class="text-xl font-bold text-white mb-2">Clone the Repository</h3>
                <div class="bg-gray-900 rounded p-4 font-mono text-sm text-green-400">
                  git clone https://github.com/thisguymartin/zettl.git<br />
                  cd zettl
                </div>
              </div>
            </div>
          </div>

          <div class="card">
            <div class="flex items-start space-x-4">
              <div class="flex-shrink-0 w-8 h-8 bg-zettl-purple rounded-full flex items-center justify-center text-white font-bold">
                2
              </div>
              <div class="flex-grow">
                <h3 class="text-xl font-bold text-white mb-2">Build the App</h3>
                <div class="bg-gray-900 rounded p-4 font-mono text-sm text-green-400">
                  make build
                </div>
              </div>
            </div>
          </div>

          <div class="card">
            <div class="flex items-start space-x-4">
              <div class="flex-shrink-0 w-8 h-8 bg-zettl-purple rounded-full flex items-center justify-center text-white font-bold">
                3
              </div>
              <div class="flex-grow">
                <h3 class="text-xl font-bold text-white mb-2">Set API Key (Optional)</h3>
                <div class="bg-gray-900 rounded p-4 font-mono text-sm text-green-400">
                  export OPENAI_API_KEY="sk-your-key-here"
                </div>
                <p class="text-gray-400 mt-2 text-sm">
                  AI features require an OpenAI API key. Skip this for basic note-taking.
                </p>
              </div>
            </div>
          </div>

          <div class="card">
            <div class="flex items-start space-x-4">
              <div class="flex-shrink-0 w-8 h-8 bg-zettl-purple rounded-full flex items-center justify-center text-white font-bold">
                4
              </div>
              <div class="flex-grow">
                <h3 class="text-xl font-bold text-white mb-2">Run Zettl</h3>
                <div class="bg-gray-900 rounded p-4 font-mono text-sm text-green-400">
                  ./bin/cli
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="text-center mt-12">
          <router-link to="/docs" class="btn-primary text-lg">
            Read Full Documentation →
          </router-link>
        </div>
      </div>
    </section>

    <!-- CTA Section -->
    <section class="py-20 bg-gradient-to-r from-zettl-purple to-purple-600">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h2 class="text-4xl md:text-5xl font-bold text-white mb-6">
          Ready to Transform Your Note-Taking?
        </h2>
        <p class="text-xl text-purple-100 mb-8 max-w-2xl mx-auto">
          Join developers who are using AI to organize and interact with their knowledge base.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <a href="https://github.com/thisguymartin/zettl" target="_blank" class="px-8 py-4 bg-white text-zettl-purple rounded-lg font-bold text-lg hover:bg-gray-100 transition-colors duration-200">
            ⭐ Star on GitHub
          </a>
          <router-link to="/pricing" class="px-8 py-4 border-2 border-white text-white rounded-lg font-bold text-lg hover:bg-white hover:text-zettl-purple transition-all duration-200">
            View Pricing
          </router-link>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
// No additional logic needed
</script>
