<template>
  <div class="docs py-20">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
        <!-- Sidebar -->
        <aside class="lg:col-span-1">
          <div class="sticky top-24 card">
            <h3 class="text-lg font-bold text-white mb-4">Documentation</h3>
            <nav class="space-y-2">
              <a
                v-for="section in sections"
                :key="section.id"
                :href="`#${section.id}`"
                class="block text-gray-400 hover:text-zettl-purple transition-colors py-1"
                @click.prevent="scrollTo(section.id)"
              >
                {{ section.title }}
              </a>
            </nav>
          </div>
        </aside>

        <!-- Main Content -->
        <main class="lg:col-span-3 space-y-12">
          <!-- Getting Started -->
          <section id="getting-started" class="card">
            <h2 class="text-3xl font-bold text-white mb-6">Getting Started</h2>

            <h3 class="text-xl font-bold text-white mb-3">Installation</h3>
            <p class="text-gray-400 mb-4">
              Zettl requires Go 1.24.3 or later. Follow these steps to install:
            </p>
            <div class="bg-gray-900 rounded p-4 mb-6 font-mono text-sm overflow-x-auto">
              <pre class="text-green-400"># Clone the repository
git clone https://github.com/thisguymartin/zettl.git
cd zettl

# Build the application
make build

# Run Zettl
./bin/cli</pre>
            </div>

            <h3 class="text-xl font-bold text-white mb-3">Configuration</h3>
            <p class="text-gray-400 mb-4">
              Set up your OpenAI API key for AI features (optional):
            </p>
            <div class="bg-gray-900 rounded p-4 mb-6 font-mono text-sm overflow-x-auto">
              <pre class="text-green-400"># Method 1: Environment variable
export OPENAI_API_KEY="sk-your-key-here"

# Method 2: Config file
echo '{
  "openai_api_key": "sk-your-key-here",
  "db_path": "~/.zettl/zettl.db"
}' > ~/.zettl/config.json</pre>
            </div>
          </section>

          <!-- Usage -->
          <section id="usage" class="card">
            <h2 class="text-3xl font-bold text-white mb-6">Usage</h2>

            <h3 class="text-xl font-bold text-white mb-3">Basic Operations</h3>
            <div class="space-y-4 mb-6">
              <div>
                <h4 class="text-white font-semibold mb-2">Creating Notes</h4>
                <p class="text-gray-400 mb-2">Press <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">n</kbd> or <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">1</kbd> from the main menu.</p>
                <ul class="list-disc list-inside text-gray-400 space-y-1 ml-4">
                  <li>Type your content</li>
                  <li>Press Enter for new lines</li>
                  <li>Press Tab for indentation</li>
                  <li>Press Ctrl+S to save</li>
                </ul>
              </div>

              <div>
                <h4 class="text-white font-semibold mb-2">Searching Notes</h4>
                <p class="text-gray-400 mb-2">Press <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">s</kbd> or <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">3</kbd> for search.</p>
                <ul class="list-disc list-inside text-gray-400 space-y-1 ml-4">
                  <li>Type your search query</li>
                  <li>Results update in real-time</li>
                  <li>Press Enter to view results</li>
                  <li>Use / from note list for quick search</li>
                </ul>
              </div>
            </div>
          </section>

          <!-- AI Features -->
          <section id="ai-features" class="card bg-gradient-to-r from-purple-900 to-indigo-900">
            <h2 class="text-3xl font-bold text-white mb-6">AI Features</h2>

            <h3 class="text-xl font-bold text-white mb-3">Chat with Your Notes</h3>
            <p class="text-purple-100 mb-4">
              Access the AI chat by pressing <kbd class="px-2 py-1 bg-black bg-opacity-40 rounded text-zettl-purple">a</kbd> then <kbd class="px-2 py-1 bg-black bg-opacity-40 rounded text-zettl-purple">c</kbd>.
            </p>
            <div class="bg-black bg-opacity-40 rounded p-4 mb-6 font-mono text-sm">
              <pre class="text-green-400">[You]: Summarize my recent project ideas

