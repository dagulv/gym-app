import path from 'path';
import { sveltekit } from '@sveltejs/kit/vite';

/** @type {import('vite').UserConfig} */
const config = {
	plugins: [sveltekit()],
	resolve: {
		alias: {
			$components: path.resolve('./src/components'),
			$iconComponents: path.resolve('./src/components/icon-components'),
			$lib: path.resolve('./src/lib'),
		}
	},
	build: {
		sourcemap: false
	},
};

export default config;
