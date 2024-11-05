import { defineConfig } from 'astro/config';
import preact from '@astrojs/preact';

import tailwind from '@astrojs/tailwind'

// https://astro.build/config
export default defineConfig({
	integrations: [preact({ compat: true }), tailwind()],
	vite: {
		resolve: {
			alias: {
				react: 'preact/compat',
				'react-dom': 'preact/compat',
			},
		},
	},
});