[AI]: Based on your 5 recent notes, here are the key themes:
1. Web scraping automation tools
2. AI-powered note organization
3. Terminal UI improvements
...</pre>
            </div>

            <h3 class="text-xl font-bold text-white mb-3">Semantic Search</h3>
            <p class="text-purple-100 mb-4">
              Generate embeddings to enable semantic search:
            </p>
            <ol class="list-decimal list-inside text-purple-100 space-y-2 ml-4">
              <li>Navigate to AI Features menu (press <kbd class="px-2 py-1 bg-black bg-opacity-40 rounded text-zettl-purple">a</kbd>)</li>
              <li>Select "Generate Embeddings" (press <kbd class="px-2 py-1 bg-black bg-opacity-40 rounded text-zettl-purple">2</kbd>)</li>
              <li>Wait for embedding generation to complete</li>
              <li>Use "Semantic Search" to find similar notes</li>
            </ol>
          </section>

          <!-- Keyboard Shortcuts -->
          <section id="keyboard-shortcuts" class="card">
            <h2 class="text-3xl font-bold text-white mb-6">Keyboard Shortcuts</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <h4 class="text-white font-semibold mb-3">Main Menu</h4>
                <div class="space-y-2">
                  <div class="flex justify-between items-center">
                    <span class="text-gray-400">New Note</span>
                    <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">n / 1</kbd>
                  </div>
                  <div class="flex justify-between items-center">
                    <span class="text-gray-400">List Notes</span>
                    <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">l / 2</kbd>
                  </div>
                  <div class="flex justify-between items-center">
                    <span class="text-gray-400">Search</span>
                    <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">s / 3</kbd>
                  </div>
                  <div class="flex justify-between items-center">
                    <span class="text-gray-400">AI Features</span>
                    <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">a / 4</kbd>
                  </div>
                </div>
              </div>

              <div>
                <h4 class="text-white font-semibold mb-3">Navigation</h4>
                <div class="space-y-2">
                  <div class="flex justify-between items-center">
                    <span class="text-gray-400">Move Up</span>
                    <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">↑ / k</kbd>
                  </div>
                  <div class="flex justify-between items-center">
                    <span class="text-gray-400">Move Down</span>
                    <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">↓ / j</kbd>
                  </div>
                  <div class="flex justify-between items-center">
                    <span class="text-gray-400">Select</span>
                    <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">Enter</kbd>
                  </div>
                  <div class="flex justify-between items-center">
                    <span class="text-gray-400">Back</span>
                    <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">Esc</kbd>
                  </div>
                </div>
              </div>

              <div>
                <h4 class="text-white font-semibold mb-3">Editing</h4>
                <div class="space-y-2">
                  <div class="flex justify-between items-center">
                    <span class="text-gray-400">Save</span>
                    <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">Ctrl+S</kbd>
                  </div>
                  <div class="flex justify-between items-center">
                    <span class="text-gray-400">New Line</span>
                    <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">Enter</kbd>
                  </div>
                  <div class="flex justify-between items-center">
                    <span class="text-gray-400">Indent</span>
                    <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">Tab</kbd>
                  </div>
                  <div class="flex justify-between items-center">
                    <span class="text-gray-400">Delete</span>
                    <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">Backspace</kbd>
                  </div>
                </div>
              </div>

              <div>
                <h4 class="text-white font-semibold mb-3">General</h4>
                <div class="space-y-2">
                  <div class="flex justify-between items-center">
                    <span class="text-gray-400">Quit</span>
                    <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">q / Ctrl+C</kbd>
                  </div>
                  <div class="flex justify-between items-center">
                    <span class="text-gray-400">Quick Search</span>
                    <kbd class="px-2 py-1 bg-gray-800 rounded text-zettl-purple">/</kbd>
                  </div>
                </div>
              </div>
            </div>
          </section>

          <!-- Export/Import -->
          <section id="export-import" class="card">
            <h2 class="text-3xl font-bold text-white mb-6">Export & Import</h2>

            <h3 class="text-xl font-bold text-white mb-3">Exporting Notes</h3>
            <p class="text-gray-400 mb-4">
              Export your notes to JSON or Markdown format (coming soon in UI, currently via code):
            </p>
            <div class="bg-gray-900 rounded p-4 mb-6 font-mono text-sm overflow-x-auto">
              <pre class="text-green-400"># JSON export (with metadata)
