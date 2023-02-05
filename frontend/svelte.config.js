import sveltePreprocess from 'svelte-preprocess';
import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter({
			fallback: 'index.html',
			precompress: true
		}),

		prerender: {
			entries: []
		},
		version: {
			pollInterval: 60 * 1000
		}
	},
	preprocess: sveltePreprocess({
		sourceMap: true,
		scss: {
			prependData: `
				@import 'src/styles/config/variables.scss';
				@import 'src/styles/mixins/mixins.scss';
			`
		},
		preserve: ['ld+json']
	}),
	onwarn: (warning, handler) => {
        const { code } = warning;
        if (code === "css-unused-selector")
            return;

        handler(warning);
    }
};

export default config;