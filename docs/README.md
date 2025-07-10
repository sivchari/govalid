# govalid Documentation

This directory contains the Hugo-based documentation website for govalid.

## Structure

```
docs/
├── hugo.toml           # Hugo configuration
├── content/            # Markdown content files
│   ├── _index.md       # Home page
│   ├── getting-started.md
│   ├── validators.md
│   ├── benchmarks.md
│   └── examples.md
├── layouts/            # Hugo templates
│   ├── _default/
│   │   ├── baseof.html
│   │   ├── single.html
│   │   └── list.html
│   └── index.html
├── static/             # Static assets
│   ├── css/
│   │   └── custom.css
│   └── js/
└── public/            # Generated site (created by Hugo)
```

## Local Development

1. Install Hugo:
   ```bash
   # macOS
   brew install hugo
   
   # Linux
   wget https://github.com/gohugoio/hugo/releases/download/v0.128.0/hugo_extended_0.128.0_linux-amd64.deb
   sudo dpkg -i hugo_extended_0.128.0_linux-amd64.deb
   ```

2. Run the development server:
   ```bash
   cd docs
   hugo server -D
   ```

3. Visit http://localhost:1313

## Building for Production

```bash
cd docs
hugo --minify
```

The generated site will be in the `public/` directory.

## GitHub Pages Deployment

The site is automatically deployed to GitHub Pages when changes are pushed to the main branch. The workflow is defined in `.github/workflows/deploy-docs.yml`.

## Customization

- **Styling**: Edit `static/css/custom.css`
- **Layout**: Edit templates in `layouts/`
- **Content**: Edit markdown files in `content/`
- **Configuration**: Edit `hugo.toml`

## Features

- **Responsive design** with mobile-first approach
- **Syntax highlighting** for code blocks
- **Copy-to-clipboard** functionality for code snippets
- **Performance optimized** with minification
- **SEO friendly** with meta tags and Open Graph support
- **Fast loading** with optimized CSS and minimal JavaScript