zettl export --format json --output notes.json

# Markdown export (individual files)
zettl export --format markdown --output ./notes-export/</pre>
            </div>

            <h3 class="text-xl font-bold text-white mb-3">Importing Notes</h3>
            <div class="bg-gray-900 rounded p-4 mb-6 font-mono text-sm overflow-x-auto">
              <pre class="text-green-400"># Import from JSON
zettl import --format json --input notes.json</pre>
            </div>
          </section>

          <!-- Troubleshooting -->
          <section id="troubleshooting" class="card">
            <h2 class="text-3xl font-bold text-white mb-6">Troubleshooting</h2>

            <div class="space-y-6">
              <div>
                <h4 class="text-white font-semibold mb-2">AI Features Not Working</h4>
                <ul class="list-disc list-inside text-gray-400 space-y-1 ml-4">
                  <li>Verify OPENAI_API_KEY is set correctly</li>
                  <li>Check your OpenAI account has credits</li>
                  <li>Ensure network connectivity</li>
                  <li>Review error messages in the UI</li>
                </ul>
              </div>

              <div>
                <h4 class="text-white font-semibold mb-2">Database Errors</h4>
                <ul class="list-disc list-inside text-gray-400 space-y-1 ml-4">
                  <li>Check file permissions for ~/.zettl/</li>
                  <li>Ensure sufficient disk space</li>
                  <li>Verify SQLite installation</li>
                  <li>Database location: ~/.zettl/zettl.db</li>
                </ul>
              </div>

              <div>
                <h4 class="text-white font-semibold mb-2">Build Issues</h4>
                <ul class="list-disc list-inside text-gray-400 space-y-1 ml-4">
                  <li>Ensure Go 1.24.3+ is installed</li>
                  <li>Run: go mod tidy</li>
                  <li>Check for missing system libraries</li>
                  <li>Clear build cache: go clean -cache</li>
                </ul>
              </div>
            </div>
          </section>

          <!-- Contributing -->
          <section id="contributing" class="card">
            <h2 class="text-3xl font-bold text-white mb-6">Contributing</h2>
            <p class="text-gray-400 mb-4">
              Contributions are welcome! Here's how you can help:
            </p>
            <ol class="list-decimal list-inside text-gray-400 space-y-2 ml-4 mb-6">
              <li>Fork the repository on GitHub</li>
              <li>Create a feature branch: <code class="text-zettl-purple">git checkout -b feature-name</code></li>
              <li>Make your changes and commit: <code class="text-zettl-purple">git commit -m "Add feature"</code></li>
              <li>Push to the branch: <code class="text-zettl-purple">git push origin feature-name</code></li>
              <li>Submit a pull request</li>
            </ol>
            <a href="https://github.com/thisguymartin/zettl/issues" target="_blank" class="btn-primary">
              Report an Issue on GitHub
            </a>
          </section>
        </main>
      </div>
    </div>
  </div>
</template>

<script setup>
const sections = [
  { id: 'getting-started', title: 'Getting Started' },
  { id: 'usage', title: 'Usage' },
  { id: 'ai-features', title: 'AI Features' },
  { id: 'keyboard-shortcuts', title: 'Keyboard Shortcuts' },
  { id: 'export-import', title: 'Export & Import' },
  { id: 'troubleshooting', title: 'Troubleshooting' },
  { id: 'contributing', title: 'Contributing' },
]

const scrollTo = (id) => {
  const element = document.getElementById(id)
  if (element) {
    element.scrollIntoView({ behavior: 'smooth', block: 'start' })
  }
}
</script>
