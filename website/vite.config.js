import tailwindcss from '@tailwindcss/vite';
// Site statique: remplacer @sveltejs/adapter-auto par @svelte/adapter-static
// import adapter from '@sveltejs/adapter-auto';
import adapter from '@sveltejs/adapter-static';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

// On vérifie si la commande en cours est "build"
const isBuilding = process.argv.includes('build');

export default defineConfig({
	plugins: [
		tailwindcss(),
		sveltekit({
			compilerOptions: {
				// Force runes mode for the project, except for libraries. Can be removed in svelte 6.
				runes: ({ filename }) =>
					filename.split(/[/\\]/).includes('node_modules') ? undefined : true
			},

			// adapter-auto only supports some environments, see https://svelte.dev/docs/kit/adapter-auto for a list.
			// If your environment is not supported, or you settled on a specific environment, switch out the adapter.
			// See https://svelte.dev/docs/kit/adapters for more information about adapters.
			//Site statique
			// adapter: adapter()
			adapter: adapter({
					pages: 'docs',
					assets: 'docs',
					fallback: undefined,
					precompress: false,
					strict: true
				}),
			paths: {
				base: process.env.NODE_ENV === 'production' ? '/svelte_website_static' : '',
				relative: isBuilding
			}
		})
	]
});

