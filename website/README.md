# Zettl Marketing Website

A modern, responsive marketing website for Zettl built with Vue 3, Vite, and Tailwind CSS.

## Features

- 🎨 **Modern Design** - Clean, professional design with purple gradient theme
- 📱 **Fully Responsive** - Mobile-first design that looks great on all devices
- ⚡ **Lightning Fast** - Built with Vite for instant hot module replacement
- 🎯 **SEO Optimized** - Semantic HTML and meta tags for search engines
- 🌙 **Dark Theme** - Eye-friendly dark color scheme
- 🔗 **Vue Router** - Smooth page transitions with Vue Router

## Pages

- **Home** - Hero section, features overview, installation guide, CTA
- **Features** - Detailed feature showcase with technical specs
- **Pricing** - Transparent pricing with cost breakdown and examples
- **Docs** - Complete documentation with sidebar navigation

## Tech Stack

- **Vue 3** - Progressive JavaScript framework
- **Vite** - Next-generation frontend tooling
- **Vue Router** - Official router for Vue.js
- **Tailwind CSS** - Utility-first CSS framework
- **PostCSS** - CSS transformations

## Getting Started

### Prerequisites

- Node.js 18+
- npm or yarn

### Installation

```bash
# Navigate to website directory
cd website

# Install dependencies
npm install

# Start development server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview
```

## Project Structure

```
website/
├── public/              # Static assets
├── src/
│   ├── assets/         # Styles and images
│   │   └── main.css    # Tailwind CSS imports
│   ├── components/     # Reusable components
│   │   ├── Navigation.vue
│   │   └── Footer.vue
│   ├── views/          # Page components
│   │   ├── Home.vue
│   │   ├── Features.vue
│   │   ├── Pricing.vue
│   │   └── Docs.vue
│   ├── router/         # Vue Router configuration
│   ├── App.vue         # Root component
│   └── main.js         # Application entry point
├── index.html          # HTML template
├── vite.config.js      # Vite configuration
├── tailwind.config.js  # Tailwind configuration
└── package.json        # Dependencies
```

## Development

### Running Locally

```bash
npm run dev
```

Visit `http://localhost:5173` in your browser.

### Building for Production

```bash
npm run build
```

The built files will be in the `dist` directory.

### Deployment

The site can be deployed to any static hosting service:

- **Vercel** - Automatic deployments from Git
- **Netlify** - Drop the `dist` folder
- **GitHub Pages** - Push `dist` to gh-pages branch
- **Cloudflare Pages** - Connect your repository

#### Vercel Deployment

```bash
# Install Vercel CLI
npm i -g vercel

# Deploy
vercel
```

#### Netlify Deployment

```bash
# Build the project
npm run build

# Deploy dist folder
npx netlify deploy --prod --dir=dist
```

## Customization

### Colors

Edit `tailwind.config.js` to change the color scheme:

```js
theme: {
  extend: {
    colors: {
      'zettl-purple': '#7D56F4',  // Primary color
      'zettl-dark': '#1a1a1a',     // Background
      'zettl-gray': '#2a2a2a',     // Card background
    },
  },
},
```

### Content

All page content is in the `src/views/` directory. Edit Vue components to change content.

### Components

Reusable components are in `src/components/`. Current components:

- `Navigation.vue` - Top navigation bar
- `Footer.vue` - Footer with links

## License

Same as the parent Zettl project.

## Contributing

Contributions welcome! Please see the main project README for guidelines.

---

Built with ❤️ for Zettl
