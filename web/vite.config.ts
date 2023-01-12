import { sveltekit } from '@sveltejs/kit/vite';
import type { UserConfig } from 'vite';

const config: UserConfig = {
	plugins: [sveltekit()],
	server: {
		fs: {
			// allow: ".."
		},
		port: 3000,
		proxy: {
			'/api':{
				target: 'http://127.0.0.1:8080',
				secure: false,
			},

		},
	}	
};

export default config;
