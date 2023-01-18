<script lang="ts">
	import { page } from '$app/stores';
	import api from '$lib/api';
	import { onDestroy } from 'svelte';
	import AdminMenu from './menu/menu-page.svelte';
	import AdminPages from './pages/pages-page.svelte';

	let error = '';
	let kind = '';

	let component: any;

	function refresh(newPath: string) {
		error = '';
	}

	let unsubscribe = page.subscribe((data) => {
		kind = '';
		api.get('page?url=/admin/' + data.params.slug).then((resp) => {
			if (resp.status == 0) {
				switch (resp.data.layout) {
					case 'admin.menu':
						component = AdminMenu;
						break;
					case 'admin.pages':
						component = AdminPages;
						break;
					default:
					// layoutNode = new LayoutNode().parse(resp.data.layout);
				}
			} else {
				error = resp.message;
			}
		});
	});

	onDestroy(unsubscribe);
</script>

{#if error != ''}
	<h2>{error}</h2>
{/if}

<svelte:component this={component} />
